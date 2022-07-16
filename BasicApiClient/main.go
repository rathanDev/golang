package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "api-client/helper"
    "api-client/model"
)

func main() {

    url := helper.GetUrl()

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

    fmt.Println("")

    accountData := model.AccountData {ID : "some-id", OrganisationID: "org-id"}
    fmt.Println(accountData)


}