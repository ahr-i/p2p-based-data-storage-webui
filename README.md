# P2P Based Distributed-data-storage / Webui
It is a distributed-data-storage that can be used in a peer-to-peer (P2P) environment.   
   
Use it in conjunction with distributed-data-storage-bootstrap.   
(https://github.com/ahr-i/p2p-based-data-storage-bootstrap)

## 1. Start
### 1.1 Clone
```
git clone https://github.com/ahr-i/p2p-based-data-storage-webui.git
```

### 1.2 build
```
cd p2p-based-data-storage-webui
docker build -t dds/node .
```

### 1.3 create network
```
docker network create --subnet 200.0.0.0/16 p2p
```

### 1.4 Run
```
docker run -d --name dds_node_1 --network p2p -p 2000:2000 dds/node
```
You need to first run the bootstrap-node.   
(https://github.com/Ahr-i/p2p-based-data-storage-bootstrap)
