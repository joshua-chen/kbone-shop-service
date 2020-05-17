module user

go 1.14

require (
	commons/config v0.0.0
	commons/datamodels v0.0.0-00010101000000-000000000000 // indirect
	commons/middleware/jwt v0.0.0-00010101000000-000000000000
	commons/mvc/models v0.0.0
	github.com/ajg/form v1.5.1 // indirect
	github.com/betacraft/yaag v1.0.0
	github.com/fasthttp-contrib/websocket v0.0.0-20160511215533-1f3b11f56072 // indirect
	github.com/go-openapi/spec v0.19.8 // indirect
	github.com/go-openapi/swag v0.19.9 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/go-xorm/xorm v0.7.9
	github.com/gorilla/schema v1.1.0 // indirect
	github.com/iris-contrib/formBinder v5.0.0+incompatible // indirect
	github.com/iris-contrib/middleware/cors v0.0.0-20191219204441-78279b78a367
	github.com/iris-contrib/middleware/jwt v0.0.0-20191219204441-78279b78a367 // indirect
	github.com/iris-contrib/swagger/v12 v12.0.1
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/kataras/golog v0.0.13 // indirect
	github.com/kataras/iris/v12 v12.1.8
	github.com/klauspost/compress v1.10.5 // indirect
	github.com/mailru/easyjson v0.7.1 // indirect
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/nats-io/nats-server/v2 v2.1.7 // indirect
	github.com/onsi/ginkgo v1.12.0 // indirect
	github.com/onsi/gomega v1.10.0 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/valyala/fasthttp v1.12.0 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/yudai/pp v2.0.1+incompatible // indirect
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37 // indirect
	golang.org/x/net v0.0.0-20200513185701-a91f0712d120 // indirect
	golang.org/x/sys v0.0.0-20200515095857-1151b9dac4a9 // indirect
	golang.org/x/tools v0.0.0-20200515220128-d3bf790afa53 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	user/controllers v0.0.0
	user/datamodels v0.0.0 // indirect
	user/docs v0.0.0
	user/repositories v0.0.0-00010101000000-000000000000 // indirect
	user/services v0.0.0 // indirect
)

replace commons/config => ../../commons/config

replace commons/mvc/models => ../../commons/mvc/models

replace commons/middleware/jwt => ../../commons/middleware/jwt

replace commons/datamodels => ../../commons/datamodels

replace user/datamodels => ../user/datamodels

replace user/repositories => ../user/repositories

replace user/services => ../user/services

replace user/middleware => ../user/middleware

replace user/controllers => ../user/controllers

replace user/docs => ../user/docs
