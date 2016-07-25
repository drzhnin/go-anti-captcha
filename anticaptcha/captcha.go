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
	"time"
)

//CaptchaService handles communication with the captcha actions of the
//anti-captcha API
type CaptchaService service

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

//UploadCaptchaFromBase64 represents updaload base64 string to http://anti-captcha.com API
//and get capthca ID
func (s *CaptchaService) UploadCaptchaFromBase64(base64 string) (int, error) {
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
