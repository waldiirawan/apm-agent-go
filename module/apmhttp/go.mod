module github.com/waldiirawan/apm-agent-go/module/apmhttp/v2

require (
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.6.1
	github.com/waldiirawan/apm-agent-go/v2 v2.2.0
	golang.org/x/net v0.0.0-20211015210444-4f30a5c0130f
)

replace github.com/waldiirawan/apm-agent-go/v2 => ../..

go 1.15
