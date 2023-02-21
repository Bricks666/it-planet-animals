package shared

import "time"

type PaginationDto struct {
	From uint32 `query:"from" validate:"number,gte=0"`
	Size uint32 `query:"size" validate:"number,gt=0"`
}

type TimeSearchQueryDto struct {
	AfterDateTime  time.Time `query:"startdatetime" validate:"omitempty,iso-8601"`
	BeforeDateTime time.Time `query:"enddatetime" validate:"omitempty,iso-8601"`
}
