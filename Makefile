consult:
	docker run \
        -d \
        -p 8500:8500 \
        -p 8600:8600/udp \
        --name=badger \
        hashicorp/consul agent -server -ui -node=server-1 -bootstrap-expect=1 -client=0.0.0.0

database:
	docker run --name movieexample_db -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=movieexample -p 3306:3306 -d mysql:latest

load_data:
	docker exec -i movieexample_db mysql movieexample -h localhost -P 3306 --protocol=tcp -uroot -ppassword < schema/schema.sql

run_movie:
	go run ./movie/cmd/*.go

run_metadata:
	go run ./metadata/cmd/*.go

run_rating:
	go run ./rating/cmd/*.go
