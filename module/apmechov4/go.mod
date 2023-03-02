module github.com/waldiirawan/apm-agent-go/module/apmechov4/v2

require (
	github.com/labstack/echo/v4 v4.6.1
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.6.1
	github.com/waldiirawan/apm-agent-go/module/apmhttp/v2 v2.2.0
	github.com/waldiirawan/apm-agent-go/v2 v2.2.0
)

replace github.com/waldiirawan/apm-agent-go/v2 => ../..

replace github.com/waldiirawan/apm-agent-go/module/apmhttp/v2 => ../apmhttp

go 1.15
