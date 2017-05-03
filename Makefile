
blox-client: blox/cluster-state-service/swagger/v1/swagger.json
	swagger generate client -f $< -a blox -t $@
