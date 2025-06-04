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

func GetPagginationRequestByOtherParam(getQueryParam func(string) string, propSotr, paramName string) CrudPaginationRequest {
	return CrudPaginationRequest{
		PaginationRequest: GetPagginationRequest(getQueryParam, propSotr),
		ParamName:         getQueryParam(paramName),
	}
}

func GetPagginationRequest(getQueryParam func(string) string, propSotr string) PaginationRequest {
	return PaginationRequest{
		Limit: func() int {
			limit := getQueryParam("limit")
			if limit == "" {
				return 5
			}
			return GetNumberForString(getQueryParam("limit"))
		}(),
		Sort: getQueryParam("sort"),
		Page: func() int {
			limit := getQueryParam("page")
			if limit == "" {
				return 1
			}
			return GetNumberForString(getQueryParam("page"))
		}(),
		PropSort: propSotr,
	}
}
