package config

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

type config struct {
	Server struct {
		Address string
	}
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
	}
}

var C config

func ReadConfig() {
	Config := &C

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(rootDir())

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	spew.Dump(C)
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(filepath.Dir(d))
}
