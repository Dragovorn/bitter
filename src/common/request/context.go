package request

type LambdaContext struct {
	BodyRaw   []byte
	Body      string
	RawPath   string
	RawQuery  string
	Query     map[string]string
	Path      map[string]string
	WasBase64 bool
}

func (c *LambdaContext) UnmarshalBody(do interface{}) *LambdaHTTPResponse {
	return UnmarshalJson(c.Body, do)
}
