package CD

// cms response struct
type CmsData struct {
	HTTPCode   int    `json:"http_code"`
	StatusCode string `json:"statusCode"`
	Msg        string `json:"msg"`
}
