package config

import "gopkg.in/ini.v1"

type Config struct {
	Server struct {
		AddressUp string `ini:"address_up"`
	} `ini:"server"`
	RedisServers struct {
		Address  string `ini:"address"`
		Password string `ini:"password"`
		DB       int    `ini:"db"`
	} `ini:"redis_subject"`
	RedisPIDs struct {
		Address  string `ini:"address"`
		Password string `ini:"password"`
		DB       int    `ini:"db"`
	} `ini:"redis_pids"`
	MongoSubject struct {
		Address  string `ini:"address"`
		Username string `ini:"username"`
		Password string `ini:"password"`
	} `ini:"mongo_subject"`
}

func ReadFromFile(fileName string) (Config, error) {
	cfg, err := ini.Load(fileName)
	config := Config{}
	if err != nil {
		return Config{}, err
	}
	err = cfg.MapTo(&config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
