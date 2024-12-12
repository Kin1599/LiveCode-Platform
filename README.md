# `Итоговый проект курсов ITAM 2024`

Team Members:

1. **Дмитрий Коноплянников** - Капитан. Backend/Frontend developer
2. **Аникин Семён** – Backend developer
3. **Виталий Николаев** – UX/UI designer
4. **Виктория Сучкова** – Frontend developer

Презентация: [тык](https://drive.google.com/drive/folders/1uUtPJNK5ln2FtZSZvMjoy7bGfDDM9ShQ?usp=sharing)
Сервис: [тык](http://217.114.2.64:3000/#)

## Проект `"Платформа для лайвкодинга"`

## Предложенное решение

1. Сосздание  сессий для лайвкодинга
2. Онлайн чат в каждой сессии
3. Регистрация и авторизация
4. Сохранение кода в сессии на 24 часа

## Сборка

##### Локалькая сборка:

1. Скачать docker 
2. Склонировать репозитеорий
3. Заполнить local.yaml, получить credentials и config для Yandex Object Storage. С помощью генератора Hmac, сгенерировать HmacSecret -> вставить его в jwt.go
4. Из корня запустить команду docker compose -f docker-compose-without-traefik.yml  up --build -d
5. Наслаждаться сервисом, фронт на 3000, бек на 80 или 8080 развернут, какоц выберут, но лучше на 80

##### Сборка без docker:

1. Из папки frontend написать 
npm install
npm run dev
2. И в новом терминале из папки backend/cmd/app запустить 
go run main.go --config="../../configs/local.yaml"
