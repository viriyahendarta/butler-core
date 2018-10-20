package service

type apiMeta struct {
	ProcessTime string `json:"process_time"`
}

type apiError struct {
	Code     int      `json:"code"`
	Type     string   `json:"type"`
	Messages []string `json:"messages"`
	Reason   string   `json:"reason"`
}

type apiErrorResponse struct {
	Meta  apiMeta  `json:"meta"`
	Error apiError `json:"error"`
}

type apiSuccessResponse struct {
	Meta apiMeta     `json:"meta"`
	Data interface{} `json:"data"`
}
