FROM node:22-alpine AS dev

WORKDIR /app

ENV NODE_ENV="development"

COPY package*.json ./

RUN npm ci

COPY . .

EXPOSE 8080

ENTRYPOINT ["npm", "run", "dev"]
