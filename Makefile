
blox: $(GOPATH)/src/github.com/blox/blox/cluster-state-service/swagger/v1/swagger.json
	rm -rf $@
	swagger generate client -f $< -t $@

.PHONY: blox
