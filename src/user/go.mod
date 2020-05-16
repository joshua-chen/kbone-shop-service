module user

go 1.14

require (
	commons/config v0.0.0
	commons/mvc/models v0.0.0
	github.com/CloudyKit/fastprinter v0.0.0-20200109182630-33d98a066a53 // indirect
	github.com/ajg/form v1.5.1 // indirect
	github.com/betacraft/yaag v1.0.0
	github.com/fasthttp-contrib/websocket v0.0.0-20160511215533-1f3b11f56072 // indirect
	github.com/flosch/pongo2 v0.0.0-20200509134334-76fc00043fe1 // indirect
	github.com/gavv/monotime v0.0.0-20190418164738-30dba4353424 // indirect
	github.com/go-openapi/spec v0.19.8 // indirect
	github.com/go-openapi/swag v0.19.9 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/go-xorm/xorm v0.7.9
	github.com/gorilla/schema v1.1.0 // indirect
	github.com/iris-contrib/formBinder v5.0.0+incompatible // indirect
	github.com/iris-contrib/httpexpect v1.1.2 // indirect
	github.com/iris-contrib/middleware/cors v0.0.0-20191219204441-78279b78a367
	github.com/iris-contrib/middleware/jwt v0.0.0-20191219204441-78279b78a367 // indirect
	github.com/iris-contrib/swagger v0.0.0-20190414182803-dc27bb5ee4ec
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/juju/errors v0.0.0-20200330140219-3fe23663418f // indirect
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/kataras/golog v0.0.13 // indirect
	github.com/kataras/iris v12.1.2
	github.com/klauspost/compress v1.10.5 // indirect
	github.com/mailru/easyjson v0.7.1 // indirect
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/onsi/ginkgo v1.12.0 // indirect
	github.com/onsi/gomega v1.10.0 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/swaggo/swag v1.6.5 // indirect
	github.com/valyala/fasthttp v1.12.0 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/yudai/pp v2.0.1+incompatible // indirect
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37 // indirect
	golang.org/x/net v0.0.0-20200513185701-a91f0712d120 // indirect
	golang.org/x/sys v0.0.0-20200515095857-1151b9dac4a9 // indirect
	golang.org/x/tools v0.0.0-20200515220128-d3bf790afa53 // indirect
	google.golang.org/protobuf v1.22.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	moul.io/http2curl v1.0.0 // indirect
	user/controllers v0.0.0
	user/datamodels v0.0.0 // indirect
	user/docs v0.0.0
	user/middleware v0.0.0-00010101000000-000000000000
	user/repositories v0.0.0-00010101000000-000000000000 // indirect
	user/services v0.0.0 // indirect
)

replace commons/config => ../../commons/config

replace commons/mvc/models => ../../commons/mvc/models

replace user/datamodels => ../user/datamodels

replace user/repositories => ../user/repositories

replace user/services => ../user/services

replace user/middleware => ../user/middleware

replace user/controllers => ../user/controllers

replace user/docs => ../user/docs
