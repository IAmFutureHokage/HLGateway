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
        "/api/add-telegram": {
            "post": {
                "description": "Add a new telegram",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Telegram"
                ],
                "summary": "Add telegram",
                "parameters": [
                    {
                        "description": "Add Telegram Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AddTelegramRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.AddTelegramResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.AddTelegramRequest": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                }
            }
        },
        "dto.AddTelegramResponse": {
            "type": "object",
            "properties": {
                "telegrams": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.Telegram"
                    }
                }
            }
        },
        "dto.IcePhenomenia": {
            "type": "object",
            "properties": {
                "intensity": {
                    "type": "integer"
                },
                "phenomen": {
                    "type": "integer"
                }
            }
        },
        "dto.Telegram": {
            "type": "object",
            "properties": {
                "air_temperature": {
                    "type": "integer"
                },
                "average_reservoir_level": {
                    "type": "integer"
                },
                "datetime": {
                    "type": "string"
                },
                "delta_water_level": {
                    "type": "integer"
                },
                "downstream_level": {
                    "type": "integer"
                },
                "group_id": {
                    "type": "string"
                },
                "headwater_level": {
                    "type": "integer"
                },
                "ice_height": {
                    "type": "integer"
                },
                "ice_phenomenia_state": {
                    "type": "integer"
                },
                "ice_phenomenias": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.IcePhenomenia"
                    }
                },
                "id": {
                    "type": "string"
                },
                "inflow": {
                    "type": "number"
                },
                "is_dangerous": {
                    "type": "boolean"
                },
                "post_code": {
                    "type": "string"
                },
                "precipitation_duration": {
                    "type": "integer"
                },
                "precipitation_value": {
                    "type": "number"
                },
                "reservoir_date": {
                    "type": "string"
                },
                "reservoir_volume": {
                    "type": "number"
                },
                "reservoir_water_inflow_date": {
                    "type": "string"
                },
                "reset": {
                    "type": "number"
                },
                "snow_height": {
                    "type": "integer"
                },
                "telegram_code": {
                    "type": "string"
                },
                "water_flow": {
                    "type": "number"
                },
                "water_level_on20h": {
                    "type": "integer"
                },
                "water_level_on_time": {
                    "type": "integer"
                },
                "water_temperature": {
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "HLGateway API",
	Description:      "This is the HLGateway server API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
