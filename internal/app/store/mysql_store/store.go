package mysql_store

import (
	"github.com/Nick1821/go-mysql-excel-loader/internal/app/store"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Store - конфигурационная структура для базы данных
type Store struct {
	db       *gorm.DB
	promoRep store.LeadRep
}

// создание объекта
func NewStore() *Store {
	return &Store{}
}

// Open - инициализация подключения хранилища
func (s *Store) Open(dsn string) error {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	s.db = db

	return nil
}

// Возвращает промо
func (s *Store) GetPromoRep() store.LeadRep {

	if s.promoRep == nil {
		s.promoRep = &PromoRep{
			db: s.db,
		}
	}

	return s.promoRep
}
