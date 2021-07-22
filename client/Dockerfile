
FROM node:16

WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

COPY package*.json ./

RUN yarn config set registry https://registry.npm.taobao.org
RUN yarn install

COPY . .

CMD [ "npm", "run", "serve" ]