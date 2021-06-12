/*
 * @Date: 2021-03-22 09:45:22
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-12 23:07:27
 * @FilePath: /potato/config/config.go
 */
package config

type Config struct {
	App      App      `mapstructure:"app" json:"app" yaml:"app"`
	Server   Server   `mapstructure:"server" json:"server" yaml:"server"`
	Email    Email    `mapstructure:"email" json:"email" yaml:"server"`
	Database Database `mapstructure:"database" json:"database" yaml:"database"`
	Redis    Redis    `mapstructure:"redis" json:"redis" yaml:"redis"`
	Zap      Zap      `mapstructure:"zap" json:"zap" yaml:"zap"`
}
