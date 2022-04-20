package models

type Stock struct {
	Id          int    `gorm:"column:id"`
	ProductId   int    `gorm:"column:product_id;index:idx_product_id"`
	ProductName string `gorm:"column:product_name"`
	Stock       int    `gorm:"column:stock;index:idx_stock"`
	Lock        int    `gorm:"column:lock;index:idx_lock"`
}

func (Stock) TableName() string {
	return "stock"
}
