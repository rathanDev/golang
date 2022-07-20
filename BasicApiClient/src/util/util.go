package util

import (
	"fmt"
	"net/http"
)

func CreateBooleanPointer(val bool) *bool {
	return &val
}

func CreateStringPointer(val string) *string {
	return &val
}

func CreateNumberPointer(val int64) *int64 {
	return &val
}

func PrintHttpResponse(resp *http.Response) {
	fmt.Println("response:", resp)
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
}
