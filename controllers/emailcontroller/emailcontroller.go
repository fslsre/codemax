package emailcontroller

import (
	"codemax/controllers/controller"
	"codemax/services/servicefactory"
	"net/http"
)

// Get get method
func Get(w http.ResponseWriter, r *http.Request) {
	factory, err := servicefactory.GetServices("emailjet", r)
	d, err := factory.FireEmail()
	if err != nil {
		controller.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	controller.ResponseWithJSON(w, http.StatusOK, d)
}

// func Post() {}

// func Put() {}

// func Delete() {}
