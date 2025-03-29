package domain

import "errors"

type ReportTemplatesByName map[string]*ReportTemplate

type ConsultationType string

const (
	ConsultationTypeChildcare ConsultationType = "childcare"
	ConsultationTypeCannabis  ConsultationType = "cannabis"
)

type ReportTemplate struct {
	Base
	TemplateName     string           `json:"template_name"`
	ConsultationType ConsultationType `json:"consultation_type"`
	InputSchema      string           `json:"input_schema"`
	Template         string           `json:"template"`
}

func NewReportTemplate(templateName, inputSchema, template string, consultationType ConsultationType) (*ReportTemplate, error) {
	rt := &ReportTemplate{
		TemplateName:     templateName,
		ConsultationType: consultationType,
		InputSchema:      inputSchema,
		Template:         template,
	}

	if err := rt.validate(); err != nil {
		return nil, err
	}

	return rt, nil
}

func (rt *ReportTemplate) validate() error {
	if rt.TemplateName == "" {
		return errors.New("templateName is required")
	}

	if rt.ConsultationType == "" {
		return errors.New("consultationType is required")
	}

	if rt.Template == "" {
		return errors.New("template is required")
	}

	return nil
}
