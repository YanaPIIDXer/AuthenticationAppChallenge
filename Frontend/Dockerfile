FROM node:11.13.0-alpine

RUN mkdir /app
WORKDIR /app

RUN apk update && \
    apk upgrade && \
    npm install -g npm && \
    npm install -g @vue/cli && \
    npm install -g nuxt

COPY . /app

RUN npm install

ENV HOST 0.0.0.0
EXPOSE 3000
