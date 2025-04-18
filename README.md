# balun_microservices
Репозиторий с кодом для курса "Микросервисы, как в BigTech - 6 поток"

Инициализируем проект
```
go mod init kadz_balun
```

Сборка сервисов
```
docker image build --tag <service>_img --build-arg SERVICE=<auth/users/chats/friends> --file ./Dockerfile .
```