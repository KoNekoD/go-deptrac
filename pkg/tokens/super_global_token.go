package tokens

type SuperGlobalToken string

const (
	GLOBALS SuperGlobalToken = "GLOBALS"
	SERVER  SuperGlobalToken = "_SERVER"
	GET     SuperGlobalToken = "_GET"
	POST    SuperGlobalToken = "_POST"
	FILES   SuperGlobalToken = "_FILES"
	COOKIE  SuperGlobalToken = "_COOKIE"
	SESSION SuperGlobalToken = "_SESSION"
	REQUEST SuperGlobalToken = "_REQUEST"
	ENV     SuperGlobalToken = "_ENV"
)

func NewSuperGlobalToken(superglobalName string) SuperGlobalToken {
	return SuperGlobalToken(superglobalName)
}

func (s SuperGlobalToken) AllowedNames() []string {
	return []string{
		string(GLOBALS),
		string(SERVER),
		string(GET),
		string(POST),
		string(FILES),
		string(COOKIE),
		string(SESSION),
		string(REQUEST),
		string(ENV),
	}
}

func (s SuperGlobalToken) ToString() string {
	return "$" + string(s)
}
