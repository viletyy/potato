###
 # @Date: 2021-06-20 21:43:20
 # @LastEditors: viletyy
 # @LastEditTime: 2021-06-20 22:23:27
 # @FilePath: /potato/scripts/gen_proto.sh
### 
protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
./proto/*.proto