#使うDockerイメージ
FROM mysql:8.0.20

#日本語化
RUN apt-get update
RUN apt-get -y install locales-all
ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP:ja
ENV LC_ALL ja_JP.UTF-8

#MySQL設定ファイルをイメージ内にコピー
COPY ./mysqld_charset.cnf /etc/mysql/conf.d/mysqld_charset.cnf
#テーブル作成
COPY ./tables.sql /docker-entrypoint-initdb.d/tables.sql

#docker runに実行される
CMD ["mysqld"]