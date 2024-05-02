package configurations

type Config struct {
	Database   Database
	ServerPort string `envconfig:"SERVER_PORT" default:"80"`
}

type Database struct {
	Host     string `envconfig:"DATABASE_HOST" required:"true"`
	Port     int    `envconfig:"DATABASE_PORT" required:"true"`
	User     string `envconfig:"DATABASE_USER" required:"true"`
	Password string `envconfig:"DATABASE_PASSWORD" required:"true"`
	Name     string `envconfig:"DB_NAME" required:"true"`
}
