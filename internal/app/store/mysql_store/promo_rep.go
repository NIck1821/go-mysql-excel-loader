package mysql_store

import (
	"github.com/Nick1821/go-mysql-excel-loader/internal/app/model"
	"gorm.io/gorm"
)

type PromoRep struct {
	db *gorm.DB
}

func (rep *PromoRep) GetLead(date_start, date_end, city string, limit, offset int) ([]model.Lead, error) {
	var lead []model.Lead
	rep.db.
		Table("crm_lead").
		Select("id, firstname,lastname,fabule, phone, city").
		Where(`date BETWEEN ? and ? and city LIKE ?`, date_start, date_end, "%" + city + "%").
		Limit(limit).
		Offset(offset).
		Scan(&lead)
	return lead, nil
}

// SELECT firstname, lastname,fabule, city, intphone
// FROM `crm_lead`
// WHERE date BETWEEN '2021-08-01 00:00:00' and '2021-12-28 00:00:00' and city = 'Москва'
// LIMIT 110;
