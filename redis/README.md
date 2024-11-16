## 单节点安装redis

1. 官网下载redis安装包，地址：<https://redis.io/downloads/>
2. 执行以下命令解压
```
tar xvf redis-7.0.15.tar.gz
```
3. 进入解压后的目录，执行make命令进行编译
```
cd redis-7.0.15
make
```
4. 编译完成后进入src目录，执行make install 进行安装，PREFIX参数指定的路径即为redis安装路径
```
cd src
make install PREFIX=/usr/local/redis
```
5. 安装完成后修改redis配置文件  
   查看安装路径 /usr/local/redis/bin下是否有redis.conf文件。若没有，则需要取原解压文件下拷贝redis.conf到 /usr/local/redis/bin 目录下  
   配置文件修改内容如下，port为redis启动端口，修改时需确认该端口未被占用。
```
port 61199
protected-mode no
bind 0.0.0
logfile "/var/log/redis/redis.log"
```

6. 设置redis服务开机自启动，执行一下命令，创建redis系统服务
```
sudo vim /etc/systemd/system/redis.service
```
输入以下内容并保存
```
[Unit]
Description=Redis Server
After=network.target
 
[Service]
Type=simple
ExecStart=/usr/local/redis/bin/redis-server /usr/local/redis/bin/redis.conf
ExecStop=/usr/local/redis/bin/redis-cli shutdown

   [Install]
WantedBy=multi-user.target   
```
7. 执行以下命令，启动redis
```
sudo systemctl deamon-reload
sudo systemctl enable redis.service
sudo systemctl start redis.service
```

### 依赖
```
go get github.com/go-redis/redis
```