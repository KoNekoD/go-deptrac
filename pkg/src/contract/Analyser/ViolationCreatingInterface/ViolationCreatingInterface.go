package ViolationCreatingInterface

// ViolationCreatingInterface - Every rule that can create a Violation has to implement this interface. It is used for output processing to display what rule has been violated.
type ViolationCreatingInterface interface {
	RuleName() string
	RuleDescription() string
}
