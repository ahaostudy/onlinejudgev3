.PHONY: gen-rpc
gen-rpc:
	mkdir -p service/${svc} && kitex -module github.com/ahaostudy/onlinejudge idl/${svc}.thrift && cd service/${svc} && kitex -module github.com/ahaostudy/onlinejudge -service ${svc}service -use github.com/ahaostudy/onlinejudge/kitex_gen ../../idl/${svc}.thrift
