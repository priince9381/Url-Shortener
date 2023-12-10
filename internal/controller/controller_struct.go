package controller

type ErrorResponse struct {
	Message string `json:"message"`
}

type CreateShortURLRequest struct {
	LongURL string `json:"long_url"`
}
