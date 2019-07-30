# kubectl run php-apache --image=k8s.gcr.io/hpa-example --requests=cpu=200m --expose --port=80

kubectl autoscale deployment php-apache --cpu-percent=50 --min=1 --max=5
