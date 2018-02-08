URL="https://raw.githubusercontent.com/googleapis/googleapis/master/"

test:
	go test ./...

regenerate:
	go install github.com/gogo/protobuf/protoc-gen-gogotypes

	protoc \
	--gogotypes_out=\
	Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
	Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
	:. \
	-I=. \
	google/rpc/status.proto \
	google/rpc/error_details.proto \
	google/rpc/code.proto

update:
	go install github.com/gogo/protobuf/gogoreplace

	(cd ./google/rpc && rm status.proto; wget ${URL}/google/rpc/status.proto)
	gogoreplace \
		'option go_package = "google.golang.org/genproto/googleapis/rpc/status;status";' \
		'option go_package = "rpc";' \
		./google/rpc/status.proto

	(cd ./google/rpc && rm error_details.proto; wget ${URL}/google/rpc/error_details.proto)
	gogoreplace \
		'option go_package = "google.golang.org/genproto/googleapis/rpc/errdetails;errdetails";' \
		'option go_package = "rpc";' \
		./google/rpc/error_details.proto

	(cd ./google/rpc && rm code.proto; wget ${URL}/google/rpc/code.proto)
	gogoreplace \
		'option go_package = "google.golang.org/genproto/googleapis/rpc/code;code";' \
		'option go_package = "rpc";' \
		./google/rpc/code.proto