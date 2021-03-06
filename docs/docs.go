// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/addCertificate": {
            "patch": {
                "description": "allows you add a certificate to your skill",
                "consumes": [
                    "application/json"
                ],
                "summary": "Add a new certificate",
                "parameters": [
                    {
                        "description": "Complete the json to create new path",
                        "name": "body-json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "type": "string",
                        "description": "pathID",
                        "name": "pathid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Certificate added. Congratulations!",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "could not notify the admin",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "default": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/addSkill": {
            "patch": {
                "description": "get user careerpath",
                "consumes": [
                    "application/json"
                ],
                "summary": "add new skill to the user path",
                "parameters": [
                    {
                        "description": "Allows to complete the json for login",
                        "name": "body-json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Insert path id",
                        "name": "pathid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Skills added, the admin will check them, please wait...",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Couldn??t add the skills to the path: ",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "default": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/createPath": {
            "post": {
                "description": "allows you to create a new path if you are an admin",
                "consumes": [
                    "application/json"
                ],
                "summary": "Create new career path",
                "parameters": [
                    {
                        "description": "Complete the json to create new path",
                        "name": "body-json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Path has been created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "could not create Path",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "default": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/createRelatedUserPath": {
            "post": {
                "description": "allows you to create a new user's path relation if you are an admin",
                "produces": [
                    "application/json"
                ],
                "summary": "create the user's path relation",
                "parameters": [
                    {
                        "description": "Allows to complete the json for login",
                        "name": "body-json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "type": "string",
                        "description": "PathID",
                        "name": "pathid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "UserID",
                        "name": "userid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User successfuly related to the selected path",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "couldn??t create user path relation",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "default": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/getPath": {
            "get": {
                "description": "get user careerpath",
                "consumes": [
                    "application/json"
                ],
                "summary": "shows a complete original career path",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success show careerpath",
                        "schema": {
                            "type": "string"
                        },
                        "headers": {
                            "Token": {
                                "type": "string",
                                "description": "jwtKey"
                            }
                        }
                    },
                    "400": {
                        "description": "missing id parameter",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "error occurred looking the path",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "default": {
                        "description": "Error show careerpath",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/getUserPath": {
            "get": {
                "description": "get user careerpath",
                "consumes": [
                    "application/json"
                ],
                "summary": "shows a complete user's career path",
                "parameters": [
                    {
                        "type": "string",
                        "description": "pathID",
                        "name": "pathid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "userID",
                        "name": "userid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "status ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "user not related with given path",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "default": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "get the email and password to access",
                "consumes": [
                    "application/json"
                ],
                "summary": "Enter the system",
                "operationId": "login",
                "parameters": [
                    {
                        "description": "Allows to complete the json for login",
                        "name": "body-json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success Login",
                        "schema": {
                            "type": "string"
                        },
                        "headers": {
                            "Token": {
                                "type": "string",
                                "description": "jwtKey"
                            }
                        }
                    },
                    "400": {
                        "description": "invalid login",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "invalid login",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "default": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/modifyPassword": {
            "put": {
                "description": "get the user and password update",
                "summary": "Modify user password",
                "parameters": [
                    {
                        "description": "Allows to replace the password for login",
                        "name": "New-password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Password updated",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "something went wrong with the given data or couldn't update password",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "something went wrong with the given data or couldn't update password",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "default": {
                        "description": "Invalid Password Update",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/notifications": {
            "get": {
                "description": "Get all the notifications from the users",
                "consumes": [
                    "application/json"
                ],
                "summary": "Get notifications",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Insert the page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "status ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Coudn??t get the notifications",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "default": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/passwordRecovery": {
            "put": {
                "description": "send email at the person what forgot the password",
                "produces": [
                    "application/json"
                ],
                "summary": "recovery the password if dont remember",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        },
                        "headers": {
                            "Token": {
                                "type": "string",
                                "description": "jwtKey"
                            }
                        }
                    },
                    "400": {
                        "description": "invalid email, please write your email or the given email is not registered",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "invalid email, please write your email or the given email is not registered",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "default": {
                        "description": "error processing password recovery",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "ask for email and password for register in the app",
                "produces": [
                    "application/json"
                ],
                "summary": "creates the user register in the db",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User credentials",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Password",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "profilePhoto",
                        "name": "profilePhoto",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "surname",
                        "name": "surname",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "role",
                        "name": "role",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "user successfully registered",
                        "schema": {
                            "type": "string"
                        },
                        "headers": {
                            "Token": {
                                "type": "string",
                                "description": "jwtKey"
                            }
                        }
                    },
                    "400": {
                        "description": "something went wrong with the given data, error or the given email is already in use",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "something went wrong with the given data, error or the given email is already in use",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "default": {
                        "description": "error in register",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/validateOrModifyUserPath": {
            "put": {
                "description": "Validate and modify the user career path if you are the admin",
                "consumes": [
                    "application/json"
                ],
                "summary": "Validate and modify the user path",
                "parameters": [
                    {
                        "description": "Allows to complete the json for login",
                        "name": "body-json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "type": "string",
                        "description": "pathID",
                        "name": "pathid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "userID",
                        "name": "userid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "message",
                        "name": "message",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Path has been modified",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "could send feedback to user. ",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "default": {
                        "description": "Error",
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
