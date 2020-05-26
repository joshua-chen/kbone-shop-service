/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-17 00:04:32
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-26 18:32:06
 */
package models

type ResponseResult struct {
	Code    string      `json:"code"`
	Msg     string      `json:"msg"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func NewResponseResult(data interface{}, c string, m ...string) *ResponseResult {
	r := ResponseResult{Data: data, Code: c, Success: false}

	if e, ok := data.(error); ok {
		if m == nil {
			r.Msg = e.Error()
		}
	} else {
		r.Success = true
		r.Msg = "SUCCESS"
	}
	if len(m) > 0 {
		r.Msg = m[0]
	}

	return &r
}

func NewErrorResponseResult(data interface{}, c string, m ...string) *ResponseResult {
	r := &ResponseResult{Data: data, Code: c, Success: false}

	if e, ok := data.(error); ok {
		if m == nil {
			r.Msg = e.Error()
		}
	} else {
		r.Success = true
		r.Msg = "SUCCESS"
	}
	if len(m) > 0 {
		r.Msg = m[0]
	}

	return r
}

func NewSuccessResponseResult(data interface{}, c string, m ...string) *ResponseResult {
	r := &ResponseResult{Data: data, Code: c, Success: true}
	return r
}
