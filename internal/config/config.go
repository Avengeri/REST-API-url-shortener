package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local"` //Вытаскиваем из ямла значения и пихаем их в переменную через тег
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"10s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	User        string        `yaml:"user" env-required:"true"`
	Password    string        `yaml:"password" env-required:"true" env:"HTTP_SERVER_PASSWORD"`
}

func MustLoad() *Config { //функция которая будет загружать конфиг файл

	configPath := os.Getenv("CONFIG_PATH") //поиск файла с названием "" в окружении
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set") // если файл пустой, завершить программу и записать в лог

	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) { //проверяем наличие файла,проверяем является ли ошибка ошибкой. Выдает false если файл отсутствует
		log.Fatalf("config file does not exist: %s", configPath)
	}
	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil { //пытаемся прочитать конф. данные из файла и записать их в структуру config, ошибка, если чтение не удалось запишется в переменную err
		log.Fatalf("cannot read config: %s", err)
	}
	return &cfg
}
