/*
 * @Date: 2021-03-22 09:46:19
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-13 22:00:01
 * @FilePath: /potato/config/app.go
 */
package config

type App struct {
	Name                 string   `mapstructure:"name" json:"name" yaml:"name"`
	Version              string   `mapstructure:"version" json:"version" yaml:"version"`
	PageSize             int64    `mapstructure:"page_size" json:"page_size" yaml:"page_size"`
	JwtSecret            string   `mapstructure:"jwt_secret" json:"jwt_secret" yaml:"jwt_secret"`
	JwtExpire            int64    `mapstructure:"jwt_expire" json:"jwt_expire" yaml:"jwt_expire"`
	RunMode              string   `mapstructure:"run_mode" json:"run_mode" yaml:"run_mode"`
	UploadSavePath       string   `mapstructure:"upload_save_path" json:"upload_save_path" yaml:"upload_save_path"`
	UploadServerPath     string   `mapstructure:"upload_server_path" json:"upload_server_path" yaml:"upload_server_path"`
	UploadImageMaxSize   int64    `mapstructure:"upload_image_max_size" json:"upload_image_max_size" yaml:"upload_image_max_size"`
	UploadImageAllowExts []string `mapstructure:"upload_image_allow_exts" json:"upload_image_allow_exts" yaml:"upload_image_allow_exts"`
}
