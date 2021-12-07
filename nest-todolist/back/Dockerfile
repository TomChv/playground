FROM node:lts-alpine

RUN mkdir /app /app/logs && chown -R node: /app
USER node
WORKDIR /app

EXPOSE ${API_PORT}

COPY package.json tsconfig.json tsconfig.build.json /app/

RUN npm i

COPY nest-cli.json /app/
COPY src/ /app/

ENTRYPOINT ["npm", "run"]

RUN npm run build
CMD ["start:prod"]



