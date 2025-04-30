# balun_microservices
Репозиторий с кодом для курса "Микросервисы, как в BigTech - 6 поток"

Инициализируем проект
```
go mod init kadz_balun
# Форматирует весь проект (рекурсивно)
go fmt ./...
```

Сборка сервисов
```
docker image build --tag <service>_img --build-arg SERVICE=<auth/users/chats/friends> --file ./Dockerfile .
```

Что сделать нужно 
[ ] k8s blue/green & canary deploy
[ ] ДЗ 2
