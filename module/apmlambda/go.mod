module github.com/waldiirawan/apm-agent-go/module/apmlambda/v2

require (
	github.com/aws/aws-lambda-go v1.8.0
	github.com/waldiirawan/apm-agent-go/v2 v2.2.0
)

replace github.com/waldiirawan/apm-agent-go/v2 => ../..

go 1.15
