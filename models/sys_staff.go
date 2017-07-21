package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type SysStaff struct {
	Id            int       `orm:"column(id);pk"`
	ChName        string    `orm:"column(ch_name);size(16);null"`
	ChPy          string    `orm:"column(ch_py);size(32);null"`
	EnName        string    `orm:"column(en_name);size(32);null"`
	Username      string    `orm:"column(username);size(32)"`
	EmpNo         string    `orm:"column(emp_no);size(128);null"`
	EmpStatus     int8      `orm:"column(emp_status);null"`
	Position      string    `orm:"column(position);size(32);null"`
	Rank          string    `orm:"column(rank);size(32);null"`
	Birthday      time.Time `orm:"column(birthday);type(datetime);null"`
	IdCardNo      string    `orm:"column(id_card_no);size(32);null"`
	EntryTime     time.Time `orm:"column(entry_time);type(datetime);null"`
	Email         string    `orm:"column(email);size(200);null"`
	PostalAddress string    `orm:"column(postal_address);size(256);null"`
	PostalCode    string    `orm:"column(postal_code);size(32);null"`
	HomeAddress   string    `orm:"column(home_address);size(500);null"`
	MobileNo      string    `orm:"column(mobile_no);size(16);null"`
	Telephone     string    `orm:"column(telephone);size(16);null"`
	Sex           int8      `orm:"column(sex);null"`
	IsEnabled     int8      `orm:"column(is_enabled);null"`
	DelFlag       int8      `orm:"column(del_flag);null"`
	DelTime       time.Time `orm:"column(del_time);type(datetime);null"`
	Description   string    `orm:"column(description);size(500);null"`
	OriginId      string    `orm:"column(origin_id);size(64);null"`
	CreatedBy     string    `orm:"column(created_by);size(27)"`
	CreationDate  time.Time `orm:"column(creation_date);type(datetime)"`
	UpdatedBy     string    `orm:"column(updated_by);size(27);null"`
	UpdateDate    time.Time `orm:"column(update_date);type(datetime);null"`
}

func (t *SysStaff) TableName() string {
	return "sys_staff"
}

func init() {
	orm.RegisterModel(new(SysStaff))
}

// AddSysStaff insert a new SysStaff into database and returns
// last inserted Id on success.
func AddSysStaff(m *SysStaff) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSysStaffById retrieves SysStaff by Id. Returns error if
// Id doesn't exist
func GetSysStaffById(id int) (v *SysStaff, err error) {
	o := orm.NewOrm()
	v = &SysStaff{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSysStaff retrieves all SysStaff matches certain condition. Returns empty list if
// no records exist
func GetAllSysStaff(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SysStaff))
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

	var l []SysStaff
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

// UpdateSysStaff updates SysStaff by Id and returns error if
// the record to be updated doesn't exist
func UpdateSysStaffById(m *SysStaff) (err error) {
	o := orm.NewOrm()
	v := SysStaff{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSysStaff deletes SysStaff by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSysStaff(id int) (err error) {
	o := orm.NewOrm()
	v := SysStaff{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SysStaff{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
