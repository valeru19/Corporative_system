package services

import (
	"time"

	"bradobrei/backend/internal/models"
	"bradobrei/backend/internal/repository"
)

type ReportService struct {
	reportRepo *repository.ReportRepository
}

func NewReportService(reportRepo *repository.ReportRepository) *ReportService {
	return &ReportService{reportRepo: reportRepo}
}

// Report 2.2.1 — Список сотрудников
func (s *ReportService) GetEmployeeList() ([]models.User, error) {
	return s.reportRepo.GetEmployeeList()
}

// Report 2.2.2 — Месячная активность салонов
func (s *ReportService) GetSalonActivity(from, to time.Time) ([]repository.SalonActivityRow, error) {
	return s.reportRepo.GetSalonActivity(from, to)
}

// Report 2.2.3 — Популярность услуг
func (s *ReportService) GetServicePopularity(from, to time.Time) ([]repository.ServicePopularityRow, error) {
	return s.reportRepo.GetServicePopularity(from, to)
}

// Report 2.2.4 — Активность мастеров
func (s *ReportService) GetMasterActivity(from, to time.Time) ([]repository.MasterActivityRow, error) {
	return s.reportRepo.GetMasterActivity(from, to)
}

// Report 2.2.5 — Отзывы об ИС
func (s *ReportService) GetReviews(from, to time.Time) ([]models.Review, error) {
	return s.reportRepo.GetReviews(from, to)
}

func (s *ReportService) GetInventoryMovement(from, to time.Time, salonID uint) ([]repository.InventoryMovementRow, error) {
	return s.reportRepo.GetInventoryMovement(from, to, salonID)
}

func (s *ReportService) GetClientLoyalty(from, to time.Time) ([]repository.ClientLoyaltyRow, error) {
	return s.reportRepo.GetClientLoyalty(from, to)
}

func (s *ReportService) GetCancelledBookings(from, to time.Time) ([]repository.CancelledBookingRow, error) {
	return s.reportRepo.GetCancelledBookings(from, to)
}

func (s *ReportService) GetFinancialSummary(from, to time.Time, salonID uint) ([]repository.FinancialSummaryRow, error) {
	return s.reportRepo.GetFinancialSummary(from, to, salonID)
}
