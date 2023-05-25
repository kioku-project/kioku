package model

import (
	"github.com/kioku-project/kioku/pkg/helper"
	"gorm.io/gorm"
)

type AdmissionStatus string

const (
	Invited   AdmissionStatus = "invited"
	Requested AdmissionStatus = "requested"
)

type GroupAdmission struct {
	ID              string
	UserID          string
	GroupID         string
	Group           Group
	AdmissionStatus AdmissionStatus
}

func (a *GroupAdmission) BeforeCreate(db *gorm.DB) (err error) {
	newID, err := helper.FindFreeID(db, 10, func() (helper.PublicID, *GroupAdmission) {
		id := helper.GenerateID('A')
		return id, &GroupAdmission{ID: id.GetStringRepresentation()}
	})
	a.ID = newID.GetStringRepresentation()
	return
}
