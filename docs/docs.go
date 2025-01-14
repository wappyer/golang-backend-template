// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "项目地址",
            "url": "https://gitee.com/wappyer/golang-backend-template.git"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {},
    "tags": [
        {
            "description": "登录相关",
            "name": "login"
        },
        {
            "description": "用户信息相关",
            "name": "user"
        },
        {
            "description": "信息配置相关",
            "name": "info"
        }
    ]
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "golang后端接口项目模版",
	Description:      "#### 接口调用须知：\n**header请求头**\n```shell\nAuthorization: {{token}}  //登录凭证,token通过对应登录接口获取\nUser-Role: user  //角色, user用户/admin管理员\n```",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
