package options

import "time"

type CrawOptions struct {
	//Url     string        `yaml:"url" json:"url"`
	TimeOut time.Duration `yaml:"timeOut" json:"timeOut"`
}
