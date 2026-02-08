
# Descomplicando Kubernetes
## Day 07 — Statefulset e o service
Vamos entender como Statefulset funciona e casos comuns de uso dele como no bancos de dados, sistema de mensageria e sistema de arquivos distribuidos. 

E por último vamos entender sobre o Service que é um objeto do Kubernetes que permite exportar nossa aplicação para o mundo externo.

## Statefulset
 ### O que Statefulset ?
O StatefulSet garante o gerenciamento do deployment e do scaling de um conjunto de Pods. Diferente de Deployments e ReplicaSets, que são Stateless (sem estado), o StatefulSet é utilizado quando os Pods precisam manter estado.

Ele fornece garantias adicionais, como identidade estável para os Pods, nomes e endereços consistentes e uma ordem previsível de criação e remoção durante o deployment e o scaling.

  #### Statefulset e PV
Durante a criação de um Pod gerenciado pelo Statefulset, é criado um PVC associado ao pod, então esse PVC é vinculado a um PV.
Ou seja, quando o Pod é recriado, ele se reconecta ao mesmo PVC como consequência ao mesmo PV garantindo a persistência dos dados.

  #### Statefulset e o Headless service
Como o StatefulSet garante que cada Pod tenha uma identidade própria, e o Headless Service permite descobrir cada Pod diretamente via DNS, sem load-balancing, eles são utilizados em conjunto.
Essa combinação permite que cada Pod do StatefulSet seja acessado por um nome DNS estável, garantindo comunicação direta entre os Pods, o que é essencial para aplicações que mantêm estado, como bancos de dados distribuídos.

Exemplo:

Imagine que temos um StatefulSet chamado giropops com três réplicas e um Headless Service chamado nginx.

Os Pods criados pelo StatefulSet serão:

1. giropops-0
2. giropops-1
3. giropops-2

Cada Pod terá um endereço DNS estável, fornecido pelo Headless Service:

- giropops-0.nginx.default.svc.cluster.local (master)
- giropops-1.nginx.default.svc.cluster.local (replica)
- giropops-2.nginx.default.svc.cluster.local (replica)

Mesmo que um Pod seja recriado, seu nome permanece o mesmo, o DNS continua válido e o volume persistente associado ao Pod, garantindo a persistência dos dados.

#### Características adicionais do StatefulSet
- **Rolling Updates**: Atualizações controladas, um Pod de cada vez
- **Ordem de deployment**: Pods são criados e removidos em ordem sequencial
- **Identidade estável**: Cada Pod mantém seu nome e identidade mesmo após reinicializações

## Criando um StatefulSet

```bash
kubectl apply -f nginx-statefulset.k8s.yaml
```

### Verificando o StatefulSet
```bash
kubectl get statefulset
kubectl describe statefulset nginx
kubectl get pods -l app=nginx
```

## Criando Headless service

O nosso StatefulSet está criado temos que criar o Headless Service para que possamos acessar os Pods.

```bash
kubectl apply -f nginx-headless-service.k8s.yaml
```

### Verificando o Service
```bash
kubectl get service nginx
kubectl describe service nginx
```

### Testando a comunicação
```bash
# Verificar resolução DNS dos Pods
kubectl exec -it nginx-0 -- nslookup nginx.default.svc.cluster.local
kubectl exec -it nginx-1 -- nslookup nginx-0.nginx.default.svc.cluster.local
```

## Arquivos de manifesto
- `nginx-statefulset.k8s.yaml` - Definição do StatefulSet com 3 réplicas e volume persistente
- `nginx-headless-service.k8s.yaml` - Headless Service para descoberta de Pods

## Service

O Service é um objeto do Kubernetes que funciona como uma camada de abstração de rede.

### Tipos de Service


| Tipo          | Descrição                                                                 |
|---------------|---------------------------------------------------------------------------|
| ClusterIP     | Expõe o serviço apenas dentro do cluster (padrão)                          |
| NodePort      | Expõe o serviço em uma porta fixa de cada nó do cluster                    |
| LoadBalancer  | Expõe o serviço externamente usando um balanceador de carga do provedor   |
| ExternalName  | Mapeia o serviço para um nome DNS externo                                  |
| Headless      | Não atribui IP ao serviço, usado para descoberta direta de Pods           |

### Criando Services

#### ClusterIp

``` bash
kubectl expose deployment meu-deployment --port=80 --target-port=8080
```

Este comando criará um serviço ClusterIP que expõe o meu-deployment na porta 80, encaminhando o tráfego para a porta 8080 dos Pods deste deployment. 

**Diferença entre port e targetPort:**
- `port`: Porta onde o serviço escuta dentro do cluster
- `targetPort`: Porta onde o container está realmente rodando nos Pods

Via YAML
``` yaml
apiVersion: v1
kind: Service # Tipo do objeto, no caso, um Service
metadata:
  name: meu-service
spec:
  selector: # Seleciona os Pods que serão expostos pelo Service
    app: meu-app # Neste caso, os Pods com o label app=meu-app
  ports:
    - protocol: TCP
      port: 80 # Porta do Service
      targetPort: 8080 # Porta dos Pods
```
Este arquivo YAML irá criar um serviço que corresponde aos Pods com o label app=meu-app, e encaminha o tráfego da porta 80 do serviço para a porta 8080 dos Pods. Perceba que estamos usando o selector para definir quais Pods serão expostos pelo Service.

#### NodePort 

``` bash
kubectl expose deployment MEU_DEPLOYMENT --port=80 --type=NodePort
```

Via YAML

``` yaml
apiVersion: v1
kind: Service
metadata:
  name: meu-service
spec:
  type: NodePort # Tipo do Service
  selector:
    app: meu-app
  ports:
    - protocol: TCP
      port: 80 # Porta do Service, que será mapeada para a porta 8080 do Pod
      targetPort: 8080 # Porta dos Pods
      nodePort: 30080   # Porta do Node, que será mapeada para a porta 80 do Service
```
Lembre que a porta do Node deve estar entre 30000 e 32767 e quando não especificado o Kubernetes vai escolher uma porta aleatória

#### LoadBalancer
É a forma mais comum de expor um serviço para a internet. Ele provisiona automaticamente um balanceador de carga do provedor de nuvem.

``` bash
kubectl expose deployment meu-deployment --type=LoadBalancer --port=80 --target-port=8080
```

Via YAML

``` yaml
apiVersion: v1
kind: Service
metadata:
  name: meu-service
spec:
  type: LoadBalancer
  selector:
    app: meu-app
  ports:
    - protocol: TCP
      port: 80
      targetPort: 8080
```

### ExternalName
O tipo de serviço ExternalName é um pouco diferente dos outros tipos de serviço. Ele não expõe um conjunto de Pods, mas sim um nome de host externo. Por exemplo, você pode ter um serviço que expõe um banco de dados externo.


``` bash
kubectl create service externalname meu-service --external-name meu-db.giropops.com.br
```

Via YAML

``` yaml
apiVersion: v1
kind: Service
metadata:
  name: meu-service
spec:
  type: ExternalName
  externalName: meu-db.giropops.com.br
```

## Verificando os Services

```bash
kubectl get services
# describe service
kubectl describe service meu-service
```

## Verificando os Endpoints
Os Endpoints são uma parte importante dos Services, pois eles são responsáveis por manter o mapeamento entre o Service e os Pods que ele está expondo. Quando um Service é criado, o Kubernetes automaticamente cria um objeto Endpoints que contém os IPs dos Pods selecionados.

```bash
kubectl get endpoints -A
kubectl get endpoints meu-service
# describe endpoints
kubectl describe endpoints meu-service
```

## Removendo um Service

```bash
kubectl delete service meu-service
# Removendo via yaml
kubectl delete -f meu-service.yaml
```
