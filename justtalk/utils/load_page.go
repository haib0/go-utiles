package utils

import (
	"github.com/haib0/go-utiles/justtalk/model"
)

func LoadPage(title string, no int, size int) (*model.Page, error) {
	page := model.Page{Title: title, No: no, Size: size}

	num, err := GetMessageNum(page.Title)
	if err != nil {
		return nil, err
	}

	page.LastNo = num / page.Size

	page.PrevNo = no - 1
	page.NextNo = no + 1
	if no > page.LastNo {
		page.PrevNo = page.LastNo
		page.NextNo = -1
	}
	if no == page.LastNo {
		page.NextNo = -1
	}
	if no <= 0 {
		page.No = 0
		page.PrevNo = -1
		page.NextNo = 1
	}

	err = page.GetMessages()
	if err != nil {
		return nil, err
	}
	return &page, nil
}
