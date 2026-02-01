
# Descomplicando Kubernetes
## Day 07 — Statefulset e o service
Vamos entender como Statefulset funciona e casos comuns de uso dele como no bancos de dados, sistema de mensageria e sistema de arquivos distribuidos. 

E por ultimo vamos entender sobre o service que é um objeto do Kubernetes que permite export nossa aplicação para o mundo externo.

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
