package controller

import (
	"log"
	"net/http"

	"app/model/feinstaub"
	"app/shared/response"
	"app/shared/router"
)

// Routes
func init() {
	router.Get("/accelerations", MessdatenAccelGET)
	router.Get("/particulatematter", ParticulateMatterGET)
}

const (
	ItemCreated      = "item created"
	ItemExists       = "item already exists"
	ItemNotFound     = "item not found"
	ItemFound        = "item found"
	ItemsFound       = "items found"
	ItemsFindEmpty   = "no items to find"
	ItemUpdated      = "item updated"
	ItemDeleted      = "item deleted"
	ItemsDeleted     = "items deleted"
	ItemsDeleteEmpty = "no items to delete"

	FriendlyError = "an error occurred, please try again later"
)

// *****************************************************************************
// Read
// *****************************************************************************

func MessdatenAccelGET(w http.ResponseWriter, r *http.Request) {
	// Get all items
	data, err := feinstaub.ReadAllAccel()
	if err != nil {
		log.Println(err)
		response.SendError(w, http.StatusInternalServerError, FriendlyError)
		return
	}

	// Send json data
	response.Send(w, http.StatusOK, ItemsFound, len(data), data)
}

func ParticulateMatterGET(w http.ResponseWriter, r *http.Request) {
	// Get all items
	data, err := feinstaub.ReadAllParticulate()
	if err != nil {
		log.Println(err)
		response.SendError(w, http.StatusInternalServerError, FriendlyError)
		return
	}

	// Send json data
	response.Send(w, http.StatusOK, ItemsFound, len(data), data)
}
