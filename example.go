package gohttpclient

import "github.com/userq11/go-httpclient/gohttp"

func example() {
	client := gohttp.New()

	client.Get()
}
