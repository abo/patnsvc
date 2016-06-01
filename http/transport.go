package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/abo/patnsvc"
	"golang.org/x/net/context"
)

//RequestDecoder decode rpc request from http request
func RequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	var request patnsvc.Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

//ResponseEncoder encode rpc response to http response
func ResponseEncoder(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

//RequestEncoder encode rpc request to http request
func RequestEncoder(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

// ResponseDecoder decode rpc response from http response
func ResponseDecoder(_ context.Context, resp *http.Response) (interface{}, error) {
	var response patnsvc.Response
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}
