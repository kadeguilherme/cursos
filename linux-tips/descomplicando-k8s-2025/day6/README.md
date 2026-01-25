# Descomplicando Kubernetes
## Day 06 — Volumes

---

## Conteudo da Aula
Entender como o k8s lidar com **Persistencia de dados**.

- Volumes no k8s   
  - StorageClass
  - PV (persistente volume) 
  - PVC (persistente volume claim)

---
## O que são volumes ?
Um Volume no Kubernetes é uma abstração de armazenamento que pode ser usada pelos contêineres dentro de um Pod. Diferente do armazenamento padrão do contêiner (que é temporário e desaparece quando o contêiner é reiniciado), os Volumes no Kubernetes permitem persistir dados de forma mais confiável.
Os volumes são divididos em 2 tipos:
1) ephemeral volumes
2) persistent volumes

O **ephemeral volume** o emptyDir que são criado e apos a destruicao do pod o emptyDir sera destruido junto, ou seja o dados sao perdidos.

O **persistent volume** sao volumes que são criados e que não sao destruido caso pod seja destruido, sao persistidos, sao mantidos os dados mesmo com destruicao do pod.

Tipos de volumes:

1) Volumes enfemeros
    - emptyDir
    - downwardAPI
    - configMap
2) Volumes persistente
    - nfs
    - iscsi
    - fc (fibre channel)
    - hostPath
    - local

[Documentação sobre volumes no Kubernetes](https://kubernetes.io/docs/concepts/storage/volumes/)

## StorageClass
O StorageClass é um objeto do Kubernetes que descreve e define diferentes classes de armazenamento disponíveis no cluster. Quando um usuário ou aplicação solicita a criação de um **PersistentVolumeClaim**(PVC), o Kubernetes utiliza o StorageClass para criar automaticamente um **PersistentVolume**(PV), evitando a necessidade de criação manual do volume, seguindo as regras e parâmetros definidos nesse StorageClass.

O storageClass é definido por `provisionador` que é o responsavel por criar o PersistentVolume dinamicamente. Os provisionador podem ser internos (fornecidos pelo próprio Kubernetes) ou exeternos (fornecidos por provedores de armazenamento específicos).

> [!NOTE]
> Cada cloud provider possui seu próprio provisionador.
> O provisionador padrão para volumes locais é:
> **`kubernetes.io/no-provisioner`**


> [!IMPORTANT]
> Em volumes locais, o StorageClass com WaitForFirstConsumer garante que o PVC só seja vinculado a um PV após o schedule do Pod, assegurando que ambos estejam no mesmo node.

### Criando um StorageClass

```bash
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: giropops
provisioner: kubernetes.io/no-provisioner
reclaimPolicy: Retain
volumeBindingMode: WaitForFirstConsumer
```
---
```bash
k apply -f local-storage.yaml

k get storageClas # Ver o storageClass criado chamado giropops
```

### PersistentVolume PV
PersistentVolume (PV) é um recurso de armazenamento disponível no cluster Kubernetes, provisionado por um administrador ou dinamicamente por meio de Storage Classes.
Assim como um Node, o PV é um recurso do cluster e possui um ciclo de vida independente de qualquer Pod que o utilize.
O objeto PV encapsula os detalhes da implementação do armazenamento, que pode ser um sistema NFS, iSCSI ou um serviço de armazenamento de um provedor de nuvem. Dessa forma, esses detalhes ficam abstraídos das aplicações, permitindo que os Pods solicitem o uso do armazenamento por meio de um PersistentVolumeClaim (PVC), sem a necessidade de configurar diretamente o tipo de armazenamento utilizado.

O persistentvolume são implementados como plugins. O kubernetes suporta os seguintes plugins de PersistentVolume:
- hostPath
- local
- nfs
- iscsi
- fc
- csi
1. Depresiado mais ainda disponivel
- awsElasticBlockStore
- azureDisk
- azureFile

[Documentação sobre persistentvolume no Kubernetes](https://kubernetes.io/docs/concepts/storage/persistent-volumes/)

### Criando um PersistentVolume PV
Cria um arquivo chamado *pv.yaml*
```bash
apiVersion: v1 # Versão da API do Kubernetes
kind: PersistentVolume # Tipo de objeto que estamos criando, no caso um PersistentVolume
metadata: # Informações sobre o objeto
  name: meu-pv # Nome do nosso PV
  labels:
    storage: local
spec: # Especificações do nosso PV
  capacity: # Capacidade do PV
    storage: 1Gi # 1 Gigabyte de armazenamento
  accessModes: # Modos de acesso ao PV
    - ReadWriteOnce # Modo de acesso ReadWriteOnce, ou seja, o PV pode ser montado como leitura e escrita por um único nó
  persistentVolumeReclaimPolicy: Retain # Política de reivindicação do PV, ou seja, o PV não será excluído quando o PVC for excluído
  hostPath: # Tipo de armazenamento que vamos utilizar, no caso um hostPath
    path: "/mnt/data" # Caminho do hostPath, do nosso nó, onde o PV será criado
  storageClassName: standard # Nome da classe de armazenamento que será utilizada
```
----
```bash
kubectl apply -f pv.yaml
persistentvolume/meu-pv created
```

--- 

### PersistentVolumeClaim PVC
O PVC reividica um PV para ser usado por um container, entao o kubernetes tenta associar automaticamente um PVC a um PV compativel garantindo o armazenamento seja alocado corretamente.
Todo PVC é associado a um *storage Class* ou a um *persistent Volume*

Vamos criar nosso PVC para o PV que criamos acima. Criar um arquivo chamado *pvc.yaml*

```bash
apiVersion: v1 # versão da API do Kubernetes
kind: PersistentVolumeClaim # tipo de recurso, no caso, um PersistentVolumeClaim
metadata: # metadados do recurso
  name: meu-pvc # nome do PVC
spec: # especificação do PVC
  accessModes: # modo de acesso ao volume
    - ReadWriteOnce # modo de acesso RWO, ou seja, somente leitura e escrita por um nó
  resources: # recursos do PVC
    requests: # solicitação de recursos
      storage: 1Gi # tamanho do volume que ele vai solicitar
  storageClassName: giropops # nome da classe de armazenamento que será utilizada
  selector: # seletor de labels
    matchLabels: # labels que serão utilizadas para selecionar o PV
      storage: local # label que será utilizada para selecionar o PV
```
---
```bash
kubectl apply -f pvc.yaml
persistentvolumeclaim/meu-pvc created
```

Vamos criar um pod ngnix que seja vinculado ao PV.
Criar um arquivo chamado *pod.yaml*

```bash
apiVersion: v1
kind: Pod
metadata:
  name: nginx-pod
spec:
  containers:
  - name: nginx
    image: nginx:latest
    ports:
    - containerPort: 80
    volumeMounts:
    - name: meu-pvc
      mountPath: /usr/share/nginx/html
  volumes:
  - name: meu-pvc
    persistentVolumeClaim:
      claimName: meu-pvc
```
---
```bash
kubectl apply -f pod.yaml
pod/nginx-pod created
```

Agora podemos ver nosso PVC foi vinculado ao PV.

```bash
kubectl get pvc
```

## Final do Day-6
Durante o Day-6 você aprendeu tudo sobre volumes no Kubernetes! Durante o Day-6 você aprender o que é um StorageClass, um PV e um PVC, e mais do que isso, você aprendeu tudo isso de forma prática!
