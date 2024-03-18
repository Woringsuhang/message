package logic

import (
	"context"
	"errors"
	"fmt"
	"su/goods"
	"su/model"
)

type Server struct {
	goods.UnimplementedStreamGreeterServer
}

func (s Server) CreateSku(ctx context.Context, req *goods.CreateSkuReq) (*goods.CreateSkuRes, error) {
	id := req.GoodsId
	name := req.Name
	value := req.Value
	if id == 0 {
		return &goods.CreateSkuRes{}, errors.New("请输入商品")
	}
	_, err := model.CreatNorms(ctx, int(id), name)
	if err != nil {
		return &goods.CreateSkuRes{}, err
	}
	model.CreateNormValue(ctx, int(id), value)
	return nil, err
}

func (s Server) GoodsCreated(ctx context.Context, req *goods.CreateGoodsReq) (*goods.GoodsResp, error) {
	fmt.Println(11111)
	return &goods.GoodsResp{Msg: 1}, nil

}
