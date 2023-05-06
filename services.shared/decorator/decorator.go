package decorator

type Base struct {
	MetricsClient MetricsClient
	Logger        Logger
}

func NewBase() Base {
	return Base{
		MetricsClient: NewMetricsClient(),
		Logger:        NewLogger(),
	}
}
