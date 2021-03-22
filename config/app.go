/*
 * @Date: 2021-03-22 09:46:19
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-22 17:14:02
 * @FilePath: /potato/config/app.go
 */
package config

type App struct {
	PageSize  int64  `mapstructure:"page_size" json:"page_size" yaml:"page_size"`
	JwtSecret string `mapstructure:"jwt_secret" json:"jwt_secret" yaml:"jwt_secret"`
	RunMode   string `mapstructure:"run_mode" json:"run_mode" yaml:"run_mode"`
}
