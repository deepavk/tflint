// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsWorklinkFleetInvalidDisplayNameRule checks the pattern is valid
type AwsWorklinkFleetInvalidDisplayNameRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsWorklinkFleetInvalidDisplayNameRule returns new rule with default attributes
func NewAwsWorklinkFleetInvalidDisplayNameRule() *AwsWorklinkFleetInvalidDisplayNameRule {
	return &AwsWorklinkFleetInvalidDisplayNameRule{
		resourceType:  "aws_worklink_fleet",
		attributeName: "display_name",
		max:           100,
	}
}

// Name returns the rule name
func (r *AwsWorklinkFleetInvalidDisplayNameRule) Name() string {
	return "aws_worklink_fleet_invalid_display_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsWorklinkFleetInvalidDisplayNameRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsWorklinkFleetInvalidDisplayNameRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsWorklinkFleetInvalidDisplayNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsWorklinkFleetInvalidDisplayNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"display_name must be 100 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}