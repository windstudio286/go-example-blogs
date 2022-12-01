FROM golang:alpine
# Tạo thư mục app
RUN mkdir /app
# Đi đến thư mục
WORKDIR /app
ADD go.mod .
ADD go.sum .

RUN go mod download
ADD . .
#https://github.com/githubnemo/CompileDaemon
RUN go install -mod=mod github.com/githubnemo/CompileDaemon

EXPOSE 8000
#CompileDaemon giống như PM2 
ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main
#ENTRYPOINT [ "go", "run","main.go"]

