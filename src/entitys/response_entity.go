package entitys

type Response struct {
	StatusCode int                    `json:"statusCode"`
	StatusDesc string                 `json:"statusDesc"`
	Body       map[string]interface{} `json:"body"`
}
