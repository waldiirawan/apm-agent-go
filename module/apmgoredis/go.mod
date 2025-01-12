module github.com/waldiirawan/apm-agent-go/module/apmgoredis/v2

go 1.15

require (
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/stretchr/testify v1.6.1
	github.com/waldiirawan/apm-agent-go/v2 v2.2.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/waldiirawan/apm-agent-go/v2 => ../..
