# learn-at-infracloud
GitHub Repo to store all Notes for the tools and technologies learnet during Initial days of infracloud 
This majorly covers personal notes related to some concepts/tasks golang, docker, kubernetes, prometheus, grafana etc.

##### Table of Contents 

[Docker-Instalation](#Docker-Instalation)

# Cluster Installation (Ubuntu)
### WSL
https://medium.com/@joaoh82/setting-up-kubernetes-on-wsl-to-work-with-minikube-on-windows-10-90dac3c72fa1
### Docker-Instalation
```
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"
sudo apt-get update
sudo apt-get install -y docker-ce=18.06.1~ce~3-0~ubuntu
sudo apt-mark hold docker-ce
```

###  Installing Kubeadm, Kubelet, and Kubectl
```
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
cat << EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
deb https://apt.kubernetes.io/ kubernetes-xenial main
EOF
sudo apt-get update
sudo apt-get install -y kubelet=1.12.7-00 kubeadm=1.12.7-00 kubectl=1.12.7-00
sudo apt-mark hold kubelet kubeadm kubectl
kubeadm version
```

###  bootstraping cluster
```
sudo kubeadm init --pod-network-cidr=10.244.0.0/16
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
kubectl version

sudo kubeadm join $some_ip:6443 --token $some_token --discovery-token-ca-cert-hash $some_hash // get from output of kubeadm init and execute on non master nodes to join the cluster

kubectl get nodes
```

###  Flannel for network setup (FOR CNI)
```
(on all 3 node)
echo "net.bridge.bridge-nf-call-iptables=1" | sudo tee -a /etc/sysctl.conf
sudo sysctl -p

(on master)
kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/bc79dd1505b0c8681ece4de4c0d86c5cd2643275/Documentation/kube-flannel.yml
kubectl get nodes // status must be ready
kubectl get pods -n kube-system // verify Flannel pods are up 
```


# Add node exporter in all cluster nodes (optional, useful for prometheus)
```
adduser prometheus
```
```
cd /home/prometheus
curl -LO "https://github.com/prometheus/node_exporter/releases/download/v0.16.0/node_exporter-0.16.0.linux-amd64.tar.gz"
tar -xvzf node_exporter-0.16.0.linux-amd64.tar.gz
mv node_exporter-0.16.0.linux-amd64 node_exporter
cd node_exporter
chown prometheus:prometheus node_exporter
```

create -> /etc/systemd/system/node_exporter.service
```
[Unit]
Description=Node Exporter

[Service]
User=prometheus
ExecStart=/home/prometheus/node_exporter/node_exporter

[Install]
WantedBy=default.target
```
```
systemctl daemon-reload
systemctl enable node_exporter.service
systemctl start node_exporter.service
```


