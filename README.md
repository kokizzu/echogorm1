
# echogorm1

This is example of [how to structure your golang project](https://kokizzu.blogspot.com/2022/05/how-to-structure-layer-your-golang-project.html) article with echo and gorm (you can change it to whatever framework and persistence libraries you like, the structure should still be similar).

```
# MVC
presentation -call-> business -call-> model

# Clean
presentation -call-> business
model -injected-to-> business

presentation should only care about transport and serialization/deserialization
model should only care about DAO and persistence (can be decoupled)
business should only care about business logic

presentation can access business
business can access model

model should not ever depend on business
business should not ever depend on presentation
```

## How to start

```shell
make setup

docker-compose up

mysql -u root -p -h 127.0.0.01 -P 3306
CREATE DATABASE test1;

air
```

## Manual test example

```shell
curl -X POST -d 'email=test&password=pass' http://localhost:1323/guest/login

curl -X POST -H 'content-type: application/json' -d '{"email":"test","password":"pass"}' http://localhost:1323/guest/login

curl -X POST -d 'email=test&password=pass' http://localhost:1323/guest/register
```

## TODO
- censor logs
- add metrics
- golangci lint
- godotenv to load config from env
