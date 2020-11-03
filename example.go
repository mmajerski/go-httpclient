package main

import (
	"fmt"

	"github.com/userq11/go-httpclient/gohttp"
)

var githubHttpClient = getGithubClient()

func getGithubClient() gohttp.Client {
	client := gohttp.NewBuilder().DisableTimeouts(true).SetMaxIdleConnections(5).SetHeaders(nil).Build()

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

	// Custom response
	fmt.Println(response.Status())
	fmt.Println(response.StatusCode())
	fmt.Println(response.String())

	// Default http.Response:
	// response.Body.Close()

	// fmt.Println(response.StatusCode)

	// bytes2, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// var user2 User
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	panic(err)
	// }
	// fmt.Println(user2.FirstName)
}
