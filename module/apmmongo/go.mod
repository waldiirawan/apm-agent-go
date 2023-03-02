module github.com/waldiirawan/apm-agent-go/module/apmmongo/v2

require (
	github.com/stretchr/testify v1.6.1
	github.com/waldiirawan/apm-agent-go/v2 v2.2.0
	go.mongodb.org/mongo-driver v1.5.1
)

replace github.com/waldiirawan/apm-agent-go/v2 => ../..

go 1.15
