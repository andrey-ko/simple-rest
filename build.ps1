$ErrorActionPreference = "Stop"
docker build -t simple-rest -f .\Dockerfile.win .
docker tag simple-rest akolomentsev/simple-rest
docker push akolomentsev/simple-rest
