cd server/ && docker build --tag server . && cd ..
docker stack deploy -c service.yml keyvalue