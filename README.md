# githubDown

## 定时下载github realease到文件服务器

### 使用

可以将downLoad目录复制到根目录

配置 Caddyfile 改为自己的域名

配置 conf.yml 指定需要下载的realease 定时任务字符串

定时任务配置参考: https://crontab.guru/

### 启动服务

后台启动
```yaml
docker-compose up -d
```

删除容器

```yaml
docker rm -f down caddy
```



