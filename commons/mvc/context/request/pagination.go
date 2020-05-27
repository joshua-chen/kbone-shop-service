/*
 * @Descripttion: 
 * @version: 
 * @Author: joshua
 * @Date: 2020-05-25 16:38:18
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-27 12:12:09
 */ 
package request

import (
	"errors"

	"github.com/kataras/iris/v12"

)

// bootstraptable 分页参数
type Pagination struct {
	PageNum int //当前看的是第几页
	PageSize   int //每页显示多少条数据

	// 用于分页设置的参数
	Offset int
	Limit int

	SortName  string //用于指定的排序
	SortOrder string // desc或asc

	// 时间范围
	StartDate string
	EndDate   string

	ID int64 // 公用的特殊参数
}

func NewPagination(ctx iris.Context) (*Pagination, error) {
	pageNum, err1 := ctx.URLParamInt("pageNum")
	pageSize, err2 := ctx.URLParamInt("pageSize")
	offset := ctx.URLParamIntDefault("offset",-1)
	limit := ctx.URLParamIntDefault("limit",-1)
	sortName := ctx.URLParam("sortName")
	sortOrder := ctx.URLParam("sortOrder")
	
	if err1 != nil || err2 != nil  {
		return nil, errors.New("请求的分页参数解析错误.")
	}
	var page Pagination
	if(offset != -1 && limit!= -1){
		page = Pagination{
			SortName:   sortName,
			SortOrder:  sortOrder,
			Limit: limit,
			Offset: offset,
		}
	}else{
		page = Pagination{
			PageNum: pageNum,
			PageSize: pageSize,
			SortName:   sortName,
			SortOrder:  sortOrder,
		}
		page.set()
	}
	
	
	return &page, nil
}

// 设置分页参数
func (p *Pagination) set() {
	if p.PageNum < 1 {
		p.PageNum = 1
	}
	if p.PageSize < 1 {
		p.PageSize = 1
	}

	p.Offset = (p.PageNum - 1) * p.PageSize
	p.Limit = p.PageSize
}

 

