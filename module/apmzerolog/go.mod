module github.com/waldiirawan/apm-agent-go/module/apmzerolog/v2

require (
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.14.3
	github.com/stretchr/testify v1.6.1
	github.com/waldiirawan/apm-agent-go/v2 v2.2.0
)

replace github.com/waldiirawan/apm-agent-go/v2 => ../..

go 1.15
