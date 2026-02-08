package handlers

import (
	"encoding/json"
	"kasir-api/services"
	"net/http"
)

type ReportHandler struct {
	service *services.ReportService
}

func NewReportHandler(service *services.ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

func (handler *ReportHandler) HandleReport(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handler.GetReport(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (handler *ReportHandler) GetReport(w http.ResponseWriter, r *http.Request) {
	start_date := r.URL.Query().Get("start_date")
	end_date := r.URL.Query().Get("end_date")
	report, err := handler.service.GetReport(start_date, end_date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "Report retrieved successfully",
		"data":    report,
	})
}

func (handler *ReportHandler) HandleReportToday(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handler.GetReportToday(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (handler *ReportHandler) GetReportToday(w http.ResponseWriter, r *http.Request) {
	today := r.URL.Query().Get("today")
	report, err := handler.service.GetReportToday(today)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "Report retrieved successfully",
		"data":    report,
	})
}
