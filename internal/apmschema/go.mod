module github.com/waldiirawan/apm-agent-go/internal/apmschema/v2

go 1.15

require (
	github.com/pkg/errors v0.9.1
	github.com/santhosh-tekuri/jsonschema v1.2.4
	github.com/stretchr/testify v1.8.1
	github.com/waldiirawan/apm-agent-go/v2 v2.2.0
)

replace github.com/waldiirawan/apm-agent-go/v2 => ../../
