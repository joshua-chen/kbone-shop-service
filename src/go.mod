module shop

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
	commons/middleware/jwt/route v0.0.0-00010101000000-000000000000 // indirect
	commons/middleware/models v0.0.0-00010101000000-000000000000 // indirect
	commons/middleware/recover v0.0.0-00010101000000-000000000000
	commons/mvc v0.0.0-00010101000000-000000000000 // indirect
	commons/mvc/application v0.0.0-00010101000000-000000000000
	commons/mvc/context/request v0.0.0-00010101000000-000000000000 // indirect
	commons/mvc/context/response v0.0.0-00010101000000-000000000000
	commons/mvc/context/response/msg v0.0.0-00010101000000-000000000000 // indirect
	commons/mvc/models v0.0.0-00010101000000-000000000000 // indirect
	commons/utils v0.0.0-00010101000000-000000000000 // indirect
	commons/utils/security v0.0.0-00010101000000-000000000000 // indirect
	commons/utils/security/aes v0.0.0-00010101000000-000000000000 // indirect
	commons/utils/yaml v0.0.0-00010101000000-000000000000 // indirect
	github.com/Chronokeeper/anyxml v0.0.0-20160530174208-54457d8e98c6 // indirect
	github.com/agrison/go-tablib v0.0.0-20160310143025-4930582c22ee // indirect
	github.com/agrison/mxj v0.0.0-20160310142625-1269f8afb3b4 // indirect
	github.com/betacraft/yaag v1.0.0
	github.com/bndr/gotabulate v1.1.2 // indirect
	github.com/casbin/casbin v1.9.1 // indirect
	github.com/clbanning/mxj v1.8.4 // indirect
	github.com/go-openapi/spec v0.19.8 // indirect
	github.com/go-openapi/swag v0.19.9 // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/iris-contrib/swagger/v12 v12.0.1
	github.com/kataras/iris/v12 v12.1.8
	github.com/klauspost/compress v1.10.5 // indirect
	github.com/lib/pq v1.5.2 // indirect
	github.com/mailru/easyjson v0.7.1 // indirect
	github.com/nats-io/nats-server/v2 v2.1.7 // indirect
	github.com/syndtr/goleveldb v1.0.0 // indirect
	github.com/tealeg/xlsx v1.0.5 // indirect
	github.com/xormplus/builder v0.0.0-20200331055651-240ff40009be // indirect
	github.com/xormplus/core v0.6.3 // indirect
	github.com/xormplus/xorm v0.0.0-20200514184607-0f37421d8714 // indirect
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37 // indirect
	golang.org/x/sys v0.0.0-20200523222454-059865788121 // indirect
	golang.org/x/tools v0.0.0-20200515220128-d3bf790afa53 // indirect
	gopkg.in/flosch/pongo2.v3 v3.0.0-20141028000813-5e81b817a0c4 // indirect
	shop/datamodels v0.0.0 // indirect
	shop/docs v0.0.0
	shop/repositories v0.0.0-00010101000000-000000000000 // indirect
	shop/routes v0.0.0
	shop/routes/products v0.0.0-000101010000go00-000000000000 // indirect
	shop/routes/users v0.0.0-00010101000000-000000000000 // indirect
	shop/services v0.0.0 // indirect
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

replace commons/middleware/jwt/route => ../commons/middleware/jwt/route

replace commons/middleware/cors => ../commons/middleware/cors

replace shop/datamodels => ./datamodels

replace shop/repositories => ./repositories

replace shop/services => ./services

replace shop/middleware => ./middleware

replace shop/routes => ./routes

replace shop/routes/users => ./routes/users

replace shop/routes/products => ./routes/products

replace shop/docs => ./docs

replace github.com/xormplus/core v0.6.3 => xorm.io/core v0.6.3
