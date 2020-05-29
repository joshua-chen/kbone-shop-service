module shop/wap

go 1.14

require (
	commons/config v0.0.0
	commons/datasource v0.0.0-00010101000000-000000000000 // indirect
	commons/exception v0.0.0-00010101000000-000000000000 // indirect
	commons/middleware v0.0.0-00010101000000-000000000000
	commons/middleware/auth v0.0.0-00010101000000-000000000000 // indirect
	commons/middleware/casbin v0.0.0-00010101000000-000000000000 // indirect
	commons/middleware/cors v0.0.0-00010101000000-000000000000
	commons/middleware/jwt v0.0.0-00010101000000-000000000000 // indirect
	commons/middleware/models v0.0.0-00010101000000-000000000000 // indirect
	commons/middleware/recover v0.0.0-00010101000000-000000000000
	commons/mvc v0.0.0-00010101000000-000000000000 // indirect
	commons/mvc/application v0.0.0-00010101000000-000000000000
	commons/mvc/context v0.0.0-00010101000000-000000000000 // indirect
	commons/mvc/context/request v0.0.0-00010101000000-000000000000 // indirect
	commons/mvc/context/response v0.0.0-00010101000000-000000000000
	commons/mvc/context/response/msg v0.0.0-00010101000000-000000000000 // indirect
	commons/mvc/models v0.0.0-00010101000000-000000000000 // indirect
	commons/utils v0.0.0-00010101000000-000000000000 // indirect
	commons/utils/security v0.0.0-00010101000000-000000000000 // indirect
	commons/utils/security/aes v0.0.0-00010101000000-000000000000 // indirect
	commons/utils/yaml v0.0.0-00010101000000-000000000000 // indirect
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/betacraft/yaag v1.0.0
	github.com/casbin/casbin v1.9.1 // indirect
	github.com/clbanning/mxj v1.8.4 // indirect
	github.com/go-openapi/spec v0.19.8 // indirect
	github.com/go-openapi/swag v0.19.9 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/iris-contrib/swagger/v12 v12.0.1
	github.com/kataras/iris/v12 v12.1.8
	github.com/klauspost/compress v1.10.5 // indirect
	github.com/lib/pq v1.5.2 // indirect
	github.com/mailru/easyjson v0.7.1 // indirect
	github.com/nats-io/nats-server/v2 v2.1.7 // indirect
	github.com/swaggo/swag v1.6.5
	github.com/xormplus/core v0.6.3 // indirect
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37 // indirect
	golang.org/x/sys v0.0.0-20200523222454-059865788121 // indirect
	golang.org/x/tools v0.0.0-20200515220128-d3bf790afa53 // indirect
	shop/datamodels v0.0.0 // indirect
	shop/repositories v0.0.0-00010101000000-000000000000 // indirect
	shop/services v0.0.0 // indirect
	shop/wap/routes v0.0.0-00010101000000-000000000000
	shop/wap/routes/products v0.0.0-00010101000000-000000000000 // indirect
)

replace commons/mvc/application => ../../commons/mvc/application

replace commons/config => ../../commons/config

replace commons/datasource => ../../commons/datasource

replace commons/utils => ../../commons/utils

replace commons/utils/yaml => ../../commons/utils/yaml

replace commons/utils/security => ../../commons/utils/security

replace commons/utils/security/aes => ../../commons/utils/security/aes

replace commons/exception => ../../commons/exception

replace commons/mvc/models => ../../commons/mvc/models

replace commons/mvc/context => ../../commons/mvc/context

replace commons/mvc/context/response => ../../commons/mvc/context/response

replace commons/mvc/context/response/msg => ../../commons/mvc/context/response/msg

replace commons/mvc/context/request => ../../commons/mvc/context/request

replace commons/mvc => ../../commons/mvc

replace commons/middleware => ../../commons/middleware

replace commons/middleware/auth => ../../commons/middleware/auth

replace commons/middleware/models => ../../commons/middleware/models

replace commons/middleware/casbin => ../../commons/middleware/casbin

replace commons/middleware/jwt => ../../commons/middleware/jwt

replace commons/middleware/recover => ../../commons/middleware/recover

replace commons/mvc/route => ../../commons/mvc/route

replace commons/middleware/cors => ../../commons/middleware/cors

replace shop/datamodels => ../datamodels

replace shop/repositories => ../repositories

replace shop/services => ../services

replace shop/middleware => ../middleware

replace shop/wap/docs => ../docs

replace shop/wap/routes => ./routes

replace shop/wap/routes/products => ./routes/products

replace github.com/xormplus/core v0.6.3 => xorm.io/core v0.6.3
