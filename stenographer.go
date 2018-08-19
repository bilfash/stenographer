package stenographer

const (
	STDOUT  = "STDOUT"
	FILE    = "FILE"
	TRACE   = "TRACE"
	DEBUG   = "DEBUG"
	INFO    = "INFO"
	WARNING = "WARNING"
	ERROR   = "ERROR"
)

var logLevel = map[string]int{
	"TRACE":   1,
	"DEBUG":   2,
	"INFO":    3,
	"WARNING": 4,
	"ERROR":   5,
}

type Stenographer struct {
	kafkaProducer *producer
	logLevel      string
	app           string
	project       string
	server        string
}

func NewStenographer(app, project, server string) *Stenographer {
	return &Stenographer{
		kafkaProducer: nil,
		logLevel:      INFO,
		app:           app,
		project:       project,
		server:        server,
	}
}

func (s *Stenographer) SetUpBroker(brokerAddress []string) {
	s.kafkaProducer = newProducer(brokerAddress)
}

func (s *Stenographer) SetUpLogLevel(level string) {
	s.logLevel = level
}
