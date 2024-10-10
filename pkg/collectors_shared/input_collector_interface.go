package collectors_shared

type InputCollectorInterface interface {
	Collect() ([]string, error)
}
