package domain

import (
	"time"
	"github.com/google/uuid"
)

type RegStateEnum int
type SexEnum int
type TargetNameEnum int
type CategoryNameEnum int

const (
	RegStateAgreements RegStateEnum = iota
	RegStateName
	RegStateSex
	RegStateMainPicture
	RegStateAnotherPicture
	RegStateTarget
)

var regStateName = map[RegStateEnum]string{
	RegStateAgreements: "state_agreements",
	RegStateName: "state_name",
	RegStateSex: "state_sex",
	RegStateMainPicture: "state_main_picture",
	RegStateAnotherPicture: "state_another_picture",
	RegStateTarget: "state_target",
}

func (e RegStateEnum) String() string {
	return regStateName[e]
}

const (
	SexMale SexEnum = iota
	SexFemale
	NotSelected
)

var sexName = map[SexEnum]string{
	SexMale: "Мужчина",
	SexFemale: "Женщина",
	NotSelected: "Не выбрано",
}

func (e SexEnum) String() string {
	return sexName[e]
}

const (
	Activity CategoryNameEnum = iota
	Interests
	Psychograph
	WorldView
)

var categoryName = map[CategoryNameEnum]string{
	Activity: "Активность",
	Interests: "Интересы",
	Psychograph: "Психограф",
	WorldView: "Мировозрение",
}

func (e CategoryNameEnum) String() string {
	return categoryName[e]
}

const (
	TargetForLike TargetNameEnum = iota
	TargetForOneDay
	TargetForAllTime
)

var targetName = map[TargetNameEnum]string{
	TargetForLike: "for_like",
	TargetForOneDay: "for_one_day",
	TargetForAllTime: "for_all_time",
}

func (e TargetNameEnum) String() string {
	return targetName[e]
}

type User struct {
	ID uuid.UUID
	Email string
	RegistrationState RegStateEnum
	Name string
	BIO string
	Birthdate time.Time
	Sex SexEnum
	Location string
	Height int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Picture struct {
	ID uuid.UUID
	IsMain bool
	Approved bool
	Path string
}

type Target struct {
	ID uuid.UUID
	Name TargetNameEnum
}

type Tag struct {
	ID uuid.UUID
	Name string
	Volume int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Category struct {
	ID uuid.UUID
	Name CategoryNameEnum
	Volume int
	CreatedAt time.Time
	UpdetedAt time.Time
	DeletedAt *time.Time
}

