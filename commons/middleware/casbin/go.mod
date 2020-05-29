module commons/middleware/casbin

go 1.14

require (
	commons/config v0.0.0-00010101000000-000000000000
	commons/datasource v0.0.0-00010101000000-000000000000
	commons/middleware/jwt v0.0.0-00010101000000-000000000000
	commons/middleware/models v0.0.0-00010101000000-000000000000
	commons/mvc/context/request v0.0.0-00010101000000-000000000000 // indirect
	commons/mvc/context/response v0.0.0-00010101000000-000000000000
	commons/mvc/context/response/msg v0.0.0-00010101000000-000000000000
	commons/utils/security v0.0.0-00010101000000-000000000000 // indirect
	commons/utils/security/aes v0.0.0-00010101000000-000000000000
	github.com/ajg/form v1.5.1 // indirect
	github.com/casbin/casbin v1.9.1
	github.com/fasthttp-contrib/websocket v0.0.0-20160511215533-1f3b11f56072 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/imkira/go-interpol v1.1.0 // indirect
	github.com/jmespath/go-jmespath v0.3.0 // indirect
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/kataras/golog v0.0.15
	github.com/kataras/iris/v12 v12.1.8
	github.com/lib/pq v1.5.2
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/moul/http2curl v1.0.0 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/valyala/fasthttp v1.13.1 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/xormplus/xorm v0.0.0-20200529061552-7d0d26c6f81c
	github.com/yalp/jsonpath v0.0.0-20180802001716-5cc68e5049a0 // indirect
	github.com/yudai/gojsondiff v1.0.0 // indirect
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	github.com/yudai/pp v2.0.1+incompatible // indirect
)

replace commons/config => ../../config

replace commons/datasource => ../../datasource

replace commons/mvc/context => ../../mvc/context

replace commons/mvc/context/request => ../../mvc/context/request

replace commons/mvc/context/response => ../../mvc/context/response

replace commons/mvc/context/response/msg => ../../mvc/context/response/msg

replace commons/mvc/models => ../../mvc/models

replace commons/utils => ../../utils

replace commons/utils/security/aes => ../../utils/security/aes

replace commons/utils/security => ../../utils/security

replace commons/middleware/jwt => ../jwt

replace commons/middleware/casbin => ../casbin

replace commons/middleware/models => ../models

replace commons/middleware/auth => ../auth

replace commons/exception => ../../exception

replace commons/utils/yaml => ../../utils/yaml
