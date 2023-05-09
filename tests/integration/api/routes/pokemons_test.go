package routes

import (
	"net/http"
	"net/http/httptest"
	"pokedex/src/api/router"
	"pokedex/src/domain/service"
	"pokedex/tests/integration/resources"
	"testing"
)

func TestGetPokemonsEndpoint(t *testing.T) {
	pokemonsEndpoint := "/pokemons"
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	router := router.PokemonRouteHandler{
		PokemonService: service.PokemonService{},
	}

	server := http.NewServeMux()
	server.HandleFunc(pokemonsEndpoint, router.RoutePokemonsRequest)

	tests := []struct {
		name           string
		queryString    string
		expectedStatus int
		isAuthorized   bool
		expectedBody   string
	}{
		{name: "ReturnUnauthorizedWhenNotAuthorized", expectedStatus: http.StatusUnauthorized, expectedBody: "{\"status\":\"Unauthorized\"}", isAuthorized: false},
		{name: "ReturnBulbasaurWhenRequestingId1", queryString: "?id=1", expectedStatus: http.StatusOK, expectedBody: resources.Bulbasaur, isAuthorized: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, pokemonsEndpoint+test.queryString, nil)
			if test.isAuthorized {
				req.Header.Set("Authorization", "valid-token")
			}

			server.ServeHTTP(w, req)

			if test.expectedBody != w.Body.String() {
				t.Fatalf("Expected status code of: '%s' but got: '%s'", test.expectedBody, w.Body.String())
			}

			if test.expectedStatus != w.Code {
				t.Fatalf("Expected status code of: '%d' but got: '%d'", test.expectedStatus, w.Code)
			}
		})
	}
}
