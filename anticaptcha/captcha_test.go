package anticaptcha

import (
	"errors"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestCaptchaService_UploadCaptchaFromURL(t *testing.T) {
	httpmock.Activate()
	setup()

	defer httpmock.DeactivateAndReset()
	defer teardown()

	httpmock.RegisterResponder("POST", "http://anti-captcha.com/in.php",
		httpmock.NewStringResponder(200, "OK|123"))
	httpmock.RegisterResponder("GET", "https://s3-us-west-2.amazonaws.com/captcha-test/1045.png",
		httpmock.NewStringResponder(200, "OK"))
	id, err := client.Captcha.UploadCaptchaFromURL("https://s3-us-west-2.amazonaws.com/captcha-test/1045.png")
	assert.Equal(t, id, 123)
	assert.Equal(t, err, nil)

	httpmock.RegisterResponder("GET", "https://s3-us-west-2.amazonaws.com/captcha-test/error.png",
		httpmock.NewStringResponder(500, ""))
	id, _ = client.Captcha.UploadCaptchaFromURL("https://s3-us-west-2.amazonaws.com/captcha-test/error.png")
	assert.Equal(t, id, 0)
}

func TestCaptchaService_UploadCaptchaFromFile(t *testing.T) {
	httpmock.Activate()
	setup()

	defer httpmock.DeactivateAndReset()
	defer teardown()

	httpmock.RegisterResponder("POST", "http://anti-captcha.com/in.php",
		httpmock.NewStringResponder(200, "OK|123"))
	id, err := client.Captcha.UploadCaptchaFromFile("captcha.png")
	assert.Equal(t, id, 123)
	assert.Equal(t, err, nil)

	httpmock.RegisterResponder("POST", "http://anti-captcha.com/in.php",
		httpmock.NewStringResponder(200, "ERROR_ZERO_CAPTCHA_FILESIZE"))
	id, err = client.Captcha.UploadCaptchaFromFile("captcha.png")
	assert.Equal(t, id, 0)
	assert.Equal(t, err, errors.New("ERROR_ZERO_CAPTCHA_FILESIZE"))

	id, err = client.Captcha.UploadCaptchaFromFile("")
	assert.Equal(t, id, 0)
	assert.Equal(t, err, errors.New("File does not exist"))
}

func TestCaptchaService_UploadCaptchaFromBase64(t *testing.T) {
	httpmock.Activate()
	setup()

	defer httpmock.DeactivateAndReset()
	defer teardown()

	httpmock.RegisterResponder("POST", "http://anti-captcha.com/in.php",
		httpmock.NewStringResponder(200, "OK|123"))
	id, err := client.Captcha.UploadCaptchaFromBase64("x8akfmUYRVW9vZ/GCpG3Q/gaDUAAAAASUVORK5CYII=")
	assert.Equal(t, id, 123)
	assert.Equal(t, err, nil)

	httpmock.RegisterResponder("POST", "http://anti-captcha.com/in.php",
		httpmock.NewStringResponder(200, "ERROR_ZERO_CAPTCHA_FILESIZE"))
	id, err = client.Captcha.UploadCaptchaFromBase64("")
	assert.Equal(t, id, 0)
	assert.Equal(t, err, errors.New("ERROR_ZERO_CAPTCHA_FILESIZE"))
}

func TestCaptchaService_GetText(t *testing.T) {
	httpmock.Activate()
	setup()

	defer httpmock.DeactivateAndReset()
	defer teardown()

	httpmock.RegisterResponder("GET", "http://anti-captcha.com/res.php?key=F629EBDA-D89A-4A0E-AAA5-069761578237&action=get&id=123", httpmock.NewStringResponder(200, "OK|text"))
	text, err := client.Captcha.GetText(123)
	assert.Equal(t, text, "text")
	assert.Equal(t, err, nil)

	httpmock.RegisterResponder("GET", "http://anti-captcha.com/res.php?key=F629EBDA-D89A-4A0E-AAA5-069761578237&action=get&id=1213",
		httpmock.NewStringResponder(200, "ERROR_NO_SUCH_CAPCHA_ID"))
	text, err = client.Captcha.GetText(1213)
	assert.Equal(t, text, "")
	assert.Equal(t, err, errors.New("ERROR_NO_SUCH_CAPCHA_ID"))

}
