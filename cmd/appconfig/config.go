package appconfig

import "github.com/ilyakaznacheev/cleanenv"

type AppConfig struct {
	Hostname string `env:"HOSTNAME" env-default:"localhost:80"`
	FileName string `env:"FILENAME" env-default:"../../db.json"`
}

// Load environment variables to AppConfig instance
func LoadAppConfig() (*AppConfig, error) {
	cfg := &AppConfig{}
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
