module github.com/waldiirawan/apm-agent-go/module/apmfiber/v2

require (
	github.com/gofiber/fiber/v2 v2.18.0
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.0
	github.com/valyala/fasthttp v1.34.0
	github.com/waldiirawan/apm-agent-go/module/apmfasthttp/v2 v2.2.0
	github.com/waldiirawan/apm-agent-go/module/apmhttp/v2 v2.2.0
	github.com/waldiirawan/apm-agent-go/v2 v2.2.0
)

replace github.com/waldiirawan/apm-agent-go/v2 => ../..

replace github.com/waldiirawan/apm-agent-go/module/apmhttp/v2 => ../apmhttp

replace github.com/waldiirawan/apm-agent-go/module/apmfasthttp/v2 => ../apmfasthttp

go 1.15
