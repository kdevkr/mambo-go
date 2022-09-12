```shell
go get github.com/eclipse/paho.mqtt.golang
go get github.com/go-co-op/gocron
go get github.com/google/uuid
```

## Add username and password
```shell
docker-compose exec mosquitto mosquitto_passwd -c /mosquitto/config/passwd mosquitto
# Password:
# Reenter password:
```