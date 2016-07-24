package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"

	"github.com/andrewdruzhinin/go-anti-captcha/anticaptcha"
)

func main() {
	client := anticaptcha.NewClient("api_key") //Set your apiKey
	content, err := ioutil.ReadFile("captcha.png")
	if err != nil {
		fmt.Println(err)
	}
	str := base64.StdEncoding.EncodeToString(content)
	ID, err := client.Captcha.UploadCaptchaFromBase64(str)
	if err != nil {
		fmt.Printf("error: %v\n\n", err)
	} else {
		fmt.Printf("Captcha ID: %d\n", ID)
	}
	res, err := client.Captcha.GetText(ID)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("Captcha Text: %s\n", res)
}
