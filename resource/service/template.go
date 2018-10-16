package service

type APIMeta struct {
	ProcessTime string `json:"process_time"`
}

type APIError struct {
	Code     int      `json:"code"`
	Type     string   `json:"type"`
	Messages []string `json:"messages"`
	Reason   string   `json:"reason"`
}

type APIErrorResponse struct {
	Meta  APIMeta  `json:"meta"`
	Error APIError `json:"error"`
}

type APISuccessResponse struct {
	Meta APIMeta     `json:"meta"`
	Data interface{} `json:"data"`
}
