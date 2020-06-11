module shop

go 1.14

require (
	github.com/betacraft/yaag v1.0.0
	github.com/go-openapi/spec v0.19.8 // indirect
	github.com/go-openapi/swag v0.19.9 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/iris-contrib/swagger/v12 v12.0.1
	github.com/joshua-chen/go-commons v1.0.21
	github.com/kataras/iris/v12 v12.1.8
	github.com/klauspost/compress v1.10.5 // indirect
	github.com/mailru/easyjson v0.7.1 // indirect
	github.com/xormplus/core v0.6.3 // indirect
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37 // indirect
	golang.org/x/tools v0.0.0-20200515220128-d3bf790afa53 // indirect
	shop/datamodels v0.0.0-00010101000000-000000000000 // indirect
	shop/docs v0.0.0-00010101000000-000000000000
	shop/repositories v0.0.0-00010101000000-000000000000 // indirect
	shop/routes v0.0.0-00010101000000-000000000000
	shop/routes/products v0.0.0-00010101000000-000000000000 // indirect
	shop/routes/sessions v0.0.0-00010101000000-000000000000 // indirect
	shop/routes/users v0.0.0-00010101000000-000000000000 // indirect
	shop/services v0.0.0-00010101000000-000000000000 // indirect
	shop_wap/routes v0.0.0-00010101000000-000000000000 // indirect
	shop_wap/routes/products v0.0.0-00010101000000-000000000000 // indirect
	shop_web/routes v0.0.0-00010101000000-000000000000 // indirect
	shop_web/routes/products v0.0.0-00010101000000-000000000000 // indirect

)

replace shop/datamodels => ./datamodels

replace shop/repositories => ./repositories

replace shop/services => ./services

replace shop/middleware => ./middleware

replace shop/routes => ./routes

replace shop/routes/users => ./routes/users

replace shop/routes/products => ./routes/products

replace shop/routes/sessions => ./routes/sessions

replace shop/docs => ./docs

replace shop_wap/routes => ../shop_wap/routes

replace shop_wap/routes/products => ../shop_wap/routes/products

replace shop_web/routes => ../shop_web/routes

replace shop_web/routes/products => ../shop_web/routes/products

replace github.com/xormplus/core v0.6.3 => xorm.io/core v0.6.3
