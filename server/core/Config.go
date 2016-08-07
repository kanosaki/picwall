package core

type Config struct {
	Twitter *TwitterConfig          `yaml:"twitter"`
	Tumblr  *TumblrConfig           `yaml:"tumblr"`
	Sources map[string]SourceConfig `yaml:"sources"`
}

type TwitterConfig struct {
	ConsumerKey       string             `yaml:"consumer_key"`
	ConsumerSecret    string             `yaml:"consumer_secret"`
	DefaultAccount    *OAuthToken        `yaml:"default_account"`
	InitialThrottling *TwitterThrottling `yaml:"initial_throttling"`
}

type TwitterThrottling struct {
	Seconds    int64 `yaml:"duration_seconds"`
	BucketSize int64 `yaml:"bucket_size"`
}

type TumblrConfig struct {
	ConsumerKey    string      `yaml:"consumer_key"`
	ConsumerSecret string      `yaml:"consumer_secret"`
	DefaultAccount *OAuthToken `yaml:"default_account"`
}

type OAuthToken struct {
	Username          string `yaml:"username"`
	AccessToken       string `yaml:"access_token"`
	AccessTokenSecret string `yaml:"access_token_secret"`
}

type SourceConfig map[string]interface{}

func (sc *SourceConfig) Type() string {
	v, ok := map[string]interface{}(*sc)["type"]
	if !ok {
		panic("Invalid configuration: type of source not found")
	}
	s, ok := v.(string)
	if !ok {
		panic("Invalid configuration: type of source should be string")
	}
	return s
}

func (sc *SourceConfig) Get(key string) interface{} {
	v, ok := map[string]interface{}(*sc)[key]
	if !ok {
		panic("Invalid configuration: type of source not found")
	}
	return v
}
