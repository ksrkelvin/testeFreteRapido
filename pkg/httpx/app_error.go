package httpx

type AppError struct {
	Status      int
	Code        string
	Message     string
	Description string
	Err         error
}

func (e *AppError) Error() string {
	return e.Message
}
