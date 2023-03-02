module github.com/waldiirawan/apm-agent-go/module/apmpgxv5/v2

go 1.15

require (
	github.com/jackc/pgx/v5 v5.0.4
	github.com/stretchr/testify v1.8.0
	github.com/waldiirawan/apm-agent-go/module/apmsql/v2 v2.2.0
	github.com/waldiirawan/apm-agent-go/v2 v2.2.0
)

replace github.com/waldiirawan/apm-agent-go/v2 => ../..

replace github.com/waldiirawan/apm-agent-go/module/apmsql/v2 => ../apmsql
