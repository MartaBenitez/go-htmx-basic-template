package results

import "htmx.try/m/v2/pkg/domain/dto"

type BusinessLineAnswer struct {

	Header interface{} `json:"header"`
	Body Codigo `json:"body"`
}

type Codigo struct {
	
}

type BaseToSave struct {
	SectionsToEdit     dto.Sections
	Business_line_data dto.BusinessLineData
	Coverage_data      []dto.CoverageData
}

func NewBaseToSave(sections dto.Sections, business dto.BusinessLineData, coverage []dto.CoverageData) BaseToSave {
	return BaseToSave{
		SectionsToEdit:     sections,
		Business_line_data: business,
		Coverage_data:      coverage,
	}
}