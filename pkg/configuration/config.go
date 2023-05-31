package configuration

import (
  "fmt"
  "github.com/spf13/viper"
)

type XmldConfig struct {
  Dsn                 string
}

func NewXmldConfig() XmldConfig {
  viper.AddConfigPath(".")
  viper.SetConfigName("xmlctl.conf")
  viper.SetConfigType("toml")
  err := viper.ReadInConfig()
  if err != nil {
    fmt.Println("error reading in config: ", err)
  }
  viper.SetDefault("database.dsn", "postgresql://bersen:bersen@localhost/xmldb")
  dsn := viper.GetString("database.dsn")

  return XmldConfig{
    Dsn: dsn,
  }
}
