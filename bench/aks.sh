# az aks nodepool update \
#   --resource-group aks \
#   --cluster-name ttt \
#   --name pool2 \
#   --enable-cluster-autoscaler \
#   --min-count 0 \
#   --max-count 2


az aks update \
  --resource-group aks \
  --name ttt \
  --cluster-autoscaler-profile scan-interval=10s scale-down-delay-after-add=1m scale-down-unneeded-time=1m scale-down-utilization-threshold=0.7 max-graceful-termination-sec=60s

