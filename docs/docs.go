// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "email": "murali.c@ensurity.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/addnftsale": {
            "post": {
                "description": "This API will put NFTs for sale",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "NFT"
                ],
                "summary": "Add NFTs",
                "parameters": [
                    {
                        "description": "NFT Detials",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/core.NFTSaleReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/commit-data-token": {
            "post": {
                "description": "This API will create data token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Data Tokens"
                ],
                "summary": "Create Data Token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "DID",
                        "name": "did",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Batch ID",
                        "name": "batchID",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/create-data-token": {
            "post": {
                "description": "This API will create data token",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "multipart/form-data"
                ],
                "tags": [
                    "Data Tokens"
                ],
                "summary": "Create Data Token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User/Entity ID",
                        "name": "UserID",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "User/Entity Info",
                        "name": "UserInfo",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Committer DID",
                        "name": "CommitterDID",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Batch ID",
                        "name": "BacthID",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "File Info is json string {",
                        "name": "FileInfo",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "File to be committed",
                        "name": "FileContent",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/createnft": {
            "post": {
                "description": "This API will create new NFT",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "multipart/form-data"
                ],
                "tags": [
                    "NFT"
                ],
                "summary": "Create NFT",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User/Entity Info",
                        "name": "UserInfo",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "File Info is json string {",
                        "name": "FileInfo",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "File to be committed",
                        "name": "FileContent",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/deploy-smart-contract": {
            "post": {
                "description": "This API will deploy smart contract Token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Smart Contract"
                ],
                "summary": "Deploy Smart Contract",
                "operationId": "deploy-smart-contract",
                "parameters": [
                    {
                        "description": "Deploy smart contract",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.DeploySmartContractSwaggoInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/execute-smart-contract": {
            "post": {
                "description": "This API will Execute smart contract Token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Smart Contract"
                ],
                "summary": "Execute Smart Contract",
                "parameters": [
                    {
                        "description": "Execute smart contrct and add details to chain",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.ExecuteSmartContractSwaggoInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/fetch-smart-contract": {
            "post": {
                "description": "This API will Fetch smart contract",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Smart Contract"
                ],
                "summary": "Fetch Smart Contract",
                "operationId": "fetch-smart-contract",
                "parameters": [
                    {
                        "description": "Fetch smart contract",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.FetchSmartContractSwaggoInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/generate-smart-contract": {
            "post": {
                "description": "This API will Generate smart contract Token",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "multipart/form-data"
                ],
                "tags": [
                    "Smart Contract"
                ],
                "summary": "Generate Smart Contract",
                "parameters": [
                    {
                        "type": "string",
                        "description": "DID",
                        "name": "did",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "location of binary code hash",
                        "name": "binaryCodePath",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "location of raw code hash",
                        "name": "rawCodePath",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "location of schema code hash",
                        "name": "schemaFilePath",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/get-account-info": {
            "get": {
                "description": "For a mentioned DID, check the account balance",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Check account balance",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User DID",
                        "name": "did",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/get-by-comment": {
            "get": {
                "description": "Retrieves the details of a transaction based on its comment.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Get transaction details by Transcation Comment",
                "operationId": "get-by-comment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Comment to identify the transaction",
                        "name": "Comment",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/get-by-did": {
            "get": {
                "description": "Retrieves the details of a transaction based on dID and date range.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Get transaction details by dID",
                "operationId": "get-by-did",
                "parameters": [
                    {
                        "type": "string",
                        "description": "DID of sender/receiver",
                        "name": "DID",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Filter by role as sender or receiver",
                        "name": "Role",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Start date of the date range (format: YYYY-MM-DD)",
                        "name": "StartDate",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "End date of the date range (format: YYYY-MM-DD)",
                        "name": "EndDate",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/get-by-txnId": {
            "get": {
                "description": "Retrieves the details of a transaction based on its ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Get transaction details by Transcation ID",
                "operationId": "get-txn-details-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The ID of the transaction to retrieve",
                        "name": "txnID",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/get-data-token": {
            "get": {
                "description": "This API will get all data token belong to the did",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Data Tokens"
                ],
                "summary": "Get Data Token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "DID",
                        "name": "did",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/get-smart-contract-token-chain-data": {
            "post": {
                "description": "This API will return smart contract token chain data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Smart Contract"
                ],
                "summary": "Get Smart Contract Token Chain Data",
                "operationId": "get-smart-contract-token-chain-data",
                "parameters": [
                    {
                        "description": "Returns Smart contract token chain Execution Data",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.GetSmartContractTokenChainDataSwaggoInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/getallnft": {
            "post": {
                "description": "This API will get all NFTs of the DID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "NFT"
                ],
                "summary": "Get ALL NFTs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.NFTTokens"
                        }
                    }
                }
            }
        },
        "/api/initiate-rbt-transfer": {
            "post": {
                "description": "This API will initiate RBT transfer to the specified dID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Initiate RBT Transfer",
                "operationId": "initiate-rbt-transfer",
                "parameters": [
                    {
                        "description": "Intitate RBT transfer",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.RBTTransferRequestSwaggoInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/register-callback-url": {
            "post": {
                "description": "This API will register call back url for when updated come for smart contract token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Smart Contract"
                ],
                "summary": "Get Smart Contract Token Chain Data",
                "operationId": "register-callback-url",
                "parameters": [
                    {
                        "description": "Register call back URL",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.RegisterCallBackURLSwaggoInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/signature-response": {
            "post": {
                "description": "This API is used to supply the password for the node along with the ID generated when Initiate RBT transfer is called.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Signature Response",
                "operationId": "signature-response",
                "parameters": [
                    {
                        "description": "Send input for requested signature",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.SignatureResponseSwaggoInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/start": {
            "get": {
                "description": "It will setup the core if not done before",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Basic"
                ],
                "summary": "Start Core",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/subscribe-smart-contract": {
            "post": {
                "description": "This API endpoint allows subscribing to a smart contract.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Smart Contract"
                ],
                "summary": "Subscribe to Smart Contract",
                "parameters": [
                    {
                        "description": "Subscribe to input contract",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.NewSubscriptionSwaggoInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "core.NFTSale": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "core.NFTSaleReq": {
            "type": "object",
            "properties": {
                "did": {
                    "type": "string"
                },
                "tokens": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/core.NFTSale"
                    }
                },
                "type": {
                    "type": "integer"
                }
            }
        },
        "model.BasicResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "result": {},
                "status": {
                    "type": "boolean"
                }
            }
        },
        "model.NFTStatus": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "token_status": {
                    "type": "integer"
                }
            }
        },
        "model.NFTTokens": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "result": {},
                "status": {
                    "type": "boolean"
                },
                "tokens": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.NFTStatus"
                    }
                }
            }
        },
        "server.DeploySmartContractSwaggoInput": {
            "type": "object",
            "properties": {
                "comment": {
                    "type": "string"
                },
                "deployerAddr": {
                    "type": "string"
                },
                "quorumType": {
                    "type": "integer"
                },
                "rbtAmount": {
                    "type": "number"
                },
                "smartContractToken": {
                    "type": "string"
                }
            }
        },
        "server.ExecuteSmartContractSwaggoInput": {
            "type": "object",
            "properties": {
                "comment": {
                    "type": "string"
                },
                "executorAddr": {
                    "type": "string"
                },
                "quorumType": {
                    "type": "integer"
                },
                "smartContractData": {
                    "type": "string"
                },
                "smartContractToken": {
                    "type": "string"
                }
            }
        },
        "server.FetchSmartContractSwaggoInput": {
            "type": "object",
            "properties": {
                "smartContractToken": {
                    "type": "string"
                }
            }
        },
        "server.GetSmartContractTokenChainDataSwaggoInput": {
            "type": "object",
            "properties": {
                "latest": {
                    "type": "boolean"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "server.NewSubscriptionSwaggoInput": {
            "type": "object",
            "properties": {
                "smartContractToken": {
                    "type": "string"
                }
            }
        },
        "server.RBTTransferRequestSwaggoInput": {
            "type": "object",
            "properties": {
                "comment": {
                    "type": "string"
                },
                "receiver": {
                    "type": "string"
                },
                "sender": {
                    "type": "string"
                },
                "tokenCOunt": {
                    "type": "number"
                },
                "type": {
                    "type": "integer"
                }
            }
        },
        "server.RegisterCallBackURLSwaggoInput": {
            "type": "object",
            "properties": {
                "callbackurl": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "server.SignatureResponseSwaggoInput": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "mode": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "SessionToken": {
            "type": "apiKey",
            "name": "Session-Token",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.9",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Rubix Core",
	Description:      "Rubix core API to control & manage the node.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
