# Instalando cluster Kubernetes com 3 nodes no Ubuntu 22.04.
- 1 control plane node
- 2 worker nodes

## Formas de instalar o cluster Kubernetes
- kubeadm: E a ferramenta oficial do Kubernetes para instalar clusters. É a forma mais recomendada e utilizada. E é a que vamos utilizar nesse tutorial.
- k3s e minikube: Ideal para ambientes de desenvolvimento e testes. Não é recomendada para produção.

## Getting Started

> [!Note]

Ha varios metodos e ferramentas para instalar o Kubernetes. Por exemplo: 
- kubeadm
- Cluster API: Uma ferramenta que permite criar clusters Kubernetes em diferentes provedores de nuvem e ambientes on-premises.
- kops: Uma ferramenta que permite criar clusters Kubernetes em diferentes provedores de nuvem e ambientes on-premises.
- kubespray: Uma ferramenta que permite criar clusters Kubernetes em diferentes provedores de nuvem e ambientes on-premises.

link: https://kubernetes.io/docs/setup/production-environment/tools/

## Instalando o Kubernetes com kubeadm
Esse guiar e para o kubernetes 1.37, que é a versão mais recente do Kubernetes.

> [!IMPORTANT]
> A instalacao do kubeadm e feita via binarios que use dynamic linking e assume que o sistema tem o glibc instalado.
Isso e a razao de assume que maioria das distribuicoes linux (como Ubuntu, Debian, CentOS, Fedora, etc) tem o glibc instalado por padrao.).

> O alpine linux e uma excecao, pois ele usa musl libc, que e uma implementacao diferente da glibc. Por isso, o kubeadm nao funciona no alpine linux. O gcompat e uma biblioteca que fornece compatibilidade com glibc para o musl libc, permitindo que aplicativos que dependem de glibc sejam executados em sistemas baseados em musl libc, como o Alpine Linux. No entanto, o gcompat e uma solucao temporaria e nao e recomendada para ambientes de producao. A melhor abordagem e usar uma distribuicao linux que seja oficialmente suportada pelo Kubernetes.

### Verificando seu versao do OS 
O cluster k8s depende de software que sao features do kernel, mas nao se limita container runtime, kubelet e container Network Interface (CNI)
Para ajudar a verificar se o seu sistema operacional e compativel com o Kubernetes, voce pode usar o comando abaixo para verificar a versao kernel e suportado.

```bash
kubeadm init --ignore-preflight-errors=SystemVerification
```
### Verificando MAC address e product_uuid sao unicos para o node

```bash
ip link 
#or
ifconfig -a
```

```bash
cat /sys/class/dmi/id/product_uuid
```
### Verificando network adapters
Se ter mais que 1 network adapter, o kubeadm pode escolher o errado. Para evitar isso, voce pode add IP route para o node, ou desativar o network adapter que nao sera usado.

```bash
# Se nao tiver gateway, voce pode add IP route para o node.
ip route add <IP_ADDRESS> dev <INTERFACE>
#or se tiver gateway, voce pode add IP route para o node.
ip route add <IP_ADDRESS> via <GATEWAY_IP> dev <INTERFACE>

ip route
```


