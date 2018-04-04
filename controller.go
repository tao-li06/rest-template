package main

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
)

type Controller struct {
	DB *gorm.DB
}

func (this *Controller) Init() error {
	if this.DB == nil {
		var err error
		this.DB, err = gorm.Open(Config.DB.Type, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Config.DB.User, Config.DB.Passowrd, Config.DB.Address, Config.DB.Port, Config.DB.Name))
		if err != nil {
			return err
		}
	}
	this.DB.AutoMigrate(&ModelExample{})
	return nil
}

func (this *Controller) GetModel(w http.ResponseWriter, r *http.Request) {
	// Some code
}
