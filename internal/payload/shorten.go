package payload

type (
	CreateShortenRequest struct {
		OriginalURL string `json:"original_url" binding:"required,url"`
	}

	CreateShortenResponse struct {
		OriginalURL string `json:"original_url"`
		ShortURL    string `json:"short_url"`
	}
)

type (
	GetShortenRequest  struct{}
	GetShortenResponse struct {
		OriginalURL string `json:"original_url"`
	}
)

type (
	GetStatsRequest  struct{}
	GetStatsResponse struct {
		OriginalURL string `json:"original_url"`
		ClickCount  uint64 `json:"click_count"`
	}
)
