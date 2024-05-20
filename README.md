# wh-api

**wh-api** - сервис для работы с товарами на
множестве складов.

## Начало работы

Для запуска приложения вам потребуются Git, Docker и Makefile. По умолчанию, бэкэнд работает на порту 8088, а Postgres - на порту 5432.

### Проверка занятости портов
```
sudo lsof -i :5432

sudo lsof -i :8088
```
Если порты заняты, их необходимо освободить:
```
sudo kill -9 {PID}
```

## Инструкция по разворачиванию
1. Склонируйте репозиторий
```bash
git clone git@github.com:kisnikita/wh-api.git
```
2. Выполните команду сброки (если необходимо, используйте `sudo`)
```
make build
```
3. Выполните команду запуска
```
make up
```
Эта команда выполнит следующие действия:
- Запустит контейнер Postgres. 
- Запустит контейнер с миграциями, выполнит их, а затем остановит.
- Запустит контейнер бэкэнда.

**Теперь сервис готов к использованию.**

## Коллекция в Postman

Описанием методов и тестовые запросы к ним:

https://documenter.getpostman.com/view/35019742/2sA3QmFb5v
