package service

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"rest-gorm/model"
)

func CreateEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	employee := model.Employee{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&employee); err != nil {
		toError(w, http.StatusBadRequest, err.Error())
	}
	defer r.Body.Close()
	if err := db.Save(&employee).Error; err != nil {
		toError(w, http.StatusInternalServerError, err.Error())
	}
	toJson(w, http.StatusOK, employee)
}

func GetEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	employee := getEmployee(db, vars["name"], w)
	if employee == nil {
		return
	}
	toJson(w, http.StatusOK, employee)
}

func getEmployee(db *gorm.DB, name string, w http.ResponseWriter) *model.Employee {
	employee := model.Employee{}

	if err := db.First(&employee, model.Employee{Name: name}).Error; err != nil {
		toError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &employee
}
func toJson(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

func toError(w http.ResponseWriter, status int, message string) {
	toJson(w, status, map[string]string{"error": message})
}
