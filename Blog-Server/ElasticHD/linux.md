linux常用操作防火墙命令

systemctl status firewalld  查看防火墙状态
systemctl stop firewalld 关闭防火墙
systemctl start firewalld 开启防火墙
firewall-cmd --list-all 查看防火墙开放的端口
sudo firewall-cmd --add-port=8080/tcp --permanent 添加某一个端口，例如添加8080端口，加上--permanent永久生效，没有此参数重启后失效
firewall-cmd --reload 重启防火墙，当添加了新的端口后，一定要重启防火墙，否则不生效
firewall-cmd --zone=public --query-port=3306/tcp 查看某个端口是否被开放，例如查询3306端口是否被开放放行
firewall-cmd --zone=public --remove-port=3306/tcp --permanent 关闭端口，例如关闭3306端口放行
