module github.com/waldiirawan/apm-agent-go/module/apmgometrics/v2

require (
	github.com/rcrowley/go-metrics v0.0.0-20181016184325-3113b8401b8a
	github.com/stretchr/testify v1.6.1
	github.com/waldiirawan/apm-agent-go/v2 v2.2.0
)

replace github.com/waldiirawan/apm-agent-go/v2 => ../..

go 1.15
