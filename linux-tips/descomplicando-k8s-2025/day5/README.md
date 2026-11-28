# Resumo
Essa doc e um guia para iniciar um cluster Kubernetes, que sera composto 3 nodes no Ubuntu 22.04.
- 1 control plane node
- 2 worker nodes

No decorrer da doc vai ser apresentado conceitos sobre ferramentas de instalacao do kubernetes, pre-requisitos e etc.

## Ferramentas para criar um cluster Kubernetes
### Ferramentas para ambiente locais
- k3s: A CNCF e responsavel dele, empacotada em um único binário e com baixo consumo de recursos. Usada em produção, principalmente em edge, IoT e ambientes com pouco hardware.
- minikube: Cria um cluster local na sua máquina, com os nodes rodando em VMs, dependendo do driver (Docker, KVM, VirtualBox etc.).

### Ferramentas para ambiente de producao
- kubeadm: E a ferramenta oficial do Kubernetes para instalar clusters. É a forma mais recomendada e utilizada. E é a que vamos utilizar nesse tutorial.
- Cluster API: Uma ferramenta que permite criar clusters Kubernetes em diferentes provedores de nuvem e ambientes on-premises.
- kops: Uma ferramenta que permite criar clusters Kubernetes em diferentes provedores de nuvem e ambientes on-premises.
- kubespray: Uma ferramenta que permite criar clusters Kubernetes em diferentes provedores de nuvem e ambientes on-premises.

Fonte: [Doc oficial](https://kubernetes.io/docs/setup/production-environment/tools/)

## Sumario da doc
1. Pré-requisitos e instalação do kubeadm, kubelet e kubectl
2. [Container runtime](container-runtime.md)
3. Instalando kubeadm, kubelet e kubectl
4. [Criando o cluster com kubeadm](criando-cluster.md)

### Pré-requisitos

> [!IMPORTANT]
> A instalacao do kubeadm e feita via binarios que usa o **dynamic linking** e assume que o sistema tem o **glibc instalado**.
A maioria das distribuicoes linux (como Ubuntu, Debian, CentOS, Fedora, etc) tem o glibc instalado por padrao.).

> [!IMPORTANT]
> O alpine linux é uma exceção, pois ele usa **musl libc**, que e uma implementação diferente da glibc. Por isso, o kubeadm nao funciona no alpine linux. O gcompat e uma biblioteca que fornece compatibilidade com glibc para o musl libc, permitindo que aplicativos que dependem de glibc sejam executados em sistemas baseados em musl libc, como o Alpine Linux. No entanto, o gcompat e uma solucao temporaria e nao e recomendada para ambientes de producao. A melhor abordagem e usar uma distribuicao linux que seja oficialmente suportada pelo Kubernetes.

> [!NOTE]
> Esse guiar e para o kubernetes **1.37**

#### Verificando a versao do kernel
Verifique a versão do kernel:

```bash
uname -r
```

Para ver as versões do kernel suportadas pelo kubeadm, acesse a [Lista LTS do kernel](https://www.kernel.org/category/releases.html).

#### Verificando MAC address e product_uuid sao unicos para o node

```bash
ip link 
#or
ifconfig -a
```

```bash
cat /sys/class/dmi/id/product_uuid
```

Se voce copiou as VMs, verifique se o product_uuid e unico para cada VM.

#### Verificando network adapters
Se ter mais que 1 network adapter, o kubeadm pode escolher o errado. Para evitar isso, voce pode add IP route para o node, ou desativar o network adapter que nao sera usado.

```bash
# Se nao tiver gateway, voce pode add IP route para o node.
ip route add <IP_ADDRESS> dev <INTERFACE>
#or se tiver gateway, voce pode add IP route para o node.
ip route add <IP_ADDRESS> via <GATEWAY_IP> dev <INTERFACE>

ip route
```

#### Portas utilizadas no Kubernetes
Essas portas devem estar abertas para que o cluster Kubernetes funcione corretamente. Se estiver usando um firewall, certifique-se de permitir o tráfego nessas portas.

##### Control Plane
| Protocolo | Direcao | Portas | Finalidade | Usado por |
| :--- | :--- | :--- | :--- | :--- |
| TCP | Inbound | 6443 | Kubernetes API Server | Todos |
| TCP | Inbound | 2379-2380 | etcd server client API | kube-apiserver, etcd |
| TCP | Inbound | 10250 | Kubelet APi | local, control plane |
| TCP | Inbound | 10259 | kube-scheduler | local |
| TCP | Inbound | 10257 | kube-controller-manager | local |

> [!Note]
> As portas do etcd pode mudar se voce estiver usando um cluster etcd externo ou portas customizadas.

##### Worker node(s)
| Protocolo | Direcao | Portas | Finalidade | Usado por |
| :--- | :--- | :--- | :--- | :--- |
| TCP | Inbound | 10250 | Kubelet API | local, control plane |
| TCP | Inbound | 10256 | kube-proxy | local, Load balancers |
| TCP | Inbound | 30000-32767 | NodePort Services | Todos |
| UDP | Inbound |  30000-32767 | NodePort Services | Todos |

> [!TIP]
> Se voce estiver usando um firewall, certifique-se de permitir o tráfego nessas portas.

> [!WARNING]
> o plugin pod netorking (CNI) que voce escolher, pode ter portas adicionais que precisam ser abertas.

#### Swap
Por padrao, o Kubelet falha ao iniciar, se tiver swap no node e preciso desativar o swap. Para desativar o swap, voce pode usar o comando abaixo:

```bash
sudo sed '/swap/s/^/#/'  /etc/fstab
# reboot
sudo reboot
# Apos reiniciar, verifique se o swap esta desativado
free -h
```
Comente a linha que contem swap, ou remova a linha salve e saia do editor

### Container Runtime
Para rodar pods no Kubernetes, voce precisa de um container runtime.
Por padrão, kubernetes usa o Container Runtime Interface (CRI) para se comunicar com o container runtime. O CRI e uma interface padrao que permite que o kubelet se comunique com diferentes container runtimes.

O CRI suporta varios container runtimes, como:
- containerd: E um container runtime leve e de alto desempenho, que e usado por padrao pelo Kubernetes. Ele e recomendado para producao.
- CRI-O: E um container runtime leve e de alto desempenho, que e usado por padrao pelo Kubernetes. Ele e recomendado para producao.
- Docker Engine(cri-dockerd): E um container runtime popular, mas nao e recomendado para producao, pois ele nao e mantido pelo Kubernetes. O Docker usa o dockershim, que foi removido do Kubernetes a partir da versao 1.24. Para usar o Docker como container runtime, voce precisa instalar o cri-dockerd, que e um projeto de comunidade que fornece suporte ao Docker no Kubernetes.

Para instalar e configurar o container runtime, veja [container-runtime.md](container-runtime.md).

### Instalando kubeadm, kubelet e kubectl
Para instalar o kubeadm, kubelet e kubectl, voce pode usar o comando abaixo:

```bash
sudo apt-get update

sudo apt-get install -y apt-transport-https ca-certificates curl gpg

# If the directory `/etc/apt/keyrings` does not exist, it should be created before the curl command, read the note below.
# sudo mkdir -p -m 755 /etc/apt/keyrings
curl -fsSL https://pkgs.k8s.io/core:/stable:/v1.37/deb/Release.key | sudo gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg

# This overwrites any existing configuration in /etc/apt/sources.list.d/kubernetes.list
echo 'deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v1.37/deb/ /' | sudo tee /etc/apt/sources.list.d/kubernetes.list

# Update the package list and install kubelet, kubeadm and kubectl
sudo apt-get update
sudo apt-get install -y kubelet kubeadm kubectl
sudo apt-mark hold kubelet kubeadm kubectl
# Enable and start the kubelet service
sudo systemctl enable --now kubelet

```
