module github.com/waldiirawan/apm-agent-go/module/apmzap/v2

require (
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.6.1
	github.com/waldiirawan/apm-agent-go/v2 v2.2.0
	go.uber.org/atomic v1.3.2 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.9.1
)

replace github.com/waldiirawan/apm-agent-go/v2 => ../..

go 1.15
