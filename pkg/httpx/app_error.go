package httpx

type AppError struct {
	StatusCode  int
	Message     string
	Description string
}

func (e *AppError) Error() string {
	return e.Message
}
