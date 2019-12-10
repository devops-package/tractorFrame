package Guzzle

import (
	"bytes"
	"context"
	"fmt"
	"github.com/hashicorp/go-cleanhttp"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	// HTTPAddrEnvName defines an environment variable name which sets
	// the HTTP address if there is no -http-addr specified.
	HTTPAddrEnvName = "Guzzle_HTTP_ADDR"

	// HTTPTokenEnvName defines an environment variable name which sets
	// the HTTP token.
	HTTPTokenEnvName = "Guzzle_HTTP_TOKEN"

	// HTTPTokenFileEnvName defines an environment variable name which sets
	// the HTTP token file.
	HTTPTokenFileEnvName = "Guzzle_HTTP_TOKEN_FILE"

	// HTTPAuthEnvName defines an environment variable name which sets
	// the HTTP authentication header.
	HTTPAuthEnvName = "Guzzle_HTTP_AUTH"

	// HTTPSSLEnvName defines an environment variable name which sets
	// whether or not to use HTTPS.
	HTTPSSLEnvName = "Guzzle_HTTP_SSL"

	// HTTPCAFile defines an environment variable name which sets the
	// CA file to use for talking to Guzzle over TLS.
	HTTPCAFile = "Guzzle_CACERT"

	// HTTPCAPath defines an environment variable name which sets the
	// path to a directory of CA certs to use for talking to Guzzle over TLS.
	HTTPCAPath = "Guzzle_CAPATH"

	// HTTPClientCert defines an environment variable name which sets the
	// client cert file to use for talking to Guzzle over TLS.
	HTTPClientCert = "Guzzle_CLIENT_CERT"

	// HTTPClientKey defines an environment variable name which sets the
	// client key file to use for talking to Guzzle over TLS.
	HTTPClientKey = "Guzzle_CLIENT_KEY"

	// HTTPTLSServerName defines an environment variable name which sets the
	// server name to use as the SNI host when connecting via TLS
	HTTPTLSServerName = "Guzzle_TLS_SERVER_NAME"

	// HTTPSSLVerifyEnvName defines an environment variable name which sets
	// whether or not to disable certificate checking.
	HTTPSSLVerifyEnvName = "Guzzle_HTTP_SSL_VERIFY"

	// GRPCAddrEnvName defines an environment variable name which sets the gRPC
	// address for Guzzle connect envoy. Note this isn't actually used by the api
	// client in this package but is defined here for consistency with all the
	// other ENV names we use.
	GRPCAddrEnvName = "Guzzle_GRPC_ADDR"
)

// Config is used to configure the creation of a client
type Config struct {
	// Address is the address of the Guzzle server
	Address string

	// Scheme is the URI scheme for the Guzzle server
	Scheme string

	// Transport is the Transport to use for the http client.
	Transport *http.Transport

	// HttpClient is the client to use. Default will be
	// used if not provided.
	HttpClient *http.Client

	// HttpAuth is the auth info to use for http access.
	HttpAuth *HttpBasicAuth

	// WaitTime limits how long a Watch will block. If not provided,
	// the agent default values will be used.
	WaitTime time.Duration

	// Token is used to provide a per-request ACL token
	// which overrides the agent's default token.
	Token string

	// TokenFile is a file containing the current token to use for this client.
	// If provided it is read once at startup and never again.
	TokenFile string

	// Namespace is the name of the namespace to send along for the request
	// when no other Namespace ispresent in the QueryOptions
	Namespace string

	TLSConfig TLSConfig
}

// HttpBasicAuth is used to authenticate http client with HTTP Basic Authentication
type HttpBasicAuth struct {
	// Username to use for HTTP Basic Authentication
	Username string

	// Password to use for HTTP Basic Authentication
	Password string
}

// TLSConfig is used to generate a TLSClientConfig that's useful for talking to
// Guzzle using TLS.
type TLSConfig struct {
	// Address is the optional address of the Guzzle server. The port, if any
	// will be removed from here and this will be set to the ServerName of the
	// resulting config.
	Address string

	// CAFile is the optional path to the CA certificate used for Guzzle
	// communication, defaults to the system bundle if not specified.
	CAFile string

	// CAPath is the optional path to a directory of CA certificates to use for
	// Guzzle communication, defaults to the system bundle if not specified.
	CAPath string

	// CertFile is the optional path to the certificate for Guzzle
	// communication. If this is set then you need to also set KeyFile.
	CertFile string

	// KeyFile is the optional path to the private key for Guzzle communication.
	// If this is set then you need to also set CertFile.
	KeyFile string

	// InsecureSkipVerify if set to true will disable TLS host verification.
	InsecureSkipVerify bool
}

// request is used to help build up a request
type request struct {
	config *Config
	method string
	url    *url.URL
	params url.Values
	body   io.Reader
	header http.Header
	obj    interface{}
	ctx    context.Context
}


type Client struct {
	config Config
}

// NewClient returns a new client
func NewClient(config *Config) (*Client, error) {
	// bootstrap the config
	defConfig := DefaultConfig()

	if len(config.Address) == 0 {
		config.Address = defConfig.Address
	}

	if len(config.Scheme) == 0 {
		config.Scheme = defConfig.Scheme
	}

	if config.Transport == nil {
		config.Transport = defConfig.Transport
	}

	if config.TLSConfig.Address == "" {
		config.TLSConfig.Address = defConfig.TLSConfig.Address
	}

	if config.TLSConfig.CAFile == "" {
		config.TLSConfig.CAFile = defConfig.TLSConfig.CAFile
	}

	if config.TLSConfig.CAPath == "" {
		config.TLSConfig.CAPath = defConfig.TLSConfig.CAPath
	}

	if config.TLSConfig.CertFile == "" {
		config.TLSConfig.CertFile = defConfig.TLSConfig.CertFile
	}

	if config.TLSConfig.KeyFile == "" {
		config.TLSConfig.KeyFile = defConfig.TLSConfig.KeyFile
	}

	if !config.TLSConfig.InsecureSkipVerify {
		config.TLSConfig.InsecureSkipVerify = defConfig.TLSConfig.InsecureSkipVerify
	}

	if config.HttpClient == nil {
		var err error
		config.HttpClient, err = NewHttpClient(config.Transport, config.TLSConfig)
		if err != nil {
			return nil, err
		}
	}

	parts := strings.SplitN(config.Address, "://", 2)
	if len(parts) == 2 {
		switch parts[0] {
		case "http":
			config.Scheme = "http"
		case "https":
			config.Scheme = "https"
		case "unix":
			trans := cleanhttp.DefaultTransport()
			trans.DialContext = func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", parts[1])
			}
			config.HttpClient = &http.Client{
				Transport: trans,
			}
		default:
			return nil, fmt.Errorf("Unknown protocol scheme: %s", parts[0])
		}
		config.Address = parts[1]
	}

	// If the TokenFile is set, always use that, even if a Token is configured.
	// This is because when TokenFile is set it is read into the Token field.
	// We want any derived clients to have to re-read the token file.
	if config.TokenFile != "" {
		data, err := ioutil.ReadFile(config.TokenFile)
		if err != nil {
			return nil, fmt.Errorf("Error loading token file: %s", err)
		}

		if token := strings.TrimSpace(string(data)); token != "" {
			config.Token = token
		}
	}
	if config.Token == "" {
		config.Token = defConfig.Token
	}

	return &Client{config: *config}, nil
}

// toHTTP converts the request to an HTTP request
func (r *request) toHTTP() (*http.Request, error) {
	// Encode the query parameters
	r.url.RawQuery = r.params.Encode()

	// Check if we should encode the body
	if r.body == nil && r.obj != nil {
		b, err := encodeBody(r.obj)
		if err != nil {
			return nil, err
		}
		r.body = b
	}

	// Create the HTTP request
	req, err := http.NewRequest(r.method, r.url.RequestURI(), r.body)
	if err != nil {
		return nil, err
	}

	req.URL.Host = r.url.Host
	req.URL.Scheme = r.url.Scheme
	req.Host = r.url.Host
	req.Header = r.header

	// Setup auth
	if r.config.HttpAuth != nil {
		req.SetBasicAuth(r.config.HttpAuth.Username, r.config.HttpAuth.Password)
	}
	if r.ctx != nil {
		return req.WithContext(r.ctx), nil
	}

	return req, nil
}

// doRequest runs a request with our client
func (c *Client) NewDoRequest(r *request) (time.Duration, *http.Response, error) {
	req, err := r.toHTTP()
	if err != nil {
		return 0, nil, err
	}
	start := time.Now()
	resp, err := c.config.HttpClient.Do(req)
	diff := time.Since(start)


	return diff, resp, err
}

// newRequest is used to create a new request
func (c *Client) DoNewRequest(method, path string) *request {
	r := &request{
		config: &c.config,
		method: method,
		url: &url.URL{
			Scheme: c.config.Scheme,
			Host:   c.config.Address,
			Path:   path,
		},
		params: make(map[string][]string),
		header: make(http.Header),
	}
	if c.config.Namespace != "" {
		r.params.Set("ns", c.config.Namespace)
	}
	if c.config.WaitTime != 0 {
		r.params.Set("wait", durToMsec(r.config.WaitTime))
	}
	if c.config.Token != "" {
		r.header.Set("X-Consul-Token", r.config.Token)
	}
	return r
}

// requireOK is used to wrap doRequest and check for a 200
func RequireOK(d time.Duration, resp *http.Response, e error) (time.Duration, *http.Response, error) {
	defaultResp := &Response{}
	if e != nil {
		if resp != nil {
			resp.Body.Close()
		}
		defaultResp.StatusCode = resp.StatusCode

		return d, nil, e
	}
	if resp.StatusCode != 200 {
		return d, nil, generateUnexpectedResponseCodeError(resp)
	}
	return d, resp, nil
}

// generateUnexpectedResponseCodeError consumes the rest of the body, closes
// the body stream and generates an error indicating the status code was
// unexpected.
func generateUnexpectedResponseCodeError(resp *http.Response) error {
	var buf bytes.Buffer
	io.Copy(&buf, resp.Body)
	resp.Body.Close()

	return fmt.Errorf("Unexpected response code: %d (%s)", resp.StatusCode, buf.Bytes())
}

func requireNotFoundOrOK(d time.Duration, resp *http.Response, e error) (bool, time.Duration, *http.Response, error) {
	if e != nil {
		if resp != nil {
			resp.Body.Close()
		}
		return false, d, nil, e
	}
	switch resp.StatusCode {
	case 200:
		return true, d, resp, nil
	case 404:
		return false, d, resp, nil
	default:
		return false, d, nil, generateUnexpectedResponseCodeError(resp)
	}
}

