// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsAppautoscalingScheduledActionInvalidScalableDimensionRule checks the pattern is valid
type AwsAppautoscalingScheduledActionInvalidScalableDimensionRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsAppautoscalingScheduledActionInvalidScalableDimensionRule returns new rule with default attributes
func NewAwsAppautoscalingScheduledActionInvalidScalableDimensionRule() *AwsAppautoscalingScheduledActionInvalidScalableDimensionRule {
	return &AwsAppautoscalingScheduledActionInvalidScalableDimensionRule{
		resourceType:  "aws_appautoscaling_scheduled_action",
		attributeName: "scalable_dimension",
		enum: []string{
			"ecs:service:DesiredCount",
			"ec2:spot-fleet-request:TargetCapacity",
			"elasticmapreduce:instancegroup:InstanceCount",
			"appstream:fleet:DesiredCapacity",
			"dynamodb:table:ReadCapacityUnits",
			"dynamodb:table:WriteCapacityUnits",
			"dynamodb:index:ReadCapacityUnits",
			"dynamodb:index:WriteCapacityUnits",
			"rds:cluster:ReadReplicaCount",
			"sagemaker:variant:DesiredInstanceCount",
			"custom-resource:ResourceType:Property",
			"comprehend:document-classifier-endpoint:DesiredInferenceUnits",
			"lambda:function:ProvisionedConcurrency",
		},
	}
}

// Name returns the rule name
func (r *AwsAppautoscalingScheduledActionInvalidScalableDimensionRule) Name() string {
	return "aws_appautoscaling_scheduled_action_invalid_scalable_dimension"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppautoscalingScheduledActionInvalidScalableDimensionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppautoscalingScheduledActionInvalidScalableDimensionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppautoscalingScheduledActionInvalidScalableDimensionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppautoscalingScheduledActionInvalidScalableDimensionRule) Check(runner *tflint.Runner) error {
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
					`scalable_dimension is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
