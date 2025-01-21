# 格式化swagger注释
swag fmt > ./docs/log.log &
# 根据项目中注释生成swagger文档
swag init --md ./docs > ./docs/log.log &
# 启动main
go run ./main.go -env dev