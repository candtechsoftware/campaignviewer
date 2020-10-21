package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"github.com/candtechsoftware/campaignapp/models"
)

// Initial Data To Test
var campaigns = []*models.Campaign{
	&models.Campaign{
		ID:        1,
		Name:      "Super Cool Campaign",
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	&models.Campaign{
		ID:        2,
		Name:      "Super Awesome Campaign",
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

// @route GET /campaign
// @return list of campaigns
// GetCampaign
func GetCampaign(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(campaigns)
}

// @route POST /campaign
// @body Campaign Object
// @return Single Campaign
// CreateCampaign creates a new campaign
func CreateCampaign(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var campaign models.Campaign
	_ = json.NewDecoder(r.Body).Decode(&campaign)
	newCampaign := models.Campaign{
		ID:        uint(len(campaigns) + 1), // Arbitray to keep Id Incremented but DB would normally handle
		Name:      campaign.Name,
		IsActive:  campaign.IsActive,
		CreatedAt: time.Now(), // DB would normally Handle this
		UpdatedAt: time.Now(), // DB would normally Handle this
	}

	log.Println(newCampaign)

	campaigns = append(campaigns, &newCampaign)
	json.NewEncoder(w).Encode(newCampaign)
}

// UpdateCampaign
// @route POST /campaign/{id}
// @params ID of the campaign
// @return Update Campaign
func UpdateCampaign(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, campaign := range campaigns {
		id, err := strconv.ParseUint(params["id"], 10, 64)
		if err != nil {
			log.Println(err)
		}

		if campaign.ID == uint(id) {
			campaigns = append(campaigns[:idx], campaigns[idx+1:]...)
			var newCampaign models.Campaign
			_ = json.NewDecoder(r.Body).Decode(&newCampaign)
			newCampaign.ID = uint(id)
			newCampaign.CreatedAt = campaign.CreatedAt
			newCampaign.UpdatedAt = time.Now()
			campaigns = append(campaigns, &newCampaign)
			json.NewEncoder(w).Encode(newCampaign)
			return
		}
	}
}

// DeleteCampaign
// @route DELETE /campaign/{id}
// @params ID Campaign
// @return
func DeleteCampaign(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicaiton/json")
	params := mux.Vars(r)
	for idx, campaign := range campaigns {
		id, err := strconv.ParseUint(params["id"], 10, 64)
		if err != nil {
			log.Println(err)
		}
		if campaign.ID == uint(id) {
			campaigns = append(campaigns[:idx], campaigns[idx+1:]...)
			json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
			return
		}
	}
	// If it is not found in the campaigns list
	json.NewEncoder(w).Encode(map[string]string{"status": "couldn't find campaign"})

}
