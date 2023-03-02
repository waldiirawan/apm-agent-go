module github.com/waldiirawan/apm-agent-go/module/apmchi/v2

require (
	github.com/go-chi/chi v1.5.1
	github.com/stretchr/testify v1.6.1
	github.com/waldiirawan/apm-agent-go/module/apmhttp/v2 v2.2.0
	github.com/waldiirawan/apm-agent-go/v2 v2.2.0
)

replace github.com/waldiirawan/apm-agent-go/v2 => ../..

replace github.com/waldiirawan/apm-agent-go/module/apmhttp/v2 => ../apmhttp

go 1.15
