package handler

import (
	"Grocery-Shopping-Category-Module/src/app/model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

func GetCategories(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var result = authenticate(w, r)
	if result!=true{
		respondError(w, http.StatusUnauthorized,"Invalid User or Token Expired")
		return
	}
	categories:=[] model.Categories{}
 db.Set("gorm:auto_preload", true).Find(&categories)
 respondJSON(w,http.StatusOK,categories)
}
func CreateCategory(db *gorm.DB,w http.ResponseWriter, r *http.Request){
	var result = authenticate(w, r)
	if result!=true{
		respondError(w, http.StatusUnauthorized,"Invalid User or Token Expired")
		return
	}
	category:= model.Categories{}
	decoder := json.NewDecoder(r.Body)
	if err:=decoder.Decode(&category); err!=nil{
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()
	err :=db.Create(&category)
	if  err.Error!=nil{
		fmt.Println(err.Error.Error())
		respondError(w, http.StatusBadRequest, err.Error.Error())
		return
	}
	respondJSON(w, http.StatusCreated,category)
}
func GetCategory(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var result = authenticate(w, r)
	if result!=true{
		respondError(w, http.StatusUnauthorized,"Invalid User or Token Expired")
		return
	}
	vars := mux.Vars(r)

	id := vars["id"]
	i, err := strconv.ParseUint(id, 10, 64)
	if err!=nil{
		fmt.Println(err.Error())
		respondError(w, http.StatusBadRequest, err.Error())
	}
	user := getCategoryOr404(db, uint(i), w, r)
	if user == nil {
		return
	}
	respondJSON(w, http.StatusOK, user)
}
func getCategoryOr404(db *gorm.DB, id uint, w http.ResponseWriter, r *http.Request) *model.Categories {
	user := model.Categories{}
	if err := db.Set("gorm:auto_preload", true).First(&user,id).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &user
}
