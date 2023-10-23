default:
	docker build . -t ngp-server:latest

run:
	docker run -d --name ngp-server-instance -e PORT=9000 -p 9000:9000 ibaku-server:latest

stop:
	docker stop ngp-server-instance

clean:
	docker rm ngp-server-instance

clean-image:
	docker image rm ngp-server:latest --force
