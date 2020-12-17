sudo docker run -d --name gatedb -p3308:3306 -v $(pwd)/gatedb:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=dkfcpfk0609! -e MYSQL_DATABASE=gatedb gatedb:latest
