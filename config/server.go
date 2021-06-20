/*
 * @Date: 2021-03-22 09:54:07
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-20 19:45:31
 * @FilePath: /potato/config/server.go
 */
package config

type Server struct {
	Port           int64  `mapstructure:"port" json:"port" yaml:"port"`
	ReadTimeout    int64  `mapstructure:"read_timeout" json:"read_timeout" yaml:"read_timeout"`
	WriteTimeout   int64  `mapstructure:"write_timeout" json:"write_timeout" yaml:"write_timeout"`
	TracerHostPort string `mapstructure:"tracer_host_port" json:"tracer_host_port" yaml:"tracer_host_port"`
}
