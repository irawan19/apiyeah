// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/data/jeniskelamin": {
            "get": {
                "description": "ambil jenis kelamin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Data"
                ],
                "summary": "ambil jenis kelamin",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/data/konfigurasiaplikasi": {
            "get": {
                "description": "ambil konfigurasi aplikasi",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Data"
                ],
                "summary": "ambil konfigurasi aplikasi",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/data/metatag": {
            "get": {
                "description": "ambil meta tag",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Data"
                ],
                "summary": "ambil meta tag",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/data/pembayaran": {
            "get": {
                "description": "ambil pembayaran",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Data"
                ],
                "summary": "ambil pembayaran",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/data/sosialmedia": {
            "get": {
                "description": "ambil sosial media",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Data"
                ],
                "summary": "ambil sosial media",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/data/statuspembayaran": {
            "get": {
                "description": "ambil status pembayaran",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Data"
                ],
                "summary": "ambil status pembayaran",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/event": {
            "get": {
                "description": "ambil semua event",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event"
                ],
                "summary": "ambil semua event",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/event/cekticket": {
            "post": {
                "description": "cek ticket",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event"
                ],
                "summary": "cek ticket",
                "parameters": [
                    {
                        "type": "string",
                        "description": "booking code",
                        "name": "booking_code",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/event/detail/{id}": {
            "get": {
                "description": "ambil detail event",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event"
                ],
                "summary": "ambil detail event",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id event",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/event/registrasi": {
            "post": {
                "description": "registrasi",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event"
                ],
                "summary": "registrasi",
                "parameters": [
                    {
                        "description": "registrasi",
                        "name": "registrasi",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/event/registrasi/buktipembayaran": {
            "post": {
                "description": "bukti pembayaran registrasi",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event"
                ],
                "summary": "bukti pembayaran registrasi",
                "parameters": [
                    {
                        "type": "string",
                        "description": "booking code",
                        "name": "booking_code",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "bukti pembayaran",
                        "name": "bukti_pembayaran_registrasi_events",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/event/registrasi/pembayaran": {
            "post": {
                "description": "pembayaran registrasi",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event"
                ],
                "summary": "pembayaran registrasi",
                "parameters": [
                    {
                        "type": "string",
                        "description": "booking code",
                        "name": "booking_code",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "id pembayaran",
                        "name": "id_pembayarans",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "id status pembayaran",
                        "name": "id_status_pembayarans",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/event/verifikasikedatangan": {
            "post": {
                "description": "verifikasi kedatangan",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event"
                ],
                "summary": "verifikasi kedatangan",
                "parameters": [
                    {
                        "type": "string",
                        "description": "booking code",
                        "name": "booking_code",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "kode scanner event",
                        "name": "kode_scanner_events",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "api.yeah.biz.id",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "API Yeah",
	Description:      "API Using Golang",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
