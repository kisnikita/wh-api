# Аргументация выбора библиотек

## Список пакетов
### 1. `github.com/joho/godotenv`
**Назначение**: загружает переменные среды из файла .env в систему.

**Аргументы**:
- Сокрытие данных в переменных окружения.
- Использование переменных окружения в контейнерах.
### 2. `github.com/sirupsen/logrus`
**Назначение**: Пакет logrus является структурированным логгером для Go, который имеет множество настроек для информативного отображения логов.

**Аргументы**:
- Поддерживает структурированное логирование, улучшая читаемость и анализ логов.
- Гибкая конфигурация и поддержка различных выходных форматов и уровней логирования.

### 3. "github.com/jmoiron/sqlx"
**Назначение**: Расширение стандартной библиотеки database/sql для работы с базами данных.

**Аргументы**:
- Упрощает выполнение запросов и обработку результатов, позволяя использовать маппинг на основе аннотаций полей структур, тем самым облегчая получение данных из БД.
- Поддерживает большое количество разнообразных методов, полезных для работы с БД, позволяя писать менее громоздкие функции.
- Легковесность в сравнении с конкурентами в виде gorm и тд.
### 4. `github.com/lib/pq`
**Назначение**: Пакет pq является чистым драйвером Postgres для пакета database/sql на Go.

**Аргументы**:
- Желание использовать sqlx

### 5. `github.com/spf13/viper`
**Назначение**: Удобная работа с конфигами.

**Аргументы**:
- Читаемость 
- Лёгкость работы благодаря отличной документации
- Следование практикам 12 factor app

### 6. `github.com/google/jsonapi`
**Назначение**: Маршал и анмаршал.

**Аргументы**:
- Обязательное требование по ТЗ

### 7. `github.com/gin-gonic/gin`
**Назначение**: Веб-фреймворк для создания веб-приложений и API.

**Аргументы**:
- Самый читаемый веб-фреймворк на go(субъективно)
- Производительность
- Наличие опыта на данном вреймворке, что уменьшает время разработки