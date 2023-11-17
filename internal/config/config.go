package config

import (
	"time"

	"github.com/spf13/viper"
)

const (
	defaultHTTPPort     = "8080"
	EnvLocal 			= "local"
	Prod				= "prod"
)

type  (
	Config struct {
		App			AppConfig
		Environment	string
		MSql		MSqlConfig
		Http		HttpConfig
		Google		GoogleConfig
		Auth		AuthConfig
	}

	AppConfig struct {
		Name		string
	}

	MSqlConfig struct {
		User		string
		Password	string
		Host		string
		Name		string
		Charset		string
	}

	HttpConfig struct {
		Host               string        `mapstructure:"host"`
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
	}

	GoogleConfig struct {
		ClientId	string
		ClientSecret string
		RedirectUrl	string
		State string
	}

	AuthConfig struct {
		Secret	string
		//TokenExpiry time.Duration
		TokenExpiry int
		RefreshExpiry time.Duration
	}
)


func Init() (*Config, error) {
	var cfg Config
	
	if err := parse(); err != nil {
		return nil, err
	}

	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	setEnvironment(&cfg)

	return &cfg, nil
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("msql", &cfg.MSql); err != nil {
		return err
	}

	return nil
}

func setEnvironment(cfg *Config) {
	// App
	cfg.App.Name = "go-gc-community"
	
	// MySQL
	cfg.MSql.User = viper.GetString("msql.user")
	cfg.MSql.Password = viper.GetString("msql.password")
	cfg.MSql.Host = viper.GetString("msql.host")
	cfg.MSql.Name = viper.GetString("msql.name")
	cfg.MSql.Charset = viper.GetString("msql.charset")

	// Http
	cfg.Http.Port = defaultHTTPPort

	// Google Oauth
	cfg.Google.ClientId = viper.GetString("google.client_id")
	cfg.Google.ClientSecret = viper.GetString("google.client_secret")
	cfg.Google.RedirectUrl = viper.GetString("google.redirect")
	cfg.Google.State = viper.GetString("google.state")

	// JWT
	cfg.Auth.Secret = viper.GetString("auth.secret")
	//cfg.Auth.TokenExpiry = viper.GetDuration("auth.token_expiry")
	cfg.Auth.TokenExpiry = viper.GetInt("auth.token_expiry")
	cfg.Auth.RefreshExpiry = viper.GetDuration("auth.refresh_expiry")
}

func parse() error {
	viper.AddConfigPath("config/")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.MergeInConfig()
}