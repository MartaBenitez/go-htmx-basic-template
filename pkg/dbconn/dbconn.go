package dbconn

import (
	"htmx.try/m/v2/pkg/domain"
	"htmx.try/m/v2/pkg/domain/dto"
)

type InMemoryDB struct {
	data      map[string]domain.InterfaceResponseFull
	responses map[string][]dto.BusinessLineData
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		data:      make(map[string]domain.InterfaceResponseFull),
		responses: make(map[string][]dto.BusinessLineData),
	}
}

func (db *InMemoryDB) GetData(key string) (domain.InterfaceResponseFull, bool) {
	val, ok := db.data[key]
	return val, ok
}

func (db *InMemoryDB) SetData(key string, value domain.InterfaceResponseFull) {
	db.data[key] = value
}

func (db *InMemoryDB) DeleteData(key string) {
	delete(db.data, key)
}

func (db *InMemoryDB) GetResponses(key string) ([]dto.BusinessLineData, bool) {
	val, ok := db.responses[key]
	return val, ok
}

func (db *InMemoryDB) SetResponse(key string, value dto.BusinessLineData) {
	db.responses[key] = append(db.responses[key], value)
}

func (db *InMemoryDB) DeleteResponses(key string) {
	delete(db.responses, key)
}
