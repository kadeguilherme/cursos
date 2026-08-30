# Descomplicando Kubernetes - Day 4

## Replicasets e Deamonsets
### O que é um ReplicaSet?
Um ReplicaSet é um recurso do Kubernetes que garante que um número especificado de réplicas de um Pod esteja em execução a qualquer momento. Ele é responsável por criar e gerenciar os Pods, garantindo que o estado desejado seja mantido. Se um Pod falhar ou for excluído , o ReplicaSet criará automaticamente um novo Pod para substituí-lo.
Deployments são a forma recomendada de gerenciar ReplicaSets, pois eles fornecem recursos adicionais, como atualizações e rollback.

deployments -> ReplicaSets -> Pods

´´´bash
k get replicasets

´´´
### O que é um DaemonSet?
Um DaemonSet é um recurso do Kubernetes que garante que uma cópia de um Pod seja executada em todos (ou alguns) nós do cluster. Ele é útil para executar serviços que precisam ester presentes em todos os nodes.


Alguns exemeplos de uso de DaemonSets:
- Coletar logs de todos os nós do cluster.
- Executar agentes de monitoramento como o Prometheus Node Exporter ou o Fluentd.
- Execucão de um proxy de redes em todos os nós do cluster, como o Calico ou o Weave Net.
- execução de agentes de segurança, como o Falco ou o Sysdig.


