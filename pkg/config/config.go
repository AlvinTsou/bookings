package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	UseCache       bool
	TemplatesCache map[string]*template.Template
	InfoLog        *log.Logger
	InProduction   bool // set to true when in production
	Session        *scs.SessionManager
}
