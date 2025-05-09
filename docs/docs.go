// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag/v2"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "components": {"schemas":{"main.sample":{"properties":{"id":{"type":"string"},"name":{"type":"string"},"title":{"type":"string"}},"type":"object"}}},
    "info": {"contact":{"email":"contact@krateoplatformops.io","name":"Krateo Support","url":"https://krateo.io"},"description":"{{escape .Description}}","license":{"name":"Apache 2.0","url":"http://www.apache.org/licenses/LICENSE-2.0.html"},"termsOfService":"http://swagger.io/terms/","title":"{{.Title}}","version":"{{.Version}}"},
    "externalDocs": {"description":"OpenAPI","url":"https://swagger.io/resources/open-api/"},
    "paths": {"/samples":{"get":{"description":"Get all Samples","operationId":"get-samples","responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/main.sample"}}},"description":"OK"}},"summary":"Get all Samples"},"post":{"description":"Add a new sample to the list","operationId":"post-sample","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/main.sample"}}},"description":"Sample to add","required":true},"responses":{"201":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/main.sample"}}},"description":"Created"},"400":{"content":{"application/json":{"schema":{"additionalProperties":{"type":"string"},"type":"object"}}},"description":"Bad Request"},"409":{"content":{"application/json":{"schema":{"additionalProperties":{"type":"string"},"type":"object"}}},"description":"Conflict"}},"summary":"Create a new Sample"}},"/samples/{name}":{"delete":{"description":"Delete a sample by its name","operationId":"delete-sample-by-name","parameters":[{"description":"Sample Name","in":"path","name":"name","required":true,"schema":{"type":"string"}}],"responses":{"200":{"content":{"application/json":{"schema":{"additionalProperties":{"type":"string"},"type":"object"}}},"description":"OK"},"404":{"content":{"application/json":{"schema":{"additionalProperties":{"type":"string"},"type":"object"}}},"description":"Not Found"}},"summary":"Delete a Sample by Name"},"get":{"description":"Get a single sample by its name","operationId":"get-sample-by-name","parameters":[{"description":"Sample Name","in":"path","name":"name","required":true,"schema":{"type":"string"}}],"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/main.sample"}}},"description":"OK"},"404":{"content":{"application/json":{"schema":{"additionalProperties":{"type":"string"},"type":"object"}}},"description":"Not Found"}},"summary":"Get a Sample by Name"},"put":{"description":"Update an existing sample","operationId":"put-sample-by-name","parameters":[{"description":"Sample Name","in":"path","name":"name","required":true,"schema":{"type":"string"}}],"requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/main.sample"}}},"description":"Updated Sample","required":true},"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/main.sample"}}},"description":"OK"},"400":{"content":{"application/json":{"schema":{"additionalProperties":{"type":"string"},"type":"object"}}},"description":"Bad Request"},"404":{"content":{"application/json":{"schema":{"additionalProperties":{"type":"string"},"type":"object"}}},"description":"Not Found"}},"summary":"Update a Sample by Name"}}},
    "openapi": "3.1.0",
    "servers": [
        {"url":"http://sample-web-server:8080"}
    ]
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Title:            "Sample Web Server API",
	Description:      "Sample Web Server API for managing samples.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
