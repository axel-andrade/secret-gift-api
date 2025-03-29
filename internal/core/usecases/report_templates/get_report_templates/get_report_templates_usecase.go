package get_report_templates

type GetReportTemplatesUC struct {
	Gateway GetReportTemplatesGateway
}

func BuildGetReportTemplatesUC(g GetReportTemplatesGateway) *GetReportTemplatesUC {
	return &GetReportTemplatesUC{Gateway: g}
}

func (bs *GetReportTemplatesUC) Execute(input GetReportTemplatesInputDTO) (*GetReportTemplatesOutputDTO, error) {
	reportTemplates, err := bs.Gateway.GetReportTemplatesByConsultationType(input.ConsultationType)

	if err != nil {
		return nil, err
	}

	return &GetReportTemplatesOutputDTO{reportTemplates}, nil
}
