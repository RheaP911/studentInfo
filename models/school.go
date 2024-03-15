package models

import (
	"github.com/uadmin/uadmin"
)

type School struct {
	uadmin.Model
	Name string `uadmin:"required;search"`
	Code string `uadmin:"search;display_name:Code Name"`
	Logo string `uadmin:"required;image"`
	WebsiteLink string `uadmin:"list_exclude;display_name:Website Link"`
	Website string `uadmin:"link"`
}

func (sc School) Validate() (errMsg map[string]string) {
	errMsg = map[string]string{}
	schools := School{}
	if uadmin.Count(&schools, "name = ? AND id <> ?", sc.Name, sc.ID) != 0 {
		errMsg["Name"] = "This school is already registered in the system."
	}
	return
}

func (sw *School) Save() {
	link := sw.WebsiteLink
	sw.Website = link
	uadmin.Save(sw)
	
}
