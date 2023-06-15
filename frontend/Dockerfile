FROM node:latest

RUN npm i -g pnpm

ENV MODE development

WORKDIR /kanban-frontend
COPY package.json pnpm-lock.yaml ./
RUN pnpm install

EXPOSE 3333
EXPOSE 8081
EXPOSE 24678

COPY src/* ./src/
COPY public/* ./public/
COPY *.json ./
COPY *.yaml ./
COPY *.js ./
COPY *.cjs ./ 

ENTRYPOINT ["pnpm", "run", "dev"]

