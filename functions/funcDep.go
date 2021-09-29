package functions

import (
	"WebService/db"
	"WebService/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func DepartmentAdd(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	department := models.Department{}
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&department)
	if err != nil {
		ReturnError(err, w)
		return
	}
	errDb := db.DataBase.Create(&department).Error
	if errDb != nil {
		ReturnError(errDb, w)
		return
	}
	w.Write([]byte("Department successfully added ID(" + strconv.Itoa(int(department.ID)) + ")"))
}

func DepartmentUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	editParams := models.Department{}
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&editParams)
	if err != nil {
		ReturnError(err, w)
		return
	}
	var department models.Department
	idStr := ps.ByName("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ReturnError(err, w)
	}
	db.DataBase.Find(&department, "departments.id = ?", id)
	if department.ID == 0 {
		http.Error(w, "Department not found", http.StatusNotFound)
		return
	}
	err = db.DataBase.Model(&department).Updates(editParams).Error
	if err != nil {
		ReturnError(err, w)
	} else {
		w.Write([]byte("Department successfully changed ID(" + strconv.Itoa(int(id)) + ")"))
	}
}

func DepartmentRemove(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	department := models.Department{}
	idStr := ps.ByName("id")
	userID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ReturnError(err, w)
		return
	}

	db.DataBase.Delete(&department, "departments.id = ?", userID)
	if department.ID == 0 {
		http.Error(w, "Department deleted or not found", http.StatusNotFound)
		return
	}
}
