package handlers

import (
	"gorm/db"
	"gorm/models"
	"net/http"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {

	users := models.Users{}
	db.Database.Find(&users)
	senData(rw, users, http.StatusOK)
}

func GetUser(rw http.ResponseWriter, r *http.Request) {

}

func CreateUser(rw http.ResponseWriter, r *http.Request) {

}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {

}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {

}

// func getUserByRequest(r *http.Request) (models.User, error) {
// 	//Obtener ID
// 	vars := mux.Vars(r)
// 	userId, _ := strconv.Atoi(vars["id"])

// 	if user, err := models.GetUser(userId); err != nil {
// 		return *user, err
// 	} else {
// 		return *user, nil
// 	}
// }
