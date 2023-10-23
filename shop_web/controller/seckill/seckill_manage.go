package seckill

import (
	"context"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_product_srv/proto/seckill"
	"strconv"
)

var microSecClient shop_product_srv.SecKillsService

func RegisterSecConsul() {
	consulReg := consul.NewRegistry()
	service := micro.NewService(
		micro.Address("127.0.0.1:6002"),
		micro.Registry(consulReg),
	)
	microSecClient = shop_product_srv.NewSecKillsService("shop_product_srv", service.Client())
}

// 获取活动列表
func GetSeckillList(c *gin.Context) {
	currentPage := c.DefaultQuery("currentPage", "1")
	cpage, _ := strconv.ParseInt(currentPage, 10, 32)
	pageSize := c.DefaultQuery("pageSize", "10")
	psize, _ := strconv.ParseInt(pageSize, 10, 32)
	RegisterSecConsul()
	rep, err := microSecClient.SecKillList(context.TODO(), &shop_product_srv.SecKillsRequest{
		CurrentPage: int32(cpage),
		PageSize:    int32(psize),
	})
	if err != nil {
		c.JSON(http.StatusOK, rep)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":         200,
		"msg":          rep.Msg,
		"seckills":     rep.Seckills,
		"total":        rep.Total,
		"current_page": rep.Current,
		"page_size":    rep.Pagesize,
	})
}

// 添加活动
func SecKillAdd(c *gin.Context) {
	name := c.PostForm("name")
	price := c.PostForm("price")
	price_float, _ := strconv.ParseFloat(price, 32)
	num := c.PostForm("num")
	num_int, _ := strconv.ParseInt(num, 10, 32)
	pid := c.PostForm("pid")
	pid_int, _ := strconv.ParseInt(pid, 10, 32)
	start_time := c.PostForm("start_time")
	end_time := c.PostForm("end_time")
	RegisterSecConsul()
	rep, err := microSecClient.SecKillAdd(context.TODO(), &shop_product_srv.AddSecKillRequest{
		Name:      name,
		Price:     float32(price_float),
		Num:       int32(num_int),
		Pid:       int32(pid_int),
		StartTime: start_time,
		EndTime:   end_time,
	})
	if err != nil {
		c.JSON(http.StatusOK, rep)
		return
	}
	c.JSON(http.StatusOK, rep)
}

// 删除活动
func SecKillDel(c *gin.Context) {
	id := c.PostForm("id")
	sid, _ := strconv.ParseInt(id, 10, 32)
	RegisterSecConsul()
	rep, err := microSecClient.SecKillDel(context.TODO(), &shop_product_srv.DelSecKillRequest{
		Id: int32(sid),
	})
	if err != nil {
		c.JSON(http.StatusOK, rep)
		return
	}
	c.JSON(http.StatusOK, rep)
}

// 根据id显示活动
func GetSecKillById(c *gin.Context) {
	id := c.Query("id")
	sid, _ := strconv.ParseInt(id, 10, 32)
	RegisterSecConsul()
	rep, err := microSecClient.SecKillById(context.TODO(), &shop_product_srv.DelSecKillRequest{
		Id: int32(sid),
	})
	if err != nil {
		c.JSON(http.StatusOK, rep)
		return
	}
	var seckill shop_product_srv.SecKillEditData
	seckill.Id = rep.Id
	seckill.Name = rep.Name
	seckill.Price = rep.Price
	seckill.Num = rep.Num
	seckill.Pid = rep.Pid
	seckill.Pname = rep.Pname
	seckill.StartTime = rep.StartTime
	seckill.EndTime = rep.EndTime
	c.JSON(http.StatusOK, gin.H{
		"code":    rep.Code,
		"msg":     rep.Msg,
		"seckill": seckill,
		"prono":   rep.Prono,
	})
}

// 编辑活动
func UpSecKill(c *gin.Context) {

}
