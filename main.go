package main

import (
	"fmt"
	"net/url"
	"os"
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/iancoleman/strcase"
)

func main() {
	loader := openapi3.NewSwaggerLoader()
	url, err := url.Parse("https://static-assets.frame.io/devrel/swagger.json")
	check(err)
	swagger, err := loader.LoadSwaggerFromURI(url)
	check(err)
	path, err := os.Getwd()
	check(err)

	var hasTime, hasUUID bool
	for schemaname, schema := range swagger.Components.Schemas {
		hasTime = false
		hasUUID = false
		filename := fmt.Sprintf("%s/go-frame-io/structs/%s.go", path, schemaname)
		f, err := os.Create(filename)
		check(err)
		defer f.Close()

		// get propertynames and sort them
		propertyNames := make([]string, 0)

		for propertyname := range schema.Value.Properties {
			propertyNames = append(propertyNames, propertyname)
		}

		sort.Strings(propertyNames)

		var properties strings.Builder

		for _, propertyname := range propertyNames {
			property := schema.Value.Properties[propertyname]
			goType := GoType(property)
			if goType == "uuid.UUID" {
				hasUUID = true
			}
			if goType == "time.Time" {
				hasTime = true
			}
			fmt.Fprintf(&properties, "%s %+v\n", strcase.ToCamel(propertyname), goType)
		}

		fmt.Fprint(f, "package structs\n\n")
		if hasUUID || hasTime {
			fmt.Fprint(f, "import (")

			if hasTime {
				fmt.Fprintf(f, `
				"time"
				`)
			}
			if hasUUID {
				fmt.Fprintf(f, `
				"github.com/google/uuid"
				`)
			}
			fmt.Fprint(f, ")\n")

		}
		fmt.Fprintf(f, "type %s struct {\n", schemaname)
		fmt.Fprint(f, properties.String())
		fmt.Fprint(f, "}\n")

	}
}

func GoType(schema *openapi3.SchemaRef) string {
	if schema.Value != nil {
		switch schema.Value.Type {
		case "integer":
			return "int"
		case "string":
			return StringType(schema.Value)
		case "boolean":
			return "bool"
		case "number":
			return "float64"
		case "array":
			return "[]interface{}"
		case "object":
			return ObjectType(schema)
		case "":
			return "*interface{}"
		default:
			panic(schema.Ref)
		}
	} else {
		panic(schema.Ref)
	}
}

func ObjectType(schema *openapi3.SchemaRef) string {
	replaced := strings.ReplaceAll(schema.Ref, "#/components/schemas/", "")
	if replaced != "" {
		return fmt.Sprintf("*%s", replaced)
	} else {
		return "interface{}"
	}
}
func StringType(schema *openapi3.Schema) string {
	switch schema.Format {
	case "date-time":
		return "time.Time"
	case "uuid":
		return "uuid.UUID"
	case "":
		if len(schema.Enum) > 0 {
			vals := make([]string, 0)
			for _, e := range schema.Enum {
				vals = append(vals, e.(string))
			}
			concat := strings.Join(vals, ",")
			return fmt.Sprintf("*string //%s", concat)
		} else {
			return "*string"
		}
	default:
		fmt.Printf("DONT KNOW %s", schema.Format)
		return "*string"
	}

}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
