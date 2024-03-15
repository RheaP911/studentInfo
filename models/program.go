package models

import "github.com/uadmin/uadmin"

type Program struct {
	uadmin.Model
	Name string `uadmin:"required;search"`
	Program string `uadmin:"required;search;display_name:Code Name"`
	Major string 
	Years uint `uadmin:"default_value: ;required;pattern:^[1-9]*$;pattern_msg:Your input must be a number."`
}

func (p Program) Validate() (errMsg map[string]string) {
	errMsg = map[string]string{}

	programs := Program{}
	if uadmin.Count(&programs, "name = ? AND id <> ? AND major = ?", p.Name, p.ID, p.Major) != 0 {
		errMsg["Name"] = "This program is already registered in the system."
	}
	return
}
