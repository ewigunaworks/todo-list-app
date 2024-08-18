package helpers

import "math"

type Paginate struct {
	CurrentPage int64 `json:"CurrentPage"`
	NextPage    int64 `json:"NextPage"`
	TotalPage   int64 `json:"TotalPage"`
	PerPage     int32 `json:"PerPage"`
	TotalRow    int64 `json:"TotalRow"`
}

func Pagination(page, limit, totalRecords int) (*Paginate, int) {
	var offset int

	if page == 0 {
		page = 1
	}

	if limit == 0 || limit > 100 {
		limit = 10
	}

	totalPage := int(math.Ceil(float64(totalRecords) / float64(limit)))

	if page != 0 {
		offset = (page * limit) - limit

	} else {
		offset = 0
	}

	return &Paginate{
		CurrentPage: int64(page),
		NextPage:    int64(page + 1),
		TotalPage:   int64(totalPage),
		PerPage:     int32(limit),
		TotalRow:    int64(totalRecords),
	}, offset

}
