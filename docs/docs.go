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
        "/paper-wallet-core-service/users": {
            "get": {
                "description": "Retrieve a list of all users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user-controller"
                ],
                "summary": "Get all users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.User"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Update the details of an existing user based on the provided user ID and data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user-controller"
                ],
                "summary": "Update an existing user",
                "parameters": [
                    {
                        "description": "User Data for Update",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.UpdateUserRequestDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.User"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new user with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user-controller"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.CreateUserRequestDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/user.CreateUserResponseDto"
                        }
                    }
                }
            }
        },
        "/paper-wallet-core-service/users/{id}": {
            "get": {
                "description": "Get details of a user by their ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user-controller"
                ],
                "summary": "Get user by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.User"
                        }
                    }
                }
            }
        },
        "/paper-wallet-core-service/wallet/withdraw": {
            "post": {
                "description": "Withdraw funds from the user's wallet",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet-controller"
                ],
                "summary": "Withdraw funds from wallet",
                "parameters": [
                    {
                        "description": "Withdraw Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/withdraw.WithdrawRequestDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Withdrawal Successful",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.User": {
            "type": "object",
            "required": [
                "balance",
                "currency",
                "name",
                "scale"
            ],
            "properties": {
                "balance": {
                    "type": "number"
                },
                "createdBy": {
                    "type": "string"
                },
                "createdDate": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "isDeleted": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "scale": {
                    "type": "integer"
                },
                "updatedBy": {
                    "type": "string"
                },
                "updatedDate": {
                    "type": "string"
                },
                "version": {
                    "type": "integer"
                }
            }
        },
        "user.CreateUserRequestDto": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/domain.User"
                }
            }
        },
        "user.CreateUserResponseDto": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/domain.User"
                }
            }
        },
        "user.UpdateUserRequestDto": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/domain.User"
                }
            }
        },
        "withdraw.WithdrawRequestDto": {
            "type": "object",
            "required": [
                "amount",
                "userId"
            ],
            "properties": {
                "amount": {
                    "type": "number"
                },
                "userId": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
