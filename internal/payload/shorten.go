package payload

type (
	CreateShortenRequest struct {
		OriginalURL string `json:"original_url"`
	}

	CreateShortenResponse struct {
	}
)

type (
	GetShortenRequest  struct{}
	GetShortenResponse struct{}
)

type (
	GetStatsRequest  struct{}
	GetStatsResponse struct{}
)
