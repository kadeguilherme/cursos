
# Descomplicando Kubernetes
## Day 07 — Statefulset e o service
Vamos entender como Statefulset funciona e casos comuns de uso dele como no bancos de dados, sistema de mensageria e sistema de arquivos distribuidos. 

E por ultimo vamos entender sobre o service que é um objeto do Kubernetes que permite export nossa aplicação para o mundo externo.

## Statefulset
 ### O que Statefulset ?
O StatefulSet garante o gerenciamento do deployment e do scaling de um conjunto de Pods. Diferente de Deployments e ReplicaSets, que são stateless (sem estado), o StatefulSet é utilizado quando os Pods precisam manter estado.

Ele fornece garantias adicionais, como identidade estável para os Pods, nomes e endereços consistentes e uma ordem previsível de criação e remoção durante o deployment e o scaling.

  #### Statefulset e PV
Durante a criaçãoade um Pod gerenciado pelo Statefulset, é criado um PVC associado ao pod, etao esse PVC é vinculado a um PV.
Ou seja, quando o Pod é reciado, ele se reconecta ao mesmo PVC como consequencia ao memos PV garantindo a persistencia dos dados.

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
