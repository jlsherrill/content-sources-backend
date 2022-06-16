// Package api GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package api

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "https://www.apache.org/licenses/LICENSE-2.0"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/repositories/": {
            "get": {
                "description": "get repositories",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "repositories"
                ],
                "summary": "List Repositories",
                "operationId": "listRepositories",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.RepositoryCollectionResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "create a repository",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "repositories"
                ],
                "summary": "Create Repository",
                "operationId": "createRepository",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.RepositoryRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": ""
                    }
                }
            }
        },
        "/repositories/:uuid": {
            "delete": {
                "tags": [
                    "repositories"
                ],
                "summary": "Delete a repository",
                "operationId": "deleteRepository",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Identifier of the Repository",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/repositories/:uuid/rpms": {
            "get": {
                "description": "get repositories RPMS",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "repositories"
                ],
                "summary": "List Repositories RPMS",
                "operationId": "listRepositoriesRpms",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.RepositoryRpmCollectionResponse"
                        }
                    }
                }
            }
        },
        "/repositories/{uuid}": {
            "get": {
                "description": "Get information about a Repository",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "repositories"
                ],
                "summary": "Get Repository",
                "operationId": "getRepository",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Identifier of the Repository",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "put": {
                "description": "Fully update a repository",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "repositories"
                ],
                "summary": "Update Repository",
                "operationId": "fullUpdateRepository",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Identifier of the Repository",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.RepositoryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "patch": {
                "description": "Partially Update a repository",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "repositories"
                ],
                "summary": "Partial Update Repository",
                "operationId": "partialUpdateRepository",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Identifier of the Repository",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.RepositoryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "api.Links": {
            "type": "object",
            "properties": {
                "first": {
                    "description": "Path to first page of results",
                    "type": "string"
                },
                "last": {
                    "description": "Path to last page of results",
                    "type": "string"
                },
                "next": {
                    "description": "Path to next page of results",
                    "type": "string"
                },
                "prev": {
                    "description": "Path to previous page of results",
                    "type": "string"
                }
            }
        },
        "api.RepositoryCollectionResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "Requested Data",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.RepositoryResponse"
                    }
                },
                "links": {
                    "description": "Links to other pages of results",
                    "$ref": "#/definitions/api.Links"
                },
                "meta": {
                    "description": "Metadata about the request",
                    "$ref": "#/definitions/api.ResponseMetadata"
                }
            }
        },
        "api.RepositoryRequest": {
            "type": "object",
            "properties": {
                "distribution_arch": {
                    "description": "Architecture to restrict client usage to",
                    "type": "string",
                    "example": "x86_64"
                },
                "distribution_versions": {
                    "description": "Versions to restrict client usage to",
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "['7'",
                        "'8']"
                    ]
                },
                "name": {
                    "type": "string"
                },
                "url": {
                    "description": "URL of the remote yum repository",
                    "type": "string"
                }
            }
        },
        "api.RepositoryResponse": {
            "type": "object",
            "properties": {
                "account_id": {
                    "description": "Account ID of the owner",
                    "type": "string",
                    "readOnly": true
                },
                "distribution_arch": {
                    "description": "Architecture to restrict client usage to",
                    "type": "string",
                    "example": "x86_64"
                },
                "distribution_versions": {
                    "description": "Versions to restrict client usage to",
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "['7'",
                        "'8']"
                    ]
                },
                "name": {
                    "type": "string"
                },
                "org_id": {
                    "description": "Organization ID of the owner",
                    "type": "string",
                    "readOnly": true
                },
                "url": {
                    "description": "URL of the remote yum repository",
                    "type": "string"
                },
                "uuid": {
                    "type": "string",
                    "readOnly": true
                }
            }
        },
        "api.RepositoryRpm": {
            "type": "object",
            "properties": {
                "arch": {
                    "description": "The architecture that this package belong to",
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "epoch": {
                    "description": "Epoch is a way to define weighted dependencies based\non version numbers. It's default value is 0 and this\nis assumed if an Epoch directive is not listed in the RPM SPEC file.\nhttps://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/8/html/packaging_and_distributing_software/advanced-topics#packaging-epoch_epoch-scriplets-and-triggers",
                    "type": "integer"
                },
                "identity": {
                    "description": "Retrieve the identity header from x-rh-identity",
                    "type": "string"
                },
                "name": {
                    "description": "The rpm package name",
                    "type": "string"
                },
                "release": {
                    "description": "The release for this package",
                    "type": "string"
                },
                "summary": {
                    "type": "string"
                },
                "uuid": {
                    "description": "RPM id",
                    "type": "string"
                },
                "version": {
                    "description": "The version for this package",
                    "type": "string"
                }
            }
        },
        "api.RepositoryRpmCollectionResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "Requested Data",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.RepositoryRpm"
                    }
                },
                "links": {
                    "description": "Links to other pages of results",
                    "$ref": "#/definitions/api.Links"
                },
                "meta": {
                    "description": "Metadata about the request",
                    "$ref": "#/definitions/api.ResponseMetadata"
                }
            }
        },
        "api.ResponseMetadata": {
            "type": "object",
            "properties": {
                "count": {
                    "description": "Total count of results",
                    "type": "integer"
                },
                "limit": {
                    "description": "Limit of results used for the request",
                    "type": "integer"
                },
                "offset": {
                    "description": "Offset into results used for the request",
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "RhIdentity": {
            "type": "apiKey",
            "name": "x-rh-identity",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v1.0.0",
	Host:             "api.example.com",
	BasePath:         "/api/content_sources/v1.0/",
	Schemes:          []string{},
	Title:            "ContentSourcesBackend",
	Description:      "API of the Content Sources application on [console.redhat.com](https://console.redhat.com)\n",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
