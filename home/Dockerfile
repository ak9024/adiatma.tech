FROM node:latest

WORKDIR /app/home

COPY . .

RUN yarn install

RUN yarn build

COPY . .

ENTRYPOINT [ "yarn", "preview", "--host" ]
