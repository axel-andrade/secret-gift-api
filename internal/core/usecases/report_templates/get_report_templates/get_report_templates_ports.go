package get_report_templates

import "github.com/axel-andrade/secret-gift-api/internal/core/domain"

type GetReportTemplatesGateway interface {
	GetReportTemplatesByConsultationType(consultationType domain.ConsultationType) ([]domain.ReportTemplate, error)
}

type GetReportTemplatesInputDTO struct {
	ConsultationType domain.ConsultationType `json:"consultation_type"`
}

type GetReportTemplatesOutputDTO struct {
	ReportTemplates []domain.ReportTemplate
}
