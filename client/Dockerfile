# 開発環境
FROM node:14.15.4-alpine as build-stage
WORKDIR /app
COPY . .
RUN yarn install
RUN yarn add axios vue-axios
RUN npm install --save axios vue-axios
RUN npm install vue-router@3.5.3 --save --legacy-peer-deps   