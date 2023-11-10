package dbconn

import (
	"htmx.try/m/v2/pkg/domain"
)

type InMemoryDB struct {
	data map[string]domain.InterfaceResponseFull
	base map[string][]domain.BaseToSave
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		data: make(map[string]domain.InterfaceResponseFull),
		base: make(map[string][]domain.BaseToSave),
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

func (db *InMemoryDB) GetBases(key string) ([]domain.BaseToSave, bool) {
	val, ok := db.base[key]
	return val, ok
}

func (db *InMemoryDB) SetBase(key string, value domain.BaseToSave) {
	db.base[key] = append(db.base[key], value)
}

func (db *InMemoryDB) DeleteBases(key string) {
	delete(db.base, key)
}
