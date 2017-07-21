package models

import "time"

type SysFunction struct {
	Id_RENAME    string    `orm:"column(id);size(27)"`
	MenuId       string    `orm:"column(menu_id);size(27)"`
	Name         string    `orm:"column(name);size(32);null"`
	Code         string    `orm:"column(code);size(64);null"`
	Seq          int8      `orm:"column(seq);null"`
	CreatedBy    string    `orm:"column(created_by);size(27)"`
	CreationDate time.Time `orm:"column(creation_date);type(datetime)"`
	UpdatedBy    string    `orm:"column(updated_by);size(27);null"`
	UpdateDate   time.Time `orm:"column(update_date);type(datetime);null"`
}
