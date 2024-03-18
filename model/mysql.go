package model

import (
	"context"
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Goods struct {
	gorm.Model
	Name  string
	Image string
}
type Norm struct {
	gorm.Model
	GoodsId int
	Name    string
}
type Norms struct {
	gorm.Model
	NormId int
	Value  string
}

func connect() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3305)/day01?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

type handle func(ctx context.Context, dbs *gorm.DB) (interface{}, error)

func initMysql(ctx context.Context, handle handle) (interface{}, error) {
	db := connect()
	i, err := handle(context.Background(), db)
	if err != nil {
		return nil, err
	}

	dbs, _ := db.DB()
	dbs.Close()
	return i, nil
}
func Create(ctx context.Context, name, image string) (int, error) {
	id, err := initMysql(ctx, func(ctx context.Context, dbs *gorm.DB) (interface{}, error) {
		var good Goods
		dbs.Where("name=?", name).First(&good)
		if good.ID != 0 {
			return nil, errors.New("商品名称已存在")
		}
		good.Name = name
		good.Image = image
		err := dbs.Create(&good).Error
		dbs.Where("name=?", name).First(&good)
		return good.ID, err
	})

	return id.(int), err
}
func GetNormId(ctx context.Context, goodsId int) []int {
	product, err := initMysql(ctx, func(ctx context.Context, dbs *gorm.DB) (interface{}, error) {
		var norm []Norm
		var id []int
		err := dbs.Where("goods_id=?", goodsId).Find(&norm).Error
		for _, v := range norm {
			id = append(id, int(v.ID))
		}

		return id, err
	})
	if err == nil {
		return nil
	}
	return product.([]int)
}
func CreatNorm(ctx context.Context, name string, goodsId int) error {
	_, err := initMysql(ctx, func(ctx context.Context, dbs *gorm.DB) (interface{}, error) {
		var norm Norm
		norm.Name = name
		norm.GoodsId = goodsId
		return nil, dbs.Create(&norm).Error
	})
	return err
}
func CreatNorms(ctx context.Context, goodsId int, name []string) (int, error) {
	ids, err := initMysql(ctx, func(ctx context.Context, dbs *gorm.DB) (interface{}, error) {
		switch len(name) {
		case 1:
			err := CreatNorm(ctx, name[0], goodsId)
			if err != nil {
				return 0, err
			}
		case 2:
			for i := 0; i < len(name); i++ {
				err := CreatNorm(ctx, name[i], goodsId)
				if err != nil {
					return 0, err
				}
			}
		case 3:
			for i := 0; i < len(name); i++ {
				err := CreatNorm(ctx, name[i], goodsId)
				if err != nil {
					return 0, err
				}
			}
		default:
			return 0, errors.New("最多只能写三个规格")
		}

		return 0, errors.New("至少填一个")
	})
	return ids.(int), err

}
func CreateNormValue(ctx context.Context, goodsId int, value []string) error {
	_, err := initMysql(ctx, func(ctx context.Context, dbs *gorm.DB) (interface{}, error) {
		var norms Norms
		id := GetNormId(ctx, goodsId)
		if len(id) == 1 {
			for _, v := range value {
				norms.NormId = id[0]
				norms.Value = v
				err := dbs.Create(&norms).Error
				return nil, err
			}
		}
		for i := 0; i < len(id); i++ {
			for _, v1 := range value[i] {
				norms.NormId = id[i]
				norms.Value = string(v1)
				err := dbs.Create(&norms).Error
				return nil, err
			}
		}
		return nil, errors.New("请先输入规格")
	})
	return err
}
