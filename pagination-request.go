package eliotlibs

import (
	"strings"
)

type CrudPaginationRequest struct {
	PaginationRequest
	ParamName string
}

type PaginationRequest struct {
	Limit    int
	Sort     string
	Page     int
	PropSort string
}

func (p *PaginationRequest) GetSort() string {
	if strings.ToLower(p.Sort) != "asc" && strings.ToLower(p.Sort) != "desc" {
		p.Sort = "desc"
	}

	if p.PropSort == "" {
		p.PropSort = "id"
	}

	return p.PropSort + " " + p.Sort
}

func GetPaginationRequestByOtherParam(getQueryParam func(string) string, propSotr, paramName string) CrudPaginationRequest {
	return CrudPaginationRequest{
		PaginationRequest: GetPaginationRequest(getQueryParam, propSotr),
		ParamName:         getQueryParam(paramName),
	}
}

func GetPaginationRequest(get func(string) string, propSotr string) PaginationRequest {
	limit := GetNumberForString(get("limit"))
	if limit <= 0 {
		limit = 5
	}

	page := GetNumberForString(get("page"))
	if page <= 0 {
		page = 1
	}

	sort := get("sort")
	if sort == "" {
		sort = "created_at DESC" // default opcional
	}

	return PaginationRequest{
		Limit:    limit,
		Page:     page,
		Sort:     sort,
		PropSort: propSotr,
	}
}
