package models

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/uadmin/uadmin"
)

type Relation int

func (Relation) Mother() Relation {
	return 1
}
func (Relation) Father() Relation {
	return 2
}
func (Relation) Sibling() Relation {
	return 3
}
func (Relation) Relatives() Relation {
	return 4
}
func (Relation) Others() Relation {
	return 5
}

type Year int

func (Year) FirstYear() Year {
	return 1
}
func (Year) SecondYear() Year {
	return 2
}
func (Year) ThirdYear() Year {
	return 3
}
func (Year) FourthYear() Year {
	return 4
}
func (Year) FifthYear() Year {
	return 5
}

type Student struct {
	uadmin.Model
	SRCode   string    `uadmin:"read_only;display_name:Student Number;search"`
	Name     string    `uadmin:"help:(ex. Dela Cruz, Juan A.);search"`
	Address  string    `uadmin:"list_exclude;help:BLK/Building, Street, Barangay, Municipality, Province;search"`
	Birthday time.Time `uadmin:"list_exclude"`
	Contact  string    `uadmin:"display_name:Contact#;pattern:^[0-9]*$;pattern_msg:Your input must be a number."`
	Email    string
	Year     Year `uadmin:"required;display_name:Year Level;search"`

	Program   Program `uadmin:"required"`
	ProgramID uint

	School   School `uadmin:"required"`
	SchoolID uint
	Code     string `uadmin:"read_only;display_name:School Code;list_exclude"`

	Parent   string   `uadmin:"display_name: Parent/Guardian Name;;list_exclude"`
	Relation Relation `uadmin:"display_name: Relation to Student;list_exclude"`
	HomeNum  string   `uadmin:"display_name:Home Phone #;list_exclude;pattern:^[0-9]*$;pattern_msg:Your input must be a number."`

	Photo string `uadmin:"image;webcam;help:Upload a 1x1 photo"`
}

func (s Student) Validate() (errMsg map[string]string) {
	errMsg = map[string]string{}

	students := Student{}
	if uadmin.Count(&students, "name = ? AND id <> ?", s.Name, s.ID) != 0 {
		errMsg["Name"] = "This student is already registered in the system."
	}
	return
}

func (sNum *Student) Save() {
	students := Student{}
	studNum := sNum.SRCode

	uadmin.Preload(sNum)
	codeName := sNum.School.Code

	sNum.Code = codeName

	if studNum == "" {
		AYear := int(sNum.Year)
		currentYear := time.Now().Year() % 100
		yearString := strconv.Itoa(currentYear - AYear)

		baseCount := uadmin.Count(&students, "year = ? AND code = ?", sNum.Year, sNum.Code)
		recentCount := baseCount + 1

		alphabets := "ABCDEFGHJKLMNPQRSTUVWXYZ"
		randAlphString := ""
		for i := 0; i < 2; i++ {
			randomIndex := rand.Intn(len(alphabets))
			randAlphString += string(alphabets[randomIndex])
		}

		width := 5
		uniqueNum := fmt.Sprintf("%0*d", width, recentCount)

		sNum.SRCode = yearString + "-" + codeName + uniqueNum + randAlphString
	}
	uadmin.Save(sNum)
}
