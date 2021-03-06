// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsKinesisStreamInvalidEncryptionTypeRule checks the pattern is valid
type AwsKinesisStreamInvalidEncryptionTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsKinesisStreamInvalidEncryptionTypeRule returns new rule with default attributes
func NewAwsKinesisStreamInvalidEncryptionTypeRule() *AwsKinesisStreamInvalidEncryptionTypeRule {
	return &AwsKinesisStreamInvalidEncryptionTypeRule{
		resourceType:  "aws_kinesis_stream",
		attributeName: "encryption_type",
		enum: []string{
			"NONE",
			"KMS",
		},
	}
}

// Name returns the rule name
func (r *AwsKinesisStreamInvalidEncryptionTypeRule) Name() string {
	return "aws_kinesis_stream_invalid_encryption_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsKinesisStreamInvalidEncryptionTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsKinesisStreamInvalidEncryptionTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsKinesisStreamInvalidEncryptionTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsKinesisStreamInvalidEncryptionTypeRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					`encryption_type is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
