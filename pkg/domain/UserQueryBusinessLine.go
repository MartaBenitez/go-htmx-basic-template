package domain

import "htmx.try/m/v2/pkg/domain/dto"

type UserQueryBusinessLine struct {
	Sections_to_edit map[string]string
	Base             dto.BusinessLineData
}