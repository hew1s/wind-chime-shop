package controller

import (
	"context"
	"errors"
	"fmt"
	"shop_seckill_srv/data_source"
	"shop_seckill_srv/models"
	"shop_seckill_srv/proto/secpro"
	"time"
)

type SecKill struct{}

func (s *SecKill) FrontSecKill(ctx context.Context, in *shop_seckill_srv.SecKillRequest, out *shop_seckill_srv.SecKillResponse) error {
	id := in.Id
	now_time := time.Now()
	//秒杀逻辑
	//1.限制购买时间  当前时间 大于等于 start_time 并且小于等于 结束时间
	//2.数量限制
	//3.生成订单
	//4.限制 用户购买只能购买一个

	var seckill models.SecKills
	result := data_source.Db.Where("id = ?", id).Find(&seckill)
	if result.Error != nil {
		out.Code = 500
		out.Msg = "下单失败"
		return errors.New("查询错误")
	}
	fmt.Println("result", result.Error)
	//判断开始时间
	ret_start := result.Where("start_time >?", now_time).Find(&seckill)
	fmt.Println("ret_start", ret_start.Error)
	fmt.Println(ret_start.Error == nil)
	if ret_start.RowsAffected != 0 {
		out.Code = 500
		out.Msg = "活动未开始"
		return nil
	}
	//判断结束时间
	ret_end := result.Where("end_time <=?", now_time).Find(&seckill)
	fmt.Println("ret_end", ret_end.Error)
	if ret_end.RowsAffected != 0 {
		out.Code = 500
		out.Msg = "活动已结束"
		return ret_end.Error
	}
	//判断是否有库存
	if seckill.Num < 1 {
		out.Code = 500
		out.Msg = "当前商品无货"
		return result.Error
	}
	ret := result.Where("num > 0").Update("num", seckill.Num-1)
	if ret.Error != nil {
		out.Code = 500
		out.Msg = "下单失败"
		return ret.Error
	}

	//var seckill models.SecKills
	////1.判断开始时间
	//ret_start := data_source.Db.Where("id = ?", id).Where("start_time >?", now_time).Find(&seckill)
	//if ret_start.RowsAffected != 0 {
	//	fmt.Println("+++++++++")
	//	out.Code = 500
	//	out.Msg = "活动未开始"
	//	return ret_start.Error
	//}
	////2.判断结束时间
	//ret_end := ret_start.Where("end_time <=?", now_time).Find(&seckill)
	//if ret_end.RowsAffected != 0 {
	//	out.Code = 500
	//	out.Msg = "活动已结束"
	//	return ret_end.Error
	//}
	//
	////3.是否有库存
	//result := ret_end.Where("num >0").Find(&seckill)
	//if result.RowsAffected != 0 {
	//	out.Code = 500
	//	out.Msg = "当前商品无货"
	//	return result.Error
	//}
	//ret := result.Where("num > 0").Update("num", seckill.Num-1)
	//if ret.RowsAffected == 0 {
	//	out.Code = 500
	//	out.Msg = "下单失败"
	//	return ret.Error
	//}
	out.Code = 200
	out.Msg = "下单成功"
	return nil
}
