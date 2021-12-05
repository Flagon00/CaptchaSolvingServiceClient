package captcha

import (
	"encoding/json"
	"net/http"
	"net/url"
	"errors"
	"bytes"
	"fmt"
)

type CaptchaServiceClient struct {
	ApiKey		string
	ApiAdress	*url.URL
	HttpClient	*http.Client
}

// Preparation client to use package
func Client(secure bool, provider string, apikey string) (*CaptchaServiceClient, error){
	var(
		parseURL 	*url.URL
		err		error
	)

	switch secure{
		case false:
			parseURL, err = url.Parse(fmt.Sprintf("http://%s/", provider))
		default:
			parseURL, err = url.Parse(fmt.Sprintf("https://%s/", provider))
	}
	if err != nil{
		return nil, err
	}

	return &CaptchaServiceClient{
		ApiKey: 	apikey,
		ApiAdress:	parseURL,
		HttpClient: http.DefaultClient,
	}, nil
}

// A method create job and return the ID
func (c *CaptchaServiceClient) CreatTask(data map[string]interface{}) (int, error){
	payload, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}

	// Formalize the POST request
	req, err := http.NewRequest("POST", c.ApiAdress.ResolveReference(&url.URL{Path: "/createTask"}).String(), bytes.NewBuffer(payload))
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	// Decode the response
	var responseBody CreateTaskResponse
	json.NewDecoder(resp.Body).Decode(&responseBody)

	if responseBody.ErrorID != 0{
		return 0, errors.New(fmt.Sprint("Error solving captcha: ", responseBody.ErrorCode))
	}

	return responseBody.TaskID, nil
}

// A method to check answer for job
func (c *CaptchaServiceClient) CheckResult(taskID int) (*Solution, bool, error){
	// Preparation form data
	data := map[string]interface{}{
		"clientKey":	c.ApiKey,
		"taskId":	taskID,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, true, errors.New(fmt.Sprint("Error solving captcha: ", err))
	}

	req, err := http.NewRequest("POST", c.ApiAdress.ResolveReference(&url.URL{Path: "/getTaskResult"}).String(), bytes.NewBuffer(payload))
	if err != nil {
		return nil, true, errors.New(fmt.Sprint("Error solving captcha: ", err))
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, true, errors.New(fmt.Sprint("Error solving captcha: ", err))
	}
	defer resp.Body.Close()

	// Decode the response
	var responseBody GetTaskResponse
	json.NewDecoder(resp.Body).Decode(&responseBody)

	// Interpret the answer
	switch responseBody.Status{
		case "processing":
			return nil, false, nil
		default:
			if responseBody.ErrorID != 0{
				return nil, true, errors.New(fmt.Sprint("Error solving captcha: ", responseBody.ErrorCode))
			}
	}

	return &responseBody.Solution, true, nil
}
