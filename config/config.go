package config

import "github.com/spf13/viper"

type App struct {
	Port string `json:"port"`
	Env  string `json:"env"`

	JwtSecretKey string `json:"jwt_secret_key"`
	JwtIssuer    string `json:"jwt_issuer"`
}

type PsqlDB struct {
	Host           string `json:"host"`
	Port           string `json:"port"`
	User           string `json:"user"`
	Password       string `json:"password"`
	DBName         string `json:"db_name"`
	DBMaxOpenConns int    `json:"db_max_open_conns"`
	DBMaxIdleConns int    `json:"db_max_idle_conns"`
}

type Config struct {
	App    App    `json:"app"`
	PsqlDB PsqlDB `json:"psql_db"`
}

func LoadConfig() (*Config, error) {
	return &Config{
		App: App{
			Port: viper.GetString("APP_PORT"),
			Env:  viper.GetString("APP_ENV"),

			JwtSecretKey: viper.GetString("JWT_SECRET_KEY"),
			JwtIssuer:    viper.GetString("JWT_ISSUER"),
		},
		PsqlDB: PsqlDB{
			Host:           viper.GetString("PSQL_DB_HOST"),
			Port:           viper.GetString("PSQL_DB_PORT"),
			User:           viper.GetString("PSQL_DB_USER"),
			Password:       viper.GetString("PSQL_DB_PASSWORD"),
			DBName:         viper.GetString("PSQL_DB_NAME"),
			DBMaxOpenConns: viper.GetInt("PSQL_DB_MAX_OPEN_CONNS"),
			DBMaxIdleConns: viper.GetInt("PSQL_DB_MAX_IDLE_CONNS"),
		},
	}, nil
}