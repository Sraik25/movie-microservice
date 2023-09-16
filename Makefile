consult:
	docker run \
        -d \
        -p 8500:8500 \
        -p 8600:8600/udp \
        --name=badger \
        hashicorp/consul agent -server -ui -node=server-1 -bootstrap-expect=1 -client=0.0.0.0

run_movie:
	go run ./movie/cmd/*.go

run_metadata:
	go run ./metadata/cmd/*.go

run_rating:
	go run ./rating/cmd/*.go
