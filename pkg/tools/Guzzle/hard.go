package Guzzle

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)


// DefaultConfig returns a default configuration for the client. By default this
// will pool and reuse idle connections to Guzzle. If you have a long-lived
// client object, this is the desired behavior and should make the most efficient
// use of the connections to Guzzle. If you don't reuse a client object, which
// is not recommended, then you may notice idle connections building up over
// time. To avoid this, use the DefaultNonPooledConfig() instead.
func DefaultConfig() *Config {
	return defaultConfig(DefaultPooledTransport)
}

// defaultConfig returns the default configuration for the client, using the
// given function to make the transport.
func defaultConfig(transportFn func() *http.Transport) *Config {
	config := &Config{
		Address:   "127.0.0.1:8500",
		Scheme:    "http",
		Transport: transportFn(),
	}

	if addr := os.Getenv(HTTPAddrEnvName); addr != "" {
		config.Address = addr
	}

	if tokenFile := os.Getenv(HTTPTokenFileEnvName); tokenFile != "" {
		config.TokenFile = tokenFile
	}

	if token := os.Getenv(HTTPTokenEnvName); token != "" {
		config.Token = token
	}

	if auth := os.Getenv(HTTPAuthEnvName); auth != "" {
		var username, password string
		if strings.Contains(auth, ":") {
			split := strings.SplitN(auth, ":", 2)
			username = split[0]
			password = split[1]
		} else {
			username = auth
		}

		config.HttpAuth = &HttpBasicAuth{
			Username: username,
			Password: password,
		}
	}

	if ssl := os.Getenv(HTTPSSLEnvName); ssl != "" {
		enabled, err := strconv.ParseBool(ssl)
		if err != nil {
			log.Printf("[WARN] client: could not parse %s: %s", HTTPSSLEnvName, err)
		}

		if enabled {
			config.Scheme = "https"
		}
	}

	if v := os.Getenv(HTTPTLSServerName); v != "" {
		config.TLSConfig.Address = v
	}
	if v := os.Getenv(HTTPCAFile); v != "" {
		config.TLSConfig.CAFile = v
	}
	if v := os.Getenv(HTTPCAPath); v != "" {
		config.TLSConfig.CAPath = v
	}
	if v := os.Getenv(HTTPClientCert); v != "" {
		config.TLSConfig.CertFile = v
	}
	if v := os.Getenv(HTTPClientKey); v != "" {
		config.TLSConfig.KeyFile = v
	}
	if v := os.Getenv(HTTPSSLVerifyEnvName); v != "" {
		doVerify, err := strconv.ParseBool(v)
		if err != nil {
			log.Printf("[WARN] client: could not parse %s: %s", HTTPSSLVerifyEnvName, err)
		}
		if !doVerify {
			config.TLSConfig.InsecureSkipVerify = true
		}
	}

	return config
}

// NewHttpClient returns an http client configured with the given Transport and TLS
// config.
func NewHttpClient(transport *http.Transport, tlsConf TLSConfig) (*http.Client, error) {
	client := &http.Client{
		Transport: transport,
	}

	// TODO (slackpad) - Once we get some run time on the HTTP/2 support we
	// should turn it on by default if TLS is enabled. We would basically
	// just need to call http2.ConfigureTransport(transport) here. We also
	// don't want to introduce another external dependency on
	// golang.org/x/net/http2 at this time. For a complete recipe for how
	// to enable HTTP/2 support on a transport suitable for the API client
	// library see agent/http_test.go:TestHTTPServer_H2.

	if transport.TLSClientConfig == nil {
		tlsClientConfig, err := SetupTLSConfig(&tlsConf)

		if err != nil {
			return nil, err
		}

		transport.TLSClientConfig = tlsClientConfig
	}

	return client, nil
}

// setQueryOptions is used to annotate the request with
// additional query options
func (r *request) setQueryOptions(q *QueryOptions) {
	if q == nil {
		return
	}
	if q.Namespace != "" {
		r.params.Set("ns", q.Namespace)
	}
	if q.Datacenter != "" {
		r.params.Set("dc", q.Datacenter)
	}
	if q.AllowStale {
		r.params.Set("stale", "")
	}
	if q.RequireConsistent {
		r.params.Set("consistent", "")
	}
	if q.WaitIndex != 0 {
		r.params.Set("index", strconv.FormatUint(q.WaitIndex, 10))
	}
	if q.WaitTime != 0 {
		r.params.Set("wait", durToMsec(q.WaitTime))
	}
	if q.WaitHash != "" {
		r.params.Set("hash", q.WaitHash)
	}
	if q.Token != "" {
		r.header.Set("X-Consul-Token", q.Token)
	}
	if q.Near != "" {
		r.params.Set("near", q.Near)
	}
	if q.Filter != "" {
		r.params.Set("filter", q.Filter)
	}
	if len(q.NodeMeta) > 0 {
		for key, value := range q.NodeMeta {
			r.params.Add("node-meta", key+":"+value)
		}
	}
	if q.RelayFactor != 0 {
		r.params.Set("relay-factor", strconv.Itoa(int(q.RelayFactor)))
	}
	if q.LocalOnly {
		r.params.Set("local-only", fmt.Sprintf("%t", q.LocalOnly))
	}
	if q.Connect {
		r.params.Set("connect", "true")
	}
	if q.UseCache && !q.RequireConsistent {
		r.params.Set("cached", "")

		cc := []string{}
		if q.MaxAge > 0 {
			cc = append(cc, fmt.Sprintf("max-age=%.0f", q.MaxAge.Seconds()))
		}
		if q.StaleIfError > 0 {
			cc = append(cc, fmt.Sprintf("stale-if-error=%.0f", q.StaleIfError.Seconds()))
		}
		if len(cc) > 0 {
			r.header.Set("Cache-Control", strings.Join(cc, ", "))
		}
	}

	r.ctx = q.ctx
}

