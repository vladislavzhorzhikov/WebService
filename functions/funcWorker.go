package functions

import (
	"WebService/db"
	"WebService/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func HomePage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "Web service smartway workers")
}

func WorkerAdd(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	worker := models.Worker{}
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&worker)
	if err != nil {
		ReturnError(err, w)
		return
	}
	errDb := db.DataBase.Create(&worker).Error
	if errDb != nil {
		ReturnError(errDb, w)
		return
	}
	w.Write([]byte("Worker successfully added ID(" + strconv.Itoa(int(worker.ID)) + ")"))
}

func WorkerUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	editParams := models.Worker{}
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&editParams)
	if err != nil {
		ReturnError(err, w)
		return
	}
	var worker models.Worker
	idStr := ps.ByName("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ReturnError(err, w)
	}
	db.DataBase.Find(&worker, "workers.id = ?", id)
	if worker.ID == 0 {
		http.Error(w, "Worker not found", http.StatusNotFound)
		return
	}
	err = db.DataBase.Model(&worker).Updates(editParams).Error
	if err != nil {
		ReturnError(err, w)
	} else {
		w.Write([]byte("Worker successfully changed ID(" + strconv.Itoa(int(id)) + ")"))
	}
}

func WorkerRemove(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	worker := models.Worker{}
	idStr := ps.ByName("id")
	userID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ReturnError(err, w)
		return
	}

	db.DataBase.Delete(&worker, "workers.id = ?", userID)
	if worker.ID == 0 {
		http.Error(w, "Worker deleted or not found", http.StatusNotFound)
		return
	}
}

func FindFromCompany(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	worker := []models.Worker{}
	idStr := ps.ByName("id")
	companyID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ReturnError(err, w)
		return
	}
	db.DataBase.Joins("Department").Joins("Passport").Find(&worker, "workers.company_id = ?", companyID)
	if len(worker) == 0 {
		http.Error(w, "Workers not found", http.StatusNotFound)
		return
	}
	enc := json.NewEncoder(w)
	w.Header().Set("content-type", "application/json")
	enc.Encode(worker)
}

func FindFromDepartment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	department := models.Department{}
	worker := []models.Worker{}
	depName := ps.ByName("name")
	db.DataBase.Find(&department, "departments.name = ?", depName)
	db.DataBase.Joins("Department").Joins("Passport").Find(&worker, "workers.department_id = ?", department.ID)
	if len(worker) == 0 {
		http.Error(w, "Department is empty", http.StatusNotFound)
		return
	}
	if department.ID == 0 {
		http.Error(w, "Department not found", http.StatusNotFound)
		return
	}
	enc := json.NewEncoder(w)
	w.Header().Set("content-type", "application/json")
	enc.Encode(worker)
}

func ReturnError(err error, w http.ResponseWriter) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
