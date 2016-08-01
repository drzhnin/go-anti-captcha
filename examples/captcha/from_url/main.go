package main

import (
	"fmt"

	"github.com/andrewdruzhinin/go-anti-captcha/anticaptcha"
)

func main() {
	client := anticaptcha.NewClient("apiKey") //Set your apiKey
	client.CaptchaAdditionalParams.EnableRegsense()
	ID, err := client.Captcha.UploadCaptchaFromURL("https://s3-us-west-2.amazonaws.com/captcha-test/1045.png")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("Captcha ID: %d\n", ID)
		res, err := client.Captcha.GetText(ID)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		fmt.Printf("Captcha Text: %s\n", res)
	}

}
