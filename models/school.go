package models

import (
	"github.com/uadmin/uadmin"
)

type School struct {
	uadmin.Model
	Name string `uadmin:"required;search"`
	Code string `uadmin:"search;display_name:Code Name"`
	Logo string `uadmin:"required;image"`
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
	code := sw.Code
	switch code {
	case "ISM":
		sw.Website = "https://www.ismanila.org/"
	case "UST":
		sw.Website = "https://www.ust.edu.ph/"
	case "Ateneo":
		sw.Website = "https://www.ateneo.edu/"
	case "DLSU":
		sw.Website = "https://www.dlsu.edu.ph/"
	case "UB":
		sw.Website = "https://ub.edu.ph/"
	case "PUP":
		sw.Website = "https://www.pup.edu.ph/"
	case "LPU":
		sw.Website = "https://www.lpu.edu.ph/"
	default:
		sw.Website = "https://batstateu.edu.ph/"
	}
	uadmin.Save(sw)
	
}
