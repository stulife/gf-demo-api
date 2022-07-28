# GoFrame demo 

第一个goFrame架手架: 
- swaggerUi.
- token验证.
- 自定义上下文对像.
- 业务码处理.
- 重写中间件全局返回码.
- 不同环境配置文件.
- 验证码、登陆操作

jenkins 发布
- export BUILD_ID=dontKillMe
- go build
- nohup ./gf-demo-api >/dev/null 2>&1 &
- sleep 10s

jenkins docker方式构建
- export TAG=gf-demo-api PORT=8000
- docker-compose build
- docker-compose up -d

demo:
- http://101.35.146.71:8000/swagger/