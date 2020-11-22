package user

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "credential"
}

type handler struct {
	db *gorm.DB
}

func New(db *gorm.DB) http.Handler {
	return &handler{db: db}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var user User

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {

	}

	if err := h.db.Create(&user).Error; err != nil {

	}

}

func NewHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User

		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {

		}

		if err := db.Create(&user).Error; err != nil {

		}

	}
}
