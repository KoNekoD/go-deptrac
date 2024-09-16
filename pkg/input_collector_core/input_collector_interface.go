package input_collector_core

type InputCollectorInterface interface {
	Collect() ([]string, error)
}
