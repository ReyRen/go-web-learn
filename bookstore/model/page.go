package model

type Page struct {
	Books       []*Book
	PageNo      int64 // current page no
	PageSize    int64 // count displayed in each page
	TotalPageNo int64 // total pages
	TotalRecord int64 // total records
}

// isHasPrev
func (page *Page) IsHasPrev() bool {
	return page.PageNo > 1
}

// isHasNext
func (page *Page) IsHasNext() bool {
	return page.PageNo < page.TotalPageNo
}

// getPrevPageNo
func (page *Page) GetPrevPageNo() int64 {
	if page.IsHasPrev() {
		return page.PageNo - 1
	} else {
		return 1
	}
}

// getNextPageNo
func (page *Page) GetNextPageNo() int64 {
	if page.IsHasNext() {
		return page.PageNo + 1
	} else {
		return page.TotalPageNo
	}
}
