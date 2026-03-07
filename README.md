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
./custom-gcl run ./...
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
        type: "module"
        description: Checks log messages.
```
**Конфигурация .custom-gcl.yml**
```yml
version: v2.10.1
plugins:
  - module: 'github.com/gtimofej0303/loglinter'
    version: v0.1.0
```

## Поддерживаемые логгеры
Анализатор распознаёт вызовы методов у следующих идентификаторов:

`slog` — стандартная библиотека `log/slog`

`zap` — `go.uber.org/zap`

`log` — стандартная библиотека `log`

`logger` — любая переменная с именем `logger`

Поддерживаемые методы: `Info`, `Error`, `Warn`, `Debug`, `Fatal`, `Panic` и их форматированные варианты (`Infof`, `Errorf`, `Warnf`, `Debugf`).

## Примеры ошибок
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

> Примечание:
слово passport не считается нарушением — проверка на " pass " (с пробелами) исключает его из ложных срабатываний.

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
    version: v0.1.0
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
        type: "module"
        description: Checks log messages.
```

### 4. Собрать кастомный бинарник golangci-lint с плагином
```bash
golangci-lint custom
```
Это создаст исполняемый файл ./custom-gcl в текущей директории.

### 5. Запустить линтер
```bash
./custom-gcl run ./...
```

> Примечание:
команда golangci-lint custom читает .custom-gcl.yml и автоматически подтягивает плагин как Go-модуль — устанавливать loglinter отдельно не нужно.

## Кастомные паттерны

Вы можете задать свои запрещённые слова и regexp-паттерны в `.golangci.yml`:

```yaml
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
          words:
              - "secret" #Пример
          patterns:
            - "credit.?card" #Пример
```