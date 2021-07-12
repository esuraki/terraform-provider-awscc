// +build ignore

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"os"
	"path"
	"sort"
	"strings"
	"text/template"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/depgraph"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/provider/generators/resource/codegen"
	"github.com/iancoleman/strcase"
	"github.com/mitchellh/cli"
)

var (
	cfTypeSchemaFile = flag.String("cfschema", "", "CloudFormation resource type schema file; required")
	packageName      = flag.String("package", "", "override package name for generated code")
	tfResourceType   = flag.String("resource", "", "Terraform resource type; required")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "\tmain.go [flags] -resource <TF-resource-type> -cfschema <CF-type-schema-file> <generated-file>\n\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 || *tfResourceType == "" || *cfTypeSchemaFile == "" {
		flag.Usage()
		os.Exit(2)
	}

	destinationPackage := os.Getenv("GOPACKAGE")
	if *packageName != "" {
		destinationPackage = *packageName
	}

	filename := args[0]

	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	generator := NewGenerator(ui, *tfResourceType, *cfTypeSchemaFile)

	if err := generator.Generate(destinationPackage, filename); err != nil {
		ui.Error(fmt.Sprintf("error generating Terraform %s resource: %s", *tfResourceType, err))
		os.Exit(1)
	}
}

type Generator struct {
	cfTypeSchemaFile string
	tfResourceType   string
	ui               cli.Ui
}

func NewGenerator(ui cli.Ui, tfResourceType, cfTypeSchemaFile string) *Generator {
	return &Generator{
		cfTypeSchemaFile: cfTypeSchemaFile,
		tfResourceType:   tfResourceType,
		ui: &cli.BasicUi{
			Reader:      os.Stdin,
			Writer:      os.Stdout,
			ErrorWriter: os.Stderr,
		},
	}
}

func (g *Generator) Infof(format string, a ...interface{}) {
	g.ui.Info(fmt.Sprintf(format, a...))
}

// Generate generates the resource's type factory into the specified file.
func (g *Generator) Generate(packageName, filename string) error {
	g.Infof("generating Terraform resource code for %q from %q into %q", g.tfResourceType, g.cfTypeSchemaFile, filename)

	// Create target directory.
	dirname := path.Dir(filename)
	err := os.MkdirAll(dirname, 0755)

	if err != nil {
		return fmt.Errorf("error creating target directory %s: %w", dirname, err)
	}

	resource, err := NewResource(g.tfResourceType, g.cfTypeSchemaFile)

	if err != nil {
		return fmt.Errorf("error reading CloudFormation resource schema for %s: %w", g.tfResourceType, err)
	}

	// Generate code for the CloudFormation definitions.

	definitionNames, err := DefinitionNames(resource.CfResource.Definitions)

	if err != nil {
		return fmt.Errorf("error determining CloudFormation definition names: %w", err)
	}

	sb := strings.Builder{}
	schemaGenerator := codegen.Generator{
		CfResource: resource.CfResource,
		Writer:     &sb,
	}

	for _, definitionName := range definitionNames {
		definition, ok := resource.CfResource.Definitions[definitionName]

		if !ok {
			return fmt.Errorf("missing CloudFormation definition: %s", definitionName)
		}

		if propertyType := definition.Type.String(); propertyType != cfschema.PropertyTypeObject {
			return fmt.Errorf("CloudFormation definition (%s) is of unsupported type: %s", definitionName, propertyType)
		}

		if len(definition.Properties) == 0 {
			return fmt.Errorf("CloudFormation definition (%s) has no properties", definitionName)
		}

		if err := schemaGenerator.AppendCfDefinition(definitionName, definition.Properties); err != nil {
			return err
		}
	}

	subpropertyAttributes := sb.String()

	// Generate code for the CloudFormation root properties.

	sb.Reset()

	rootDefinitionName := "(Root)"

	schemaGenerator.AppendCfDefinition(rootDefinitionName, resource.CfResource.Properties)

	rootPropertyAttributes := sb.String()

	templateData := TemplateData{
		CloudFormationTypeName:             *resource.CfResource.TypeName,
		FactoryFunctionName:                strcase.ToLowerCamel(resource.TfType),
		PackageName:                        packageName,
		RootPropertyAttributes:             rootPropertyAttributes,
		RootPropertyAttributesVariableName: codegen.CfDefinitionTfAttributesVariableName(rootDefinitionName),
		SchemaVersion:                      1,
		SubpropertyAttributes:              subpropertyAttributes,
		TerraformTypeName:                  resource.TfType,
	}

	tmpl, err := template.New("function").Parse(templateBody)

	if err != nil {
		return fmt.Errorf("error parsing function template: %w", err)
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, templateData)

	if err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	generatedFileContents, err := format.Source(buffer.Bytes())

	if err != nil {
		g.Infof("%s", buffer.String())
		return fmt.Errorf("error formatting generated source code: %w", err)
	}

	f, err := os.Create(filename)

	if err != nil {
		return fmt.Errorf("error creating file (%s): %w", filename, err)
	}

	defer f.Close()

	_, err = f.Write(generatedFileContents)

	if err != nil {
		return fmt.Errorf("error writing to file (%s): %w", filename, err)
	}

	return nil
}

var templateBody = `
// Code generated by generators/resource/main.go; DO NOT EDIT.

package {{ .PackageName }}

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("{{ .TerraformTypeName }}", {{ .FactoryFunctionName }})
}

// {{ .FactoryFunctionName }} returns the Terraform {{ .TerraformTypeName }} resource type.
// This Terraform resource type corresponds to the CloudFormation {{ .CloudFormationTypeName }} resource type.
func {{ .FactoryFunctionName }}(ctx context.Context) (tfsdk.ResourceType, error) {
	// Subproperty definitions.
	{{ .SubpropertyAttributes }}

	// Root property definition.
	{{ .RootPropertyAttributes }}

	// Resource instance unique identifier.
	{{ .RootPropertyAttributesVariableName }}["identifier"] = schema.Attribute{
		Type:        types.StringType,
		Computed:    true,
		Description: "The resource instance's unique identifier.",
	}

	schema := schema.Schema{
		Version:    {{ .SchemaVersion }},
		Attributes: {{ .RootPropertyAttributesVariableName }},
	}

	resourceType := generic.NewResourceType(
		"{{ .CloudFormationTypeName }}",
		"{{ .TerraformTypeName }}",
		schema,
	)

	tflog.Debug(ctx, "Generated schema for %s:\n\n%v", "{{ .TerraformTypeName }}", schema)

	return resourceType, nil
}
`

type TemplateData struct {
	CloudFormationTypeName             string
	FactoryFunctionName                string
	PackageName                        string
	RootPropertyAttributes             string
	RootPropertyAttributesVariableName string
	SchemaVersion                      int64
	SubpropertyAttributes              string
	TerraformTypeName                  string
}

type Resource struct {
	CfResource *cfschema.Resource
	TfType     string
}

func NewResource(resourceType, cfTypeSchemaFile string) (*Resource, error) {
	resourceSchema, err := cfschema.NewResourceJsonSchemaPath(cfTypeSchemaFile)

	if err != nil {
		return nil, fmt.Errorf("error reading CloudFormation Resource Type Schema: %w", err)
	}

	resource, err := resourceSchema.Resource()

	if err != nil {
		return nil, fmt.Errorf("error parsing CloudFormation Resource Type Schema: %w", err)
	}

	return &Resource{
		CfResource: resource,
		TfType:     resourceType,
	}, nil
}

// DefinitionNames returns the CloudFormation definition names in the order that code should be generated.
func DefinitionNames(definitions map[string]*cfschema.Property) ([]string, error) {
	definitionNames := make([]string, 0)

	for definitionName := range definitions {
		definitionNames = append(definitionNames, definitionName)
	}

	// Sort the definition names to reduce generated code diffs.
	sort.Strings(definitionNames)

	dg := depgraph.New()

	for _, defName := range definitionNames {
		definition := definitions[defName]

		dg.AddNode(defName)

		if ref := definition.Ref; ref != nil {
			if refDefName := ref.Field(); refDefName != "" {
				dg.AddNode(refDefName)
				dg.AddDependency(defName, refDefName)
			}

			continue
		}

		switch definition.Type.String() {
		case cfschema.PropertyTypeArray:
			if items := definition.Items; items != nil {
				if ref := items.Ref; ref != nil {
					if refDefName := ref.Field(); refDefName != "" {
						dg.AddNode(refDefName)
						dg.AddDependency(defName, refDefName)
					}
				}
			}
		case cfschema.PropertyTypeObject:
			for _, prop := range definition.Properties {
				if ref := prop.Ref; ref != nil {
					if refDefName := ref.Field(); refDefName != "" {
						dg.AddNode(refDefName)
						dg.AddDependency(defName, refDefName)
					}

					continue
				}

				switch prop.Type.String() {
				case cfschema.PropertyTypeArray:
					if items := prop.Items; items != nil {
						if ref := items.Ref; ref != nil {
							if refDefName := ref.Field(); refDefName != "" {
								dg.AddNode(refDefName)
								dg.AddDependency(defName, refDefName)
							}
						}
					}
				}
			}
		}
	}

	return dg.OverallOrder()
}