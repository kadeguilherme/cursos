# Container Runtime
> [!Note]
> Dockershim foi removido do Kubernetes a partir da versao 1.24.
Voce precisa instalar o container runtime dentro de cada node do cluster. O container runtime e responsavel por executar os containers dentro dos pods.
## Instalando e pre-requisitos da configuração
### Configuracao Network
O kernel do linux nao permite pacotes IPv4 seja roteados entre interfaces de rede, por padrao. Para habilitar o roteamento de pacotes IPv4 entre interfaces de rede, voce precisa habilitar o encaminhamento de pacotes IPv4 no kernel do linux. Para habilitar o encaminhamento de pacotes IPv4 no kernel do linux, voce pode usar o comando abaixo:
```bash
cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.ipv4.ip_forward = 1
EOF
# apply the changes without reboot
sudo sysctl --system
# verify the changes
sysctl net.ipv4.ip_forward
```
### cgroup drivers
No Linux, control groups (cgroups) sao usados para limitar e isolar o uso de recursos (CPU, memoria, disco, etc.) de um grupo de processos.
Ha dois tipos de cgroup drivers:
- cgroupfs
- systemd
> ![Note]
> O cgroup driver do kubelet e do container runtime devem ser iguais.
#### systemd como cgroup driver
Para usar o systemd como cgroup driver, voce precisa editar o ```kubeletconfiguration``` opcao ```cgroupDriver``` para ```systemd```.
```bash
apiVersion: kubelet.config.k8s.io/v1beta1
kind: KubeletConfiguration
...
cgroupDriver: systemd
```
### Verificando a configuracao kubelet
- iniciar proxy server usando:
```bash
kubectl proxy
```
- verificar a configuracao do kubelet usando:
```bash
curl -X GET http://127.0.0.1:8001/api/v1/nodes/<node-name>/proxy/configz | jq .
```
## Containerd
Para instalar o [containerd](https://github.com/containerd/containerd/blob/main/docs/getting-started.md).
### Configurando o systemd cgroup driver
Para usar o systemd como cgroup driver, voce precisa editar o arquivo de configuracao do containerd em ```/etc/containerd/config.toml```:
Containerd version 1.x:
```toml
[plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc]
  ...
  [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc.options]
    SystemdCgroup = true
```
Containerd version 2.x:
```toml
[plugins.'io.containerd.cri.v1.runtime'.containerd.runtimes.runc]
  ...
  [plugins.'io.containerd.cri.v1.runtime'.containerd.runtimes.runc.options]
    SystemdCgroup = true
```
Reinicie o containerd para aplicar as alteracoes:
```bash
sudo systemctl restart containerd
```
## CRI-O
Para instalar o [CRI-O](https://github.com/cri-o/packaging/blod/main/README.md).
CRI-O usa o systemd como cgroup driver por padrao.
## Docker Engine
Para instalar o [Docker Enginei](https://docs.docker.com/engine/install/).
