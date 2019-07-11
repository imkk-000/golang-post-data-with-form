package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func checkPanic(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func login(voucherCode string) {
	calledURL := "http://10.0.2.1:8002/index.php?zone=cpzone"
	formData := url.Values{}
	formData.Set("auth_user", "")
	formData.Set("auth_pass", "")
	formData.Set("auth_voucher", voucherCode)
	formData.Set("redirurl", "http://www.apple.com/library/test/success.html")
	formData.Set("accept", "Accept")

	response, err := http.PostForm(calledURL, formData)
	checkPanic(err)
	responseBody, err := ioutil.ReadAll(response.Body)
	checkPanic(err)
	fmt.Println(string(responseBody))
}

func main() {
	voucherCode := flag.String("voucher_code", "", "Set voucher code when you bought something at Cafe")
	flag.Parse()
	if len(*voucherCode) > 0 {
		login(*voucherCode)
	}
}
