## 定时下载github realease到文件服务器

### 特点

多仓库配置

定时任务

caddy自动https

相比nginx的文件服务，美观，便于安装


### 使用

可以将downLoad目录复制到根目录

配置 Caddyfile 改为自己的域名

配置 conf.yml 指定需要下载的realease 定时任务字符串

定时任务配置参考: https://crontab.guru/

### 启动服务

安装docker
```yaml
curl -fsSL get.docker.com -o get-docker.sh && sh get-docker.sh --mirror Aliyun&&systemctl enable docker&&systemctl start docker
```

安装docker-compose
```yaml
curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose &&chmod +x /usr/local/bin/docker-compose
```

后台启动
```yaml
docker-compose up -d
```

查看日志

```yaml
docker-compose logs -f 
```


删除容器

```yaml
docker rm -f down caddy
```


