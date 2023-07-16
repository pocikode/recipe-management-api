package pkg

type Pagination struct {
	Page       int               `json:"page,omitempty;query:page"`
	Limit      int               `json:"limit,omitempty;query:limit"`
	TotalRows  int64             `json:"total_rows"`
	TotalPages int               `json:"total_pages"`
	Rows       interface{}       `json:"rows"`
	Sort       string            `json:"sort,omitempty;query:sort"`
	Keyword    string            `json:"-"`
	Filters    map[string]string `json:"-"`
}

func (p *Pagination) GetPage() int {
	if p.Page <= 0 {
		p.Page = 1
	}

	return p.Page
}

func (p *Pagination) GetLimit() int {
	if p.Limit <= 0 {
		p.Limit = 10
	}

	return p.Limit
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Id desc"
	}
	return p.Sort
}

func (p *Pagination) GetKeyword() string {
	return p.Keyword
}

func (p *Pagination) ToResponse() *Pagination {
	p.Page = p.GetPage()
	p.Limit = p.GetLimit()
	p.Sort = p.GetSort()

	return p
}
