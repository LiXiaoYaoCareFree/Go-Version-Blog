## Basic Usage


```
Step1: Download the corresponding elasticHD version，unzip xxx_elasticHd_xxx.zip
Step2: chmod 0777 ElasticHD
Step3: ./ElasticHD -p 127.0.0.1:9800 

```

```
netstat -tulnp | grep 9800
kill -9 <PID>

启动es
cd /opt/blogx/Blog-Server/init/deploy/
chmod 777 /opt/es/data
docker compose up -d

远程访问时使用 SSH 端口转发
ssh -L 9800:localhost:9800 root@192.168.179.135:22
cd /opt/BlogX/Blog-Server/Blog-Server/ElasticHD/
./ElasticHD -p 127.0.0.1:9800 
DELETE  http://192.168.179.135:9200/article_models/_doc/2

进入 MySQL 容器
docker exec -it mysql-master bash
mysql -u root -proot

docker exec -it mysql-slave bash
mysql -u root -proot

systemctl daemon-reload
systemctl restart docker



SET NAMES utf8mb4;

git commit --amend -m "新的提交信息"

修改mysql服务的时区
ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

运行ollama
/opt/BlogX/Blog-Server/Blog-Server/bin

docker exec -it ollama-container /bin/bash

docker run -d -v ollama:/root/.ollama -p 11434:11434 --name ollama ollama/ollama

```
