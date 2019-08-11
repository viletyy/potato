// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-08-12 01:35:45.643682 +0800 CST m=+0.068783392

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/businesses": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "businesses"
                ],
                "summary": "业务系统列表",
                "responses": {
                    "200": {
                        "description": "{\"code\" : 200, \"data\" : {}, \"msg\" : \"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "businesses"
                ],
                "summary": "新增业务系统",
                "parameters": [
                    {
                        "type": "string",
                        "description": "业务系统 名称",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "业务系统 描述",
                        "name": "desc",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "业务系统 云端id",
                        "name": "c_id",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\" : 200, \"msg\" : \"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/businesses/{id}": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "businesses"
                ],
                "summary": "删除业务系统",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "业务系统 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\" : 200, \"msg\" : \"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "businesses"
                ],
                "summary": "修改业务系统",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "业务系统 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "业务系统 名称",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "业务系统 描述",
                        "name": "desc",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "业务系统 云端id",
                        "name": "c_id",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\" : 200, \"msg\" : \"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/meta_databases": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "meta_databases"
                ],
                "summary": "数据源列表",
                "responses": {
                    "200": {
                        "description": "{\"code\" : 200, \"basic\" : {}, \"msg\": \"ok\" }",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "meta_databases"
                ],
                "summary": "新增数据源",
                "parameters": [
                    {
                        "type": "string",
                        "description": "数据源 名称",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "数据源 地址",
                        "name": "host",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "数据源 端口号",
                        "name": "port",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "数据源 数据库名称",
                        "name": "db_name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "数据源 用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "数据源 密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "数据源 备注",
                        "name": "comment",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "系统厂商 id",
                        "name": "vendor_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "业务系统 id",
                        "name": "business_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 200, data: {}, \"msg\" : \"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/meta_databases/{id}": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "meta_databases"
                ],
                "summary": "删除数据源",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "数据源 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\" : 200, \"msg\" : \"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "meta_databases"
                ],
                "summary": "修改数据源",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "数据源 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "数据源 名称",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "数据源 地址",
                        "name": "host",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "数据源 端口号",
                        "name": "port",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "数据源 数据库名称",
                        "name": "db_name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "数据源 用户名",
                        "name": "username",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "数据源 密码",
                        "name": "password",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "数据源 备注",
                        "name": "comment",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "系统厂商 id",
                        "name": "vendor_id",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "业务系统 id",
                        "name": "business_id",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\" : 200, \"msg\" : \"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/meta_databases/{id}/meta_tables": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "meta_tables"
                ],
                "summary": "元数据列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "数据源 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\" : 200, \"basic\" : {}, \"msg\": \"ok\" }",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/vendors": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vendors"
                ],
                "summary": "系统厂商列表",
                "responses": {
                    "200": {
                        "description": "{\"code\" : 200, \"data\" : {}, \"msg\" : \"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vendors"
                ],
                "summary": "新增系统厂商",
                "parameters": [
                    {
                        "type": "string",
                        "description": "系统厂商 名称",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "系统厂商 云端id",
                        "name": "c_id",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\" : 200, \"msg\" : \"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/vendors/{id}": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vendors"
                ],
                "summary": "删除系统厂商",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "系统厂商 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\" : 200, \"msg\" : \"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vendors"
                ],
                "summary": "修改系统厂商",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "系统厂商 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "系统厂商 名称",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "系统厂商 云端id",
                        "name": "c_id",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\" : 200, \"msg\" : \"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "/api",
	Schemes:     []string{},
	Title:       "Potato Api",
	Description: "This is a data_govern use golang",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
