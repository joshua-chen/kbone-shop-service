module shop_wap

go 1.14

require (
	github.com/flosch/pongo2 v0.0.0-20200518135938-dfb43dbdc22a // indirect
	github.com/gavv/monotime v0.0.0-20190418164738-30dba4353424 // indirect
	github.com/go-openapi/spec v0.19.8 // indirect
	github.com/go-openapi/swag v0.19.9 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/gorilla/schema v1.1.0 // indirect
	github.com/iris-contrib/formBinder v5.0.0+incompatible // indirect
	github.com/iris-contrib/httpexpect v1.1.2 // indirect
	github.com/iris-contrib/swagger/v12 v12.0.1
	github.com/joshua-chen/go-commons v1.0.21
	github.com/kataras/iris/v12 v12.1.8
	github.com/klauspost/compress v1.10.5 // indirect
	github.com/mailru/easyjson v0.7.1 // indirect
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37 // indirect
	golang.org/x/sys v0.0.0-20200523222454-059865788121 // indirect
	golang.org/x/tools v0.0.0-20200515220128-d3bf790afa53 // indirect
	moul.io/http2curl v1.0.0 // indirect
	shop/datamodels v0.0.0-00010101000000-000000000000 // indirect
	shop/repositories v0.0.0-00010101000000-000000000000 // indirect
	shop/services v0.0.0-00010101000000-000000000000 // indirect
	shop_wap/docs v0.0.0-00010101000000-000000000000
	shop_wap/routes v0.0.0-00010101000000-000000000000
	shop_wap/routes/products v0.0.0-00010101000000-000000000000 // indirect

)

replace shop/datamodels => ../shop/datamodels

replace shop/repositories => ../shop/repositories

replace shop/services => ../shop/services

replace shop/routes => ../shop/routes

replace shop_wap/docs => ./docs

replace shop_wap/routes => ./routes

replace shop_wap/routes/products => ./routes/products

replace github.com/xormplus/core v0.6.3 => xorm.io/core v0.6.3
