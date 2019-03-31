package handlers

import (
	"net/http"

	cache "github.com/patrickmn/go-cache"
)

func (handler *AWSHandler) CostAndUsageHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("cost_usage_history")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeCostAndUsage(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "ce:GetCostAndUsage is missing")
		} else {
			handler.cache.Set("cost_usage_history", response.History, cache.DefaultExpiration)
			respondWithJSON(w, 200, response.History)
		}
	}
}

func (handler *AWSHandler) CurrentCostHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("cost_usage_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeCostAndUsage(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "ce:GetCostAndUsage is missing")
		} else {
			handler.cache.Set("cost_usage_total", response.Total, cache.DefaultExpiration)
			respondWithJSON(w, 200, response.Total)
		}
	}
}
