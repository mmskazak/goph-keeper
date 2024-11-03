# Goph-Keeper

Goph-Keeper — это серверное приложение, предназначенное для безопасного хранения паролей и файлов.
Приложение предоставляет API для регистрации пользователей, а также управления паролями и файлами.

---
## Доступ к приложению

### Регистрация

- **Запрос**:  
  
    `POST /register`  
  
    Тело запроса:
    ```json
    {
        "username": "your_username", // Уникальное имя
        "password": "your_password"  // Пароль для входа в приложение
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

    `POST /login`

    Тело запроса
    ```json
    {
        "username": "your_username", // Уникальное имя
        "password": "your_password"  // Пароль для входа в приложение
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

###  Удалить пароль
- **Запрос**:

  `GET /pwd/delete/{id}`

  - **Ответ**:

  Успешный ответ
  ```json
  {
    "status": "success",
    "message": "Password deleted successfully."
  }
  ```

###  Обновить пароль
- **Запрос**:

  `GET /pwd/update/{id}`

  - **Тело запроса**:
   ```json
  {
    "title": "My Email Updated",
    "password": "new_super_secret_password",
    "description": "Updated email account password"
  }
  ```
  - **Ответ**:

    Успешный ответ
    ```json
    {
    "status": "success",
    "message": "Password updated successfully."
    }
    ```
    
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
## Управление файлами

### Сохранить файл
- **Запрос**:

  `POST /file/save`

  Тело запроса
    ```json
    {
      "title": "My Document",
      "description": "Important document",
      "file_data": "<base64_encoded_file_data>"
    }
    ```

- **Ответ**:

  Успешный ответ
  ```json
  {
    "status": "success",
    "message": "File saved successfully."
  }
  ```
### Получить файл
  - **Запрос**:

    `GET /file/get/{id}`

    - **Ответ**:

    Успешный ответ
    ```json
    {
      "file_id": "123",
      "file_data": "<base64_encoded_file_data>"
    }
    ```
### Удалить файл
  - **Запрос**:
  
  `GET /file/delete/{id}`

  - **Ответ**:

  Успешный ответ
   ```json
  {
    "status": "success",
    "message": "File deleted successfully."
  }
   ```
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