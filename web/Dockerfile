FROM node:22-alpine as dev

WORKDIR /app

ENV NODE_ENV="development"
# webpack 4 calls crypto md4, removed from the OpenSSL 3 default provider
ENV NODE_OPTIONS="--openssl-legacy-provider"

COPY package*.json ./

RUN npm ci

COPY . .

EXPOSE 8080

ENTRYPOINT ["npm", "run", "serve"]
