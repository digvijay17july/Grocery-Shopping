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

func GetProducts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var result = authenticate(w, r)
	if result!=true{
		respondError(w, http.StatusUnauthorized,"Invalid User or Token Expired")
		return
	}
	product:=[] model.Product{}
 db.Set("gorm:auto_preload", true).Find(&product)
 respondJSON(w,http.StatusOK,product)
}
func CreateProduct(db *gorm.DB,w http.ResponseWriter, r *http.Request){
	var result = authenticate(w, r)
	if result!=true{
		respondError(w, http.StatusUnauthorized,"Invalid User or Token Expired")
		return
	}
	product:= model.Product{}
	decoder := json.NewDecoder(r.Body)
	if err:=decoder.Decode(&product); err!=nil{
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()
		err :=db.Create(&product)
	if  err.Error!=nil{
		fmt.Println(err.Error.Error())
		respondError(w, http.StatusBadRequest, err.Error.Error())
		return
	}

	respondJSON(w, http.StatusCreated,product)
}
func GetProduct(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
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
	user := getProductOr404(db, uint(i), w, r)
	if user == nil {
		return
	}
	respondJSON(w, http.StatusOK, user)
}
func getProductOr404(db *gorm.DB, id uint, w http.ResponseWriter, r *http.Request) *model.Product {

	user := model.Product{}
	if err := db.Set("gorm:auto_preload", true).First(&user,id).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &user
}
