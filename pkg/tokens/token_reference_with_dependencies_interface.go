package tokens

type TokenReferenceWithDependenciesInterface interface {
	TokenReferenceInterface
	GetDependencies() []*DependencyToken
}
