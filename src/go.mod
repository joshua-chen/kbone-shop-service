module shop

go 1.14

require (
	commons/config v0.0.0
	commons/datasource v0.0.0-00010101000000-000000000000 // indirect
	commons/middleware/cors v0.0.0-00010101000000-000000000000
	commons/middleware/jwt v0.0.0-00010101000000-000000000000
	commons/middleware/jwt/route v0.0.0-00010101000000-000000000000 // indirect
	commons/mvc v0.0.0-00010101000000-000000000000
	commons/mvc/models v0.0.0-00010101000000-000000000000 // indirect
	commons/mvc/recover v0.0.0-00010101000000-000000000000
	commons/mvc/response v0.0.0-00010101000000-000000000000 // indirect
	commons/utils/yaml v0.0.0-00010101000000-000000000000 // indirect
	github.com/betacraft/yaag v1.0.0
	github.com/go-openapi/spec v0.19.8 // indirect
	github.com/go-openapi/swag v0.19.9 // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/go-xorm/core v0.6.3 // indirect
	github.com/go-xorm/xorm v0.7.9 // indirect
	github.com/gorilla/schema v1.1.0 // indirect
	github.com/iris-contrib/formBinder v5.0.0+incompatible // indirect
	github.com/iris-contrib/swagger/v12 v12.0.1
	github.com/kataras/iris/v12 v12.1.8
	github.com/klauspost/compress v1.10.5 // indirect
	github.com/mailru/easyjson v0.7.1 // indirect
	github.com/nats-io/nats-server/v2 v2.1.7 // indirect
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37 // indirect
	golang.org/x/tools v0.0.0-20200515220128-d3bf790afa53 // indirect
	shop/datamodels v0.0.0 // indirect
	shop/docs v0.0.0
	shop/repositories v0.0.0-00010101000000-000000000000 // indirect
	shop/routes v0.0.0
	shop/services v0.0.0 // indirect
)

replace commons/config => ../commons/config

replace commons/datasource => ../commons/datasource

replace commons/utils/yaml => ../commons/utils/yaml

replace commons/mvc/models => ../commons/mvc/models

replace commons/mvc/recover => ../commons/mvc/recover

replace commons/mvc/response => ../commons/mvc/response

replace commons/mvc => ../commons/mvc

replace commons/middleware/jwt => ../commons/middleware/jwt

replace commons/middleware/jwt/route => ../commons/middleware/jwt/route

replace commons/middleware/cors => ../commons/middleware/cors

replace shop/datamodels => ./datamodels

replace shop/repositories => ./repositories

replace shop/services => ./services

replace shop/middleware => ./middleware

replace shop/routes => ./routes

replace shop/docs => ./docs

replace github.com/go-xorm/core v0.6.3 => xorm.io/core v0.6.3
