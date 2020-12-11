## Deploy locally

Install Docker & Docker Compose
```
docker build -t test .
docker run -p 0.0.0.0:80:80 test
```
Go to http://localhost
```
docker-compose up
```

## Install Nodejs
https://www.digitalocean.com/community/tutorials/how-to-install-node-js-on-ubuntu-20-04

## How-to MVC nodejs

https://github.com/Symfomany/express-mvc/tree/master/models

## Deploy to Remote Host VM

```
terraform init
terraform plan
terrafor apply
```
