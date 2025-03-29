package update_partial_report_template

import (
	"context"
	"fmt"
	"log"

	err_msg "github.com/axel-andrade/secret-gift-api/internal/core/domain/constants/errors"
)

type UpdatePartialReportTemplateUC struct {
	Gateway UpdatePartialReportTemplateGateway
}

func BuildUpdatePartialReportTemplateUC(g UpdatePartialReportTemplateGateway) *UpdatePartialReportTemplateUC {
	return &UpdatePartialReportTemplateUC{Gateway: g}
}

func (bs *UpdatePartialReportTemplateUC) Execute(ctx context.Context, input UpdatePartialReportTemplateInputDTO) error {
	reportTemplate, err := bs.Gateway.GetReportTemplateByID(input.ReportTemplateID)
	if err != nil {
		log.Printf("Error getting report template: %s", err)
		return err
	}

	if reportTemplate == nil {
		log.Printf("Report template not found: %s", input.ReportTemplateID)
		return fmt.Errorf(err_msg.REPORT_TEMPLATE_NOT_FOUND)
	}

	if input.TemplateName != "" {
		reportTemplate.TemplateName = input.TemplateName
	}

	if input.ConsultationType != "" {
		reportTemplate.ConsultationType = input.ConsultationType
	}

	if input.InputSchema != "" {
		reportTemplate.InputSchema = input.InputSchema
	}

	if input.Template != "" {
		reportTemplate.Template = input.Template
	}

	if _, err = bs.Gateway.UpdateReportTemplate(*reportTemplate); err != nil {
		log.Printf("Error updating report template: %s", err)
		return err
	}

	if err := bs.Gateway.DeleteCachedReportTemplates(ctx, reportTemplate.ConsultationType); err != nil {
		log.Printf("Error deleting cached report templates: %s", err)
		return err
	}

	log.Printf("Report template updated and cache cleared: %s", input.ReportTemplateID)
	return nil
}
