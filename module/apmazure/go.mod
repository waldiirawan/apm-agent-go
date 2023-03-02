module github.com/waldiirawan/apm-agent-go/module/apmazure/v2

go 1.15

require (
	github.com/Azure/azure-pipeline-go v0.2.3
	github.com/Azure/azure-storage-blob-go v0.14.0
	github.com/Azure/azure-storage-file-go v0.8.0
	github.com/Azure/azure-storage-queue-go v0.0.0-20191125232315-636801874cdd
	github.com/stretchr/testify v1.7.0
	github.com/waldiirawan/apm-agent-go/module/apmhttp/v2 v2.2.0
	github.com/waldiirawan/apm-agent-go/v2 v2.2.0
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897 // indirect
)

replace github.com/waldiirawan/apm-agent-go/v2 => ../..

replace github.com/waldiirawan/apm-agent-go/module/apmhttp/v2 => ../apmhttp
