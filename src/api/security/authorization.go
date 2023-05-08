package security

import (
	"errors"
	"net/http"
)

type AuthorizationResponse struct {
	Valid   bool
	Message error
}

func AuthorizeRequest(r *http.Request) *AuthorizationResponse {
	if authorizationToken := r.Header.Get("Authorization"); authorizationToken == "valid-token" {
		return &AuthorizationResponse{
			Valid:   true,
			Message: nil,
		}
	}

	return &AuthorizationResponse{
		Valid:   false,
		Message: errors.New("unauthorized token"),
	}
}
