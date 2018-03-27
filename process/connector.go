package process

import (
	"github.com/jinzhu/gorm"
)

type EndPoint struct {
	gorm.Model
	Id     int
	Name   string
	Params string `sql:"type:json"`
}

type Connector struct {
	gorm.Model
	Id       string
	SourceId int
	TargetId int
	Source   EndPoint
	Target   EndPoint
}

// func (con *Connector) toJson() string {
//
//   json, err := json.Marshal(con.toMap())
//
//   if err != nil {
//     dlog.Errlog.Println("Error when Marshalling %s, %s", con, err)
//     log.Fatal(err)
//   }
//
//   return string(json)
//
// }

func (ep *EndPoint) toMap() map[string]string {
	return map[string]string{
		"params": ep.Params,
		"name":   ep.Name,
	}
}
