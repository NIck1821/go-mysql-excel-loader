package mysql_store

import (
	"github.com/Nick1821/go-mysql-excel-loader/internal/app/model"
	"gorm.io/gorm"
)

type PromoRep struct {
	db *gorm.DB
}

func (rep *PromoRep) GetLead(date_start, date_end, city, region string, limit, offset int) ([]model.Lead, error) {
	var lead []model.Lead
	rep.db.
		Table("crm_lead").
		Select("id, firstname,lastname,fabule, intphone, city").
		Where(`date BETWEEN ? and ? and (city LIKE ? or city LIKE ?)`, date_start, date_end, "%" + city + "%", "%" + region + "%").
		Limit(limit).
		Offset(offset).
		Scan(&lead)
	return lead, nil
}