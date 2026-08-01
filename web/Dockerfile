FROM node:22-alpine AS dev

WORKDIR /app

ENV NODE_ENV="development"

COPY package*.json ./

RUN npm ci

COPY . .

EXPOSE 8080

ENTRYPOINT ["npm", "run", "dev"]

FROM node:22-alpine AS build

WORKDIR /app

COPY package*.json ./

RUN npm ci

COPY . .

ARG MP_VERSION=dev

RUN MP_VERSION=$MP_VERSION npm run build

FROM nginx:1.29-alpine AS prod

ARG MP_VERSION=dev

ENV MP_VERSION=$MP_VERSION

COPY --from=build /app/dist /usr/share/nginx/html

COPY nginx.conf /etc/nginx/conf.d/default.conf

COPY config-entrypoint.sh /docker-entrypoint.d/40-menuplanner-config.sh

RUN chmod +x /docker-entrypoint.d/40-menuplanner-config.sh

EXPOSE 80
