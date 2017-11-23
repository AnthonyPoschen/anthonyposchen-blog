goservices=(
	frontend
	user
	blog
)

tsFrontendservices=(
	frontend
)
#admin section as its own SPA seperate from main blog.
tsAdminServices=(
)

for service in ${goservices[@]} 
do
	# generate the golang file
	echo "Generating Go Service: ${service}" 
	protoc -I proto/ proto/${service}.proto --go_out=plugins=grpc:proto
	mv ./proto/${service}.pb.go ./services/${service}/
done

for service in ${tsFrontendservices[@]}
do
	echo "Generating TS Service: ${service}"
	# generate the typescript files
	protoc \
	--plugin=protoc-gen-ts=./services/frontend/web/node_modules/.bin/protoc-gen-ts \
	--js_out=import_style=commonjs,binary:proto \
	--ts_out=service=true:proto \
	-I ./proto \
	proto/${service}.proto
	
	mv ./proto/${service}_pb* ./services/${service}/web/src/services/
	rm ./proto/${service}_pb* 2>/dev/null
done