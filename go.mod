module admin

go 1.12

require (
	github.com/DataDog/zstd v1.4.4 // indirect
	github.com/Joker/jade v1.0.0 // indirect
	github.com/Shopify/goreferrer v0.0.0-20181106222321-ec9c9a553398 // indirect
	github.com/ajg/form v1.5.1 // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v1.60.315 // indirect
	github.com/apache/thrift v0.13.0
	github.com/aymerick/raymond v2.0.2+incompatible // indirect
	github.com/buger/jsonparser v0.0.0-20191204142016-1a29609e0929 // indirect
	github.com/bwmarrin/snowflake v0.3.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/eknkc/amber v0.0.0-20171010120322-cdade1c07385 // indirect
	github.com/fastly/go-utils v0.0.0-20180712184237-d95a45783239 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/flosch/pongo2 v0.0.0-20190707114632-bbf5a6c351f4 // indirect
	github.com/gavv/monotime v0.0.0-20190418164738-30dba4353424 // indirect
	github.com/go-errors/errors v1.0.1 // indirect
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/google/wire v0.4.0
	github.com/gorilla/schema v1.1.0 // indirect
	github.com/imkira/go-interpol v1.1.0 // indirect
	github.com/iris-contrib/blackfriday v2.0.0+incompatible // indirect
	github.com/iris-contrib/formBinder v5.0.0+incompatible // indirect
	github.com/iris-contrib/go.uuid v2.0.0+incompatible // indirect
	github.com/iris-contrib/httpexpect v0.0.0-20180314041918-ebe99fcebbce // indirect
	github.com/jehiah/go-strftime v0.0.0-20171201141054-1d33003b3869 // indirect
	github.com/jinzhu/gorm v1.9.11
	github.com/jonboulle/clockwork v0.1.0 // indirect
	github.com/json-iterator/go v1.1.7 // indirect
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/kataras/golog v0.0.0-20190624001437-99c81de45f40 // indirect
	github.com/kataras/iris v11.1.1+incompatible
	github.com/kataras/pio v0.0.0-20190103105442-ea782b38602d // indirect
	github.com/klauspost/compress v1.9.5 // indirect
	github.com/lestrrat/go-envload v0.0.0-20180220120943-6ed08b54a570 // indirect
	github.com/lestrrat/go-file-rotatelogs v0.0.0-20180223000712-d3151e2a480f // indirect
	github.com/lestrrat/go-strftime v0.0.0-20180220042222-ba3bf9c1d042 // indirect
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/microcosm-cc/bluemonday v1.0.2 // indirect
	github.com/moul/http2curl v1.0.0 // indirect
	github.com/nacos-group/nacos-sdk-go v0.0.0-20191128082542-fe1b325b125c
	github.com/ryanuber/columnize v2.1.0+incompatible // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/tebeka/strftime v0.1.3 // indirect
	github.com/tidwall/pretty v1.0.0 // indirect
	github.com/toolkits/concurrent v0.0.0-20150624120057-a4371d70e3e3 // indirect
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/yalp/jsonpath v0.0.0-20180802001716-5cc68e5049a0 // indirect
	github.com/yudai/gojsondiff v1.0.0 // indirect
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	github.com/yudai/pp v2.0.1+incompatible // indirect
	go.mongodb.org/mongo-driver v1.2.0
	gopkg.in/yaml.v2 v2.2.4
)

replace (
	github.com/mongodb/mongo-go-driver => go.mongodb.org/mongo-driver v1.2.0
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190605123033-f99c8df09eb5
	golang.org/x/net => github.com/golang/net v0.0.0-20190603091049-60506f45cf65
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190602015325-4c4f7f33c9ed
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190606050223-4d9ae51c2468
	google.golang.org/grpc => github.com/grpc/grpc-go v1.21.1
)
