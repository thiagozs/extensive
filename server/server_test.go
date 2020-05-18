package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method,
	path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestServerHandlerEndpoint(t *testing.T) {

	tests := []struct {
		endpoint string
		expected string
		code     int
	}{
		{endpoint: "/1", code: 200, expected: `{"extenso":"um"}`},
		{endpoint: "/2", code: 200, expected: `{"extenso":"dois"}`},
		{endpoint: "/-1", code: 200, expected: `{"extenso":"menos um"}`},
		{endpoint: "/0", code: 200, expected: `{"extenso":"zero"}`},
		{endpoint: "/-0", code: 200, expected: `{"extenso":"zero"}`},
		{endpoint: "/aaA", code: 400, expected: `{"error":"invalid content"}`},
		{endpoint: "/--", code: 400, expected: `{"error":"invalid content"}`},
		{endpoint: "/@", code: 400, expected: `{"error":"invalid content"}`},
		{endpoint: "/%20A", code: 400, expected: `{"error":"invalid content"}`},
	}

	server := New(":8080", true)
	server.RegisterRoutes()

	//s := gin.Default()
	//s.GET("/:number", server.GetFromAPI())

	for _, test := range tests {

		t.Run(test.endpoint, func(t *testing.T) {

			w := performRequest(server.Engine, "GET", test.endpoint)

			var response map[string]interface{}
			errUmarshal := json.Unmarshal(w.Body.Bytes(), &response)
			body, errMarshal := json.Marshal(response)

			bstr := strings.TrimSpace(string(body))

			assert.Nil(t, errUmarshal)
			assert.Nil(t, errMarshal)
			assert.Equal(t, bstr, test.expected)
			assert.Equal(t, w.Code, test.code)

		})
	}

}
