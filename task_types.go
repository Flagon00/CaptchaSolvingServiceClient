package captcha

import (
	"errors"
	"time"
)

// A method that creates job with regular image captcha in base64 and gives you the answer
func (c *CaptchaServiceClient) RegularCaptcha(base64Image string, timeout time.Duration) (string, error){
	// Preparation form data
	data := map[string]interface{}{
		"clientKey":	c.ApiKey,
		"task":			map[string]interface{}{
			"type":		"ImageToTextTask",
			"body":		base64Image,
		},
	}

	// Create task to solve
	jobID, err := c.CreatTask(data)
	if err != nil{
		return "", err
	}

	// Check the status of job every 5 seconds
	ping := time.NewTicker(5 * time.Second)
	timeoutBreak := time.NewTimer(timeout * time.Second)

	for {
		select{
			case <-ping.C:
				answer, ready, err := c.CheckResult(jobID)
				if err != nil{
					return "", err
				}

				if ready{
					return answer.TextSolution, nil
				}
			case <-timeoutBreak.C:
				return "", errors.New("Waiting for captcha result timeout")
		}
	}

	// If the job takes too long
	return "", errors.New("Waiting for captcha result timeout")
}

// A method that creates job with Google reCaptcha and gives you the answer
func (c *CaptchaServiceClient) ReCaptchaV2(website, siteKey, sKey string, isInvisible bool, timeout time.Duration) (string, error){
	// Preparation form data
	data := map[string]interface{}{
		"clientKey":	c.ApiKey,
		"task":			map[string]interface{}{
			"type":		"NoCaptchaTaskProxyless",
			"websiteURL":	website,
			"websiteKey":	siteKey,
		},
	}

	// sKey is not required
	if sKey != ""{ data["task"].(map[string]interface{})["recaptchaDataSValue"] = sKey }

	// isInvisible is not required
	if isInvisible{ data["task"].(map[string]interface{})["isInvisible"] = true }

	// Create task to resolve
	jobID, err := c.CreatTask(data)
	if err != nil{
		return "", err
	}

	// Check the status of job every 5 seconds
	ping := time.NewTicker(5 * time.Second)
	timeoutBreak := time.NewTimer(timeout * time.Second)

	for {
		select {
		case <-ping.C:
			answer, ready, err := c.CheckResult(jobID)
			if err != nil{
				return "", err
			}

			if ready{
				return answer.RecaptchaV2Solution, nil
			}
		case <-timeoutBreak.C:
			return "", errors.New("Waiting for captcha result timeout")
		}
	}

	// If the job takes too long
	return "", errors.New("Waiting for captcha result timeout")
}

// A method that creates job with Google reCaptcha and gives you the answer
func (c *CaptchaServiceClient) ReCaptchaV2Enterprise(website, siteKey, sKey string, timeout time.Duration) (string, error){
	// Preparation form data
	data := map[string]interface{}{
		"clientKey":	c.ApiKey,
		"task":			map[string]interface{}{
			"type":		"RecaptchaV2EnterpriseTaskProxyless",
			"websiteURL":	website,
			"websiteKey":	siteKey,
			"enterprisePayload":	map[string]interface{}{
				"s":	sKey,
			},
		},
	}

	// sKey is not required
	if sKey != ""{ data["task"].(map[string]interface{})["enterprisePayload"].(map[string]interface{})["s"] = sKey }

	// Create task to resolve
	jobID, err := c.CreatTask(data)
	if err != nil{
		return "", err
	}

	// Check the status of job every 5 seconds
	ping := time.NewTicker(5 * time.Second)
	timeoutBreak := time.NewTimer(timeout * time.Second)

	for {
		select {
		case <-ping.C:
			answer, ready, err := c.CheckResult(jobID)
			if err != nil{
				return "", err
			}

			if ready{
				return answer.RecaptchaV2Solution, nil
			}
		case <-timeoutBreak.C:
			return "", errors.New("Waiting for captcha result timeout")
		}
	}

	// If the job takes too long
	return "", errors.New("Waiting for captcha result timeout")
}
