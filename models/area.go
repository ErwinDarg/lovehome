package models

type Area struct {
	Id int
	contury string `orm:size(100)`
}
