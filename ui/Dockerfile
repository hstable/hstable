FROM node:lts-alpine
WORKDIR /ui
RUN apk add tzdata 
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN apk del tzdata
ADD . .
RUN yarn config set sass_binary_site http://cdn.npm.taobao.org/dist/node-sass -g
RUN yarn
RUN yarn build
CMD yarn start
