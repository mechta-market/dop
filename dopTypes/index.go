package dopTypes

import (
	"time"
)

type ErrRep struct {
	ErrorCode string            `json:"error_code"`
	Desc      string            `json:"desc,omitempty"`
	Fields    map[string]string `json:"fields,omitempty"`
}

type ListParams struct {
	Cols           []string `json:"cols" form:"cols"`
	Page           int64    `json:"page" form:"page"`
	PageSize       int64    `json:"page_size" form:"page_size"`
	WithTotalCount bool     `json:"with_total_count" form:"with_total_count"`
	OnlyCount      bool     `json:"only_count" form:"only_count"`
	SortName       string   `json:"sort_name" form:"sort_name"`
	Sort           []string `json:"sort" form:"sort"`
}

type ListRep struct {
	Results any `json:"results"`
}

type PaginatedListRep struct {
	Page       int64 `json:"page"`
	PageSize   int64 `json:"page_size"`
	TotalCount int64 `json:"total_count"`

	Results any `json:"results"`
}

type CreateRep struct {
	Id any `json:"id"`
}

type PeriodPars struct {
	TsGTE *time.Time `json:"ts_gte" form:"ts_gte"`
	TsLTE *time.Time `json:"ts_lte" form:"ts_lte"`
}

type ChartVByTimeSt struct {
	Ts time.Time `json:"ts"`
	V  int64     `json:"v"`
}
