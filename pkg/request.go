/*
 * @Date: 2021-03-22 23:46:49
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-10 15:18:09
 * @FilePath: /potato/pkg/request.go
 */
package pkg

type PageInfo struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type SearchResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}
