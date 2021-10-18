FROM node:14-alpine
WORKDIR /app

COPY ./package.json package.json

RUN yarn install --production

COPY . .

CMD [ "node", "src/index.js" ]