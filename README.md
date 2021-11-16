# Backend-разработка на Go. Уровень 2

---

### For install grpc in go
```
https://grpc.io/docs/languages/go/quickstart/
```

### Build docker image
```
docker build -t lapitskyss/go_backend_2_lesson2:v1.0.0 .
docker push lapitskyss/go_backend_2_lesson2:v1.0.0
```

### Deploy to minikube commands
```
kubectl apply -f deployment.yaml
kubectl get pods

kubectl apply -f service.yaml
kubectl apply -f ingress.yaml

minikube service k8s-go-app-srv --url
minikube service k8s-go-app-srv-grpc --url
```