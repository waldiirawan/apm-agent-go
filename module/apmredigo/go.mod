module github.com/waldiirawan/apm-agent-go/module/apmredigo/v2

require (
	github.com/gomodule/redigo v1.8.2
	github.com/stretchr/testify v1.6.1
	github.com/waldiirawan/apm-agent-go/v2 v2.2.0
)

replace github.com/waldiirawan/apm-agent-go/v2 => ../..

go 1.15
