package rootcert

type Config struct {
	// CAFile is a path to a PEM-encoded certificate file or bundle. Takes
	// precedence over CAPath.
	CAFile string

	// CAPath is a path to a directory populated with PEM-encoded certificates.
	CAPath string
}