module github.com/waldiirawan/apm-agent-go/module/apmgorilla/v2

require (
	github.com/gorilla/context v1.1.1 // indirect
	github.com/gorilla/mux v1.6.2
	github.com/stretchr/testify v1.6.1
	github.com/waldiirawan/apm-agent-go/module/apmhttp/v2 v2.2.0
	github.com/waldiirawan/apm-agent-go/v2 v2.2.0
)

replace github.com/waldiirawan/apm-agent-go/v2 => ../..

replace github.com/waldiirawan/apm-agent-go/module/apmhttp/v2 => ../apmhttp

go 1.15
