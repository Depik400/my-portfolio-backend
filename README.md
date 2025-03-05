# Мой Портфолио Бэкенд (Написано использую copilot)

![Размер репозитория GitHub](https://img.shields.io/github/repo-size/Depik400/my-portfolio-backend)
![Язык программирования GitHub](https://img.shields.io/github/languages/top/Depik400/my-portfolio-backend)
![Последний коммит GitHub](https://img.shields.io/github/last-commit/Depik400/my-portfolio-backend)
![Проблемы GitHub](https://img.shields.io/github/issues/Depik400/my-portfolio-backend)

Этот репозиторий содержит бэкенд для моего портфолио веб-сайта, предназначенный в основном для обслуживания статических файлов.

## Содержание

- [Установка](#установка)
- [Использование](#использование)
- [Вклад](#вклад)
- [Лицензия](#лицензия)

## Установка

Чтобы получить локальную копию и запустить её, выполните следующие шаги.

### Необходимые условия

- Docker

### Клонирование репозитория

```bash
git clone https://github.com/Depik400/my-portfolio-backend.git
cd my-portfolio-backend
```

### Построение Docker-образа

```bash
docker build -t my-portfolio-backend .
```

### Запуск Docker-контейнера

```bash
docker run -p 80:80 my-portfolio-backend
```

## Использование

После запуска Docker-контейнера, вы можете получить доступ к бэкенду портфолио по адресу `http://localhost:80`. Он будет обслуживать статические файлы для вашего веб-сайта портфолио.

## Вклад

Вклад - это то, что делает сообщество с открытым исходным кодом таким удивительным местом для обучения, вдохновения и создания. Любой ваш вклад **очень ценится**.

1. Форкните проект
2. Создайте свою ветку функции (`git checkout -b feature/YourFeature`)
3. Зафиксируйте свои изменения (`git commit -m 'Добавьте некоторую функцию'`)
4. Отправьте ветку (`git push origin feature/YourFeature`)
5. Откройте запрос на слияние

## Лицензия

Распространяется по лицензии MIT. См. `LICENSE` для получения дополнительной информации.

## Контакт

Depik400 - [Профиль на GitHub](https://github.com/Depik400)
