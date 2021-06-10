/*
 * @Date: 2021-03-22 09:46:19
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-10 18:24:42
 * @FilePath: /potato/config/app.go
 */
package config

type App struct {
	PageSize   int64  `mapstructure:"page_size" json:"page_size" yaml:"page_size"`
	JwtSecret  string `mapstructure:"jwt_secret" json:"jwt_secret" yaml:"jwt_secret"`
	JwtIssuser string `mapstructure:"jwt_issuser" json:"jwt_issuser" yaml:"jwt_issuser"`
	JwtExpire  int64  `mapstructure:"jwt_expire" json:"jwt_expire" yaml:"jwt_expire"`
	RunMode    string `mapstructure:"run_mode" json:"run_mode" yaml:"run_mode"`
}
