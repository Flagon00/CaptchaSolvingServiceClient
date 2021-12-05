package captcha

import(
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"testing"
	"net/url"
	"strings"
	"bufio"
	"os"
)

const(
	https		= true
	provider	= "api.example.com"
	apiKey		= "api-key"
)

type DataToTest struct{
	pathFile	string
	expected	string
}

func TestRegularCaptcha(t *testing.T){
	var input = []DataToTest{
		{"img/test_captcha_1.jpg", "pvpg78"},
		{"img/test_captcha_2.jpg", "8zjcmw"},
		{"img/test_captcha_3.jpg", "y4vuj"},
	}

	for _, target := range input{
		// Open image file
		imgFile, err := os.Open(target.pathFile)
		if err != nil {
			t.Fatal("Expected no error, but got:", err)
		}

	  	// Read image file
		reader := bufio.NewReader(imgFile)
		content, err := ioutil.ReadAll(reader)
		if err != nil {
			t.Fatal("Expected no error, but got:", err)
		}

		// Encode as base64
		encodedImg := base64.StdEncoding.EncodeToString(content)

		client, err := Client(https, provider, apiKey)
		if err != nil{
			t.Fatal("Expected no error, but got:", err)
		}

		resolve, err := client.RegularCaptcha(encodedImg, 60)
		if err != nil{
			t.Fatal("Expected no error, but got:", err)
		}

		if resolve == ""{
			t.Errorf("Expected %s, but got blank", target.expected)
		} else if resolve != target.expected{
			t.Errorf("Expected output to be %s, but got %s", target.expected, resolve)
		}
	}
}

func TestReCaptchaV2(t *testing.T){
	var(
		inputTarget	= "https://www.google.com/recaptcha/api2/demo"
		targetSitekey	= "6Le-wvkSAAAAAPBMRTvw0Q4Muexq9bi0DJwx_mJ-"
	)

	client, err := Client(https, provider, apiKey)
	if err != nil{
		t.Fatal("Expected no error, but got:", err)
	}

	resolve, err := client.ReCaptchaV2(inputTarget, targetSitekey, "", false, 60)
	if err != nil{
		t.Fatal("Expected no error, but got:", err)
	}

	if resolve == ""{
		t.Fatal("Expected resolve, but got blank")
	}

	resp, err := http.PostForm(inputTarget, url.Values{"g-recaptcha-response": {resolve}})
	if err != nil {
		t.Fatal("Expected no error, but got:", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Expected no error, but got:", err)
	}

	if !strings.Contains(string(body), "Verification Success... Hooray!"){
		t.Error("The captcha answer received is invalid")
	}
}
