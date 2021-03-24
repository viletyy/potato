/*
 * @Date: 2021-03-24 10:18:40
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-24 10:30:27
 * @FilePath: /potato/utils/helper.go
 */
package utils

import "strconv"

func ToString(i interface{}) string {
	switch i.(type) {
	case string:
		return i.(string)
	case int:
		return strconv.Itoa(i.(int))
	case int64:
		return strconv.FormatInt(i.(int64), 10)
	}
	return ""
}

func ToInt(i interface{}) int {
	switch i.(type) {
	case string:
		result, err := strconv.Atoi(i.(string))
		if err != nil {
			return 0
		}
		return result
	case int64:
		return int(i.(int64))
	}
	return 0
}
