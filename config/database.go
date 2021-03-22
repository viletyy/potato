/*
 * @Date: 2021-03-22 09:56:48
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-22 10:00:26
 * @FilePath: /potato/config/database.go
 */
package config

type Database struct {
	Type        string `mapstructure:"type" json:"type" yaml:"app"`
	User        string `mapstructure:"user" json:"user" yaml:"user"`
	Password    string `mapstructure:"password" json:"password" yaml:"password"`
	Host        string `mapstructure:"host" json:"host" yaml:"host"`
	Port        int64  `mapstructure:"port" json:"port" yaml:"port"`
	Name        string `mapstructure:"name" json:"name" yaml:"name"`
	TablePrefix string `mapstructure:"table_prefix" json:"table_prefix" yaml:"table_prefix"`
}
