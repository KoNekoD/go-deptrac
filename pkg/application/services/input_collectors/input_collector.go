package input_collectors

type InputCollector interface {
	Collect() ([]string, error)
}
