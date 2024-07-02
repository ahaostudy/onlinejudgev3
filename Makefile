.PHONY: gen-rpc
gen-rpc:
	mkdir -p app/${svc} && kitex -module github.com/ahaostudy/onlinejudge idl/${svc}.thrift && cd app/${svc} && kitex -module github.com/ahaostudy/onlinejudge -service ${svc}service -use github.com/ahaostudy/onlinejudge/kitex_gen -I ../../idl/ ../../idl/${svc}.thrift
