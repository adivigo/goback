package vehiclecontroller

import (
	"net/http"

	"github.com/adivigo/goback/helper"
	"github.com/adivigo/goback/models"
)

var ResponseJson = helper.ResponseJson
var ResponseError = helper.ResponseError

func Index(w http.ResponseWriter, r *http.Request) {
	var vehicles []models.Vehicle

	if err := models.DB.Find(&vehicles).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJson(w, http.StatusOK, vehicles)
}

// func Show(w http.ResponseWriter, r *http.Request) {

// }

// func Create(w http.ResponseWriter, r *http.Request) {

// }

// func Update(w http.ResponseWriter, r *http.Request) {

// }

// func Delete(w http.ResponseWriter, r *http.Request) {

// }
