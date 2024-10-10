package enums

type RuleTypeEnum string

const (
	TypeViolation        RuleTypeEnum = "violation"
	TypeSkippedViolation RuleTypeEnum = "skipped_violation"
	TypeUncovered        RuleTypeEnum = "uncovered"
	TypeAllowed          RuleTypeEnum = "allowed"
)
