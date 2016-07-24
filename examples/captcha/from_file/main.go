package main

import (
	"fmt"

	"github.com/andrewdruzhinin/go-anti-captcha/anticaptcha"
)

func main() {
	client := anticaptcha.NewClient("api_key") //Set your apiKey
	captchaID, err := client.Captcha.UploadCaptchaFromFile("captcha.png")
	if err != nil {
		fmt.Printf("error: %v\n\n", err)
	} else {
		fmt.Printf("Captcha ID: %d\n", captchaID)
	}
	result, err := client.Captcha.GetText(captchaID)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("Captcha text: %s\n", result)
}
