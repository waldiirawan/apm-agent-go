module github.com/waldiirawan/apm-agent-go/module/apmgorm/v2

require (
	github.com/jinzhu/gorm v1.9.10
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.0
	github.com/waldiirawan/apm-agent-go/module/apmsql/v2 v2.2.0
	github.com/waldiirawan/apm-agent-go/v2 v2.2.0
)

replace github.com/waldiirawan/apm-agent-go/v2 => ../..

replace github.com/waldiirawan/apm-agent-go/module/apmsql/v2 => ../apmsql

go 1.15
