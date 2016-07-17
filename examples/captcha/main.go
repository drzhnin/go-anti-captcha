package main

import (
	"fmt"
	"time"

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
	for {
		result, err := client.Captcha.GetText(captchaID)
		if err != nil {
			if result == "CAPCHA_NOT_READY" {
				time.Sleep(time.Second * 3)
				continue
			}
			fmt.Printf("error: %v\n\n", err)
			break
		} else {
			fmt.Printf("Captcha text: %s\n", result)
			break
		}
	}

}
