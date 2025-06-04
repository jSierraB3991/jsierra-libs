package eliotlibs

import "strings"

type Paggination struct {
	Limit      int         `json:"limit"`
	Page       int         `json:"page"`
	TotalRows  int64       `json:"rows"`
	TotalPages int         `json:"pages"`
	Data       interface{} `json:"data"`
	Sort       string      `json:"-"`
}

func (p *Paggination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Paggination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Paggination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Paggination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Id desc"
	}
	return p.Sort
}

type PagginationParam struct {
	Where string
	Data  interface{}
}

type PreloadParams struct {
	Preload          string
	PagginationParam PagginationParam
}

func PaginationReqToModel(page, limit int, datForWhere interface{}, sort, propSort string) Paggination {
	return Paggination{
		Limit: limit,
		Page:  page,
		Data:  datForWhere,
		Sort: func(sort, propSort string) string {
			if strings.ToLower(sort) != "asc" && strings.ToLower(sort) != "desc" {
				sort = "desc"
			}

			if propSort == "" {
				propSort = "id"
			}

			return propSort + " " + sort
		}(sort, propSort),
	}
}
