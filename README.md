## docker

docker build -t comments-api

docker run -it -p 8080:8080 comments-api

## kind

brew install kind

kind create cluster

kubectl cluster-info --context kind-kind

kubectl get pods

kubectl get services

envsubst < config/deployment.yml > temp.yml

docker build -t pdevted/comments-api .

docker push pdevted/comments-api:lastest

envsubst < config/deployment.yml | kubectl apply -f -

kubectl get pods

ex)
kubectl logs comments-api-27823472843d-2km4n


-----

kubectl apply -f config/service.yml

kubectl get service

kubectl get enpoints

kubectl post-forward service/comments-api 8080:8080

