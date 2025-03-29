package postgres_models

type GiftModel struct {
	BaseModel
	Type       string  `gorm:"type:varchar(50)" json:"type"`
	CPF        string  `gorm:"type:varchar(14);not null" json:"cpf"`
	Amount     float64 `gorm:"type:decimal(10,2)" json:"amount"`
	PixKeyType string  `gorm:"type:varchar(50)" json:"pix_key_type"`
	PixKey     string  `gorm:"type:varchar(255)" json:"pix_key"`
	Message    string  `gorm:"type:text" json:"message"`
	Status     string  `gorm:"type:varchar(50);not null" json:"status"`
}

func (GiftModel) TableName() string {
	return "gifts"
}
