package main

import (
	"WebService/db"
	"WebService/functions"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	db.MigrateDB()
	router := httprouter.New()
	router.GET("/", functions.HomePage)
	router.GET("/users/delete/:id", functions.WorkerRemove)
	router.GET("/users/findfromcompany/:id", functions.FindFromCompany)
	router.GET("/users/findfromdepartment/:name", functions.FindFromDepartment)
	router.POST("/users/", functions.WorkerAdd)
	router.PUT("/users/:id", functions.WorkerUpdate)
	router.POST("/departments/", functions.DepartmentAdd)
	router.PUT("/departments/:id", functions.DepartmentUpdate)
	router.GET("/departments/delete/:id", functions.DepartmentRemove)

	http.ListenAndServe(":8080", router)
}
