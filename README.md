# web_complier
在线代码运行。

## 介绍
- 在线编译器
- 实现语言： golang + docker + Docker Engine API


使用前需先安装 docker，事先下载运行代码的镜像并重命名，名字与配置文件相对应。
启动容器时，指定映射，与配置文件相对应
```
docker pull python:3 && docker tag python:3 python3
docker pull rust
docker pull golang
```

## 测试

1. 运行代码
```bash
curl -H "Content-Type: application/json" -X POST -d '{"lang":"golang","code":"package main\n\nimport (\n\t\"fmt\"\n)\n\nfunc main() {\n\tfmt.Println(\"Hello, 世界\")\n}","input":""}' http://127.0.0.1:9999/complier/v1/run
```
2. 保存&分享代码
```bash
curl -H "Content-Type: application/json" -X POST -d '{"lang":"golang","code":"package main\n\nimport (\n\t\"fmt\"\n)\n\nfunc main() {\n\tfmt.Println(\"Hello, 世界\")\n}","input":""}' http://127.0.0.1:9999/complier/v1/share
```
3. 获取保存代码
```bash
curl http://127.0.0.1:9999/complier/v1/code\?gid\="cc4u6vebc58n4vg8i7mg"
```
