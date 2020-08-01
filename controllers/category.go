package controllers

import (
	"bookstore/models"
	"fmt"

	"github.com/astaxie/beego/orm"
)

// CategoryList controller
func (controller *MainController) CategoryList() {

	fmt.Println("category list")
	o := orm.NewOrm()
	var cats []*models.Category
	num, err := o.QueryTable("category").All(&cats)
	if err == orm.ErrNoRows {
		fmt.Println("No category found.")
	} else if err == orm.ErrMissPK {
		fmt.Println("No primary key found.")
	} else {
		fmt.Println(num)
		//fmt.Println(cats)
	}
	controller.Data["cats"] = cats
	controller.TplName = "CategoryList.html"

}

// CategoryAdd controller
func (controller *MainController) CategoryAdd() {
	controller.TplName = "CategoryAdd.html"
}

// CategoryEdit controller
func (controller *MainController) CategoryEdit() {

	id, err := controller.GetInt("id")
	o := orm.NewOrm()
	cat := models.Category{ID: id}
	err = o.Read(&cat)
	if err == orm.ErrNoRows {
		fmt.Println("No category found.")
	} else if err == orm.ErrMissPK {
		fmt.Println("No primary key found.")
	} else {
		fmt.Println(cat.ID, cat.Name)
	}

	controller.Data["ID"] = cat.ID
	controller.Data["Name"] = cat.Name
	controller.TplName = "CategoryEdit.html"

}

// CategorySave controller
func (controller *MainController) CategorySave() {

	id, err := controller.GetInt("id")
	name := controller.GetString("name")
	if err != nil {
		fmt.Println(id)
	}

	if id != 0 {
		o := orm.NewOrm()
		cat := models.Category{ID: id}
		err2 := o.Read(&cat)
		if err2 == orm.ErrNoRows {
			fmt.Println("No category found.")
		} else if err2 == orm.ErrMissPK {
			fmt.Println("No primary key found.")
		} else {
			fmt.Println(cat.ID, cat.Name)
		}
		cat.Name = name

		o.Update(&cat)

	} else {

		o := orm.NewOrm()
		cat := new(models.Category)
		cat.Name = name
		fmt.Println("save")
		fmt.Println(cat.Name)
		o.Insert(cat)

	}

	controller.Redirect("/category/", 302)

}

// CategoryDelete controller
func (controller *MainController) CategoryDelete() {

	id, err := controller.GetInt("id")
	o := orm.NewOrm()
	cat := models.Category{ID: id}
	err = o.Read(&cat)
	if err == orm.ErrNoRows {
		fmt.Println("No category found.")
	} else if err == orm.ErrMissPK {
		fmt.Println("No primary key found.")
	} else {
		fmt.Println(cat.ID, cat.Name)
	}

	num, err := o.Delete(&cat)
	if err == nil {
		fmt.Println(num)
	}
	controller.Redirect("/category/", 302)

}
