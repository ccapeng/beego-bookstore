package controllers

import (
	"bookstore/models"
	"fmt"

	"github.com/astaxie/beego/orm"
)

// BookList controller
func (controller *MainController) BookList() {

	o := orm.NewOrm()
	var books []*models.Book
	//qs := o.QueryTable("book")
	//qs.RelatedSel("category")
	//num, err := o.QueryTable("book").RelatedSet("category").All(&books)
	// num, err := o.QueryTable("book").
	// 	RelatedSel("category").
	// 	RelatedSel("publisher").
	// 	RelatedSel("author").
	// 	All(&books)
	//num, err := o.QueryTable("book").RelatedSel().All(&books)
	num, err := o.QueryTable("book").All(&books)
	//num, err := o.Raw("select a.i_d, a.title, b.name, c.name, d.last_name, d.first_name from book a left join category b on a.category_id = b.i_d left join publisher c on a.category_id = c.i_d left join author d on a.author_id = d.i_d").QueryRows(&books)

	//num, err := o.QueryTable("book").RelatedSel("category").All(&books)
	//num, err := o.QueryTable("book").All(&books)
	//num, err := qs.All(&books)
	if err == orm.ErrNoRows {
		fmt.Println("No book found.")
	} else if err == orm.ErrMissPK {
		fmt.Println("No primary key found.")
	} else {
		fmt.Println(num)
		fmt.Println(books)
	}

	//fmt.Printf("%+v\n", books)

	//fmt.Println("len", len(books))
	o2 := orm.NewOrm()
	for _, book := range books {
		fmt.Printf("%+v\n", book)
		o2.LoadRelated(book, "category")
		o2.LoadRelated(book, "publisher")
		o2.LoadRelated(book, "author")
	}
	controller.Data["books"] = books
	controller.TplName = "BookList.html"

}

func loadFormData(controller *MainController) {

	o := orm.NewOrm()

	var cats []*models.Category
	num, err := o.QueryTable("category").All(&cats)
	if err == orm.ErrNoRows {
		fmt.Println("No category found.")
	} else if err == orm.ErrMissPK {
		fmt.Println("No primary key found.")
	} else {
		fmt.Println(num)
		fmt.Println(cats)
	}
	controller.Data["cats"] = cats

	var publishers []*models.Publisher
	num2, err2 := o.QueryTable("publisher").All(&publishers)
	if err2 == orm.ErrNoRows {
		fmt.Println("No publisher found.")
	} else if err2 == orm.ErrMissPK {
		fmt.Println("No primary key found.")
	} else {
		fmt.Println(num2)
		fmt.Println(publishers)
	}
	controller.Data["publishers"] = publishers

	var authors []*models.Author
	num3, err3 := o.QueryTable("author").All(&authors)
	if err3 == orm.ErrNoRows {
		fmt.Println("No author found.")
	} else if err3 == orm.ErrMissPK {
		fmt.Println("No primary key found.")
	} else {
		fmt.Println(num3)
		fmt.Println(authors)
	}
	controller.Data["authors"] = authors

}

// BookAdd controller
func (controller *MainController) BookAdd() {

	loadFormData(controller)
	controller.TplName = "BookAdd.html"
}

// BookEdit controller
func (controller *MainController) BookEdit() {

	loadFormData(controller)

	id, err := controller.GetInt("id")
	o := orm.NewOrm()
	book := models.Book{ID: id}
	err = o.Read(&book)
	if err == orm.ErrNoRows {
		fmt.Println("No book found.")
	} else if err == orm.ErrMissPK {
		fmt.Println("No primary key found.")
	} else {
		fmt.Println(book.ID, book.Title)
	}

	// book := new(models.Book)
	// book.ID = id
	// err := o.Read(&book)

	controller.Data["ID"] = book.ID
	controller.Data["Title"] = book.Title
	controller.Data["categoryId"] = book.Category.ID
	controller.Data["publisherId"] = book.Publisher.ID
	controller.Data["authorId"] = book.Author.ID

	controller.TplName = "BookEdit.html"

}

// BookSave controller
func (controller *MainController) BookSave() {

	id, err := controller.GetInt("id")
	if err == nil {
		fmt.Println(id)
	}
	title := controller.GetString("title")
	categoryID, errCat := controller.GetInt("category")
	if errCat == nil {
		fmt.Println(categoryID)
	}
	publisherID, errPub := controller.GetInt("publisher")
	if errPub == nil {
		fmt.Println(publisherID)
	}
	authorID, errAut := controller.GetInt("author")
	if errAut == nil {
		fmt.Println(authorID)
	}

	o := orm.NewOrm()
	category := models.Category{ID: categoryID}
	readErr := o.Read(&category)
	if readErr == nil {
		fmt.Println(category)
	}
	publisher := models.Publisher{ID: publisherID}
	readErr2 := o.Read(&publisher)
	if readErr2 == nil {
		fmt.Println(publisher)
	}
	author := models.Author{ID: authorID}
	readErr3 := o.Read(&author)
	if readErr3 == nil {
		fmt.Println(author)
	}

	if id == 0 {

		o := orm.NewOrm()
		book := new(models.Book)
		book.Title = title
		book.Category = &category
		book.Publisher = &publisher
		book.Author = &author
		o.Insert(book)

	} else {

		o := orm.NewOrm()
		book := models.Book{ID: id}
		err := o.Read(&book)
		if err == orm.ErrNoRows {
			fmt.Println("No book found.")
		} else if err == orm.ErrMissPK {
			fmt.Println("No primary key found.")
		} else {
			fmt.Println(book.ID, book.Title)
		}

		book.Title = title
		book.Category = &category
		book.Publisher = &publisher
		book.Author = &author

		o.Update(book)

	}

	controller.Redirect("/book/", 302)

}

// func (this *MainController) BookUpdate() {
// 	this.activeContent("logout")
// 	this.DelSession("acme")
// 	this.Redirect("/home", 302)
// }

// func (this *MainController) BookDelete() {
// 	this.activeContent("book/delete")
// 	session := this.GetSession("book")
// 	if session == nil {
// 		this.Redirect("/book/", 302)
// 		return
// 	}
// 	this.DelSession("acme")

// 	_, err = o.Delete(&user)
// 	if err == nil {
// 		flash.Notice("Your account is deleted.")
// 		flash.Store(&this.Controller)
// 		this.DelSession("acme")
// 		this.Redirect("/notice", 302)
// 	} else {
// 		flash.Error("Internal error")
// 		flash.Store(&this.Controller)
// 		return
// 	}

// 	this.Redirect("/book", 302)
// }
