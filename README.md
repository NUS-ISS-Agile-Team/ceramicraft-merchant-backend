
# 注意事项

我先用中文记录一些，提交时再完善并修改为英文。

- 使用go 1.24.0
- 在根目录下执行：go run ./cmd/main.go 来启动项目
- 在项目根目录下使用：swag init -g cmd/main.go 生成swagger文档
- 使用zap生成日志，默认存储在项目根目录下的test.log文件中
