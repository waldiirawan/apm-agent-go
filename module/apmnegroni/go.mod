module github.com/waldiirawan/apm-agent-go/module/apmnegroni/v2

go 1.15

require (
	github.com/stretchr/testify v1.6.1
	github.com/urfave/negroni v1.0.0
	github.com/waldiirawan/apm-agent-go/module/apmhttp/v2 v2.2.0
	github.com/waldiirawan/apm-agent-go/v2 v2.2.0
)

replace github.com/waldiirawan/apm-agent-go/v2 => ../..

replace github.com/waldiirawan/apm-agent-go/module/apmhttp/v2 => ../apmhttp
