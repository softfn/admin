package system

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type SysMenu struct {
	Id           int       `orm:"column(id);pk"`
	ParentId     string    `orm:"column(parent_id);size(27);null"`
	Name         string    `orm:"column(name);size(32)"`
	Url          string    `orm:"column(url);size(500);null"`
	Seq          int8      `orm:"column(seq);null"`
	Visible      int8      `orm:"column(visible);null"`
	Tooltip      string    `orm:"column(tooltip);size(64);null"`
	Style        string    `orm:"column(style);size(500);null"`
	Description  string    `orm:"column(description);size(500);null"`
	CreatedBy    string    `orm:"column(created_by);size(27)"`
	CreationDate time.Time `orm:"column(creation_date);type(datetime)"`
	UpdatedBy    string    `orm:"column(updated_by);size(27);null"`
	UpdateDate   time.Time `orm:"column(update_date);type(datetime);null"`
}

func (t *SysMenu) TableName() string {
	return "sys_menu"
}

func init() {
	orm.RegisterModel(new(SysMenu))
}

// AddSysMenu insert a new SysMenu into database and returns
// last inserted Id on success.
func AddSysMenu(m *SysMenu) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSysMenuById retrieves SysMenu by Id. Returns error if
// Id doesn't exist
func GetSysMenuById(id int) (v *SysMenu, err error) {
	o := orm.NewOrm()
	v = &SysMenu{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSysMenu retrieves all SysMenu matches certain condition. Returns empty list if
// no records exist
func GetAllSysMenu(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SysMenu))
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

	var l []SysMenu
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

// UpdateSysMenu updates SysMenu by Id and returns error if
// the record to be updated doesn't exist
func UpdateSysMenuById(m *SysMenu) (err error) {
	o := orm.NewOrm()
	v := SysMenu{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSysMenu deletes SysMenu by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSysMenu(id int) (err error) {
	o := orm.NewOrm()
	v := SysMenu{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SysMenu{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
