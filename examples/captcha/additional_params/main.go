package main

import (
	"fmt"

	"github.com/andrewdruzhinin/go-anti-captcha/anticaptcha"
)

func main() {
	client := anticaptcha.NewClient("APIKey")
	client.CaptchaAdditionalParams.EnablePhrate()
	client.CaptchaAdditionalParams.EnableAllowEmpty()
	client.CaptchaAdditionalParams.AddComment("What color is the sky?")
	fmt.Println(client.CaptchaAdditionalParams.LoadParams())
}
