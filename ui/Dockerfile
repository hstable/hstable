FROM node:lts-alpine
WORKDIR /ui
ADD . .
RUN yarn
RUN yarn generate
CMD yarn start
