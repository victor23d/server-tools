go mod vendor
docker build . -t victor23d/exporter
docker push victor23d/exporter
docker run -p 80:80 victor23d/exporter
