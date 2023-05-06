package config

type MongoWaitlist struct {
	Host       string `env:"MONGO_WAITLIST_HOST" envDefault:"localhost"`
	Port       string `env:"MONGO_WAITLIST_PORT" envDefault:"27017"`
	Username   string `env:"MONGO_WAITLIST_USERNAME" envDefault:""`
	Password   string `env:"MONGO_WAITLIST_PASSWORD" envDefault:""`
	Database   string `env:"MONGO_WAITLIST_DATABASE" envDefault:"account"`
	Collection string `env:"MONGO_WAITLIST_COLLECTION" envDefault:"accounts"`
	Query      string `env:"MONGO_WAITLIST_QUERY" envDefault:""`
}

type I18n struct {
	Fallback string   `env:"I18N_FALLBACK_LANGUAGE" envDefault:"en"`
	Dir      string   `env:"I18N_DIR" envDefault:"./src/locales"`
	Locales  []string `env:"I18N_LOCALES" envDefault:"en,tr"`
}

type Server struct {
	Host  string `env:"SERVER_HOST" envDefault:"localhost"`
	Port  int    `env:"SERVER_PORT" envDefault:"3000"`
	Group string `env:"SERVER_GROUP" envDefault:"account"`
}

type HttpHeaders struct {
	AllowedOrigins   string `env:"CORS_ALLOWED_ORIGINS" envDefault:"*"`
	AllowedMethods   string `env:"CORS_ALLOWED_METHODS" envDefault:"GET,POST,PUT,DELETE,OPTIONS"`
	AllowedHeaders   string `env:"CORS_ALLOWED_HEADERS" envDefault:"*"`
	AllowCredentials bool   `env:"CORS_ALLOW_CREDENTIALS" envDefault:"true"`
	Domain           string `env:"HTTP_HEADER_DOMAIN" envDefault:"*"`
}

type Topics struct {
	Notify NotifyTopics
}

type NotifyTopics struct {
	SendMail     string `env:"STREAMING_TOPIC_NOTIFY_EMAIL"`
	SendTelegram string `env:"STREAMING_TOPIC_NOTIFY_TELEGRAM"`
}

type Nats struct {
	Url     string   `env:"NATS_URL" envDefault:"nats://localhost:4222"`
	Streams []string `env:"NATS_STREAMS" envDefault:""`
}

type App struct {
	Protocol string `env:"PROTOCOL" envDefault:"http"`
	DB       struct {
		Waitlist MongoWaitlist
	}
	Server      Server
	HttpHeaders HttpHeaders
	I18n        I18n
	Topics      Topics
	Nats        Nats
}
