// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Shoyo",
            "url": "https://github.com/shoyo10/jubobe"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/orders": {
            "post": {
                "description": "create a order",
                "parameters": [
                    {
                        "description": "order fields",
                        "name": "reqBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.createOrderReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "order id",
                        "schema": {
                            "$ref": "#/definitions/http.createOrderResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.HTTPError"
                        }
                    }
                }
            }
        },
        "/patients": {
            "get": {
                "description": "list all patients",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.listPatientsResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "errors.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "message": {}
            }
        },
        "http.createOrderReq": {
            "type": "object",
            "properties": {
                "Message": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 1
                },
                "PatientId": {
                    "type": "integer",
                    "minimum": 1
                }
            }
        },
        "http.createOrderResp": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/http.createOrderRespData"
                }
            }
        },
        "http.createOrderRespData": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "http.listPatientsResp": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/http.listPatientsRespData"
                    }
                }
            }
        },
        "http.listPatientsRespData": {
            "type": "object",
            "properties": {
                "Id": {
                    "type": "integer"
                },
                "Name": {
                    "type": "string"
                },
                "OrderId": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9090",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Jubobe API Document",
	Description:      "This is jubo backend api document.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
