package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
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

func main() {

    url := getUrl()

    response, err := http.Get(url)

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(responseData))

}