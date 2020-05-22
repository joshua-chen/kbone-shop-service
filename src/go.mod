module shop

go 1.14

require (
	commons/config v0.0.0
	commons/middleware/cors v0.0.0-00010101000000-000000000000
	commons/middleware/jwt v0.0.0-00010101000000-000000000000
	commons/middleware/mvc/models v0.0.0-00010101000000-000000000000 // indirect
	commons/mvc v0.0.0-00010101000000-000000000000
	commons/utils/yaml v0.0.0-00010101000000-000000000000 // indirect
	github.com/ajg/form v1.5.1 // indirect
	github.com/betacraft/yaag v1.0.0
	github.com/fasthttp-contrib/websocket v0.0.0-20160511215533-1f3b11f56072 // indirect
	github.com/go-openapi/spec v0.19.8 // indirect
	github.com/go-openapi/swag v0.19.9 // indirect
	github.com/gorilla/schema v1.1.0 // indirect
	github.com/iris-contrib/formBinder v5.0.0+incompatible // indirect
	github.com/iris-contrib/swagger/v12 v12.0.1
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/kataras/iris/v12 v12.1.8
	github.com/klauspost/compress v1.10.5 // indirect
	github.com/mailru/easyjson v0.7.1 // indirect
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/nats-io/nats-server/v2 v2.1.7 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/valyala/fasthttp v1.12.0 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/yudai/pp v2.0.1+incompatible // indirect
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37 // indirect
	golang.org/x/tools v0.0.0-20200515220128-d3bf790afa53 // indirect
	shop/controllers v0.0.0
	shop/datamodels v0.0.0 // indirect
	shop/docs v0.0.0
	shop/repositories v0.0.0-00010101000000-000000000000 // indirect
	shop/services v0.0.0 // indirect
)

replace commons/config => ../commons/config

replace commons/utils/yaml => ../commons/utils/yaml

replace commons/mvc/models => ../commons/mvc/models

replace commons/mvc => ../commons/mvc

replace commons/middleware/mvc/models => ../commons/middleware/mvc/models

replace commons/middleware/jwt => ../commons/middleware/jwt

replace commons/middleware/jwt/route => ../commons/middleware/jwt/route

replace commons/middleware/cors => ../commons/middleware/cors

replace shop/datamodels => ./datamodels

replace shop/repositories => ./repositories

replace shop/services => ./services

replace shop/middleware => ./middleware

replace shop/route => ./route

replace shop/controllers => ./controllers

replace shop/docs => ./docs
