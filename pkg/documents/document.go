package documents

import "time"

type DocumentList struct {
	Documents []Document `yaml:"documents"`
}

type Document struct {
	Id         string    `yaml:"string"`
	Name       string    `yaml:"name"`
	Syntax     string    `yaml:"syntax"`
	IsPublic   bool      `yaml:"is_public"`
	PublicPath string    `yaml:"public_path"`
	Created    time.Time `yaml:"created"`
	Updated    time.Time `yaml:"updated"`
	Content    string    `yaml:"content,flow"`
}
