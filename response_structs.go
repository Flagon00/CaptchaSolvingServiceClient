package captcha

type CreateTaskResponse struct{
	ErrorID				int	`json:"errorId"`
	ErrorCode			string	`json:"errorCode"`
	ErrorDescription		string	`json:"errorDescription"`
	TaskID				int	`json:"taskId"`
}

type GetTaskResponse struct{
	ErrorID				int		`json:"errorId"`
	ErrorCode			string		`json:"errorCode"`
	ErrorDescription		string		`json:"errorDescription"`
	Status				string		`json:"status"`
	Solution			Solution	`json:"solution"`
}

type Solution struct{
	TextSolution			string	`json:"text"`
	RecaptchaV2Solution		string	`json:"gRecaptchaResponse"`
}
