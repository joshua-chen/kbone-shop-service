module shop

go 1.14

require (
	github.com/Chronokeeper/anyxml v0.0.0-20160530174208-54457d8e98c6 // indirect
	github.com/agrison/go-tablib v0.0.0-20160310143025-4930582c22ee // indirect
	github.com/agrison/mxj v0.0.0-20160310142625-1269f8afb3b4 // indirect
	github.com/ajg/form v1.5.1 // indirect
	github.com/betacraft/yaag v1.0.0 // indirect
	github.com/bndr/gotabulate v1.1.2 // indirect
	github.com/casbin/casbin v1.9.1 // indirect
	github.com/clbanning/mxj v1.8.4 // indirect
	github.com/fasthttp-contrib/websocket v0.0.0-20160511215533-1f3b11f56072 // indirect
	github.com/flosch/pongo2 v0.0.0-20200518135938-dfb43dbdc22a // indirect
	github.com/gavv/monotime v0.0.0-20190418164738-30dba4353424 // indirect
	github.com/go-openapi/spec v0.19.8 // indirect
	github.com/go-openapi/swag v0.19.9 // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/gorilla/schema v1.1.0 // indirect
	github.com/iris-contrib/formBinder v5.0.0+incompatible // indirect
	github.com/iris-contrib/httpexpect v1.1.2 // indirect
	github.com/iris-contrib/middleware/cors v0.0.0-20191219204441-78279b78a367 // indirect
	github.com/iris-contrib/middleware/jwt v0.0.0-20191219204441-78279b78a367 // indirect
	github.com/iris-contrib/swagger/v12 v12.0.1 // indirect
	github.com/jmespath/go-jmespath v0.3.0 // indirect
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/kataras/golog v0.0.15 // indirect
	github.com/klauspost/compress v1.10.5 // indirect
	github.com/lib/pq v1.5.2 // indirect
	github.com/mailru/easyjson v0.7.1 // indirect
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/nats-io/nats-server/v2 v2.1.7 // indirect
	github.com/onsi/ginkgo v1.12.2 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/syndtr/goleveldb v1.0.0 // indirect
	github.com/tealeg/xlsx v1.0.5 // indirect
	github.com/valyala/fasthttp v1.13.1 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/xormplus/builder v0.0.0-20200331055651-240ff40009be // indirect
	github.com/xormplus/core v0.6.3 // indirect
	github.com/xormplus/xorm v0.0.0-20200529061552-7d0d26c6f81c // indirect
	github.com/yudai/pp v2.0.1+incompatible // indirect
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37 // indirect
	golang.org/x/sys v0.0.0-20200523222454-059865788121 // indirect
	golang.org/x/tools v0.0.0-20200515220128-d3bf790afa53 // indirect
	gopkg.in/flosch/pongo2.v3 v3.0.0-20141028000813-5e81b817a0c4 // indirect
	moul.io/http2curl v1.0.0 // indirect
)

replace commons/mvc/application => ../commons/mvc/application

replace commons/config => ../commons/config

replace commons/datasource => ../commons/datasource

replace commons/utils => ../commons/utils

replace commons/utils/yaml => ../commons/utils/yaml

replace commons/utils/security => ../commons/utils/security

replace commons/utils/security/aes => ../commons/utils/security/aes

replace commons/exception => ../commons/exception

replace commons/mvc/models => ../commons/mvc/models

replace commons/mvc/route => ../commons/mvc/route

replace commons/mvc/context => ../commons/mvc/context

replace commons/mvc/context/response => ../commons/mvc/context/response

replace commons/mvc/context/response/msg => ../commons/mvc/context/response/msg

replace commons/mvc/context/request => ../commons/mvc/context/request

replace commons/mvc => ../commons/mvc

replace commons/middleware => ../commons/middleware

replace commons/middleware/auth => ../commons/middleware/auth

replace commons/middleware/models => ../commons/middleware/models

replace commons/middleware/casbin => ../commons/middleware/casbin

replace commons/middleware/jwt => ../commons/middleware/jwt

replace commons/middleware/recover => ../commons/middleware/recover

replace commons/middleware/cors => ../commons/middleware/cors

replace shop/datamodels => ./datamodels

replace shop/repositories => ./repositories

replace shop/services => ./services

replace shop/middleware => ./middleware

replace shop/routes => ./routes

replace shop/routes/users => ./routes/users

replace shop/routes/products => ./routes/products

replace shop/docs => ./docs

replace shop/wap/routes => ../shop_wap/routes

replace shop/wap/routes/products => ../shop_wap/routes/products

replace shop/web/routes => ../shop_wap/routes

replace shop/web/routes/products => ../shop_wap/routes/products

replace github.com/xormplus/core v0.6.3 => xorm.io/core v0.6.3
