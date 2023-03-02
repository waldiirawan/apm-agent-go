module github.com/waldiirawan/apm-agent-go/module/apmbeego/v2

require (
	github.com/astaxie/beego v1.12.3
	github.com/stretchr/testify v1.7.0
	github.com/waldiirawan/apm-agent-go/module/apmhttp/v2 v2.2.0
	github.com/waldiirawan/apm-agent-go/module/apmsql/v2 v2.2.0
	github.com/waldiirawan/apm-agent-go/v2 v2.2.0
)

replace github.com/waldiirawan/apm-agent-go/v2 => ../..

replace github.com/waldiirawan/apm-agent-go/module/apmhttp/v2 => ../apmhttp

replace github.com/waldiirawan/apm-agent-go/module/apmsql/v2 => ../apmsql

go 1.15
