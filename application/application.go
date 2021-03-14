package application

import (
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
	"rest-gorm/model"
	"gorm.io/driver/mysql"
	"github.com/spf13/viper"
	"rest-gorm/service"
)

type Application struct {
	DB     *gorm.DB
	Router *mux.Router
}

func NewContainer() (*Application, error) {
	dbUri := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		viper.Get("db.username"),
		viper.Get("db.password"),
		viper.Get("db.name"),
		viper.Get("db.charset"))
	log.Print("Try to connect " + dbUri)
	db, err := gorm.Open(mysql.Open(dbUri))
	if err != nil {
		log.Fatal("Problem connecting", err)
	}
	db = model.DBMigrate(db)
	router := mux.NewRouter()
	ap := &Application{db, router}
	ap.init()
	log.Print("Application initialized")
	return ap, err
}

func (ap *Application) init() {
	ap.GET("/employee/find", ap.getEmployee)
	ap.POST("/employee/create", ap.createEmployee)
	log.Fatal(http.ListenAndServe(":"+viper.GetString("port"), ap.Router))
}

func (ap *Application) GET(path string, f func(w http.ResponseWriter, r *http.Request)) {
	ap.Router.HandleFunc(path, f)
}
func (ap *Application) POST(path string, f func(w http.ResponseWriter, r *http.Request)) {
	ap.Router.HandleFunc(path, f)
}
func (ap *Application) PUT(path string, f func(w http.ResponseWriter, r *http.Request)) {
	ap.Router.HandleFunc(path, f)
}
func (ap *Application) DELETE(path string, f func(w http.ResponseWriter, r *http.Request)) {
	ap.Router.HandleFunc(path, f)
}

func (ap *Application) createEmployee(w http.ResponseWriter, r *http.Request) {
	service.CreateEmployee(ap.DB, w, r)
}
func (ap *Application) getEmployee(w http.ResponseWriter, r *http.Request) {
	service.GetEmployee(ap.DB, w, r)
}
