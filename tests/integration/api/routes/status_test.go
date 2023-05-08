package routes

import (
	"net/http"
	"net/http/httptest"
	"pokedex/src/api/router"
	"testing"
)

func TestStatusEndpoint(t *testing.T) {
	statusEndpoint := "/status"
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	server := http.NewServeMux()
	server.HandleFunc(statusEndpoint, router.RouteStatusRequest)

	tests := []struct {
		name           string
		method         string
		isAuthorized   bool
		expectedStatus int
	}{
		{name: "ReturnUnauthorizedWhenNotAuthorized", method: http.MethodGet, expectedStatus: http.StatusUnauthorized, isAuthorized: false},
		{name: "ReturnsOKWhenAuthorized", method: http.MethodGet, expectedStatus: http.StatusOK, isAuthorized: true},
		{name: "ReturnUnsupportedWhenUsingPUTHttpMethod", method: http.MethodPut, expectedStatus: http.StatusNotImplemented, isAuthorized: true},
		{name: "ReturnUnsupportedWhenUsingPOSTHttpMethod", method: http.MethodPost, expectedStatus: http.StatusNotImplemented, isAuthorized: true},
		{name: "ReturnUnsupportedWhenUsingPATCHHttpMethod", method: http.MethodPatch, expectedStatus: http.StatusNotImplemented, isAuthorized: true},
		{name: "ReturnUnsupportedWhenUsingDELETEHttpMethod", method: http.MethodDelete, expectedStatus: http.StatusNotImplemented, isAuthorized: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest(test.method, statusEndpoint, nil)
			if test.isAuthorized {
				req.Header.Set("Authorization", "valid-token")
			}

			server.ServeHTTP(w, req)

			if test.expectedStatus != w.Code {
				t.Fatalf("Expected status code of: '%d' but got: '%d'", test.expectedStatus, w.Code)
			}
		})
	}
}
