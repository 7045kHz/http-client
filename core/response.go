package core

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type Response struct {
	Status     string
	StatusCode int
	Headers    http.Header
	Body       []byte
}

func (r *Response) Bytes() []byte {
	return r.Body
}

func (r *Response) String() string {
	return string(r.Body)
}

func (r *Response) UnmarshalJson(target interface{}) error {
	return json.Unmarshal(r.Bytes(), target)
}

func (r *Response) UnmarshalXml(target interface{}) error {
	return xml.Unmarshal(r.Bytes(), target)
}
