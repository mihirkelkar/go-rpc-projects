1) Build the dev docker container environment :
`docker compose build -t <containername>:<version> .`

2) Run docker container
`docker run -v <pwd>:<app-dir> -it <containername>:<version>`
