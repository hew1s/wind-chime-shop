package controller

import (
	"context"
	"shop_product_srv/data_source"
	"shop_product_srv/models"
	"shop_product_srv/proto/product"
	"time"
)

type Products struct{}

// 获取商品列表
func (p *Products) ProductList(ctx context.Context, in *shop_product_srv.ProductsRequest, out *shop_product_srv.ProductsResponse) error {
	currentPage := in.CurrentPage
	pageSize := in.PageSize
	var products []models.Products
	var count int32
	count = int32(data_source.Db.Find(&products).RowsAffected)
	result := data_source.Db.Limit(pageSize).Offset(pageSize * (currentPage - 1)).Find(&products)
	if result.Error != nil {
		out.Code = 500
		out.Msg = "服务错误，请重试!"
		return result.Error
	}
	var productList []*shop_product_srv.Product
	if result.RowsAffected < 1 {
		out.Code = 500
		out.Msg = "未查询到商品数据"
		return result.Error
	}
	for _, v := range products {
		var product shop_product_srv.Product
		product.Id = int32(v.Id)
		product.Name = v.Name
		product.Price = v.Price
		product.Num = int32(v.Num)
		product.Desc = v.Desc
		product.Unit = v.Unit
		product.Pic = v.Pic
		product.CreateTime = v.CreateTime.Format("2006-01-02 15:04:05")
		productList = append(productList, &product)
	}
	out.Code = 200
	out.Msg = "查询成功"
	out.Current = currentPage
	out.Pagesize = pageSize
	out.Products = productList
	out.Total = count
	return nil
}

// 添加商品
func (p *Products) ProductAdd(ctx context.Context, in *shop_product_srv.ProductAddRequest, out *shop_product_srv.ProductAddResponse) error {
	name := in.Name
	price := in.Price
	num := in.Num
	unit := in.Unit
	pic_path := in.Pic
	desc := in.Desc
	product := models.Products{
		Name:       name,
		Price:      price,
		Num:        int(num),
		Unit:       unit,
		Pic:        pic_path,
		Desc:       desc,
		CreateTime: time.Now(),
	}
	result := data_source.Db.Create(&product)
	if result.Error != nil {
		out.Code = 500
		out.Msg = "添加商品失败"
		return result.Error
	}
	out.Code = 200
	out.Msg = "添加商品成功"
	return nil
}

// 删除商品
func (p *Products) ProductDel(ctx context.Context, in *shop_product_srv.ProductDelRequest, out *shop_product_srv.ProductAddResponse) error {
	id := in.Id
	var product models.Products
	product.Id = int(id)
	err := data_source.Db.Delete(&product).Error
	if err != nil {
		out.Code = 500
		out.Msg = "删除失败"
		return err
	}
	out.Code = 200
	out.Msg = "删除成功"
	return nil
}

// 根据id获取商品信息
func (p *Products) GetProduct(ctx context.Context, in *shop_product_srv.ProductDelRequest, out *shop_product_srv.GetProductResponse) error {
	id := in.Id
	var pro models.Products
	result := data_source.Db.Where("id =?", id).First(&pro)
	if result.Error != nil {
		out.Code = 500
		out.Msg = "未查询到该商品信息"
		return result.Error
	}
	out.Code = 200
	out.Msg = "success"
	out.Id = int32(pro.Id)
	out.Name = pro.Name
	out.Price = pro.Price
	out.Num = int32(pro.Num)
	out.Unit = pro.Unit
	out.Desc = pro.Desc
	out.Pic = pro.Pic
	return nil
}

// 编辑商品
func (p *Products) EditProduct(ctx context.Context, in *shop_product_srv.ProductEditRequest, out *shop_product_srv.ProductAddResponse) error {
	id := in.Id
	name := in.Name
	price := in.Price
	num := in.Num
	unit := in.Unit
	desc := in.Desc
	var pro models.Products
	result := data_source.Db.Model(&pro).Omit("pic").Where("id =?", id).Updates(map[string]interface{}{
		"name":  name,
		"price": price,
		"num":   num,
		"unit":  unit,
		"desc":  desc,
		"pic":   "",
	})
	if result.Error != nil {
		out.Code = 500
		out.Msg = "更新失败"
		return result.Error
	}
	out.Code = 200
	out.Msg = "更新商品数据成功"
	return nil
}
