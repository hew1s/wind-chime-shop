package controller

import (
	"context"
	"shop_product_srv/data_source"
	"shop_product_srv/models"
	"shop_product_srv/proto/seckill"
	"shop_product_srv/utils"
	"time"
)

type SecKills struct{}

// 获取活动列表
func (s *SecKills) SecKillList(ctx context.Context, in *shop_product_srv.SecKillsRequest, out *shop_product_srv.SecKillsResponse) error {
	currentpage := in.CurrentPage
	pagesize := in.PageSize
	var seckills []models.SecKills
	var count int32
	count = int32(data_source.Db.Find(&seckills).RowsAffected)
	result := data_source.Db.Limit(pagesize).Offset(pagesize * (currentpage - 1)).Find(&seckills)
	if result.Error != nil {
		out.Code = 500
		out.Msg = "服务错误，请重试!"
		return result.Error
	}
	var seckillList []*shop_product_srv.SecKill
	if result.RowsAffected < 1 {
		out.Code = 500
		out.Msg = "未查询到活动数据"
		return result.Error
	}
	for _, v := range seckills {
		var seckill shop_product_srv.SecKill
		seckill.Id = int32(v.Id)
		seckill.Name = v.Name
		seckill.Price = v.Price
		seckill.Num = int32(v.Num)
		var pro models.Products
		data_source.Db.Where("id = ?", int32(v.PId)).First(&pro)
		seckill.Pname = pro.Name
		seckill.Pid = int32(v.PId)
		seckill.StartTime = v.StartTime.Format("2006-01-02 15:04:05")
		seckill.EndTime = v.EndTime.Format("2006-01-02 15:04:05")
		seckill.CreateTime = v.CreateTime.Format("2006-01-02 15:04:05")
		seckillList = append(seckillList, &seckill)
	}
	out.Code = 200
	out.Msg = "查询成功"
	out.Current = currentpage
	out.Pagesize = pagesize
	out.Seckills = seckillList
	out.Total = count
	return nil
}

// 添加活动
func (s *SecKills) SecKillAdd(ctx context.Context, in *shop_product_srv.AddSecKillRequest, out *shop_product_srv.AddSecKillResponse) error {
	start_time := in.StartTime
	stime, _ := time.ParseInLocation("2006-01-02 15:04:05", start_time, time.Local)
	end_time := in.EndTime
	etime, _ := time.ParseInLocation("2006-01-02 15:04:05", end_time, time.Local)
	var seckill models.SecKills
	seckill.Name = in.Name
	seckill.Price = in.Price
	seckill.Num = int(in.Num)
	seckill.PId = int(in.Pid)
	seckill.StartTime = stime
	seckill.EndTime = etime
	seckill.Status = 0
	seckill.CreateTime = time.Now()
	result := data_source.Db.Create(&seckill)
	if result.Error != nil {
		out.Code = 500
		out.Msg = "添加失败"
		return result.Error
	}
	out.Code = 200
	out.Msg = "添加成功"
	return nil
}

// 删除活动
func (s *SecKills) SecKillDel(ctx context.Context, in *shop_product_srv.DelSecKillRequest, out *shop_product_srv.AddSecKillResponse) error {
	id := in.Id
	var seckill models.SecKills
	err := data_source.Db.Where("id = ?", id).Delete(&seckill).Error
	if err != nil {
		out.Code = 500
		out.Msg = "删除失败"
		return err
	}
	out.Code = 200
	out.Msg = "删除成功"
	return nil
}

// 根据id显示活动
func (s *SecKills) SecKillById(ctx context.Context, in *shop_product_srv.DelSecKillRequest, out *shop_product_srv.SecKillEditData) error {
	id := in.Id
	var seckill models.SecKills
	err := data_source.Db.Where("id = ?", id).First(&seckill).Error
	if err != nil {
		out.Code = 500
		out.Msg = "服务错误"
		return err
	}
	var prono []models.Products
	err = data_source.Db.Not(map[string]interface{}{"id": seckill.PId}).Find(&prono).Error
	if err != nil {
		out.Code = 500
		out.Msg = "服务错误"
		return err
	}
	var productno []*shop_product_srv.Productno
	for _, v := range prono {
		var pro shop_product_srv.Productno
		pro.Id = int32(v.Id)
		pro.Name = v.Name
		productno = append(productno, &pro)
	}
	out.Code = 200
	out.Msg = "success"
	out.Id = int32(seckill.Id)
	out.Name = seckill.Name
	out.Price = seckill.Price
	out.Num = int32(seckill.Num)
	out.Pid = int32(seckill.PId)
	var pro models.Products
	data_source.Db.Where("id = ?", int32(seckill.PId)).First(&pro)
	out.Pname = pro.Name
	out.Prono = productno
	out.StartTime = seckill.StartTime.Format("2006-01-02 15:04:05")
	out.EndTime = seckill.EndTime.Format("2006-01-02 15:04:05")
	return nil
}

// 编辑活动
func (s *SecKills) SecKillEdit(ctx context.Context, in *shop_product_srv.SecKillEditRequest, out *shop_product_srv.AddSecKillResponse) error {
	id := in.Id
	name := in.Name
	price := in.Price
	num := in.Num
	pid := in.Pid
	starttime := in.StartTime
	stime, _ := time.ParseInLocation("2006-01-02 15:04:05", starttime, time.Local)
	endtime := in.EndTime
	etime, _ := time.ParseInLocation("2006-01-02 15:04:05", endtime, time.Local)
	var seckill models.SecKills
	err := data_source.Db.Model(&seckill).Omit("create_time").Where("id =?", id).Updates(map[string]interface{}{
		"name":       name,
		"price":      price,
		"num":        num,
		"p_id":       pid,
		"start_time": stime,
		"end_time":   etime,
	}).Error
	if err != nil {
		out.Code = 500
		out.Msg = "修改失败"
		return err
	}
	out.Code = 200
	out.Msg = "success"
	return nil
}

// 前端活动展示
func (s *SecKills) FrontSecKills(ctx context.Context, in *shop_product_srv.FrontSecKillRequest, out *shop_product_srv.FrontSecKillResponse) error {
	tomorrow := utils.AddHour(24)
	currentpage := in.CurrentPage
	pagesize := in.PageSize
	var seckills []models.SecKills
	var count int32
	count = int32(data_source.Db.Find(&seckills).RowsAffected)
	err := data_source.Db.Where("start_time <=?", tomorrow).Where("status=?", 0).Limit(pagesize).Offset(pagesize * (currentpage - 1)).Find(&seckills).Error
	if err != nil {
		out.Code = 500
		out.Msg = "查询失败"
		return err
	}
	var seckilllist []*shop_product_srv.SecKill
	for _, v := range seckills {
		var seckill shop_product_srv.SecKill
		seckill.Id = int32(v.Id)
		seckill.Name = v.Name
		var pro models.Products
		data_source.Db.Where("id = ?", v.PId).First(&pro)
		seckill.Price = v.Price
		seckill.Pprice = pro.Price
		seckill.Pdesc = pro.Desc
		seckill.Pic = pro.Pic
		seckilllist = append(seckilllist, &seckill)
	}
	out.Code = 200
	out.Msg = "success"
	out.SeckillList = seckilllist
	out.Currentpage = currentpage
	out.Pagesize = pagesize
	totalpage := count / pagesize
	yu := count % pagesize
	if yu == 0 {
		out.Totalpage = totalpage
	} else {
		out.Totalpage = totalpage + 1
	}
	return nil
}

// 活动详情展示
func (s *SecKills) FrontSeckillDetail(ctx context.Context, in *shop_product_srv.DelSecKillRequest, out *shop_product_srv.FrontSeckillDetailResponse) error {
	id := in.Id
	var seckill models.SecKills
	err := data_source.Db.Where("id =?", id).First(&seckill).Error
	if err != nil {
		out.Code = 500
		out.Msg = "查询错误"
		return err
	}
	var seckillDetail shop_product_srv.SecKill
	seckillDetail.Id = int32(seckill.Id)
	seckillDetail.Name = seckill.Name
	seckillDetail.Num = int32(seckill.Num)
	seckillDetail.Price = seckill.Price
	seckillDetail.Pid = int32(seckill.PId)
	var pro models.Products
	data_source.Db.Where("id = ?", seckill.PId).First(&pro)
	seckillDetail.Pic = pro.Pic
	seckillDetail.Pprice = pro.Price
	seckillDetail.Pdesc = pro.Desc
	seckillDetail.Unit = pro.Unit
	seckillDetail.StartTime = seckill.StartTime.Format("2006-01-02 15:04:04")
	seckillDetail.EndTime = seckill.EndTime.Format("2006-01-02 15:04:04")
	out.Code = 200
	out.Msg = "success"
	out.Seckilldetail = &seckillDetail
	return nil
}
