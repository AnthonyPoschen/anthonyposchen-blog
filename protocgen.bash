services=(
	helloworld
)
for service in ${services[@]} 
do
	echo generated: ${service}
	# generate the golang file
	protoc -I proto/ proto/${service}.proto --go_out=plugins=grpc:proto
	# generate the typescript files
	protoc \
	--plugin=protoc-gen-ts=./frontend/node_modules/.bin/protoc-gen-ts \
	--js_out=import_style=commonjs,binary:proto \
	--ts_out=service=true:proto \
	-I ./proto \
	proto/${service}.proto
	
	mv ./proto/${service}.pb.go ./services/${service}/
	mv ./proto/${service}_pb* ./frontend/src/services/
done