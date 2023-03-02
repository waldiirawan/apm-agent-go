module github.com/waldiirawan/apm-agent-go/module/apmawssdkgo/v2

go 1.15

require (
	github.com/aws/aws-sdk-go v1.38.14
	github.com/stretchr/testify v1.7.0
	github.com/waldiirawan/apm-agent-go/module/apmhttp/v2 v2.2.0
	github.com/waldiirawan/apm-agent-go/v2 v2.2.0
)

replace github.com/waldiirawan/apm-agent-go/v2 => ../..

replace github.com/waldiirawan/apm-agent-go/module/apmhttp/v2 => ../apmhttp
