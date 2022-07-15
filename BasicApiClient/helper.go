package main

import (
	"fmt"
)

const baseUrl string = "http://localhost:8080"
var version string = "/v1";
const accountUrl string = "/organisation/accounts"

func getUrl() string {

	version = "/v1"

	fmt.Printf("BaseUrl is %v", baseUrl)

	var url = baseUrl + version + accountUrl
	fmt.Println(url)

	return url
}