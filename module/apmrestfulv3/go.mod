module github.com/waldiirawan/apm-agent-go/module/apmrestfulv3/v2

require (
	github.com/emicklei/go-restful/v3 v3.8.0
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/stretchr/testify v1.6.1
	github.com/waldiirawan/apm-agent-go/module/apmhttp/v2 v2.2.0
	github.com/waldiirawan/apm-agent-go/v2 v2.2.0
)

replace github.com/waldiirawan/apm-agent-go/v2 => ../..

replace github.com/waldiirawan/apm-agent-go/module/apmhttp/v2 => ../apmhttp

go 1.15
