package examples

import (
	"errors"
	"net/http"
	"testing"

	"github.com/userq11/go-httpclient/gohttp"
)

func TestGetEndpoints(t *testing.T) {
	t.Run("TestErrorFetchingFromGithub", func(t *testing.T) {
		mock := gohttp.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			Error:  errors.New("timeout getting github endpoints"),
		}

		endpoints, err := GetEndpoints()
	})

	t.Run("TestErrorUnmarshalResponseBody", func(t *testing.T) {
		mock := gohttp.Mock{
			Method:       http.MethodGet,
			Url:          "https://api.github.com",
			ResponseBody: `{"current_user_url": 123}`,
		}

		endpoints, err := GetEndpoints()
	})

	t.Run("TestNoError", func(t *testing.T) {
		mock := gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": "https://api.github.com/user"}`,
		}

		endpoints, err := GetEndpoints()
	})
}
