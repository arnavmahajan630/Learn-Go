/* 
PURPOSE:
- Load env variables ONCE
- Validate them
- Expose strongly-typed config
*/

package config

import (
	"strings"

	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	DB_ADDR   string `koanf:"db_addr" validate:"required"`
	DB_PASS   string `koanf:"db_pass" validate:"required"`
	APP_PORT  string `koanf:"app_port" validate:"required"`
	DOMAIN    string `koanf:"domain" validate:"required"`
	API_QUOTA int    `koanf:"api_quota" validate:"gt=0"`
}

func Load() (*Config, error) {
	k := koanf.New(".")

	err := k.Load(
		env.Provider("APP_", ".", func(s string) string {
			return strings.ToLower(strings.TrimPrefix(s, "APP_"))
		}),
		nil,
	)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}

	if err := k.Unmarshal("", cfg); err != nil {
		return nil, err
	}

	validate := validator.New()
	if err := validate.Struct(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
