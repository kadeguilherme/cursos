# Descomplicando Kubernetes - Day 1

## O que é um container?
Container é uma unidade leve e portável que empacota uma aplicação e suas dependências, executando em um ambiente isolado no sistema operacional. Esse isolamento é possível graças a duas tecnologias do kernel Linux:

- **cgroups (control groups)**: controlam e limitam o uso de recursos do sistema, como CPU, memória, I/O de disco e PID.
- **Namespaces**: isolam a visão que o processo tem do sistema, como PID, rede, mount, UTS, IPC e usuário.

Enquanto os cgroups definem **quanto** recurso o container pode usar, os namespaces definem **o que** o processo enxerga.

## O que é um container engine?
É a ferramenta responsável por gerenciar o ciclo de vida dos containers: criar, executar, parar, remover, além de gerenciar imagens, volumes e redes. Exemplos: **Docker** e **Podman**.

O container engine depende de um container runtime para interagir com o kernel.

## O que é um container runtime?
É o software responsável por executar o container de fato. Ele implementa as especificações do OCI (Open Container Initiative) e se comunica com o kernel para criar e rodar os containers.

### Tipos de container runtime
- **Baixo nível (low-level)**: responsável por criar e executar o container diretamente com o kernel. Exemplos: **runC**, **crun**.
- **Alto nível (high-level)**: gerencia o ciclo de vida dos containers e implementa o CRI (Container Runtime Interface), usado pelo kubelet do Kubernetes. Exemplos: **containerd**, **CRI-O**.
- **Sandbox**: oferece uma camada extra de isolamento e segurança, executando cada container em uma VM leve. Exemplos: **gVisor**, **Kata Containers**.

## O que é Kubernetes?
Kubernetes (K8s) é um orquestrador de containers que automatiza o deploy, a escalabilidade e o gerenciamento de aplicações em container. Ele agrupa containers em **Pods**, distribui a carga entre nós, realiza auto-scaling, health checks, rollouts e rollbacks de forma declarativa.

## O que são workers e control plane no Kubernetes?
O Kubernetes possui dois planos principais: o **Control Plane** (plano de controle) e os **Worker Nodes** (nós de trabalho).

### Control Plane
Responsável por gerenciar e controlar o cluster como um todo. Seus componentes são:

- **etcd**: banco de dados chave-valor distribuído que armazena todo o estado e configuração do cluster.
- **kube-apiserver**: ponto central de comunicação do cluster. Expõe a API do Kubernetes, valida, autentica e autoriza requisições, e persiste o estado no etcd.
- **kube-scheduler**: responsável por decidir **em qual nó** cada Pod será executado, levando em conta recursos disponíveis, afinidades, taints e tolerations.
- **kube-controller-manager**: executa os controladores que mantêm o estado desejado do cluster, como Deployment, ReplicaSet, Node Controller, Endpoint Controller, etc.
- **cloud-controller-manager** (opcional): integra o cluster com APIs de provedores de nuvem (AWS, GCP, Azure), gerenciando recursos como load balancers e volumes.

### Worker Nodes
Responsáveis por executar as aplicações em container. Cada worker node contém:

- **kubelet**: agente que roda em cada nó, responsável por garantir que os containers dos Pods estejam em execução conforme definido. Ele se comunica com o kube-apiserver.
- **kube-proxy**: responsável pelas regras de rede do cluster, mantendo a comunicação entre Pods e Services via iptables ou IPVS.
- **Container runtime**: o software que executa os containers (containerd, CRI-O, etc.), conforme explicado acima.

### pods, replicas, deployment e service
- pods
- replicas
- dployment 
- service

referencia:
[concepts](https://kubernetes.io/docs/concepts/extend-kubernetes)

