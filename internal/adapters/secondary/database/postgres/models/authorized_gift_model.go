package postgres_models

import (
	"time"

	"github.com/google/uuid"
)

type AuthorizedGiftModel struct {
	ID             uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	GiftID         uuid.UUID  `gorm:"type:uuid;not null" json:"gift_id"`
	Gift           GiftModel  `gorm:"foreignKey:GiftID;constraint:OnDelete:CASCADE" json:"gift"` // relacionamento
	ExpirationDate *time.Time `gorm:"type:timestamp" json:"expiration_date"`
	CreatedAt      time.Time  `gorm:"type:timestamp;autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time  `gorm:"type:timestamp;autoUpdateTime" json:"updated_at"`
}

func (AuthorizedGiftModel) TableName() string {
	return "authorized_gifts"
}
