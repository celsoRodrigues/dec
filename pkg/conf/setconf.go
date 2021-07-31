package conf

import (
	"io"

	"github.com/spf13/viper"
)

//ViperConfig is our exported interface
type ViperConfig interface {
	SetConfigName(in string)
	SetConfigType(in string)
	Set(key, val string)
	SetConfigFile(in string)
	AddConfigPath(in string)
	ReadInConfig() error
	//ReadConfig(in io.Reader) error
	WriteConfig() error
	GetString(key string) string
	GetStringMapString(key string) map[string]string
	ConfigFileUsed() string
	IsSet(key string) bool
}

type viperConfig struct {
	MyConf                  *viper.Viper
	ConfigFileNotFoundError *viper.ConfigFileNotFoundError
}

func (c *viperConfig) SetConfigName(name string) {
	c.MyConf.SetConfigName(name)
}
func (c viperConfig) SetConfigType(ctype string) {
	c.MyConf.SetConfigType(ctype)
}
func (c viperConfig) SetConfigFile(in string) {
	c.MyConf.SetConfigFile(in)

}
func (c viperConfig) AddConfigPath(path string) {
	c.MyConf.AddConfigPath(path)
}
func (c viperConfig) ReadInConfig() error {
	err := c.MyConf.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
func (c viperConfig) ReadConfig(in io.Reader) error {
	err := c.MyConf.ReadConfig(in)
	if err != nil {
		return err
	}
	return nil
}
func (c viperConfig) WriteConfig() error {
	err := c.MyConf.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}
func (c viperConfig) Set(key, val string) {
	c.MyConf.Set(key, val)
}
func (c viperConfig) GetString(key string) string {
	val := c.MyConf.GetString(key)
	return val
}
func (c viperConfig) GetStringMapString(key string) map[string]string {
	val := c.MyConf.GetStringMapString(key)
	return val
}
func (c viperConfig) ConfigFileUsed() string {
	return c.MyConf.ConfigFileUsed()
}

func (c viperConfig) IsSet(key string) bool {
	return c.MyConf.IsSet(key)
}

//NewConfigWithOpt created a new configuration asset
func NewConfigWithOpt(configName, configType, configPath string) (ViperConfig, error) {

	confAsset := viperConfig{
		MyConf: viper.New(),
	}
	confAsset.SetConfigName(configName)
	confAsset.SetConfigType(configType)
	confAsset.AddConfigPath(configPath)

	return &confAsset, nil

}

//NewConfig created a new configuration asset
func NewConfig() (ViperConfig, error) {

	confAsset := viperConfig{
		MyConf: viper.New(),
	}

	return &confAsset, nil

}
