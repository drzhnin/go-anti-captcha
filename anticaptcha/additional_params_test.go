package anticaptcha

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdditionalParams_toMap(t *testing.T) {
	setup()
	defer teardown()
	params := CaptchaParams{
		Phrase:      1,
		Regsense:    1,
		Calc:        1,
		IsRussian:   1,
		HeaderAcao:  1,
		AllowEmpty:  1,
		Numeric:     1,
		MinLength:   5,
		MaxLength:   10,
		SoftID:      123,
		CaptchaType: "",
		Comment:     "",
	}

	p := params.toMap("t")
	m := map[string]string{"regsense": "1", "is_russian": "1", "header_acao": "1", "min_len": "5", "soft_id": "123", "phrase": "1", "calc": "1", "allow_empty": "1", "numeric": "1", "max_len": "10"}
	assert.Equal(t, p, m)
}

func TestAdditionalParams_LoadParams(t *testing.T) {
	setup()
	defer teardown()
	client.CaptchaAdditionalParams.EnablePhrate()
	client.CaptchaAdditionalParams.EnableCalc()
	client.CaptchaAdditionalParams.EnableAllowEmpty()
	client.CaptchaAdditionalParams.EnableHeaderAcao()
	client.CaptchaAdditionalParams.EnableIsRussian()
	client.CaptchaAdditionalParams.EnableRegsense()
	client.CaptchaAdditionalParams.SetCaptchaType("recaptcha2_44")
	client.CaptchaAdditionalParams.SetMaxCaptchaLength(10)
	client.CaptchaAdditionalParams.SetMinCaptchaLength(2)
	client.CaptchaAdditionalParams.AddComment("Added test comment")
	client.CaptchaAdditionalParams.SetSoftID(123)
	client.CaptchaAdditionalParams.CaptchaContainsNumericOnly()
	params := client.CaptchaAdditionalParams.LoadParams()
	expectedParamsPart1 := map[string]string{"phrase": "1", "regsense": "1", "numeric": "1", "min_len": "2", "soft_id": "123", "type": "recaptcha2_44", "comment": "Added test comment", "calc": "1", "is_russian": "1", "header_acao": "1", "allow_empty": "1", "max_len": "10"}
	assert.Equal(t, params, expectedParamsPart1)
	client.CaptchaAdditionalParams.CaptchaNotContainAnyDigits()
	params = client.CaptchaAdditionalParams.LoadParams()
	expectedParamsPart2 := map[string]string{"phrase": "1", "regsense": "1", "numeric": "2", "min_len": "2", "soft_id": "123", "type": "recaptcha2_44", "comment": "Added test comment", "calc": "1", "is_russian": "1", "header_acao": "1", "allow_empty": "1", "max_len": "10"}
	assert.Equal(t, params, expectedParamsPart2)

	client.CaptchaAdditionalParams.CaptchaContainsAnyCharacters()
	params = client.CaptchaAdditionalParams.LoadParams()
	expectedParamsPart3 := map[string]string{"phrase": "1", "regsense": "1", "min_len": "2", "soft_id": "123", "type": "recaptcha2_44", "comment": "Added test comment", "calc": "1", "is_russian": "1", "header_acao": "1", "allow_empty": "1", "max_len": "10"}
	assert.Equal(t, params, expectedParamsPart3)

	client.CaptchaAdditionalParams.DisablePhrate()
	client.CaptchaAdditionalParams.DisableRegsense()
	client.CaptchaAdditionalParams.DisableCalc()
	client.CaptchaAdditionalParams.DisableIsRussian()
	client.CaptchaAdditionalParams.DisableHeaderAcao()
	client.CaptchaAdditionalParams.DisableAllowEmpty()
	client.CaptchaAdditionalParams.SetMinCaptchaLength(21)
	client.CaptchaAdditionalParams.SetMaxCaptchaLength(21)
	client.CaptchaAdditionalParams.AddComment("qwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiop")
	params = client.CaptchaAdditionalParams.LoadParams()
	expectedParamsPart4 := map[string]string{"soft_id": "123", "type": "recaptcha2_44"}
	assert.Equal(t, params, expectedParamsPart4)

	client.CaptchaAdditionalParams.SetMinCaptchaLength(19)
	client.CaptchaAdditionalParams.SetMaxCaptchaLength(8)
	params = client.CaptchaAdditionalParams.LoadParams()
	expectedParamsPart5 := map[string]string{"min_len": "19", "max_len": "8", "soft_id": "123", "type": "recaptcha2_44"}
	assert.Equal(t, params, expectedParamsPart5)

	client.CaptchaAdditionalParams.ResetMaxCaptchaLengthToDefault()
	client.CaptchaAdditionalParams.ResetMinCaptchaLengthToDefault()
	params = client.CaptchaAdditionalParams.LoadParams()
	expectedParamsPart6 := map[string]string{"soft_id": "123", "type": "recaptcha2_44"}
	assert.Equal(t, params, expectedParamsPart6)

}

func TestAdditionalParams_DebugParams(t *testing.T) {
	setup()
	defer teardown()
	client.CaptchaAdditionalParams.EnablePhrate()
	client.CaptchaAdditionalParams.EnableCalc()
	client.CaptchaAdditionalParams.AddComment("Added test comment")
	client.CaptchaAdditionalParams.DebugParams()
}
