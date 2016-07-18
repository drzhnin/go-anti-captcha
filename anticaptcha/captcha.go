package anticaptcha

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

//CaptchaService handles communication with the captcha actions of the
//anti-captcha API
type CaptchaService service

//Captcha response captcha struct
type Captcha struct {
	ID   int
	Text string
}

//UploadCaptchaFromFile represents updaload image(jpg, gif, png) to http://anti-captcha.com API
//and get capthca ID
func (s *CaptchaService) UploadCaptchaFromFile(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, errors.New("File does not exist")
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		return 0, err
	}
	_, err = io.Copy(part, file)
	_ = writer.WriteField("key", s.client.APIKey)
	_ = writer.WriteField("method", "post")
	err = writer.Close()
	if err != nil {
		return 0, err
	}
	req, err := s.client.NewRequest("POST", "in.php", body)
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	data, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	captcha := new(Captcha)
	splited := strings.Split(string(data), "|")

	if splited[0] == "OK" {
		captchaID := splited[1]
		captcha.ID, err = strconv.Atoi(captchaID)
		if err != nil {
			return 0, err
		}
		return captcha.ID, nil
	}

	return 0, errors.New(splited[0])
}

//GetText represents return captcha text by ID
func (s *CaptchaService) GetText(id int) (string, error) {
	reqURL := fmt.Sprintf("res.php?key=%s&action=get&id=%d", s.client.APIKey, id)

	req, err := s.client.NewRequest("GET", reqURL, nil)
	if err != nil {
		return "", err
	}

	res, err := s.client.Do(req)
	if err != nil {
		return "", err
	}

	captcha := new(Captcha)
	splited := strings.Split(string(res), "|")

	if splited[0] == "OK" {
		captchaText := splited[1]
		captcha.Text = string(captchaText)
		if err != nil {
			return "", err
		}
		return captcha.Text, nil
	}

	return splited[0], errors.New("Captcha error")

}
