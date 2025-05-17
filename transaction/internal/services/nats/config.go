package nats

type NatsConfig struct {
	URL            string `env:"NATS_URL" envDefault:"nats://localhost:4222"`
	PublishStream  string `env:"NATS_PUB_STR" envDefault:"BANKING_STREAM"`
	PublishSubject string `env:"NATS_PUB_SUBJECT" envDefault:"BANKING.TRANSACTION"`
}
