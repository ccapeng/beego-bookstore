package routers

import (
	"bookstore/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/book/", &controllers.MainController{}, "get:BookList")
	beego.Router("/book/add/", &controllers.MainController{}, "get:BookAdd")
	beego.Router("/book/edit", &controllers.MainController{}, "get:BookEdit")
	beego.Router("/book/save/", &controllers.MainController{}, "post:BookSave")

	beego.Router("/category/", &controllers.MainController{}, "get:CategoryList")
	beego.Router("/category/add/", &controllers.MainController{}, "get:CategoryAdd")
	beego.Router("/category/edit", &controllers.MainController{}, "get:CategoryEdit")
	beego.Router("/category/save/", &controllers.MainController{}, "post:CategorySave")
	beego.Router("/category/delete", &controllers.MainController{}, "get:CategoryDelete")

	beego.Router("/publisher/", &controllers.MainController{}, "get:PublisherList")
	beego.Router("/publisher/add/", &controllers.MainController{}, "get:PublisherAdd")
	beego.Router("/publisher/edit", &controllers.MainController{}, "get:PublisherEdit")
	beego.Router("/publisher/save/", &controllers.MainController{}, "post:PublisherSave")
	beego.Router("/publisher/delete", &controllers.MainController{}, "get:PublisherDelete")

	beego.Router("/author/", &controllers.MainController{}, "get:AuthorList")
	beego.Router("/author/add/", &controllers.MainController{}, "get:AuthorAdd")
	beego.Router("/author/edit", &controllers.MainController{}, "get:AuthorEdit")
	beego.Router("/author/save/", &controllers.MainController{}, "post:AuthorSave")
	beego.Router("/author/delete", &controllers.MainController{}, "get:AuthorDelete")

	// beego.Router("/book/update/", &controllers.MainController{},
	// 	"get,post:BookUpdate")
	// beego.Router("/book/delete/", &controllers.MainController{},
	// 	"get,post:BookDelete")

}
