package response

type Page struct {
	Limit int `json:"limit,omitempty"`
	Page int `json:"page,omitempty"`
	TotalElements int64 `json:"total_elements"`
	TotalPages int `json:"total_pages"`
	Content any `json:"content"`
}

func (p *Page) GetPage() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Page) GetLimit() int {
	if p.Limit <= 0 {
		p.Limit = 50
	}
	return p.Limit
}

func (p *Page) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}
