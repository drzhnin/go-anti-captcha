package anticaptcha

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
)

//CaptchaService handles communication with the captcha actions of the
//anti-captcha API
type CaptchaService service

//UploadCaptchaFromFile represents updaload image(jpg, gif, png) to http://anti-captcha.com API
//and get capthca ID
func (s *CaptchaService) UploadCaptchaFromFile(path string) (int, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return 0, errors.New("File does not exist")
	}
	strBase64 := base64.StdEncoding.EncodeToString(content)
	captchaID, err := s.uploadCaptcha(strBase64)
	if err != nil {
		return 0, err
	}
	return captchaID, err
}

//UploadCaptchaFromBase64 represents updaload base64 string to http://anti-captcha.com API
//and get capthca ID
func (s *CaptchaService) UploadCaptchaFromBase64(base64 string) (int, error) {
	captchaID, err := s.uploadCaptcha(base64)
	if err != nil {
		return 0, err
	}
	return captchaID, err
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

	resOk, err := checkResponse(res)
	if err != nil {
		if fmt.Sprintf("%s", err) == "CAPCHA_NOT_READY" {
			time.Sleep(5 * time.Second)
			return s.GetText(id)
		}
		return "", err
	}

	return resOk, err
}

//checkResponse represents check response for have OK message
func checkResponse(body []byte) (string, error) {
	response := string(body)

	if strings.Contains(response, "OK") {
		list := strings.Split(response, "|")
		for i := range list {
			list[i] = strings.TrimSpace(list[i])
		}

		return list[1], nil
	}

	return "", errors.New(response)
}

//uploadCaptcha represents updaload base64 string to http://anti-captcha.com API
//and get capthca ID
func (s CaptchaService) uploadCaptcha(base64 string) (int, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_ = writer.WriteField("body", base64)
	_ = writer.WriteField("key", s.client.APIKey)
	_ = writer.WriteField("method", "base64")
	err := writer.Close()
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

	res, err := checkResponse(data)
	if err != nil {
		return 0, err
	}
	captchaID, err := strconv.Atoi(res)
	if err != nil {
		return 0, err
	}
	return captchaID, nil
}
