// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package magicmodules

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleHealthcareFhirStoreInvalidComplexDataTypeReferenceParsingRule checks the pattern is valid
type GoogleHealthcareFhirStoreInvalidComplexDataTypeReferenceParsingRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewGoogleHealthcareFhirStoreInvalidComplexDataTypeReferenceParsingRule returns new rule with default attributes
func NewGoogleHealthcareFhirStoreInvalidComplexDataTypeReferenceParsingRule() *GoogleHealthcareFhirStoreInvalidComplexDataTypeReferenceParsingRule {
	return &GoogleHealthcareFhirStoreInvalidComplexDataTypeReferenceParsingRule{
		resourceType:  "google_healthcare_fhir_store",
		attributeName: "complex_data_type_reference_parsing",
	}
}

// Name returns the rule name
func (r *GoogleHealthcareFhirStoreInvalidComplexDataTypeReferenceParsingRule) Name() string {
	return "google_healthcare_fhir_store_invalid_complex_data_type_reference_parsing"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleHealthcareFhirStoreInvalidComplexDataTypeReferenceParsingRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleHealthcareFhirStoreInvalidComplexDataTypeReferenceParsingRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleHealthcareFhirStoreInvalidComplexDataTypeReferenceParsingRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleHealthcareFhirStoreInvalidComplexDataTypeReferenceParsingRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{{Name: r.attributeName}},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		err := runner.EvaluateExpr(attribute.Expr, func(val string) error {
			validateFunc := validation.StringInSlice([]string{"COMPLEX_DATA_TYPE_REFERENCE_PARSING_UNSPECIFIED", "DISABLED", "ENABLED", ""}, false)

			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				if err := runner.EmitIssue(r, err.Error(), attribute.Expr.Range()); err != nil {
					return err
				}
			}
			return nil
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
