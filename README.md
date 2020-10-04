# bx-backend-trainee-assignment

Реализация сервера, позволяющего следить за изменением цены любого объявления на Авито:

Сервис должен предоставить HTTP метод для подписки на изменение цены. На вход метод получает - ссылку на объявление, email на который присылать уведомления.
После успешной подписки, сервис должен следить за ценой объявления и присылать уведомления на указанный email.
Если несколько пользователей подписались на одно и тоже объявление, сервис не должен лишний раз проверять цену объявления.

### ERD database
![ERDdatabase](https://github.com/demsasha4yt/bx-backend-trainee-assignment/blob/master/assets/erd_db.png)

### References

[Задание](https://github.com/avito-tech/bx-backend-trainee-assignment)
