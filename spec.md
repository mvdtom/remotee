# Support Tools Service Spec

# Overview

## Problem statement

Пользователь - служба технической поддержки. Оказывают поддержку своим клиентам используя инструменты удалённого доступа к рабочему столу.

В текущем виде рабочий процесс выглядит так:

1. Клиент обращается за тех. поддержкой
2. Сотрудник тех. поддержки даёт клиенту инструкции
3. Клиент скачивает VNC сервер
4. Клиент устанавливает/запускает VNC сервер
5. Сотрудник ТП сообщает адрес и порт для подключения
6. Клиенту необходимо ввести адрес/порт для подключения к Вьюверу сотрудника тех. поддержки
7. Сотрудник начинает оказывать тех. поддержку

Есть желание оптимизировать этот рабочий процесс, исключив из него установку сервера и настройку соединения конечным клиентом.

## Proposed Work

Решением проблемы может служить сервис, который позволяет сотрудникам тех. поддержки создавать и передавать конечным клиентам заранее сконфигурированный VNC сервер, который самостоятельно после запуска будет подключаться к нужному сотруднику тех. поддержки.

Таким образом, сотруднику нужно будет только передать ссылку на скачивание клиенту, а клиенту только скачать предварительно сконфигурированный сервер и запустить его.

# Success Criteria

К моменту готовности проекта сотрудники ТП должен иметь возможность инициировать создание pre-configured сервера и получении ссылки на его скачивание, которую он передаст конечному пользователю.

Конечный пользователь имеет возможность скачать сгенерированный сервер, который после запуска автоматически подключится к Вьюверу сотрудника ТП.

Сервис должен реализовывать возможности администрирования для управления пользователями (генерирующих сервера)  и регуляции их активности.

# User stories

## Получатель тех. поддержки

- имея ссылку на скачивание pre-configured сервера, скачивает и запускает его
- при скачивании видит страницу-предупреждение о том что он скачивая и запуская сервер передаёт контроль за своим PC сотруднику такой-то организации на такой-то IP адрес
- в процессе работы, то есть когда ему оказывается ТП имеет возможность в любой момент приостановить или прервать контроль за своим компьютером
- не может запустить сервер повторно после истечения его времени жизни

## Оказывающий тех. поддержку

- указывая свой IP-адрес и порт, генерирует pre-configured сервер и получает ссылку, по которой данный сервер может быть скачан получателем тех. поддержки
- генерирует другие серверы для каждого нового получающего тех. поддержку

## Менеджер службы тех. поддержки

/todo

## Администратор сервиса

/todo

# Scope

## Requirements (for first stage)

На данный момент сервис должен по запросу на основе базовой версии VNC сервера, модифицируя EXE файл, создавать pre-configured VNC сервер и выдавать ссылку на его скачивание.

По сути нам нужно изменить 4 байта IP-адреса и 2 байта номера порта. Плюс 4 байта - время жизни.

После модификации новый сгенерированный файл сервиса должен быть подписан цифровой подписью.

В конфигурации сервиса должны быть указаны:

- путь к базовому exe файлу
- офсеты записи конфигурационных данных
- путь к сертификату подписи

По полученной ссылке должен скачиваться сервер, который настроен для автоматического подключения к указанному адресу - порту.

Решение должно запускаться на Linux и быть подготовлено для деплоя в докер-контейнере.

## Future Work

- будет введён учёт пользователей и организаций
- реализация механизма подписок
- усиление безопасности и контроля путём указанием в pre-configured сервере только ID, а всю остальную информацию он будет получать через запросы из Сервиса
- Предупреждающая страничка получающего тех. поддержку перед скачиванием
- специальная версия сервера с соответствующим UI

# Alternative Considered

/todo

# Follow-up Tasks

- [ ]  Подготовить каркасный проект
- [ ]  Реализовать отдачу файла
- [ ]  Реализовать генерацию модифицированного файла по запросу
- [ ]  ...

# Полезные ссылки

[Certificates in .NET Core on Linux and Docker](https://medium.com/gsoft-tech/certificates-in-net-core-on-linux-and-docker-29b3d5f09cd6)