SOURCE_FILE=main.go
TARGET_FILE=main
# 缺省任务 不要换行
all: clean build

build:
	go build -o $(TARGET_FILE) $(SOURCE_FILE)
	./$(TARGET_FILE)

clean:
	rm -rf $(TARGET_FILE)

test:
    go test -v main.go

run:
	go run main.go