# Golang Version
FROM golang:latest

# Working Directory 설정
WORKDIR /node

# 필요한 패키지 설치
#RUN apt-get update && \
#    apt-get install -y vim sqlite3 iproute2 && \
#    rm -rf /var/lib/apt/lists/*

# 기본 File 복사
COPY go.mod go.sum ./
RUN go mod download

# 컨테이너 내부의 Working Directory로 Code 복사
COPY . .

# Build
RUN go build -o main .

# 연결 Port
EXPOSE 2000

# Execute
CMD ["./main"]
