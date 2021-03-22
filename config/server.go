/*
 * @Date: 2021-03-22 09:54:07
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-22 09:56:38
 * @FilePath: /potato/config/server.go
 */
package config

type Server struct {
	HttpPort     int64 `mapstructure:"http_port" json:"http_port" yaml:"http_port"`
	ReadTimeout  int64 `mapstructure:"read_timeout" json:"read_timeout" yaml:"read_timeout"`
	WriteTimeout int64 `mapstructure:"write_timeout" json:"write_timeout" yaml:"write_timeout"`
}
