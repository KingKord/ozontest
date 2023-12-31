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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/original": {
            "get": {
                "description": "get URL back from short URL",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "URL"
                ],
                "summary": "get URL back",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Short URL",
                        "name": "ShortURL",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful returned",
                        "schema": {
                            "$ref": "#/definitions/main.jsonResponse"
                        }
                    }
                }
            }
        },
        "/shorten": {
            "post": {
                "description": "save URL To DB and get short URL back",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "URL"
                ],
                "summary": "URL shortener",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Long URL",
                        "name": "LongURL",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful shorten",
                        "schema": {
                            "$ref": "#/definitions/main.jsonResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.jsonResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "test task FrontEnd",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
