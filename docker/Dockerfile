FROM debian:bookworm-slim@sha256:90522eeb7e5923ee2b871c639059537b30521272f10ca86fdbbbb2b75a8c40cd

# 必要なパッケージをインストール
RUN apt update && apt install -y --no-install-recommends \
    wget \
    curl \
    make \
    ca-certificates

# Go をインストール
RUN wget https://go.dev/dl/go1.24.3.linux-amd64.tar.gz && \
    rm -rf /usr/local/go && \
    tar -C /usr/local -xzf go1.24.3.linux-amd64.tar.gz && \
    rm go1.24.3.linux-amd64.tar.gz

# 環境変数設定
ENV GOROOT=/usr/local/go
ENV PATH=$GOROOT/bin:/usr/local/bin:$PATH

# golangci-lint をインストール
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
    | sh -s -- -b /usr/local/bin v2.1.6

# 動作確認
RUN go version && make --version && golangci-lint --version