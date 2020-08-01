package controllers

import (
	"bookstore/models"
	"fmt"

	"github.com/astaxie/beego/orm"
)

// PublisherList controller
func (controller *MainController) PublisherList() {

	fmt.Println("publisher list")
	o := orm.NewOrm()
	var publishers []*models.Publisher
	num, err := o.QueryTable("publisher").All(&publishers)
	if err == orm.ErrNoRows {
		fmt.Println("No publishers found.")
	} else if err == orm.ErrMissPK {
		fmt.Println("No primary key found.")
	} else {
		fmt.Println(num)
	}
	controller.Data["publishers"] = publishers
	controller.TplName = "PublisherList.html"

}

// PublisherAdd controller
func (controller *MainController) PublisherAdd() {
	controller.TplName = "PublisherAdd.html"
}

// PublisherEdit controller
func (controller *MainController) PublisherEdit() {

	id, err := controller.GetInt("id")
	o := orm.NewOrm()
	publisher := models.Publisher{ID: id}
	err = o.Read(&publisher)
	if err == orm.ErrNoRows {
		fmt.Println("No publishers found.")
	} else if err == orm.ErrMissPK {
		fmt.Println("No primary key found.")
	} else {
		fmt.Println(publisher.ID, publisher.Name)
	}

	controller.Data["ID"] = publisher.ID
	controller.Data["Name"] = publisher.Name
	controller.TplName = "PublisherEdit.html"

}

// PublisherSave controller
func (controller *MainController) PublisherSave() {

	id, err := controller.GetInt("id")
	name := controller.GetString("name")
	if err != nil {
		fmt.Println(id)
	}

	if id != 0 {
		o := orm.NewOrm()
		publisher := models.Publisher{ID: id}
		err2 := o.Read(&publisher)
		if err2 == orm.ErrNoRows {
			fmt.Println("No publisher found.")
		} else if err2 == orm.ErrMissPK {
			fmt.Println("No primary key found.")
		} else {
			fmt.Println(publisher.ID, publisher.Name)
		}
		publisher.Name = name

		o.Update(&publisher)

	} else {

		o := orm.NewOrm()
		publisher := new(models.Publisher)
		publisher.Name = name
		fmt.Println("save")
		fmt.Println(publisher.Name)
		o.Insert(publisher)

	}

	controller.Redirect("/publisher/", 302)

}

// PublisherDelete controller
func (controller *MainController) PublisherDelete() {

	id, err := controller.GetInt("id")
	o := orm.NewOrm()
	publisher := models.Publisher{ID: id}
	err = o.Read(&publisher)
	if err == orm.ErrNoRows {
		fmt.Println("No publisher found.")
	} else if err == orm.ErrMissPK {
		fmt.Println("No primary key found.")
	} else {
		fmt.Println(publisher.ID, publisher.Name)
	}

	num, err := o.Delete(&publisher)
	if err == nil {
		fmt.Println(num)
	}
	controller.Redirect("/publisher/", 302)

}
