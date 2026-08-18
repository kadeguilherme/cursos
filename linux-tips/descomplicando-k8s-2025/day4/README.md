# Descomplicando Kubernetes - Day 4

## Replicasets e Deamonsets
### O que é um ReplicaSet?
Um ReplicaSet é um recurso do Kubernetes que garante que um número especificado de réplicas de um Pod esteja em execução a qualquer momento. Ele é responsável por criar e gerenciar os Pods, garantindo que o estado desejado seja mantido. Se um Pod falhar ou for excluído , o ReplicaSet criará automaticamente um novo Pod para substituí-lo.
Deployments são a forma recomendada de gerenciar ReplicaSets, pois eles fornecem recursos adicionais, como atualizações e rollback.

deployments -> ReplicaSets -> Pods

