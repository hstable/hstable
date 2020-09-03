FROM node:lts-alpine
WORKDIR /ui
ADD . .
RUN yarn config set sass_binary_site http://cdn.npm.taobao.org/dist/node-sass -g
RUN yarn
RUN yarn generate
CMD yarn start
