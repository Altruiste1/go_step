package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)
type CashFlow struct {
	Id               uint `gorm:"primary_key"`
	CreatedAt        string
	UpdatedAt        string
	UserId           uint
	ParkLotId        uint        //用户收支相对的商户
	MoneyValue       float64      `gorm:"type:decimal(19,2);null"` //商家收钱为正数，付给用户为负数
	MoneyNotified    float64      `gorm:"type:decimal(19,2);null"` //支付渠道回调的钱数
	CashReasonId     uint         `gorm:"not null;default:0"`      //停车收款 付款,奖励等，
	CashChannelId    uint         //1:支付宝,2:微信,等， 转账的源地址
	CashChannelDstId uint         `gorm:"not null;default:0"` //1:支付宝,2:微信,billPackFund 等,转账的目的地址
	OrderId          uint         //如果CashReasonId是停车收款,对应订单号

	ChargeType       uint                `gorm:"type:int;default:0"`
	TradeNumber      string              `gorm:"type:varchar(100);null"`
	TenantId         uint                `gorm:"type:int;index;default:0;"`
	TenantUserId     uint                `gorm:"type:int;index;default:0;"`
	State            string              `gorm:"type:enum('paying','payed','cancel','disable');null"`
	Remark           string              `gorm:"type:varchar(255);null"`
	PaperCouponCode  string              `gorm:"type:varchar(255);null"`
	CouponMoney      float64             `gorm:"type:decimal(19,2);default:0.00"`
	BillCouponId     uint                `gorm:"type:int;default:0;"`
	PayTradeNumber   string              `gorm:"type:varchar(255);null"`
	CashReasonRemark string              `gorm:"-"`
	SceneType        string              `gorm:"-"` //h5支付使用
	PlateNo          string              `gorm:"-"` //无感支付使用
	PlateColor       int                 `gorm:"-"` //无感支付使用
	PlateNumber      string              `gorm:"-"`
	StartTime        string              `gorm:"-"`
	ParkId           string              `gorm:"-"`
}
func main()  {
	filePath:="users.csv"
	f,err:=os.Open(filePath)
	if err!=nil{
		log.Fatal(err)
		return
	}
	r:=csv.NewReader(f)
	res,err:=r.ReadAll()
	if err!=nil{
		log.Fatal(err)
		return
	}
	wxOpenId:=make([]string,0,len(res))
	for _,v:=range res{
		wxOpenId=append(wxOpenId,v[36])
	}
	for _,v:=range wxOpenId{
		fmt.Println(v)
	}
}

