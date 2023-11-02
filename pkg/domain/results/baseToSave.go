package results

import "htmx.try/m/v2/pkg/domain/dto"

type BaseToSave struct {
	Business_line_data *dto.BusinessLineData
	Coverage_data      *[]dto.CoverageData
}
