module github.com/waldiirawan/apm-agent-go/module/apmlogrus/v2

require (
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.2.0
	github.com/stretchr/testify v1.6.1
	github.com/waldiirawan/apm-agent-go/v2 v2.2.0
	golang.org/x/crypto v0.0.0-20191206172530-e9b2fee46413 // indirect
)

replace github.com/waldiirawan/apm-agent-go/v2 => ../..

go 1.15
