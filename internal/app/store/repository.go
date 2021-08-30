package store

import "github.com/Nick1821/go-mysql-excel-loader/internal/app/model"

type LeadRep interface {
	GetLead(date_start, date_end, city, region string, limit, offset int) ([]model.Lead, error)
}
