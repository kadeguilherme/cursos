# Descomplicando Kubernetes - Day 4
- Replicasets
- Deamonsets
- Probes

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

### Criando um DaemonSets
``` bash
kubectl apply -f ./node-exporter-daemonset.k8s.yaml
```

### Probes
As Probes são mecanismos de verificação de saúde dos Pods no Kubernetes. Elas permitem que o Kubernetes monitore o estado dos contêineres e tome ações apropriadas com base nos resultados dessas verificações.
Existem três tipos principais de Probes:
1. **Liveness Probe**: Verifica se o contêiner está em execução.
2. **Readiness Probe**: Verifica se o contêiner está pronto para receber tráfego.
3. **Startup Probe**: Verifica se o contêiner iniciou corretamente.
 
#### Liveness Probe
A Liveness Probe é usada para determinar se um contêiner está em execução. Se a Liveness Probe falhar, o Kubernetes reiniciará o contêiner. Isso é útil para detectar situações em que um contêiner está em um estado de falha e precisa ser reiniciado.
Exemplo de configuração de Liveness Probe em um Pod:
```bash
kubectl apply -f nginx-liveness.k8s.yaml
```

### Readiness Probe
A Readiness Probe é usada para determinar se um contêiner está pronto para receber tráfego de rede. Se a Readiness Probe falhar, o Kubernetes removerá o Pod do serviço de balanceamento de carga, evitando que ele receba solicitações até que esteja pronto novamente.
Exemplo de configuração de Readiness Probe em um Pod:
```bash
kubectl apply -f nginx-readiness.k8s.yaml
```
### Startup Probe
A Startup Probe é usada para determinar se um contêiner iniciou corretamente. Ela é útil para contêineres que podem demorar mais para iniciar do que o tempo limite padrão das Liveness e Readiness Probes. Se a Startup Probe falhar, o Kubernetes reiniciará o contêiner. Isso ajuda a evitar reinicializações prematuras de contêineres que estão apenas demorando para iniciar.
Exemplo de configuração de Startup Probe em um Pod:
```bash
kubectl apply -f nginx-startup.k8s.yaml
```

