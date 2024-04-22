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
        "/exam-prep/generate": {
            "post": {
                "description": "Generate exam preparation material based on the provided topics and prompt",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "exam_prep"
                ],
                "summary": "Generate exam preparation material",
                "parameters": [
                    {
                        "description": "Generate exam preparation material criteria",
                        "name": "generateDTO",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.GenerateExamPrepDTO"
                        }
                    },
                    {
                        "type": "string",
                        "default": "Bearer",
                        "description": "Authorization header",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Markdown formatted exam preparation material",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/experiments/add": {
            "post": {
                "description": "Add a new experiment to the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "experiments"
                ],
                "summary": "Add a new experiment",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer",
                        "description": "Authorization header",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "New experiment data",
                        "name": "addDTO",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.ExperimentDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.ExperimentDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/experiments/chapter/{chapterID}": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Get a list of experiments for a specific chapter",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "experiments"
                ],
                "summary": "Get experiments for a chapter",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Chapter ID",
                        "name": "chapterID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "Bearer",
                        "description": "Authorization header",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dtos.ExperimentDTO"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/experiments/content/{experimentID}": {
            "get": {
                "description": "Get the content for a specific experiment",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "experiments"
                ],
                "summary": "Get content for an experiment",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer",
                        "description": "Authorization header",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Experiment ID",
                        "name": "experimentID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/experiments/delete/{experimentID}": {
            "delete": {
                "description": "Delete an experiment from the database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "experiments"
                ],
                "summary": "Delete an experiment",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer",
                        "description": "Authorization header",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Experiment ID",
                        "name": "experimentID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/experiments/generate": {
            "post": {
                "description": "Generate experiments for a specific chapter based on provided criteria",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "experiments"
                ],
                "summary": "Generate experiments for a chapter",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer",
                        "description": "Authorization header",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Generate experiment criteria",
                        "name": "generateDTO",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.GenerateExperimentDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dtos.ExperimentDTO"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/experiments/subject/{subjectID}": {
            "get": {
                "description": "Get a list of experiments for a specific subject",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "experiments"
                ],
                "summary": "Get experiments for a subject",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer",
                        "description": "Authorization header",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Subject ID",
                        "name": "subjectID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dtos.ExperimentDTO"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/experiments/update/{experimentID}": {
            "put": {
                "description": "Update an existing experiment in the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "experiments"
                ],
                "summary": "Update an existing experiment",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer",
                        "description": "Authorization header",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Experiment ID",
                        "name": "experimentID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated experiment data",
                        "name": "updateDTO",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.ExperimentDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.ExperimentDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/quiz/generate": {
            "post": {
                "description": "Generate a quiz based on the provided topics and prompt",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "quiz"
                ],
                "summary": "Generate a quiz",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer",
                        "description": "Authorization header",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Generate quiz criteria",
                        "name": "generateDTO",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.GenerateQuizDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dtos.Quiz"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/admin/register": {
            "post": {
                "description": "Register a new admin user with the provided information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Register a new admin user",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer",
                        "description": "Authorization header",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Admin user registration information",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.RegisterDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.UserDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/delete/{id}": {
            "delete": {
                "description": "Delete a user by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Delete a user",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer",
                        "description": "Authorization header",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
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
                            "$ref": "#/definitions/dtos.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "Sign in a user with the provided credentials",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Sign in a user",
                "parameters": [
                    {
                        "description": "User sign-in information",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.SignInDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.UserDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/logout": {
            "post": {
                "description": "Logout a user by invalidating the refresh token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Logout a user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/me": {
            "get": {
                "description": "Retrieve the profile details of the currently authenticated user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user profile details",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer",
                        "description": "Authorization header",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "user profile details",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "description": "Register a new user with the provided information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User registration information",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.RegisterDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.UserDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/update/{id}": {
            "put": {
                "description": "Update the information of the currently authenticated user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update user information",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer",
                        "description": "Authorization header",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User update information",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.UpdateUserDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.UserDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.UserRole": {
            "type": "string",
            "enum": [
                "teacher",
                "student",
                "admin"
            ],
            "x-enum-varnames": [
                "Teacher",
                "Student",
                "Admin"
            ]
        },
        "dtos.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "dtos.ExperimentDTO": {
            "type": "object",
            "properties": {
                "chapter_id": {
                    "type": "integer"
                },
                "content_link": {
                    "type": "string"
                },
                "experiment_name": {
                    "type": "string"
                },
                "topic_id": {
                    "type": "integer"
                }
            }
        },
        "dtos.GenerateExamPrepDTO": {
            "type": "object",
            "properties": {
                "prompt": {
                    "type": "string"
                },
                "readTime": {
                    "type": "integer"
                },
                "topics": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "dtos.GenerateExperimentDTO": {
            "type": "object",
            "properties": {
                "chapterID": {
                    "type": "integer"
                },
                "promptMessage": {
                    "type": "string"
                },
                "subjectID": {
                    "type": "integer"
                }
            }
        },
        "dtos.GenerateQuizDTO": {
            "type": "object",
            "properties": {
                "number_of_quizzes": {
                    "type": "integer"
                },
                "prompt": {
                    "type": "string"
                },
                "topics": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "dtos.Option": {
            "type": "object",
            "properties": {
                "is_answer": {
                    "type": "boolean"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "dtos.Quiz": {
            "type": "object",
            "properties": {
                "explanation": {
                    "type": "string"
                },
                "options": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dtos.Option"
                    }
                },
                "question": {
                    "type": "string"
                }
            }
        },
        "dtos.RegisterDTO": {
            "type": "object",
            "required": [
                "email",
                "password",
                "role",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dtos.SignInDTO": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dtos.SuccessResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "dtos.UpdateUserDTO": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dtos.UserDTO": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "refreshToken": {
                    "type": "string"
                },
                "role": {
                    "$ref": "#/definitions/domain.UserRole"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "learning-app-idt8.onrender.com",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Brilliant Learning APP API",
	Description:      "This is the API for the Brilliant application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
