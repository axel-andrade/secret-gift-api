package json_schemas

import (
	"embed"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

//go:embed schemas/*
var schemaFS embed.FS

func LoadJSONSchema(fileName string) (gojsonschema.JSONLoader, error) {
	schemaFile, err := schemaFS.ReadFile(fmt.Sprintf("schemas/%s_schema.json", fileName))
	if err != nil {
		return nil, fmt.Errorf("error reading JSON Schema file: %w", err)
	}

	schemaLoader := gojsonschema.NewStringLoader(string(schemaFile))

	if containsRef(schemaFile) {
		defsFile, err := schemaFS.ReadFile("schemas/defs.json")
		if err != nil {
			return nil, fmt.Errorf("error reading defs.json file: %w", err)
		}

		defsLoader := gojsonschema.NewStringLoader(string(defsFile))

		resolvedSchema, err := resolveRefs(schemaLoader, defsLoader)
		if err != nil {
			return nil, fmt.Errorf("error resolving refs: %w", err)
		}

		return resolvedSchema, nil
	}

	return schemaLoader, nil
}

func containsRef(schemaFile []byte) bool {
	var schema map[string]interface{}
	if err := json.Unmarshal(schemaFile, &schema); err != nil {
		return false
	}

	return hasRefInSection(schema, "body") || hasRefInSection(schema, "params") || hasRefInSection(schema, "query") || hasRefInSection(schema, "headers")
}

func hasRefInSection(schema map[string]interface{}, section string) bool {
	properties, exists := schema["properties"].(map[string]interface{})
	if !exists {
		return false
	}

	sectionData, exists := properties[section].(map[string]interface{})
	if !exists {
		return false
	}

	sectionProperties, exists := sectionData["properties"].(map[string]interface{})
	if !exists {
		return false
	}

	for _, value := range sectionProperties {
		if propertyMap, ok := value.(map[string]interface{}); ok {
			if _, exists := propertyMap["$ref"].(string); exists {
				return true
			}
		}
	}

	return false
}

func resolveRefs(schemaLoader gojsonschema.JSONLoader, defsLoader gojsonschema.JSONLoader) (gojsonschema.JSONLoader, error) {
	schemaJson, err := schemaLoader.LoadJSON()
	if err != nil {
		return nil, fmt.Errorf("error loading schema JSON: %w", err)
	}

	defsJson, err := defsLoader.LoadJSON()
	if err != nil {
		return nil, fmt.Errorf("error loading defs JSON: %w", err)
	}

	if err := replaceRefs(schemaJson.(map[string]interface{}), defsJson.(map[string]interface{})); err != nil {
		return nil, fmt.Errorf("error replacing refs: %w", err)
	}

	updatedSchema, err := json.Marshal(schemaJson)
	if err != nil {
		return nil, fmt.Errorf("error marshalling updated schema: %w", err)
	}

	return gojsonschema.NewStringLoader(string(updatedSchema)), nil
}

func replaceRefs(schema map[string]interface{}, defs map[string]interface{}) error {
	if err := replaceRefsInSection(schema, defs, "body"); err != nil {
		return err
	}

	if err := replaceRefsInSection(schema, defs, "params"); err != nil {
		return err
	}

	if err := replaceRefsInSection(schema, defs, "query"); err != nil {
		return err
	}

	if err := replaceRefsInSection(schema, defs, "headers"); err != nil {
		return err
	}

	return nil
}

func replaceRefsInSection(schema map[string]interface{}, defs map[string]interface{}, section string) error {
	properties, exists := schema["properties"].(map[string]interface{})
	if !exists {
		return nil
	}

	sectionData, exists := properties[section].(map[string]interface{})
	if !exists {
		return nil
	}

	sectionProperties, exists := sectionData["properties"].(map[string]interface{})
	if !exists {
		return nil
	}

	for key, value := range sectionProperties {
		if propertyMap, ok := value.(map[string]interface{}); ok {
			if ref, exists := propertyMap["$ref"].(string); exists {
				defKey := strings.TrimPrefix(ref, "defs#/definitions/")
				if def, defExists := defs["definitions"].(map[string]interface{})[defKey]; defExists {
					sectionProperties[key] = def
				} else {
					return fmt.Errorf("definition %s not found", defKey)
				}
			}
		}
	}

	return nil
}
