package configs

import (
	"github.com/caarlos0/env"
)


var BootConfig = struct {
	DB_Host                     string `env:"DB_HOST" envDefault:"localhost"`
	DB_Port                     string `env:"DB_PORT" envDefault:"3306"`
	DB_User                     string `env:"DB_USERNAME"`
	DB_Password                 string `env:"DB_PASSWORD"`
	DB_Name		                string `env:"DB_NAME" envDefault:"providers"`
	DB_Type                 	string `env:"DB_TYPE" envDefault:"mysql"`
}{}

func init() {
	env.Parse(&BootConfig)
}