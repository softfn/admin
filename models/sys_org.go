package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type SysOrg struct {
	Id           int       `orm:"column(id);pk"`
	ParentId     string    `orm:"column(parent_id);size(27);null"`
	VirtualCode  string    `orm:"column(virtual_code);size(64);null"`
	Code         string    `orm:"column(code);size(32);null"`
	Name         string    `orm:"column(name);size(128)"`
	FullName     string    `orm:"column(full_name);size(512);null"`
	Type         int8      `orm:"column(type);null"`
	DepChief     string    `orm:"column(dep_chief);size(64);null"`
	Email        string    `orm:"column(email);size(30);null"`
	Tel          string    `orm:"column(tel);size(12);null"`
	Address      string    `orm:"column(address);size(512);null"`
	Seq          int16     `orm:"column(seq);null"`
	DelFlag      int8      `orm:"column(del_flag);null"`
	DelTime      time.Time `orm:"column(del_time);type(datetime);null"`
	Description  string    `orm:"column(description);size(500);null"`
	OriginId     string    `orm:"column(origin_id);size(64);null"`
	CreatedBy    string    `orm:"column(created_by);size(27)"`
	CreationDate time.Time `orm:"column(creation_date);type(datetime)"`
	UpdatedBy    string    `orm:"column(updated_by);size(27);null"`
	UpdateDate   time.Time `orm:"column(update_date);type(datetime);null"`
}

func (t *SysOrg) TableName() string {
	return "sys_org"
}

func init() {
	orm.RegisterModel(new(SysOrg))
}

// AddSysOrg insert a new SysOrg into database and returns
// last inserted Id on success.
func AddSysOrg(m *SysOrg) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSysOrgById retrieves SysOrg by Id. Returns error if
// Id doesn't exist
func GetSysOrgById(id int) (v *SysOrg, err error) {
	o := orm.NewOrm()
	v = &SysOrg{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSysOrg retrieves all SysOrg matches certain condition. Returns empty list if
// no records exist
func GetAllSysOrg(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SysOrg))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []SysOrg
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateSysOrg updates SysOrg by Id and returns error if
// the record to be updated doesn't exist
func UpdateSysOrgById(m *SysOrg) (err error) {
	o := orm.NewOrm()
	v := SysOrg{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSysOrg deletes SysOrg by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSysOrg(id int) (err error) {
	o := orm.NewOrm()
	v := SysOrg{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SysOrg{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
