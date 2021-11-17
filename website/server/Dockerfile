FROM node:14-alpine
RUN mkdir -p /client
WORKDIR /client

COPY . .

RUN npm i
EXPOSE 8080

CMD ["npm", "start"]