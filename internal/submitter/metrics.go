package submitter

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	Insufficience *prometheus.GaugeVec
	Errors        prometheus.Counter
}

func NewMetrics(reg prometheus.Registerer) *Metrics {
	insuf := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "metis:srs:insufficience",
		Help: "Check if balance and allowrance of the mpc and payer is sufficient to pay",
	}, []string{"address", "type", "alias"})

	errors := prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "metis:srs:errors",
			Help: "Number of errors",
		},
	)

	reg.MustRegister(errors)
	reg.MustRegister(insuf)
	return &Metrics{Insufficience: insuf, Errors: errors}
}
