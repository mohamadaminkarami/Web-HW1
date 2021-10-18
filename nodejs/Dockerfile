FROM node:14
WORKDIR /app

COPY ./package.json package.json

RUN yarn install --production

COPY . .

RUN yarn start