# Goph-Keeper

**Goph-Keeper** — это серверное приложение, предназначенное для безопасного хранения паролей и файлов. Приложение предоставляет API для регистрации пользователей, а также управления паролями и файлами. Оно поддерживает работу как через HTTP API, так и через gRPC, предоставляя гибкость для различных сценариев использования.

### Основные функции:

- **Регистрация и аутентификация**: Поддержка API для регистрации и аутентификации пользователей.
- **Безопасное хранение данных**: Функционал для безопасной обработки и хранения данных.
- **Доступ через различные протоколы**: Возможность взаимодействия с приложением через HTTP или gRPC.

### Репозиторий клиента

Репозиторий клиента доступен на GitHub: [Client Goph-Keeper](https://github.com/mmskazak/client-goph-keeper/pull/1).

### Запуск приложения

Для запуска приложения требуется подготовить базу данных PostgreSQL. <br>
Команда `make db` автоматически развернёт контейнер с базой данных PostgreSQL в Docker. <br>
Чтобы запустить приложение с дефолтными настройками достаточно перейти в директорию <br>
`/cmd/goph-keeper/` и выполнить команду `go run .`

### Ручное тестирование приложения

Если вы используете IDE Goland, то для "ручных" тестов можно воспользоваться встроенным HTTP клиентом,
в директории `/xtests` находятся файлы с готовыми http запросами. Если вы хотите воспользоваться своим
клиентом, то можете воспользоваться примерами запросов представленных ниже.

---
### API приложения

### Регистрация

- **Запрос**:  
  
    `POST /register`  
  
    Тело запроса:
    ```json
    {
        "username": "your_username",
        "password": "your_password" 
    }
    ```

-  **Ответ**:

    Успешный ответ
    ```json
    {
         "jwt": "jwt token"
    }
     ```
  
### Авторизация

- **Запрос**:

    `POST /login`

    Тело запроса
    ```json
    {
        "username": "your_username",
        "password": "your_password" 
    }
    ```

- **Ответ**:

  ```json
  {
      "jwt": "jwt token"
  }
  ```
  
---

## Управление паролями

### Сохранить пароль
- **Запрос**:

  `POST /pwd/save` 

  Тело запроса
    ```json
    {
        "title": "My Email",
        "password": "super_secret_password",
        "description": "Email account password"
    }
    ```

- **Ответ**:
    
    Отсутствие ошибки в ответе будет являться критерием успешного запроса.
  
### Получить пароль
- **Запрос**:

  `GET /pwd/get/{id}`

  - **Ответ**:

  Успешный ответ
  ```json
  {
      "id": "123",
      "title": "My Email",
      "description": "Email account password",
      "login": "login for password",
      "password": "super secret password"
  }
  ```

###  Удалить пароль
- **Запрос**:

  `GET /pwd/delete/{id}`

  - **Ответ**:
  
  Отсутствие ошибки в ответе будет являться критерием успешного запроса.
  

###  Обновить пароль
- **Запрос**:

  `GET /pwd/update/{id}`

  - **Тело запроса**:
   ```json
  {
      "id": "123",
      "title": "My Email",
      "description": "Email account password",
      "login": "login for password",
      "password": "super secret password"
  }
  ```
  - **Ответ**:

  Отсутствие ошибки в ответе будет являться критерием успешного запроса.
    
    
###  Получить все пароли
- **Запрос**:

  `GET /pwd/all`

  - **Ответ**:

    Успешный ответ
    ```json
    [
      {
          "id": "123",
          "title": "My Email",
          "password": "super_secret_password",
          "description": "Email account password"
      },
      {
          "id": "124",
          "title": "My Social Media",
          "password": "another_secret_password",
          "description": "Social media password"
      }
    ]
    ```
    
---

## Управление файлами

### Сохранить файл
- **Запрос**:

  `POST /file/save`

  Тело запроса
    ```json
    {
      "title": "My Document",
      "description": "Important document",
      "file_data": "base64_encoded_file_data"
    }
    ```

- **Ответ**:

  Отсутствие ошибки в ответе будет являться критерием успешного запроса.

 
### Получить файл
  - **Запрос**:

    `GET /file/get/{id}`

    - **Ответ**:

    Успешный ответ
    ```json
    {
      "file_id": "123",
      "file_data": "base64_encoded_file_data"
    }
    ```
### Удалить файл
  - **Запрос**:
  
  `GET /file/delete/{id}`

  - **Ответ**:

    Отсутствие ошибки в ответе будет являться критерием успешного запроса.

### Получить все файлы
- **Запрос**:

`GET /file/all`

- **Ответ**:

Успешный ответ
  ```json
  [
    {
      "file_id": "123",
      "title": "My Document",
      "description": "Important document"
    },
    {
      "file_id": "124",
      "title": "My Image",
      "description": "Profile picture"
    }
  ]
```