package models

import (
	"github.com/astaxie/beego/orm"
)

// Category object
type Category struct {
	ID   int    `orm:"auto"`
	Name string `orm:"size(32)"`
}

// Publisher object
type Publisher struct {
	ID   int    `orm:"auto"`
	Name string `orm:"size(32)"`
}

// Author object
type Author struct {
	ID        int    `orm:"auto"`
	LastName  string `orm:"size(32)"`
	FirstName string `orm:"size(32)"`
}

// Book object
type Book struct {
	ID        int        `orm:"auto"`
	Title     string     `orm:"size(255)"`
	Category  *Category  `orm:"rel(fk)"`
	Publisher *Publisher `orm:"rel(fk)"`
	Author    *Author    `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(Category))
	orm.RegisterModel(new(Publisher))
	orm.RegisterModel(new(Author))
	orm.RegisterModel(new(Book))
}
