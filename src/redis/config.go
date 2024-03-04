package redis

type Config struct {
	Port            int    `yaml:"port"`
	Url             string `yaml:"url"`
	ExpiryInMinutes int    `yaml:"expiryInMinutes"`
}
