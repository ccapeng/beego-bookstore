package controllers

import (
	"bookstore/models"
	"fmt"

	"github.com/astaxie/beego/orm"
)

// AuthorList controller
func (controller *MainController) AuthorList() {

	fmt.Println("author list")
	o := orm.NewOrm()
	var authors []*models.Author
	num, err := o.QueryTable("author").All(&authors)
	if err == orm.ErrNoRows {
		fmt.Println("No author found.")
	} else if err == orm.ErrMissPK {
		fmt.Println("No primary key found.")
	} else {
		fmt.Println(num)
	}
	controller.Data["authors"] = authors
	controller.TplName = "AuthorList.html"

}

// AuthorAdd controller
func (controller *MainController) AuthorAdd() {
	controller.TplName = "AuthorAdd.html"
}

// AuthorEdit controller
func (controller *MainController) AuthorEdit() {

	id, err := controller.GetInt("id")
	o := orm.NewOrm()
	author := models.Author{ID: id}
	err = o.Read(&author)
	if err == orm.ErrNoRows {
		fmt.Println("No author found.")
	} else if err == orm.ErrMissPK {
		fmt.Println("No primary key found.")
	} else {
		fmt.Println(author.ID, author.LastName, author.FirstName)
	}

	controller.Data["ID"] = author.ID
	controller.Data["LastName"] = author.LastName
	controller.Data["FirstName"] = author.FirstName
	controller.TplName = "AuthorEdit.html"

}

// AuthorSave controller
func (controller *MainController) AuthorSave() {

	id, err := controller.GetInt("id")
	lastName := controller.GetString("lastName")
	firstName := controller.GetString("firstName")
	if err != nil {
		fmt.Println(id)
	}

	if id != 0 {
		o := orm.NewOrm()
		author := models.Author{ID: id}
		err2 := o.Read(&author)
		if err2 == orm.ErrNoRows {
			fmt.Println("No author found.")
		} else if err2 == orm.ErrMissPK {
			fmt.Println("No primary key found.")
		} else {
			fmt.Println(author.ID, author.LastName)
		}
		author.LastName = lastName
		author.FirstName = firstName
		o.Update(&author)

	} else {

		o := orm.NewOrm()
		author := new(models.Author)
		author.LastName = lastName
		author.FirstName = firstName
		fmt.Println("save")
		fmt.Println(author.LastName)
		o.Insert(author)

	}

	controller.Redirect("/author/", 302)

}

// AuthorDelete controller
func (controller *MainController) AuthorDelete() {

	id, err := controller.GetInt("id")
	o := orm.NewOrm()
	author := models.Author{ID: id}
	err = o.Read(&author)
	if err == orm.ErrNoRows {
		fmt.Println("No author found.")
	} else if err == orm.ErrMissPK {
		fmt.Println("No primary key found.")
	} else {
		fmt.Println(author.ID, author.LastName)
	}

	num, err := o.Delete(&author)
	if err == nil {
		fmt.Println(num)
	}
	controller.Redirect("/author/", 302)

}
