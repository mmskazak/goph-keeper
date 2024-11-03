# Goph-keeper

Goph-keeper — это серверная часть приложения, предназначенного для безопасного хранения паролей и файлов. Приложение предоставляет API для регистрации пользователей, а также управления паролями и файлами.

## Функциональность

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
         "status": "success",
         "message": "User registered successfully."
    }
     ```
  
  Ошибка:
  ```json
  {
       "status": "error",
       "message": "Username already exists."
  }
  ```
  
### Авторизация

- **Запрос**:

    `POST /register`

    Тело запроса
    ```json
    {
        "username": "your_username",
        "password": "your_password"
    }
    ```

- **Ответ**:

  Успешный ответ
  ```json
  {
    "status": "success",
    "token": "your_jwt_token"
  }
  ```

  Ошибка
  ```json
  {
    "status": "error",
    "message": "Invalid username or password."
  }
  ```

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

  Успешный ответ
  ```json
  {
     "status": "success",
     "message": "Password saved successfully."
  }
  ```

  Ошибка
  ```json
  {
    "status": "error",
    "message": "Invalid username or password."
  }
  ```
  
### Получить пароль
- **Запрос**:

  `GET /pwd/get/{id}`

  - **Ответ**:

    Успешный ответ
    ```json
      {
          "id": "123",
          "title": "My Email",
          "password": "super_secret_password",
          "description": "Email account password"
      }
    ```

###  пароль
- **Запрос**:

  `GET /pwd/get/{id}`

    - **Ответ**:

      Успешный ответ
      ```json
        {
            "id": "123",
            "title": "My Email",
            "password": "super_secret_password",
            "description": "Email account password"
        }
      ```