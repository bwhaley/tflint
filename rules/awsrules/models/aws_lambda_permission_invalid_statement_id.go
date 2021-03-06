// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsLambdaPermissionInvalidStatementIDRule checks the pattern is valid
type AwsLambdaPermissionInvalidStatementIDRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsLambdaPermissionInvalidStatementIDRule returns new rule with default attributes
func NewAwsLambdaPermissionInvalidStatementIDRule() *AwsLambdaPermissionInvalidStatementIDRule {
	return &AwsLambdaPermissionInvalidStatementIDRule{
		resourceType:  "aws_lambda_permission",
		attributeName: "statement_id",
		max:           100,
		min:           1,
		pattern:       regexp.MustCompile(`^([a-zA-Z0-9-_]+)$`),
	}
}

// Name returns the rule name
func (r *AwsLambdaPermissionInvalidStatementIDRule) Name() string {
	return "aws_lambda_permission_invalid_statement_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLambdaPermissionInvalidStatementIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLambdaPermissionInvalidStatementIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLambdaPermissionInvalidStatementIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLambdaPermissionInvalidStatementIDRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"statement_id must be 100 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"statement_id must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`statement_id does not match valid pattern ^([a-zA-Z0-9-_]+)$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
