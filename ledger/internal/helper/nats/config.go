package nats

type NatsConfig struct {
	URL        string `env:"NATS_URL" envDefault:"nats://localhost:4222"`
	SubStream  string `env:"NATS_SUB_STR" envDefault:"BANKING_STREAM"`
	SubSubject string `env:"NATS_SUB_SUBJECT" envDefault:"BANKING.TRANSACTION"`
}
