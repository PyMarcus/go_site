package config

import (
	"log"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	InfoLog    *log.Logger
	Production bool
	Session    *scs.SessionManager
}
