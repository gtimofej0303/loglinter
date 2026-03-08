# loglinter

![CI](https://github.com/gtimofej0303/loglinter/actions/workflows/ci.yml/badge.svg)

`loglinter` — статический анализатор Go, реализованный как плагин для [golangci-lint](https://golangci-lint.run/). Проверяет сообщения в вызовах логгеров (`slog`, `zap`, `log`, `logger`) на соответствие набору правил форматирования.

## Правила

| # | Правило | Нельзя | Можно |
|---|---------|------------------|--------------------|
| 1 | Сообщение должно начинаться со строчной буквы | `"Hello"` | `"hello"` |
| 2 | Только английские буквы | `"ошибка"` | `"error"` |
| 3 | Нет спецсимволов и emoji | `"cool!"`, `"done ✅"` | `"cool"`, `"done"` |
| 4 | Нет секретных данных (`password`, `token` и др.) | `"password changed"` | `"user data updated"` |

## Требования

- Go 1.21+
- [golangci-lint](https://golangci-lint.run/) v2.10.1+ с поддержкой module-плагинов (`custom-gcl`)


## Установка и сборка
### 1. Клонировать репозиторий
```bash
git clone https://github.com/gtimofej0303/loglinter
cd loglinter
```
### 2. Установить зависимости
```bash
go mod tidy
```

### 3. Собрать кастомный бинарник golangci-lint с плагином
```bash
# Установить утилиту custom-gcl
go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.10.1
# Собрать кастомный бинарь через .custom-gcl.yml
golangci-lint custom
```
Это создаст исполняемый файл ./custom-gcl в текущей директории.

## Использование
Запуск через кастомный golangci-lint
```bash
./custom-gcl run ./... --max-same-issues=0
```

**Конфигурация .golangci.yml**
```yml
version: "2"

linters:
  default: none
  enable:
    - loglinter
  settings:
    custom:
      loglinter:
        type: module
        description: Checks log messages.
        settings:
          enabled: true # Включение/выключение плагина
          enable_lowercase:   true # Включение/выключение правила о больших буквах
          enable_english:     true # Включение/выключение правила об английских буквах
          enable_specchars:   true # Включение/выключение правила о спецсимволах
          enable_sensitive:   true # Включение/выключение правила о чувствительных данных
          enable_custom:      true # Включение/выключение кастомных паттернов
          words: # Дополнительные чувствительные ключевые слова
            - "internal"
            - "secret"
            - "debug_mode"
          patterns: # Кастомные паттерны
            - "credit.?card"
            - "social.?security"
```

## Поддерживаемые логгеры
Анализатор распознаёт вызовы методов у следующих идентификаторов:

`slog` — стандартная библиотека `log/slog`

`zap` — `go.uber.org/zap`

`log` — стандартная библиотека `log`

`logger` — любая переменная с именем `logger`

Поддерживаемые методы: `Info`, `Error`, `Warn`, `Debug`, `Fatal`, `Panic` и их форматированные варианты (`Infof`, `Errorf`, `Warnf`, `Debugf`).

## Примеры использования
``` go
// Нарушение 1: заглавная буква
logger.Info("User logged in")
// loglinter: log message must start with a lowercase letter: "User logged in"

// Нарушение 2: не английский язык
logger.Info("пользователь вошёл")
// loglinter: log message must contain only English letters: "пользователь вошёл"

// Нарушение 3: чувствительные данные
logger.Info("user password reset")
// loglinter: log message must not contain sensitive data ("password"): "user password reset"

// Нарушение 4: спецсимволы и emoji
logger.Info("success!")
logger.Info("done ✅")
// loglinter: log message must not contain special characters: "success!"
// loglinter: log message must not contain special characters or emoji: "done ✅"

// Корректные вызовы
logger.Info("user logged in")
logger.Info("connection established")
logger.Debug("processing request")
```

## Тестирование 
```bash
# Запустить все тесты
go test ./...

# Только unit-тесты правил
go test ./pkg/analyzer/rules/...

# Интеграционный тест анализатора
go test ./pkg/analyzer/...

```

## Чувствительные ключевые слова
Анализатор флагирует сообщения, содержащие следующие слова (без учёта регистра):

`password`, `passwd`, `pass`, `token`, `apikey`, `api_key`, `secret`, `private_key`, `auth`

В файле конфигурации есть влзможность добавить свои чувствительные ключевые слова.

> Примечание:
слово passport не считается нарушением — проверка на " pass " (с пробелами) исключает его из ложных срабатываний.

## Авто-исправление
В проекте реализовано **авто-исправление**, оно работает с двумя правилами следующим образом:
- Если в сообщении первая буква **заглавная** - заменяет её на **строчную**.
- Если в сообщении присутствуют **спецсимволы или эмодзи** - удаляет их, заменяя **пустой строкой**.

Авто-исправление происходит прямо внутри кода. 
Запускается авто-исправление **следующей командой**:
```bash
./custom-gcl run --fix ./...
```

## Интеграция в проект
### 1. Установить golangci-lint
```bash
go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.10.1
```

### 2. Создать `.custom-gcl.yml` в корне проекта
```yml
version: v2.10.1
plugins:
  - module: 'github.com/gtimofej0303/loglinter'
    version: v0.2.0
```
### 3. Создать .golangci.yml
```yml
version: "2"

linters:
  default: none
  enable:
    - loglinter
  settings:
    custom:
      loglinter:
        type: module
        description: Checks log messages.
        settings:
          enabled: true
          enable_lowercase:   true
          enable_english:     true
          enable_specchars:   true
          enable_sensitive:   true
          enable_custom:      true
          words:
            - "internal"
            - "secret"
            - "debug_mode"
          patterns:
            - "credit.?card"
            - "social.?security"
```

### 4. Собрать кастомный бинарник golangci-lint с плагином
```bash
golangci-lint custom
```
Это создаст исполняемый файл ./custom-gcl в текущей директории.

### 5. Запустить линтер
```bash
./custom-gcl run ./... --max-same-issues=0
```