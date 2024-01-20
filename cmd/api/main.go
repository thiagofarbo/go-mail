package main

import (
	"go-mail/internal/domain/campaign"
	"go-mail/internal/endpoints"
	"go-mail/internal/infrastructure/db"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	handler := endpoints.Handler{}

	campaignService := campaign.Service{
		Repository: &db.CampaignRepository{},
	}

	handler := endpoints.Handler{
		CampaignService = campaignService
	}

	r.Post("/campaigns", endpoints.HandlerError(handler.CampaignPost))
	r.Get("/campaigns", endpoints.CampaignGet)

	print("Starting service on port 3000 ...")
	http.ListenAndServe(":3000", r)

}
