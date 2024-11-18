# File service

## Start

1.  Copy `.env.example` file as `.env`. Replace the values of the file. These are currently hardcoded, but would be dynamically returned by the lookup service.
    - FILEURLS: The go service wasn't able to respond to all requests, due to the fact that the file server returned 503 (unavailable). The current solution is to pass in multiple URLs - each being served by a different file server. The file-service itself then picks randomly between one of the passed in URLs.
2.  `docker build -t file-service --platform linux/amd64 .`
3.  `docker run -e PORT=3000 -p 3000:3000 file-service`

## Further considerations

- What happens on network downtime?
- Look into request and response timeouts
- grpc vs REST

## Azure instructions

1. `az login`
2. `az acr login --name <container-registry-name>`
3. `docker build -t file-service:v1 . --no-cache --platform linux/amd64`
4. `docker tag file-service:v1 <container-registry-name>.azurecr.io/file-service:v1`
5. `docker push <container-registry-name>.azurecr.io/file-service:v1`
