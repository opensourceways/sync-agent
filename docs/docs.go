// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/synchronization/{platform}/comment": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Synchronization"
                ],
                "summary": "同步 gitee 或 github 平台的 comment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "平台：gitee 或 github",
                        "name": "platform",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "需要同步的comment",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SyncComment"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "同步成功",
                        "schema": {
                            "$ref": "#/definitions/models.BaseResp"
                        }
                    },
                    "400": {
                        "description": "错误返回",
                        "schema": {
                            "$ref": "#/definitions/models.BaseResp"
                        }
                    },
                    "404": {
                        "description": "错误返回",
                        "schema": {
                            "$ref": "#/definitions/models.BaseResp"
                        }
                    }
                }
            }
        },
        "/synchronization/{platform}/issue": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Synchronization"
                ],
                "summary": "同步更新 gitee 或 github 平台的 issue",
                "parameters": [
                    {
                        "type": "string",
                        "description": "平台：gitee 或 github",
                        "name": "platform",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "需要跟新的issue信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.IssueUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "同步成功",
                        "schema": {
                            "$ref": "#/definitions/models.BaseResp"
                        }
                    },
                    "400": {
                        "description": "错误返回",
                        "schema": {
                            "$ref": "#/definitions/models.BaseResp"
                        }
                    },
                    "404": {
                        "description": "错误返回",
                        "schema": {
                            "$ref": "#/definitions/models.BaseResp"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Synchronization"
                ],
                "summary": "同步 gitee 或 github 平台的 issue",
                "parameters": [
                    {
                        "type": "string",
                        "description": "平台：gitee 或 github",
                        "name": "platform",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "需要同步的issue",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Issue"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "同步成功",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.BaseResp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.SyncIssueResult"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "错误返回",
                        "schema": {
                            "$ref": "#/definitions/models.BaseResp"
                        }
                    },
                    "404": {
                        "description": "错误返回",
                        "schema": {
                            "$ref": "#/definitions/models.BaseResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.BaseResp": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "data": {
                    "type": "object"
                },
                "msg": {
                    "type": "string",
                    "example": "请求成功/失败"
                }
            }
        },
        "models.SyncComment": {
            "type": "object",
            "required": [
                "org",
                "repo"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "number": {
                    "type": "string"
                },
                "org": {
                    "type": "string"
                },
                "repo": {
                    "type": "string"
                }
            }
        },
        "models.Issue": {
            "type": "object",
            "required": [
                "content",
                "org",
                "repo",
                "title"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "org": {
                    "type": "string"
                },
                "repo": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.IssueUpdate": {
            "type": "object",
            "required": [
                "number",
                "org",
                "repo",
                "state"
            ],
            "properties": {
                "number": {
                    "type": "string"
                },
                "org": {
                    "type": "string"
                },
                "repo": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                }
            }
        },
        "models.SyncIssueResult": {
            "type": "object",
            "required": [
                "org",
                "repo"
            ],
            "properties": {
                "link": {
                    "type": "string"
                },
                "number": {
                    "type": "string"
                },
                "org": {
                    "type": "string"
                },
                "repo": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "access-token",
            "in": "header"
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
	Version:     "0.0.1",
	Host:        "",
	BasePath:    "/v1",
	Schemes:     []string{},
	Title:       "Swagger sync-agent API",
	Description: "plugin maintenance server api doc",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
