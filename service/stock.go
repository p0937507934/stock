package service

import (
	"context"
	"stock/driver"
	"stock/models"
	. "stock/proto/proto"

	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type StockService struct {
	UnimplementedStockServiceServer
	orm *gorm.DB
}

func NewStockService() StockServiceServer {
	return &StockService{orm: driver.InitGorm()}
}

func (s *StockService) UpdateStock(ctx context.Context, req *UpdateStockRequest) (*emptypb.Empty, error) {
	stock := &models.Stock{}
	tx := s.orm.Begin()
	tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(stock, &models.Stock{ProductId: int(req.ProductId)})
	stock.Stock += int(req.Stock)
	stock.Lock += int(req.Lock)
	result := tx.Updates(stock)
	tx.Commit()
	return &emptypb.Empty{}, result.Error
}
