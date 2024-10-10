package dependencies

type DependencyList struct {
	dependencies        map[string][]DependencyInterface
	inheritDependencies map[string][]*InheritDependency
}

func NewDependencyList() *DependencyList {
	return &DependencyList{
		dependencies:        map[string][]DependencyInterface{},
		inheritDependencies: map[string][]*InheritDependency{},
	}
}

func (l *DependencyList) AddDependency(dep DependencyInterface) {
	tokenName := dep.GetDepender().ToString()
	if l.dependencies[tokenName] == nil {
		l.dependencies[tokenName] = make([]DependencyInterface, 0)
	}
	l.dependencies[tokenName] = append(l.dependencies[tokenName], dep)
}

func (l *DependencyList) AddInheritDependency(dep *InheritDependency) {
	tokenName := dep.GetDepender().ToString()
	l.inheritDependencies[tokenName] = append(l.inheritDependencies[dep.GetDepender().ToString()], dep)
}

func (l *DependencyList) GetDependenciesByClass(classLikeName string) []DependencyInterface {
	return l.dependencies[classLikeName]
}

func (l *DependencyList) GetDependenciesAndInheritDependencies() []DependencyInterface {
	buffer := make([]DependencyInterface, 0)
	for _, v := range l.dependencies {
		buffer = append(buffer, v...)
	}
	for _, v := range l.inheritDependencies {
		for _, inheritDependency := range v {
			buffer = append(buffer, inheritDependency)
		}
	}
	return buffer
}
