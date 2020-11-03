package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/userq11/go-httpclient/gohttp"
)

var githubHttpClient = getGithubClient()

func getGithubClient() gohttp.HttpClient {
	client := gohttp.New()

	client.DisableTimeouts(true)

	commonHeaders := make(http.Header)
	commonHeaders.Set("Authorization", "Bearer asdjkp12p3o4emp2")
	client.SetHeaders(commonHeaders)

	return client
}

func main() {
	getUrls()
}

type User struct {
	FirstName string `json: "first_name"`
	LastName  string `json: "last_name"`
}

func getUrls() {
	response, err := githubHttpClient.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}

func createUser(user User) {
	response, err := githubHttpClient.Post("https://api.github.com", nil, user)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}
