package InputCollectorInterface

type InputCollectorInterface interface {
	Collect() ([]string, error)
}
