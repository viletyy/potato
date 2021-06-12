/*
 * @Date: 2021-06-12 23:01:51
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-12 23:04:48
 * @FilePath: /potato/config/email.go
 */
package config

type Email struct {
	Host     string   `mapstructure:"host" json:"host" yaml:"host"`
	Port     int      `mapstructure:"port" json:"port" yaml:"port"`
	UserName string   `mapstructure:"username" json:"username" yaml:"username"`
	Password string   `mapstructure:"password" json:"password" yaml:"password"`
	IsSSL    bool     `mapstructure:"is_ssl" json:"is_ssl" yaml:"is_ssl"`
	From     string   `mapstucture:"from" json:"from" yaml:"from"`
	To       []string `mapstucture:"to" json:"to" yaml:"to"`
}
