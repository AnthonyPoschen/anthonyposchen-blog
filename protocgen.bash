goservices=(
	helloworld
)

tsFrontendservices=(
	helloworld
)
#admin section as its own SPA seperate from main blog.
tsAdminServices=(
	helloworld
)

for service in ${goservices[@]} 
do
	# generate the golang file
	protoc -I proto/ proto/${service}.proto --go_out=plugins=grpc:proto
	mv ./proto/${service}.pb.go ./services/${service}/
done

for service in ${tsFrontendservices[@]}
do

	# generate the typescript files
	protoc \
	--plugin=protoc-gen-ts=./services/frontend/web/node_modules/.bin/protoc-gen-ts \
	--js_out=import_style=commonjs,binary:proto \
	--ts_out=service=true:proto \
	-I ./proto \
	proto/${service}.proto
	
	mv ./proto/${service}_pb* ./services/frontend/web/src/services/
	rm ./proto/${service}_pb*
done