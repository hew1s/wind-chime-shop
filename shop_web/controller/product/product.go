package product

import (
	"context"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_product_srv/proto/product"
	"shop_web/utils"
	"strconv"
	"time"
)

var microProClient shop_product_srv.ProductsService

// 将注册服务封装成一个函数
func RegisterProConsul() {
	consulReg := consul.NewRegistry()
	service := micro.NewService(
		micro.Address("127.0.0.1:6002"),
		micro.Registry(consulReg),
	)
	microProClient = shop_product_srv.NewProductsService("shop_product_srv", service.Client())
}

// 获取商品列表
func GetProductList(c *gin.Context) {
	currentPage := c.DefaultQuery("currentPage", "1")
	cpage, _ := strconv.ParseInt(currentPage, 10, 32)
	pageSize := c.DefaultQuery("pageSize", "10")
	psize, _ := strconv.ParseInt(pageSize, 10, 32)
	RegisterProConsul()
	rep, err := microProClient.ProductList(context.TODO(), &shop_product_srv.ProductsRequest{
		CurrentPage: int32(cpage),
		PageSize:    int32(psize),
	})
	if err != nil {
		c.JSON(http.StatusOK, rep)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":         200,
		"msg":          "success",
		"products":     rep.Products,
		"total":        rep.Total,
		"current_page": rep.Current,
		"page_size":    rep.Pagesize,
	})
}

// 添加商品
func ProductAdd(c *gin.Context) {
	name := c.PostForm("name")
	price := c.PostForm("price")
	price_float, _ := strconv.ParseFloat(price, 32)
	num := c.PostForm("num")
	num_int, _ := strconv.ParseInt(num, 10, 32)
	unit := c.PostForm("unit")
	desc := c.PostForm("desc")
	file, err := c.FormFile("pic")
	if err != nil {
		//有错误不保存file
	}
	unix_str := strconv.FormatInt(time.Now().Unix(), 10)
	file_path := "upload/" + unix_str + file.Filename
	c.SaveUploadedFile(file, file_path)
	RegisterProConsul()
	rep, err := microProClient.ProductAdd(context.TODO(), &shop_product_srv.ProductAddRequest{
		Name:  name,
		Price: float32(price_float),
		Num:   int32(num_int),
		Unit:  unit,
		Desc:  desc,
		Pic:   file_path,
	})
	if err != nil {
		c.JSON(http.StatusOK, rep)
		return
	}
	c.JSON(http.StatusOK, rep)
}

// 删除商品
func ProductDel(c *gin.Context) {
	id := c.PostForm("id")
	pid, _ := strconv.ParseInt(id, 10, 32)
	RegisterProConsul()
	rep, err := microProClient.ProductDel(context.TODO(), &shop_product_srv.ProductDelRequest{
		Id: int32(pid),
	})
	if err != nil {
		c.JSON(http.StatusOK, rep)
		return
	}
	c.JSON(http.StatusOK, rep)
}

// 根据id获取商品信息
func GetProductById(c *gin.Context) {
	id := c.Query("id")
	pid, _ := strconv.ParseInt(id, 10, 32)
	var product struct {
		ID    int     `json:"id"`
		Name  string  `json:"name"`
		Price float32 `json:"price"`
		Num   int     `json:"num"`
		Uint  string  `json:"unit"`
		Desc  string  `json:"desc"`
	}
	RegisterProConsul()
	rep, err := microProClient.GetProduct(context.TODO(), &shop_product_srv.ProductDelRequest{
		Id: int32(pid),
	})
	if err != nil {
		c.JSON(http.StatusOK, rep)
		return
	}
	product.ID = int(rep.Id)
	product.Name = rep.Name
	product.Price = rep.Price
	product.Num = int(rep.Num)
	product.Uint = rep.Unit
	product.Desc = rep.Desc
	c.JSON(http.StatusOK, gin.H{
		"code":       rep.Code,
		"msg":        rep.Msg,
		"product":    product,
		"img_base64": utils.ImgToBase64(rep.Pic),
	})
}

// 编辑商品
func EditProduct(c *gin.Context) {
	id := c.PostForm("id")
	pid, _ := strconv.ParseInt(id, 10, 32)
	name := c.PostForm("name")
	price := c.PostForm("price")
	price_float, _ := strconv.ParseFloat(price, 32)
	num := c.PostForm("num")
	num_int, _ := strconv.ParseInt(num, 10, 32)
	unit := c.PostForm("unit")
	desc := c.PostForm("desc")
	file, err := c.FormFile("pic")
	if err != nil {
		RegisterProConsul()
		rep, err := microProClient.EditProduct(context.TODO(), &shop_product_srv.ProductEditRequest{
			Name:  name,
			Price: float32(price_float),
			Num:   int32(num_int),
			Id:    int32(pid),
			Unit:  unit,
			Desc:  desc,
		})
		if err != nil {
			c.JSON(http.StatusOK, rep)
			return
		}
		c.JSON(http.StatusOK, rep)
		return
	}
	unix_str := strconv.FormatInt(time.Now().Unix(), 10)
	file_path := "upload/" + unix_str + file.Filename
	RegisterProConsul()
	c.SaveUploadedFile(file, file_path)
	rep, err := microProClient.ProductAdd(context.TODO(), &shop_product_srv.ProductAddRequest{
		Name:  name,
		Price: float32(price_float),
		Num:   int32(num_int),
		Unit:  unit,
		Desc:  desc,
		Pic:   file_path,
	})
	if err != nil {
		c.JSON(http.StatusOK, rep)
		return
	}
	c.JSON(http.StatusOK, rep)
}
