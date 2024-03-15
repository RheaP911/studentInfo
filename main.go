package main

import (
	"net/http"

	"github.com/RheaP911/studentInfo/api"
	"github.com/RheaP911/studentInfo/models"
	"github.com/uadmin/uadmin"
)

func main() {
	uadmin.Register (
		models.Student{},
		models.Program{},
		models.School{},
	)

	uadmin.RegisterInlines(models.Program{}, map[string]string{
		"Student": "ProgramID",
	})
	uadmin.RegisterInlines(models.School{}, map[string]string{
		"Student": "SchoolID",
	})

	http.HandleFunc("/api/", uadmin.Handler(api.Handler))
    uadmin.StartServer()

	InitializeRootURL()
}

func InitializeRootURL() {
	setting := uadmin.Setting{}
	uadmin.Get(&setting, "code = ?", "uAdmin.RootURL")
	setting.ParseFormValue([]string{"/admin/"})
	setting.Save()
}

