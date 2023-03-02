module tracecontexttest

require github.com/waldiirawan/apm-agent-go/module/apmhttp/v2 v2.0.0

replace github.com/waldiirawan/apm-agent-go/v2 => ../..

replace github.com/waldiirawan/apm-agent-go/module/apmhttp/v2 => ../../module/apmhttp

go 1.13
