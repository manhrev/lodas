package config

func GetConfig(env string) Config {
	if env == "deployment" {
		return &deployConfig{}
	}
	return &localConfig{}
}

type (
	Configuration struct {
		Database Database `mapstructure:"database"`
		Grpc     Grpc     `mapstructure:"grpc"`
		Cache    Cache    `mapstructure:"cache"`
	}

	Database struct {
		Username string `mapstructure:username"`
		Password string `mapstructure:"password"`
		Domain   string `mapstructure:"domain"`
		Port     string `mapstructure:"port"`
		Name     string `mapstructure:"name"`
	}

	Grpc struct {
		Host    string `mapstructure:"host"`
		Port    string `mapstructure:"port"`
		Network string `mapstructure:"network"`
	}

	Cache struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Password string `mapstructure:"password"`
		Db       int    `mapstructure:"db"`
	}
)

type Config interface {
	New() (*Configuration, error)
}
