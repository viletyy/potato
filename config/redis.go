/*
 * @Date: 2021-03-22 10:00:42
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-22 10:02:39
 * @FilePath: /potato/config/redis.go
 */
package config

type Redis struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     int64  `mapstructure:"port" json:"port" yaml:"port"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Db       int64  `mapstructure:"db" json:"db" yaml:"db"`
}
