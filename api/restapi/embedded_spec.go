// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "BolusWizard",
    "version": "1.0"
  },
  "paths": {
    "/CorrectionRange": {
      "get": {
        "summary": "get correction range",
        "operationId": "correctionRange",
        "parameters": [
          {
            "type": "string",
            "name": "authToken",
            "in": "header",
            "required": true
          },
          {
            "type": "integer",
            "name": "Bg",
            "in": "query",
            "required": true
          },
          {
            "type": "integer",
            "name": "Carbs",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "object",
              "properties": {
                "Data": {
                  "$ref": "#/definitions/CorrectionResponse"
                },
                "Error": {
                  "type": "object"
                },
                "Status": {
                  "type": "integer"
                }
              }
            }
          },
          "default": {
            "description": "Unsuccessful operation",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          }
        }
      }
    },
    "/Duration": {
      "get": {
        "summary": "Get Duration",
        "operationId": "getDuration",
        "parameters": [
          {
            "type": "string",
            "name": "authToken",
            "in": "header",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "object",
              "properties": {
                "Data": {
                  "type": "string"
                },
                "Error": {
                  "type": "object"
                },
                "Status": {
                  "type": "integer"
                }
              }
            }
          },
          "default": {
            "description": "Unsuccessful operation",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          }
        }
      },
      "post": {
        "summary": "Create Duration",
        "operationId": "createDuration",
        "parameters": [
          {
            "type": "string",
            "name": "authToken",
            "in": "header",
            "required": true
          },
          {
            "name": "Duration",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "Duration": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "object",
              "properties": {
                "Data": {
                  "type": "object"
                },
                "Error": {
                  "type": "object"
                },
                "Status": {
                  "type": "integer"
                }
              }
            }
          },
          "default": {
            "description": "Unsuccessful operation",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          }
        }
      }
    },
    "/ISF": {
      "get": {
        "summary": "Get ISFs",
        "operationId": "getISFs",
        "parameters": [
          {
            "type": "string",
            "name": "authToken",
            "in": "header",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "object",
              "properties": {
                "Data": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/ISF"
                  }
                },
                "Error": {
                  "type": "object"
                },
                "Status": {
                  "type": "integer"
                }
              }
            }
          },
          "default": {
            "description": "Unsuccessful operation",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          }
        }
      },
      "post": {
        "summary": "Create ISF",
        "operationId": "createISFs",
        "parameters": [
          {
            "type": "string",
            "name": "authToken",
            "in": "header",
            "required": true
          },
          {
            "name": "targetRatios",
            "in": "body",
            "required": true,
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/ISF"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "object",
              "properties": {
                "Data": {
                  "type": "object"
                },
                "Error": {
                  "type": "object"
                },
                "Status": {
                  "type": "integer"
                }
              }
            }
          },
          "default": {
            "description": "Unsuccessful operation",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          }
        }
      }
    },
    "/NewCorrection": {
      "get": {
        "summary": "New Correction",
        "operationId": "newCorrection",
        "parameters": [
          {
            "type": "string",
            "name": "authToken",
            "in": "header",
            "required": true
          },
          {
            "type": "integer",
            "name": "Bg",
            "in": "query",
            "required": true
          },
          {
            "type": "integer",
            "name": "Carbs",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "object",
              "properties": {
                "Data": {
                  "$ref": "#/definitions/CorrectionResponse"
                },
                "Error": {
                  "type": "object"
                },
                "Status": {
                  "type": "integer"
                }
              }
            }
          },
          "default": {
            "description": "Unsuccessful operation",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          }
        }
      }
    },
    "/SignIn": {
      "post": {
        "summary": "Sign In",
        "operationId": "signIn",
        "parameters": [
          {
            "name": "signInData",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Sign"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "object",
              "properties": {
                "Data": {
                  "$ref": "#/definitions/Token"
                },
                "Error": {
                  "type": "object"
                },
                "Status": {
                  "type": "integer"
                }
              }
            }
          },
          "default": {
            "description": "Unsuccessful operation",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          }
        }
      }
    },
    "/SignUp": {
      "post": {
        "summary": "Sign Up",
        "operationId": "signUp",
        "parameters": [
          {
            "name": "signUpData",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Sign"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "object",
              "properties": {
                "Data": {
                  "$ref": "#/definitions/Token"
                },
                "Error": {
                  "type": "object"
                },
                "Status": {
                  "type": "integer"
                }
              }
            }
          },
          "default": {
            "description": "Unsuccessful operation",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          }
        }
      }
    },
    "/Targets": {
      "post": {
        "summary": "Create Target Ratios",
        "operationId": "createTargets",
        "parameters": [
          {
            "type": "string",
            "name": "authToken",
            "in": "header",
            "required": true
          },
          {
            "name": "targetRatios",
            "in": "body",
            "required": true,
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Target"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "object",
              "properties": {
                "Data": {
                  "type": "object"
                },
                "Error": {
                  "type": "object"
                },
                "Status": {
                  "type": "integer"
                }
              }
            }
          },
          "default": {
            "description": "Unsuccessful operation",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          }
        }
      }
    },
    "/corrections": {
      "get": {
        "summary": "Get Corrections",
        "operationId": "getCorrections",
        "parameters": [
          {
            "type": "string",
            "name": "authToken",
            "in": "header",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "object",
              "properties": {
                "Data": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/Correction"
                  }
                },
                "Error": {
                  "type": "object"
                },
                "Status": {
                  "type": "integer"
                }
              }
            }
          },
          "default": {
            "description": "Unsuccessful operation",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          }
        }
      },
      "post": {
        "summary": "Create Corrections",
        "operationId": "createCorrections",
        "parameters": [
          {
            "type": "string",
            "name": "authToken",
            "in": "header",
            "required": true
          },
          {
            "name": "corrections",
            "in": "body",
            "required": true,
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Correction"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          },
          "default": {
            "description": "Unsuccessful operation",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          }
        }
      }
    },
    "/health-check": {
      "get": {
        "summary": "Health Check",
        "operationId": "healthCheck",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/ratios": {
      "post": {
        "summary": "Create Carb Ratios",
        "operationId": "createRatios",
        "parameters": [
          {
            "type": "string",
            "name": "authToken",
            "in": "header",
            "required": true
          },
          {
            "name": "carbRatios",
            "in": "body",
            "required": true,
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/CarbRatio"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "object",
              "properties": {
                "Data": {
                  "type": "object"
                },
                "Error": {
                  "type": "object"
                },
                "Status": {
                  "description": "status code",
                  "type": "integer"
                }
              }
            }
          },
          "default": {
            "description": "Unsuccessful operation",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "CarbRatio": {
      "type": "object",
      "properties": {
        "End": {
          "type": "string"
        },
        "Key": {
          "type": "string"
        },
        "Ratio": {
          "type": "number",
          "format": "float64"
        },
        "Start": {
          "type": "string"
        }
      }
    },
    "Correction": {
      "type": "object",
      "properties": {
        "Bg": {
          "type": "number",
          "format": "float64"
        },
        "Bolus": {
          "type": "number",
          "format": "float64"
        },
        "Carbs": {
          "type": "number",
          "format": "float64"
        },
        "Key": {
          "type": "string"
        },
        "TimeStamp": {
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "CorrectionResponse": {
      "type": "object",
      "properties": {
        "ActiveInsulinReduction": {
          "type": "number",
          "format": "float64"
        },
        "BgCorrection": {
          "type": "number",
          "format": "float64"
        },
        "Bolus": {
          "type": "number",
          "format": "float64"
        },
        "CarbCorrection": {
          "type": "number",
          "format": "float64"
        }
      }
    },
    "ISF": {
      "type": "object",
      "properties": {
        "End": {
          "type": "string"
        },
        "Key": {
          "type": "string"
        },
        "Sensitivity": {
          "type": "number",
          "format": "float64"
        },
        "Start": {
          "type": "string"
        }
      }
    },
    "Response": {
      "type": "object",
      "properties": {
        "Error": {
          "type": "object"
        },
        "Status": {
          "type": "integer"
        }
      }
    },
    "Sign": {
      "type": "object",
      "properties": {
        "Password": {
          "type": "string"
        },
        "Username": {
          "type": "string"
        }
      }
    },
    "Target": {
      "type": "object",
      "properties": {
        "End": {
          "type": "string"
        },
        "Key": {
          "type": "string"
        },
        "Ratio": {
          "type": "number",
          "format": "float64"
        },
        "Start": {
          "type": "string"
        }
      }
    },
    "Token": {
      "type": "object",
      "properties": {
        "Token": {
          "type": "string"
        },
        "Username": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "BolusWizard",
    "version": "1.0"
  },
  "paths": {
    "/CorrectionRange": {
      "get": {
        "summary": "get correction range",
        "operationId": "correctionRange",
        "parameters": [
          {
            "type": "string",
            "name": "authToken",
            "in": "header",
            "required": true
          },
          {
            "type": "integer",
            "name": "Bg",
            "in": "query",
            "required": true
          },
          {
            "type": "integer",
            "name": "Carbs",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "object",
              "properties": {
                "Data": {
                  "$ref": "#/definitions/CorrectionResponse"
                },
                "Error": {
                  "type": "object"
                },
                "Status": {
                  "type": "integer"
                }
              }
            }
          },
          "default": {
            "description": "Unsuccessful operation",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          }
        }
      }
    },
    "/Duration": {
      "get": {
        "summary": "Get Duration",
        "operationId": "getDuration",
        "parameters": [
          {
            "type": "string",
            "name": "authToken",
            "in": "header",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "object",
              "properties": {
                "Data": {
                  "type": "string"
                },
                "Error": {
                  "type": "object"
                },
                "Status": {
                  "type": "integer"
                }
              }
            }
          },
          "default": {
            "description": "Unsuccessful operation",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          }
        }
      },
      "post": {
        "summary": "Create Duration",
        "operationId": "createDuration",
        "parameters": [
          {
            "type": "string",
            "name": "authToken",
            "in": "header",
            "required": true
          },
          {
            "name": "Duration",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "Duration": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "object",
              "properties": {
                "Data": {
                  "type": "object"
                },
                "Error": {
                  "type": "object"
                },
                "Status": {
                  "type": "integer"
                }
              }
            }
          },
          "default": {
            "description": "Unsuccessful operation",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          }
        }
      }
    },
    "/ISF": {
      "get": {
        "summary": "Get ISFs",
        "operationId": "getISFs",
        "parameters": [
          {
            "type": "string",
            "name": "authToken",
            "in": "header",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "object",
              "properties": {
                "Data": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/ISF"
                  }
                },
                "Error": {
                  "type": "object"
                },
                "Status": {
                  "type": "integer"
                }
              }
            }
          },
          "default": {
            "description": "Unsuccessful operation",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          }
        }
      },
      "post": {
        "summary": "Create ISF",
        "operationId": "createISFs",
        "parameters": [
          {
            "type": "string",
            "name": "authToken",
            "in": "header",
            "required": true
          },
          {
            "name": "targetRatios",
            "in": "body",
            "required": true,
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/ISF"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "object",
              "properties": {
                "Data": {
                  "type": "object"
                },
                "Error": {
                  "type": "object"
                },
                "Status": {
                  "type": "integer"
                }
              }
            }
          },
          "default": {
            "description": "Unsuccessful operation",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          }
        }
      }
    },
    "/NewCorrection": {
      "get": {
        "summary": "New Correction",
        "operationId": "newCorrection",
        "parameters": [
          {
            "type": "string",
            "name": "authToken",
            "in": "header",
            "required": true
          },
          {
            "type": "integer",
            "name": "Bg",
            "in": "query",
            "required": true
          },
          {
            "type": "integer",
            "name": "Carbs",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "object",
              "properties": {
                "Data": {
                  "$ref": "#/definitions/CorrectionResponse"
                },
                "Error": {
                  "type": "object"
                },
                "Status": {
                  "type": "integer"
                }
              }
            }
          },
          "default": {
            "description": "Unsuccessful operation",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          }
        }
      }
    },
    "/SignIn": {
      "post": {
        "summary": "Sign In",
        "operationId": "signIn",
        "parameters": [
          {
            "name": "signInData",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Sign"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "object",
              "properties": {
                "Data": {
                  "$ref": "#/definitions/Token"
                },
                "Error": {
                  "type": "object"
                },
                "Status": {
                  "type": "integer"
                }
              }
            }
          },
          "default": {
            "description": "Unsuccessful operation",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          }
        }
      }
    },
    "/SignUp": {
      "post": {
        "summary": "Sign Up",
        "operationId": "signUp",
        "parameters": [
          {
            "name": "signUpData",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Sign"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "object",
              "properties": {
                "Data": {
                  "$ref": "#/definitions/Token"
                },
                "Error": {
                  "type": "object"
                },
                "Status": {
                  "type": "integer"
                }
              }
            }
          },
          "default": {
            "description": "Unsuccessful operation",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          }
        }
      }
    },
    "/Targets": {
      "post": {
        "summary": "Create Target Ratios",
        "operationId": "createTargets",
        "parameters": [
          {
            "type": "string",
            "name": "authToken",
            "in": "header",
            "required": true
          },
          {
            "name": "targetRatios",
            "in": "body",
            "required": true,
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Target"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "object",
              "properties": {
                "Data": {
                  "type": "object"
                },
                "Error": {
                  "type": "object"
                },
                "Status": {
                  "type": "integer"
                }
              }
            }
          },
          "default": {
            "description": "Unsuccessful operation",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          }
        }
      }
    },
    "/corrections": {
      "get": {
        "summary": "Get Corrections",
        "operationId": "getCorrections",
        "parameters": [
          {
            "type": "string",
            "name": "authToken",
            "in": "header",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "object",
              "properties": {
                "Data": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/Correction"
                  }
                },
                "Error": {
                  "type": "object"
                },
                "Status": {
                  "type": "integer"
                }
              }
            }
          },
          "default": {
            "description": "Unsuccessful operation",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          }
        }
      },
      "post": {
        "summary": "Create Corrections",
        "operationId": "createCorrections",
        "parameters": [
          {
            "type": "string",
            "name": "authToken",
            "in": "header",
            "required": true
          },
          {
            "name": "corrections",
            "in": "body",
            "required": true,
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Correction"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          },
          "default": {
            "description": "Unsuccessful operation",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          }
        }
      }
    },
    "/health-check": {
      "get": {
        "summary": "Health Check",
        "operationId": "healthCheck",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/ratios": {
      "post": {
        "summary": "Create Carb Ratios",
        "operationId": "createRatios",
        "parameters": [
          {
            "type": "string",
            "name": "authToken",
            "in": "header",
            "required": true
          },
          {
            "name": "carbRatios",
            "in": "body",
            "required": true,
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/CarbRatio"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "object",
              "properties": {
                "Data": {
                  "type": "object"
                },
                "Error": {
                  "type": "object"
                },
                "Status": {
                  "description": "status code",
                  "type": "integer"
                }
              }
            }
          },
          "default": {
            "description": "Unsuccessful operation",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "CarbRatio": {
      "type": "object",
      "properties": {
        "End": {
          "type": "string"
        },
        "Key": {
          "type": "string"
        },
        "Ratio": {
          "type": "number",
          "format": "float64"
        },
        "Start": {
          "type": "string"
        }
      }
    },
    "Correction": {
      "type": "object",
      "properties": {
        "Bg": {
          "type": "number",
          "format": "float64"
        },
        "Bolus": {
          "type": "number",
          "format": "float64"
        },
        "Carbs": {
          "type": "number",
          "format": "float64"
        },
        "Key": {
          "type": "string"
        },
        "TimeStamp": {
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "CorrectionResponse": {
      "type": "object",
      "properties": {
        "ActiveInsulinReduction": {
          "type": "number",
          "format": "float64"
        },
        "BgCorrection": {
          "type": "number",
          "format": "float64"
        },
        "Bolus": {
          "type": "number",
          "format": "float64"
        },
        "CarbCorrection": {
          "type": "number",
          "format": "float64"
        }
      }
    },
    "ISF": {
      "type": "object",
      "properties": {
        "End": {
          "type": "string"
        },
        "Key": {
          "type": "string"
        },
        "Sensitivity": {
          "type": "number",
          "format": "float64"
        },
        "Start": {
          "type": "string"
        }
      }
    },
    "Response": {
      "type": "object",
      "properties": {
        "Error": {
          "type": "object"
        },
        "Status": {
          "type": "integer"
        }
      }
    },
    "Sign": {
      "type": "object",
      "properties": {
        "Password": {
          "type": "string"
        },
        "Username": {
          "type": "string"
        }
      }
    },
    "Target": {
      "type": "object",
      "properties": {
        "End": {
          "type": "string"
        },
        "Key": {
          "type": "string"
        },
        "Ratio": {
          "type": "number",
          "format": "float64"
        },
        "Start": {
          "type": "string"
        }
      }
    },
    "Token": {
      "type": "object",
      "properties": {
        "Token": {
          "type": "string"
        },
        "Username": {
          "type": "string"
        }
      }
    }
  }
}`))
}
