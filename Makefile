SOURCE_FILE=main.go
TARGET_FILE=main
# 缺省任务 不要换行
all: clean build

#定义build任务
build:
	go build -o $(TARGET_FILE) $(SOURCE_FILE)
	./$(TARGET_FILE)

#定义clean任务
clean:
	rm -rf $(TARGET_FILE)

#定义test任务
test:
	go test -v main.go
run:
	go run main.go