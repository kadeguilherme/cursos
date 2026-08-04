# Descomplicando Kubernetes - Day 2

## O que é um Pod?
Um Pod pode ter mais de um container. O Pod é a menor unidade do Kubernetes. Por exemplo, um sidecar (container auxiliar que roda junto com o container principal).

## Características do Pod

- **Efêmero** - Pods podem morrer e ser substituídos, não são permanentes
- **Compartilham rede** - todos containers no mesmo Pod compartilham o mesmo IP
- **Compartilham storage** - volumes são acessíveis por todos containers do Pod
- **Cada Pod roda em um único Node**

## Comandos

```
kubectl get pods

kubectl describe pods <nome-do-pod>
 
kubectl get pods <nome-do-pod> -o yaml

kubectl delete pod <nome-do-pod>

kubectl run <nome-do-pod> --image <nome-da-imagem> --port <porta> 

kubectl attach <nome-do-pod>
kubectl exec -ti <nome> -- bash

```

## Volume
É a forma de guardar os dados.
Nessa aula aprendemos o tipo EmptyDir.
