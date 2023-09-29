package handler

import (
	"github.com/BruggerPedro/Gopportunities.git/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InicializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
