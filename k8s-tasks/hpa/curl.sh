DOMAIN=34.85.14.202
for (( i=0;i<500;i++ ))
do
    curl $DOMAIN
    # sleep 0.0005
done
