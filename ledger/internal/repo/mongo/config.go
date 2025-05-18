package mongo

type MongoConfig struct {
	Host     string `env:"MONGO_HOST" envDefault:"localhost"`
	Port     int    `env:"MONGO_PORT" envDefault:"27017"`
	Username string `env:"MONGO_USERNAME" envDefault:"admin"`
	Password string `env:"MONGO_PASSWORD" envDefault:"admin"`
}
