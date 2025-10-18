package controllers

import (
	"fmt"
	"net/http"
	"vuka-api/pkg/config"
	"vuka-api/pkg/httpx"
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/services"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type SourceController struct {
	sourceService *services.SourceService
	rssService    *services.RssService
}

func NewSourceController() *SourceController {
	serviceManager := services.NewServices(config.GetDB())
	return &SourceController{
		sourceService: serviceManager.Source,
		rssService:    serviceManager.Rss,
	}
}

func (sc *SourceController) GetAllSources(w http.ResponseWriter, r *http.Request) {
	sources, err := sc.sourceService.GetAllSources()
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	httpx.WriteJSON(w, http.StatusOK, sources)
}

func (sc *SourceController) GetSourceByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	source, err := sc.sourceService.GetSourceByID(vars["id"])
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusNotFound)
		return
	}
	httpx.WriteJSON(w, http.StatusOK, source)
}

func (sc *SourceController) CreateSource(w http.ResponseWriter, r *http.Request) {
	var source db.Source
	if err := httpx.ParseBody(r, &source); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := sc.sourceService.CreateSource(&source); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	httpx.WriteJSON(w, http.StatusCreated, source)
}

func (sc *SourceController) UpdateSource(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var source db.Source
	if err := httpx.ParseBody(r, &source); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Parse string ID to uuid.UUID
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		httpx.WriteErrorJSON(w, "Invalid source ID", http.StatusBadRequest)
		return
	}
	source.ID = id
	if err := sc.sourceService.UpdateSource(&source); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	httpx.WriteJSON(w, http.StatusOK, source)
}

func (sc *SourceController) DeleteSource(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if err := sc.sourceService.DeleteSource(vars["id"]); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (sc *SourceController) IngestSourceFeed(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println("Syncing source")
	sourceID, err := uuid.Parse(vars["id"])
	if err != nil {
		httpx.WriteErrorJSON(w, "Invalid source ID", http.StatusBadRequest)
		return
	}

	source, err := sc.sourceService.GetSourceByID(sourceID.String())
	if err != nil {
		httpx.WriteErrorJSON(w, "Source not found", http.StatusNotFound)
		return
	}

	if source.RssFeedUrl == "" {
		httpx.WriteErrorJSON(w, "Source does not have an RSS feed URL", http.StatusBadRequest)
		return
	}

	go func() {
		if err := sc.rssService.IngestRSSFeedWithSource(source.RssFeedUrl, &sourceID); err != nil {
			// Log the error, but don't write to the response as it's in a goroutine
			fmt.Printf("Error ingesting RSS feed for source %s: %v\n", sourceID, err)
		}
	}()

	httpx.WriteJSON(w, http.StatusAccepted, map[string]string{"message": "RSS feed ingestion started"})
}
