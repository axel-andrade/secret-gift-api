package postgres_models

import (
	"time"
)

type AuthorizedGiftModel struct {
	BaseModel
	GiftID         string     `gorm:"type:uuid;not null" json:"gift_id"`
	Gift           GiftModel  `gorm:"foreignKey:GiftID;constraint:OnDelete:CASCADE" json:"gift"` // relacionamento
	ExpirationDate *time.Time `gorm:"type:timestamp" json:"expiration_date"`
}

func (AuthorizedGiftModel) TableName() string {
	return "authorized_gifts"
}
