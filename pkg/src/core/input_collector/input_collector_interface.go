package input_collector

type InputCollectorInterface interface {
	Collect() ([]string, error)
}
