FROM node:20.9.0

# 作業ディレクトリを設定
WORKDIR /usr/src/app

ENV LANG=ja_JP.UTF-8
ENV LANGUAGE=ja_JP:ja
ENV LC_ALL=ja_JP.UTF-8

# Japaneseロケールの設定
RUN apt-get update \
    && apt-get install -y locales \
    && locale-gen ja_JP.UTF-8

RUN localedef -f UTF-8 -i ja_JP ja_JP.utf8

EXPOSE 5173


