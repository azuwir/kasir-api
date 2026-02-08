package services

import (
	"kasir-api/repositories"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (service *ReportService) GetReport(start_date, end_date string) (map[string]interface{}, error) {
	return service.repo.GetReport(start_date, end_date)
}

func (service *ReportService) GetReportToday(today string) (map[string]interface{}, error) {
	return service.repo.GetReportToday(today)
}
