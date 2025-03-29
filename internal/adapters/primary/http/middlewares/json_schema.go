package middlewares

import (
	"net/http"
	"strings"

	"github.com/axel-andrade/secret-gift-api/internal/adapters/primary/http/json_schemas"
	common_ptr "github.com/axel-andrade/secret-gift-api/internal/adapters/primary/http/presenters/common"
	"github.com/gin-gonic/gin"
	"github.com/xeipuuv/gojsonschema"
)

const (
	RequestBody    = "body"
	RequestHeaders = "headers"
	RequestQuery   = "query"
	RequestParams  = "params"
)

func errorsToStructs(errors []gojsonschema.ResultError) []common_ptr.ValidateDetail {
	var validationErrors []common_ptr.ValidateDetail

	for _, err := range errors {
		validationErrors = append(validationErrors, common_ptr.ValidateDetail{
			Namespace: err.Field(),
			Tag:       err.Type(),
			Param:     err.Description(),
		})
	}

	return validationErrors
}

func ValidateRequest(schemaPath string) gin.HandlerFunc {
	schemaLoader, err := json_schemas.LoadJSONSchema(schemaPath)

	if err != nil {
		panic(err)
	}

	jsonSchemaPtr := common_ptr.JsonSchemaPresenter{}

	return func(c *gin.Context) {
		var input = make(map[string]any)
		var validationResult *gojsonschema.Result

		schemaJson, err := schemaLoader.LoadJSON()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, jsonSchemaPtr.Format(nil))
			return
		}

		properties := schemaJson.(map[string]any)["properties"].(map[string]any)
		for paramName := range properties {
			if paramName == RequestBody {
				var bodyMap = make(map[string]any)
				if err := c.BindJSON(&bodyMap); err != nil {
					c.AbortWithStatusJSON(http.StatusBadRequest, jsonSchemaPtr.Format(nil))
					return
				}
				input[RequestBody] = bodyMap
			} else if paramName == RequestHeaders {
				var headersMap = make(map[string][]string)
				for key, values := range c.Request.Header {
					headersMap[strings.ToLower(key)] = values
				}
				input[RequestHeaders] = headersMap
			} else if paramName == RequestQuery {
				var queryMap = make(map[string]string)
				for key := range c.Request.URL.Query() {
					queryMap[key] = c.Query(key)
				}
				input[RequestQuery] = queryMap
			} else if paramName == RequestParams {
				var paramsMap = make(map[string]string)
				for _, param := range c.Params {
					paramsMap[param.Key] = param.Value
				}
				input[RequestParams] = paramsMap
			}
		}

		validationResult, err = gojsonschema.Validate(
			schemaLoader,
			gojsonschema.NewGoLoader(input),
		)

		if err != nil || !validationResult.Valid() {
			c.AbortWithStatusJSON(http.StatusBadRequest, jsonSchemaPtr.Format(errorsToStructs(validationResult.Errors())))
			return
		}

		if value, ok := input[RequestBody]; ok {
			c.Set(RequestBody, value)
		}

		c.Next()
	}
}
