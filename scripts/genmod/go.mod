module genmod

require (
	github.com/pkg/errors v0.9.1
	github.com/waldiirawan/apm-agent-go/v2 v2.1.0
)

replace github.com/waldiirawan/apm-agent-go/v2 => ../..

go 1.15
