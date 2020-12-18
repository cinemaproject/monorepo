FROM node:15.4 as builder

WORKDIR /src

COPY package.json ./
COPY package-lock.json ./
COPY index.html ./
COPY gulpfile.js ./
ADD app ./app
ADD images ./images
ADD styles ./styles
ADD templates ./templates
ADD views ./views

RUN npm install
RUN npm i -g gulp

RUN gulp

FROM clearlinux/nginx:1.17

COPY --from=builder /src/dist /var/www/html
