package response

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}
