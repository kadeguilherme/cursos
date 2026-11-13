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

### Portas requeridas para o Kubernetes
Essas portas devem estar abertas para que o cluster Kubernetes funcione corretamente. Se estiver usando um firewall, certifique-se de permitir o tráfego nessas portas.

#### Control Plane
| Protocolo | Direcao | Portas | Finalidade | Usado por |
| :--- | :--- | :--- | :--- | :--- |
| TCP | Inbound | 6443 | Kubernetes API Server | Todos |
| TCP | Inbound | 2379-2380 | etcd server client API | kube-apiserver, etcd |
| TCP | Inbound | 10250 | Kubelet APi | local, control plane |
| TCP | Inbound | 10259 | kube-scheduler | local |
| TCP | Inbound | 10257 | kube-controller-manager | local |

> [!Note]
> As portas do etcd pode mudar se voce estiver usando um cluster etcd externo ou portas customizadas.

#### Worker node(s)
| Protocolo | Direcao | Portas | Finalidade | Usado por |
| :--- | :--- | :--- | :--- | :--- |
| TCP | Inbound | 10250 | Kubelet API | local, control plane |
| TCP | Inbound | 10256 | kube-proxy | local, Load balancers |
| TCP | Inbound | 30000-32767 | NodePort Services | Todos |
| UDP | Inbound |  30000-32767 | NodePort Services | Todos |

Para verificar se as portas estao abertas com netcat, voce pode usar o comando abaixo:

```bash
nc 127.0.0.1 6443 -zv -w 2
```
### Swap
Por padrao, o Kubelet falha ao iniciar, se tiver swap no node e preciso desativar o swap. Para desativar o swap, voce pode usar o comando abaixo:

```bash
 vim /etc/fstab
```
Comente a linha que contem swap, ou remova a linha salve e saia do editor


### Container Runtime
Para rodar pods no Kubernetes, voce precisa de um container runtime.
Por padrao, kubernetes usa o Container Runtime Interface (CRI) para se comunicar com o container runtime. O CRI e uma interface padrao que permite que o kubelet se comunique com diferentes container runtimes.
O CRI suporta varios container runtimes, como:
- containerd: E um container runtime leve e de alto desempenho, que e usado por padrao pelo Kubernetes. Ele e recomendado para producao.
- CRI-O: E um container runtime leve e de alto desempenho, que e usado por padrao pelo Kubernetes. Ele e recomendado para producao.
- Docker Engine(cri-dockerd): E um container runtime popular, mas nao e recomendado para producao, pois ele nao e mantido pelo Kubernetes. O Docker usa o dockershim, que foi removido do Kubernetes a partir da versao 1.24. Para usar o Docker como container runtime, voce precisa instalar o cri-dockerd, que e um projeto de comunidade que fornece suporte ao Docker no Kubernetes.

### Instalando kubeadm, kubelet e kubectl
Para instalar o kubeadm, kubelet e kubectl, voce pode usar o comando abaixo:
```bash
# This overwrites any existing configuration in /etc/apt/sources.list.d/kubernetes.list
echo 'deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v1.37/deb/ /' | sudo tee /etc/apt/sources.list.d/kubernetes.list
# Update the package list and install kubelet, kubeadm and kubectl
sudo apt-get update
sudo apt-get install -y kubelet kubeadm kubectl
sudo apt-mark hold kubelet kubeadm kubectl
# Enable and start the kubelet service
sudo systemctl enable --now kubelet
```
