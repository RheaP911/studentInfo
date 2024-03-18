package api

// import (
//     "net/http"

//     // Specify the username that you used inside github.com folder
//     "github.com/RheaP911/studentInfo/models"
//     "github.com/uadmin/uadmin"
// )

// // TodoListAPIHandler !
// func StudentListAPIHandler(w http.ResponseWriter, r *http.Request) {
//     // Fetch all records in the database
//     student := []models.Student{}
//     uadmin.All(&student)

//     // Accesses and fetches data from another model
//     for st := range student {
//         uadmin.Preload(&student[st])
//     }

//     // Return todo JSON object
//     uadmin.ReturnJSON(w, r, student)
// }