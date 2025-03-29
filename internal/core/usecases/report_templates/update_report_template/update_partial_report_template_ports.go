package update_partial_report_template

import (
	"context"

	"github.com/axel-andrade/secret-gift-api/internal/core/domain"
)

type UpdatePartialReportTemplateGateway interface {
	GetReportTemplateByID(reportTemplateID string) (*domain.ReportTemplate, error)
	UpdateReportTemplate(reportTemplate domain.ReportTemplate) (*domain.ReportTemplate, error)
	DeleteCachedReportTemplates(ctx context.Context, consultationType domain.ConsultationType) error
}

type UpdatePartialReportTemplateInputDTO struct {
	ReportTemplateID string                  `json:"id"`
	TemplateName     string                  `json:"template_name"`
	ConsultationType domain.ConsultationType `json:"consultation_type"`
	InputSchema      string                  `json:"input_schema"`
	Template         string                  `json:"template"`
}
