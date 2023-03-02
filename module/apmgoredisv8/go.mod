module github.com/waldiirawan/apm-agent-go/module/apmgoredisv8/v2

go 1.15

require (
	github.com/go-redis/redis/v8 v8.11.4
	github.com/stretchr/testify v1.7.0
	github.com/waldiirawan/apm-agent-go/v2 v2.2.0
)

replace github.com/waldiirawan/apm-agent-go/v2 => ../..
