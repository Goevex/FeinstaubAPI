package controller

import (
	"log"
	"net/http"
	"regexp"

	"app/model/feinstaub"
	"app/shared/response"
	"app/shared/router"
)

// Routes
func init() {
	router.Get("/measurements", MeasurementsGET)
}

const (
	itemsNotFound = "items not found"
	itemsFound    = "items found"

	friendlyError = "an error occurred, please try again later"
)

// *****************************************************************************
// Read
// *****************************************************************************

// MeasurementsGET sends measurements json response
func MeasurementsGET(w http.ResponseWriter, r *http.Request) {
	// Get condition GET parameter
	condition := r.URL.Query().Get("condition")

	// Validate condition
	reg, errReg := regexp.Compile("^(\\w+[<>=]+\\w+(\\s(and|or)\\s?)?)*$")
	if errReg != nil {
		log.Println(errReg)
		response.SendError(w, http.StatusInternalServerError, friendlyError)
		return
	}
	if !reg.MatchString(condition) {
		log.Println("Wrong condition statement: " + condition)
		response.SendError(w, http.StatusInternalServerError, itemsNotFound)
		return
	}

	// Get all items
	data, err := feinstaub.ReadMeasurements(condition)
	if err != nil {
		log.Println(err)
		response.SendError(w, http.StatusInternalServerError, friendlyError)
		return
	}

	// Send json data
	if len(data) == 0 {
		response.Send(w, http.StatusOK, itemsNotFound, 0, nil)
	} else {
		response.Send(w, http.StatusOK, itemsFound, len(data), data)
	}
}
