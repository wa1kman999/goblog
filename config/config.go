package config

// Config 全局配置
type Config struct {
	System *System `mapstructure:"system" json:"system" yaml:"system"`
	JWT    *JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis  *Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	Mysql  *Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}
