package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/deepakjacob/digester/domain"
	"github.com/deepakjacob/digester/service"
)

// RegistrationHandler interface
type RegistrationHandler struct {
	RegistrationService service.RegistrationService
}

// Registration handler for registration
func (rh *RegistrationHandler) Registration(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fileName := r.Form["file_name"]
	fileDate := time.Now()
	towerID := r.Form["tower_id"]
	locationID := r.Form["location_id"]
	postalCode := r.Form["postal_code"]
	areaCode := r.Form["area_code"]

	registration := &domain.Registration{
		FileName:   fileName[0],
		FileDate:   fileDate,
		TowerID:    towerID[0],
		LocationID: locationID[0],
		PostalCode: postalCode[0],
		AreaCode:   areaCode[0],
	}

	rStatus, err := rh.RegistrationService.RegisterFile(r.Context(), registration)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	rjson, err := json.Marshal(rStatus)
	if err != nil {
		http.Error(w,
			http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(rjson)
}
