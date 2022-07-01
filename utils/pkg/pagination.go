package pkg

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Param struct {
	Total int64
	Page  int
	Limit int
	Ctx   *fiber.Ctx
}

// Paginator 分页返回
type Paginator struct {
	TotalRecord int64       `json:"total_record"`
	TotalPage   int         `json:"total_page"`
	Records     interface{} `json:"records"`
	Offset      int         `json:"offset"`
	Limit       int         `json:"limit"`
	Page        int         `json:"page"`
	PrevPage    int         `json:"prev_page"`
	NextPage    int         `json:"next_page"`
}

// Paging 分页
func Paging(p *Param) *Paginator {
	page, _ := strconv.Atoi(p.Ctx.Query("page"))
	pageSize, _ := strconv.Atoi(p.Ctx.Query("page_size"))
	p.Page = page
	p.Limit = pageSize
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit == 0 {
		p.Limit = 10
	}
	var paginator Paginator
	count := p.Total
	var offset int

	if p.Page == 1 {
		offset = 0
	} else {
		offset = (p.Page - 1) * p.Limit
	}

	paginator.TotalRecord = int64(count)
	paginator.Page = p.Page

	paginator.Offset = offset
	paginator.Limit = p.Limit
	paginator.TotalPage = int(math.Ceil(float64(count) / float64(p.Limit)))

	if p.Page > 1 {
		paginator.PrevPage = p.Page - 1
	} else {
		paginator.PrevPage = p.Page
	}

	if p.Page == paginator.TotalPage {
		paginator.NextPage = p.Page
	} else {
		paginator.NextPage = p.Page + 1
	}

	p.Ctx.Append("TotalRecord", fmt.Sprintf("%d", paginator.TotalRecord))
	p.Ctx.Append("Page", fmt.Sprintf("%d", paginator.Page))
	p.Ctx.Append("Offset", fmt.Sprintf("%d", paginator.Offset))
	p.Ctx.Append("Limit", fmt.Sprintf("%d", paginator.Limit))
	p.Ctx.Append("TotalPage", fmt.Sprintf("%d", paginator.TotalPage))
	p.Ctx.Append("PrevPage", fmt.Sprintf("%d", paginator.PrevPage))
	p.Ctx.Append("NextPage", fmt.Sprintf("%d", paginator.NextPage))

	log.Println(paginator.TotalRecord)
	return &paginator
}
