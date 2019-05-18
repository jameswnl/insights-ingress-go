package config

import (
	"github.com/spf13/viper"
)

// IngressConfig represents the runtime configuration
type IngressConfig struct {
	MaxSize              int
	StageBucket          string
	RejectBucket         string
	Auth                 bool
	KafkaBrokers         []string
	KafkaGroupID         string
	KafkaAvailableTopic  string
	KafkaValidationTopic string
}

// Get returns an initialized IngressConfig
func Get() *IngressConfig {
	options := viper.New()
	options.SetDefault("MaxSize", 10*1024*1024)
	options.SetDefault("StageBucket", "available")
	options.SetDefault("RejectBucket", "rejected")
	options.SetDefault("Auth", true)
	options.SetDefault("KafkaBrokers", []string{"kafka:29092"})
	options.SetDefault("KafkaGroupID", "ingress")
	options.SetDefault("KafkaAvailableTopic", "platform.upload.available")
	options.SetDefault("KafkaValidationTopic", "platform.upload.validation")
	options.SetEnvPrefix("INGRESS")
	options.AutomaticEnv()

	return &IngressConfig{
		MaxSize:              options.GetInt("MaxSize"),
		StageBucket:          options.GetString("StageBucket"),
		RejectBucket:         options.GetString("RejectBucket"),
		Auth:                 options.GetBool("Auth"),
		KafkaBrokers:         options.GetStringSlice("KafkaBrokers"),
		KafkaGroupID:         options.GetString("KafkaGroupID"),
		KafkaAvailableTopic:  options.GetString("KafkaAvailableTopic"),
		KafkaValidationTopic: options.GetString("KafkaValidationTopic"),
	}
}
