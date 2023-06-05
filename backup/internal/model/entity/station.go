// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Station is the golang structure for table station.
type Station struct {
	Id         int         `json:"id"         ` //
	Name       string      `json:"name"       ` //
	Address    string      `json:"address"    ` //
	Coordinate string      `json:"coordinate" ` //
	CreateAt   *gtime.Time `json:"createAt"   ` //
	UpdateAt   *gtime.Time `json:"updateAt"   ` //
}
