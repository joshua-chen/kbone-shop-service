module product

go 1.14

require (
	commons/config v0.0.0
	github.com/Joker/jade v1.0.0 // indirect
	github.com/Shopify/goreferrer v0.0.0-20181106222321-ec9c9a553398 // indirect
	github.com/ajg/form v1.5.1 // indirect
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/aymerick/raymond v2.0.2+incompatible // indirect
	github.com/eknkc/amber v0.0.0-20171010120322-cdade1c07385 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/flosch/pongo2 v0.0.0-20200509134334-76fc00043fe1 // indirect
	github.com/gavv/monotime v0.0.0-20190418164738-30dba4353424 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-openapi/spec v0.19.8 // indirect
	github.com/go-openapi/swag v0.19.9 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/go-xorm/xorm v0.7.9
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/gorilla/schema v1.1.0 // indirect
	github.com/imkira/go-interpol v1.1.0 // indirect
	github.com/iris-contrib/blackfriday v2.0.0+incompatible // indirect
	github.com/iris-contrib/formBinder v5.0.0+incompatible // indirect
	github.com/iris-contrib/go.uuid v2.0.0+incompatible // indirect
	github.com/iris-contrib/httpexpect v1.1.2 // indirect
	github.com/iris-contrib/swagger v0.0.0-20190414182803-dc27bb5ee4ec
	github.com/juju/errors v0.0.0-20200330140219-3fe23663418f // indirect
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/kataras/golog v0.0.13 // indirect
	github.com/kataras/iris v11.1.1+incompatible
	github.com/klauspost/compress v1.10.5 // indirect
	github.com/mailru/easyjson v0.7.1 // indirect
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/microcosm-cc/bluemonday v1.0.2 // indirect
	github.com/onsi/ginkgo v1.12.0 // indirect
	github.com/onsi/gomega v1.10.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/ryanuber/columnize v2.1.0+incompatible // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.5
	github.com/valyala/fasthttp v1.12.0 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/yalp/jsonpath v0.0.0-20180802001716-5cc68e5049a0 // indirect
	github.com/yudai/gojsondiff v1.0.0 // indirect
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	github.com/yudai/pp v2.0.1+incompatible // indirect
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37 // indirect
	golang.org/x/net v0.0.0-20200513185701-a91f0712d120 // indirect
	golang.org/x/sys v0.0.0-20200515095857-1151b9dac4a9 // indirect
	golang.org/x/tools v0.0.0-20200515220128-d3bf790afa53 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	moul.io/http2curl v1.0.0 // indirect
	product/controllers v0.0.0
	product/datamodels v0.0.0 // indirect
	product/docs v0.0.0
	product/repositories v0.0.0-00010101000000-000000000000 // indirect
	product/services v0.0.0 // indirect
)

replace commons/config => ../../commons/config

replace product/datamodels => ../product/datamodels

replace product/repositories => ../product/repositories

replace product/services => ../product/services

replace product/controllers => ../product/controllers

replace product/docs => ../product/docs
