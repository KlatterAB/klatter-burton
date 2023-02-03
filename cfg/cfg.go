package cfg

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/ilyakaznacheev/cleanenv"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Name          string `yaml:"name"`
	Color         string `yaml:"color"`
	Notifications bool   `yaml:"notifications"`
	ID            string `yaml:"id"`
}

var Cfg Config

func InitConfig() {
	fpath := GetConfigPath()

	err := cleanenv.ReadConfig(fpath, &Cfg)
	if err != nil {
		fmt.Println("failed to read config file, creating config with default values")
		createDefaultConfig(fpath)
	}
}

func GetConfigPath() string {
	dir := os.Getenv("HOME") + "/.config/klatter-burton/"
	err := makeDirectoryIfNotExists(dir)
	if err != nil {
		fmt.Println(err)
	}

	fpath := path.Join(dir, "config.yml")

	return fpath
}

func ReloadConfig() {
	fpath := GetConfigPath()
	err := cleanenv.ReadConfig(fpath, &Cfg)
	if err != nil {
		log.Fatalln("Could not read config file.")
	}

}

func createDefaultConfig(path string) {
	Cfg = Config{
		Name:          os.Getenv("USER"),
		Color:         "#46D9FF",
		Notifications: true,
		ID:            "öö",
	}

	bytes, err := yaml.Marshal(Cfg)
	if err != nil {
		fmt.Println("failed to marshal default config values")
	}

	e := ioutil.WriteFile(path, bytes, 0644)
	if e != nil {
		fmt.Println("failed to create default config file")
		fmt.Println(e)
		panic(e)
	}
}

func makeDirectoryIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, os.ModeDir|0755)
	}
	return nil
}
