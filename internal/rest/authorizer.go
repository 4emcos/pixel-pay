package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type AuthorizerResponse struct {
	Authorization bool `json:"authorization"`
}

var (
	authorizationURL = os.Getenv("AUTHORIZER_HOST")
)

func Authorizer() error {
	if authorizationURL == "" {
		authorizationURL = "localhost"
	}
	resp, err := http.Get("http://" + authorizationURL + ":8099/authorization")

	if err != nil {
		return fmt.Errorf("authorizer call err")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("authorizer body parser err")
	}

	var authorizerResponse AuthorizerResponse
	err = json.Unmarshal(body, &authorizerResponse)

	if authorizerResponse.Authorization {
		return nil
	}

	return fmt.Errorf("authorizer err")

}
