# golang-backend-template

### 介绍
golang后端接口项目模版

### 启动项目

```shell
# 启动项目
$ ./cmd/run.sh
# swag fmt > ./docs/log.log &
# swag init --md ./docs > ./docs/log.log &
# go run ./main.go -env dev
# 具体查看启动文件`cmd/run.sh`内容

# 项目启动后会调用runnersFactory中注册的所有服务的run方法
# 目前启动:
#   mainRunner:web restful接口服务
#   docsRunner:swagger文档服务
#   pprofRunner:pprof监控服务
# 均根据配置独立端口启动
```

- 直接clone项目缺少配置文件`config/env/config.yaml`; 可通过`config/config.yaml.default`模板文件新建

### 项目目录结构

- config
    - env `对应的环境目录`
        - config.yaml `项目配置文件`
        - ecosystem.config.js `pm2启动配置文件`
    - config_default.yaml `项目配置文件模板`
- docs `文档（swagger）`
- gen `代码生成器`
- internal `项目代码`
    - api `表现层（校验传参，定义返回）`
        - controller `控制器层（路由handle直接处理请求的方法）`
        - middleware `中间件`
    - application `应用层（调用领域对象完成任务，不包含具体业务领域的代码）`
    - domain `领域层（系统的核心，负责处理业务逻辑。定义领域模型，聚合根，领域服务等处理业务）`
        - aggregate `聚合根`
        - entity `实体层，定义实体模型`
        - service `服务层，业务逻辑写在这里，注意这层不要直接操作数据库`
    - infrastructure `基础设施层（对其他层提供通用的技术支持能力，如消息通信，通用工具，配置等的实现）`
      - db
        - repository `存储层，定义数据库操作接口`
        - model `数据库表结构（数据库表均通过结构体生成，妥善修改）`
      - pkg `项目内部使用的工具包`
- log `日志`
- pkg `三方包`
- runner `服务启动脚本（主要web服务，swagger日志服务，pprof监控服务）`

### 参与开发

1. 流程：
   Fork 本仓库 -> 新建 Feat_xxx 分支 -> 提交代码 -> 新建 Pull Request

2. commit规范：
    - 格式：type: desc
    - type用于说明 commit 的类别，只允许使用下面8个标识:
        - feat：新功能（feature）
        - fix：修补功能
        - bugfix: 此项特别针对bug号，用于向测试反馈bug列表的bug修改情况
        - docs：文档（documentation）
        - style： 格式（不影响代码运行的变动）
        - refactor：重构（即不是新增功能，也不是修改bug的代码变动）
        - test：增加测试
        - chore：构建过程或辅助工具的变动
        - revert: feat(pencil): add 'graphiteWidth' option (撤销之前的commit)

