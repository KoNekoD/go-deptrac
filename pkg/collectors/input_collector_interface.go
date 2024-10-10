package collectors

type InputCollectorInterface interface {
	Collect() ([]string, error)
}
