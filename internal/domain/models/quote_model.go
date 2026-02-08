package models

import "gorm.io/gorm"

type QuoteModel struct {
	gorm.Model
	Carrier []CarrierModel `gorm:"foreignKey:QuoteID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type CarrierModel struct {
	gorm.Model
	QuoteID  uint
	Name     string  `gorm:"type:varchar(100);not null"`
	Service  string  `gorm:"type:varchar(100);not null"`
	Deadline string  `gorm:"type:varchar(20);not null"`
	Price    float64 `gorm:"not null"`
}
