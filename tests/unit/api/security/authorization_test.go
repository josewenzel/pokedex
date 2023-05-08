package security

import (
	"net/http"
	"net/http/httptest"
	"pokedex/src/api/security"
	"testing"
)

func TestAuthorization(t *testing.T) {
	tests := []struct {
		name          string
		token         string
		expectedValid bool
	}{
		{name: "ReturnsTrueIfTokenIsValid", token: "valid-token", expectedValid: true},
		{name: "ReturnsFalseIfTokenIsInvalid", token: "unauthorized-token", expectedValid: false},
		{name: "ReturnsFalseIfTokenIsMissing", token: "", expectedValid: false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			someMethod := http.MethodGet
			request := httptest.NewRequest(someMethod, "/", nil)
			request.Header.Set("Authorization", test.token)

			authorizationResponse := security.AuthorizeRequest(request)

			if authorizationResponse.Valid != test.expectedValid {
				t.Fatalf("Expected valid to be: '%v' but was: '%v'", authorizationResponse.Valid, test.expectedValid)
			}
		})
	}
}
