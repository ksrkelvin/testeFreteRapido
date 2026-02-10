package httpx

type ErrorResponse struct {
	Message     string `json:"message"`
	Description string `json:"description,omitempty"`
}
