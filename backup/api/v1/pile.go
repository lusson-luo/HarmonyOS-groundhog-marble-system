package v1

import (
	"login-demo/internal/model"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type PileListReq struct {
	g.Meta    `path:"/api/pile/list" tags:"pile" method:"get" summary:"充电桩点列表"`
	StationId int    `json:"stationId"`
	Code      string `json:"code"`
	model.PageReq
}

type PileListRes struct {
	g.Meta      `mime:"application/json" example:"string"`
	Id          int       `json:"id"`
	Code        string    `json:"code"`
	StationId   int       `json:"stationId"`
	StationName string    `json:"stationName"`
	State       int       `json:"state"`        // 0:空闲中,1:充电中,2:故障
	CreateAt    time.Time `json:"createAt"    ` //
	UpdateAt    time.Time `json:"updateAt"    ` //
}

type PileAddReq struct {
	g.Meta    `path:"/api/pile/add" tags:"pileAdd" method:"POST" summary:"添加充电桩"`
	Code      string `json:"code"`
	StationId int    `json:"stationId"`
}

type PileAddRes struct {
	g.Meta `mime:"application/json" `
}

type PileDelReq struct {
	g.Meta `path:"/api/pile/delete" tags:"pileDelete" method:"DELETE" summary:"删除充电桩"`
	Id     int `json:"id"`
}

type PileDelRes struct {
	g.Meta `mime:"application/json" `
}

type PileUpdateReq struct {
	g.Meta    `path:"/api/pile/update" tags:"pileDelete" method:"post" summary:"修改充电桩"`
	Id        int    `json:"id"`
	Code      string `json:"code"`
	StationId int    `json:"stationId"`
	State     int    `json:"state"`
}

type PileUpdateRes struct {
	g.Meta `mime:"application/json" `
}
