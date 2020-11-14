package model

import (
	"encoding/json"
	"net/http"
)

//APIResponder return struct for api
type APIResponder struct {
	Data     interface{}            `json:"data"`
	Code     int                    `json:"code"`
	Meta     map[string]interface{} `json:"meta"`
	Hostname string                 `json:"hostname"`
}

//WriteResponse return json for APIResponder
func (res *APIResponder) WriteResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res.Hostname = r.Host
	data, err := json.Marshal(res)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(data)
}
