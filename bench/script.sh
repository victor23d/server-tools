go mod vendor
docker build . -t victor23d/bench
docker push victor23d/bench
docker run -p 80:80 victor23d/bench
