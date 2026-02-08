package interfaces

type QuoteRepository interface {
	SaveQuote(data any) error
}
