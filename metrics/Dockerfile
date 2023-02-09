FROM golang:1.18

WORKDIR /app

COPY go.mod go.sum app.toml main.go ./

ENV GO111MODULE=on \
#    GOPROXY=https://goproxy.cn,direct
    GOPROXY=https://goproxy.cn,direct \
	GOPRIVATE=g.hz.netease.com

RUN git config --global url."https://easy-ngo:tzKdXSFgD7QU3FZSNszP@github.com/NetEase-Media/easy-ngo.git".insteadOf "https://github.com/NetEase-Media/easy-ngo.git"


RUN go build -o main .

EXPOSE 8888

CMD ["./main","-c","app.toml"]
