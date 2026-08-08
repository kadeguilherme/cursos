# Descomplicando Kubernetes - Day 3

## O que é um Deployment?
O Deployment é um recurso do Kubernetes que fornece atualizações declarativas para Pods e ReplicaSets. Ele permite que você descreva o estado desejado para seus Pods e o Kubernetes altera o estado real para o estado desejado de forma controlada.

## Características do deployment

- **Gerencia Pods** - cria, atualiza e deleta Pods conforme necessário
- **ReplicaSets** - garante que o número desejado de Pods esteja sempre em execução

## Comandos

```
kubectl get deployments

kubectl get deployment <nome-do-deployment> -o yaml

kubectl describe deployment <nome-do-deployment>

kubectl get replicasets
```

## Como criar um Deployment
Para criar um Deployment, você pode usar o comando `kubectl create deployment` seguido do nome do deployment e da imagem do container que deseja usar. Por exemplo:

```
kubectl create deployment meu-deployment --image=nginx
```
dry-run=client significa que o comando será executado apenas localmente, sem enviar nada para o servidor do Kubernetes. Isso é útil para verificar se o comando está correto antes de aplicá-lo.

```
kubectl create deployment meu-deployment --image=nginx --dry-run=client -o yaml >> meu-deployment.k8s.yaml
```

### Criando um deployment com arquivo YAML
Você também pode criar um Deployment usando um arquivo YAML. Aqui está um exemplo de arquivo YAML para

    criar um Deployment chamado `meu-deployment` que usa a imagem `nginx`:
    
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: meu-deployment
spec:
  replicas: 3
  selector:
    matchLabels:
      app: meu-app
  strategy:
    type: RollingUpdate
  template:
    metadata:
      labels:
        app: meu-app
    spec:
      containers:
        - name: nginx
          image: nginx
          resources:
            requests:
              memory: "64Mi"
              cpu: "250m"
            limits:
              memory: "128Mi"
              cpu: "500m"
          ports:
            - containerPort: 80
 ```
kubectl apply -f meu-deployment.k8s.yaml
kubectl delete -f meu-deployment.k8s.yaml
```
### Estrategia de atualização do Deployment

- Rolling Update é a estratégia padrão de atualização de um Deployment. Ele atualiza os Pods de forma gradual, garantindo que o número desejado de Pods esteja sempre em execução.
- Recreate é outra estratégia de atualização, que primeiro deleta todos os Pods antigos antes de criar os novos. Isso pode causar downtime, mas é útil em alguns casos.

### Rolback de um Deployment
O Kubernetes mantém um histórico de revisões do Deployment, permitindo que você faça rollback para uma versão anterior caso algo dê errado durante uma atualização. Para fazer rollback, você pode usar o comando:
    
```
kubectl rollout undo deployment <nome-do-deployment>
kubectl rollout history deployment <nome-do-deployment>
```
