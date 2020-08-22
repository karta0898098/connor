package template

// Config template ...
const Config = `package configs

import (
	"bytes"
	"github.com/karta0898098/kara/db"
	"github.com/karta0898098/kara/grpc"
	"github.com/karta0898098/kara/http"
	"github.com/karta0898098/kara/zlog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"io/ioutil"
)

// Module for export configs to fx injection
var Module = fx.Provide(
	NewConfig,
)

// Path to read config path
var Path = ""

// Config all setting to create instance
type Config struct {
	fx.Out
	Log      *zlog.Config
	Database *db.Config
	HTTP     *http.Config
	GRPC     *grpc.Config
}

// NewConfig read config and create new instance 
func NewConfig() Config {

	//set file type toml or yaml
	viper.AutomaticEnv()
	viper.SetConfigType("{{.ConfigType}}")
	var config Config
	
	//check user has input config path 
	//if config path is exist then use that config
	if Path != "" {
		buf, err := ioutil.ReadFile(Path)

		if err != nil {
			log.Error().Msgf("Error reading config file, %s", err)
		}

		if err := viper.ReadConfig(bytes.NewBuffer(buf)); err != nil {
			log.Error().Msgf("Error reading config file, %s", err)
		}

		err = viper.Unmarshal(&config)
		if err != nil {
			log.Error().Msgf("unable to decode into struct, %v", err)
			return config
		}
		return config
	}

	// user does't input any config
	// then check env export config path
	path := viper.GetString("CONFIG_PATH")
	if path == "" {
		path = "./deployments/config"
	}
	
	//check user want setting other config
	name := viper.GetString("CONFIG_NAME")
	if name == "" {
		name = "app"
	}

	viper.SetConfigName(name)
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		log.Error().Msgf("Error reading config file, %s", err)
		return config
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Error().Msgf("unable to decode into struct, %v", err)
		return config
	}
	return config
}
`
