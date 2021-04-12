FROM node:15.14-alpine as dev

WORKDIR /app

ENV NODE_ENV="development"

COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 8080

ENTRYPOINT ["npm", "run", "serve"]