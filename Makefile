.PHONY:	gen-product
gen-product:
  ##export ROOT_MOD=github.com/cloudwego/biz-demo/gomall
	@cd rpc_gen && cwgo client --type RPC --service product --module ${ROOT_MOD}/rpc_gen -I ../idl --idl ../idl/product.proto
	@cd app/product && cwgo server --type RPC --service product --module ${ROOT_MOD}/app/product --pass "use ${ROOT_MOD}/rpc_gen" -I ../../idl --idl ../../idl/product.proto

.PHONY: middleware
middleware:
	@consul agent -dev

	@mysql -u root -p, 123456
  
	@systemctl redis-server start

.PHONY: gen-frontend
gen-frontend:
	cd app/frontend && cwgo server -I ../../idl --type HTTP --service frontend --module frontend --idl ../../idl/frontend/about.proto

.PHONY: exportrootmod
exportrootmod:
	echo	"hhh"

.PHONY:	gen-email
gen-email:
	@cd rpc_gen && cwgo client --type RPC --service email --module ${ROOT_MOD}/rpc_gen -I ../idl --idl ../idl/email.proto
	@cd app/email && cwgo server --type RPC --service email --module ${ROOT_MOD}/app/email --pass "use ${ROOT_MOD}/rpc_gen" -I ../../idl --idl ../../idl/email.proto

.PHONY:	gen-eino
gen-eino:
	@cd rpc_gen && cwgo client --type RPC --service eino --module ${ROOT_MOD}/rpc_gen -I ../idl --idl ../idl/eino.proto
	@cd app/eino && cwgo server --type RPC --service eino --module ${ROOT_MOD}/app/eino --pass "use ${ROOT_MOD}/rpc_gen" -I ../../idl --idl ../../idl/eino.proto