# bx-backend-trainee-assignment

Реализация сервера, позволяющего следить за изменением цены любого объявления на Авито:

Сервис должен предоставить HTTP метод для подписки на изменение цены. На вход метод получает - ссылку на объявление, email на который присылать уведомления.
После успешной подписки, сервис должен следить за ценой объявления и присылать уведомления на указанный email.
Если несколько пользователей подписались на одно и тоже объявление, сервис не должен лишний раз проверять цену объявления.

### ERD database
![ERDdatabase](https://github.com/demsasha4yt/bx-backend-trainee-assignment/blob/master/assets/erd_db.png)

## Фрагменты кода, решающие некоторые задачи

### Работа с БД

Создание объявления и/или подписка на существующее объявление

```golang
  // AdsRepository repository
  type AdsRepository struct {
    store *Store
  }

  func txAppendEmails(ctx context.Context, tx pgx.Tx, u *models.Ad) error {
    if u.Emails != nil && len(u.Emails) > 0 {
      for _, email := range u.Emails {
        _, err := tx.Exec(
          ctx,
          `INSERT INTO ads_emails(ad_id, email_id, confirm_token, unsubscribe_token) 
            VALUES ($1, $2, $3, $4)`,
          u.ID, email.ID, email.ConfirmToken, email.UnsubscribeToken,
        )
        if err != nil {
          return err
        }
      }
    }
    return nil
  }

  // Create creates new ad in db
  func (r *AdsRepository) Create(ctx context.Context, u *models.Ad) error {
    tx, err := r.store.db.Begin(ctx)
    if err != nil {
      return err
    }
    defer tx.Rollback(ctx)

    if err := tx.QueryRow(
      ctx,
      "INSERT INTO ads (avito_id, actual) VALUES ($1, $2) RETURNING id",
      u.AvitoID,
      u.Actual,
    ).Scan(
      &u.ID,
    ); err != nil {
      return err
    }

    if err := txAppendEmails(ctx, tx, u); err != nil {
      return err
    }

    if _, err := tx.Exec(
      ctx,
      "INSERT INTO ads_prices (ad_id, price) VALUES($1, $2)",
      u.ID, u.CurrentPrice,
    ); err != nil {
      return err
    }

    return tx.Commit(ctx)
  }

  // CreateEmails append emails to ad in ads_emails table
  func (r *AdsRepository) CreateEmails(ctx context.Context, u *models.Ad) error {
    tx, err := r.store.db.Begin(ctx)
    if err != nil {
      return err
    }
    defer tx.Rollback(ctx)
    if err := txAppendEmails(ctx, tx, u); err != nil {
      return err
    }
    return tx.Commit(ctx)
  }

```

### Отправка уведомления на почту

```golang
  // EmailSender ...
  type EmailSender interface {
    Send(to string, subject string, body string) error
  }

  type emailSender struct {
    conf Config
  }

  // NewEmailSender ...
  func NewEmailSender(conf Config) EmailSender {
    return &emailSender{conf}
  }

  func (e *emailSender) Send(to string, subject string, body string) error {

    // Setup headers
    headers := make(map[string]string)
    headers["From"] = e.conf.SenderAddr
    headers["To"] = to
    headers["Subject"] = subject

    // Setup message
    message := ""
    for k, v := range headers {
      message += fmt.Sprintf("%s: %s\r\n", k, v)
    }
    message += "\r\n" + body

    // Connect to the SMTP Server
    addr := e.conf.ServerHost + ":" + e.conf.ServerPort

    auth := smtp.PlainAuth("", e.conf.Username, e.conf.Password, e.conf.ServerHost)

    // TLS config
    tlsconfig := &tls.Config{
      InsecureSkipVerify: true,
      ServerName:         e.conf.ServerHost,
    }

    // Here is the key, you need to call tls.Dial instead of smtp.Dial
    // for smtp servers running on 465 that require an ssl connection
    // from the very beginning (no starttls)
    conn, err := tls.Dial("tcp", addr, tlsconfig)
    if err != nil {
      return err
    }
    c, err := smtp.NewClient(conn, e.conf.ServerHost)
    if err != nil {
      return err
    }

    // Auth
    if err = c.Auth(auth); err != nil {
      return err
    }

    // To && From
    if err = c.Mail(e.conf.SenderAddr); err != nil {
      return err
    }
    if err = c.Rcpt(to); err != nil {
      return err
    }

    // Data
    w, err := c.Data()
    if err != nil {
      return err
    }

    _, err = w.Write([]byte(message))
    if err != nil {
      return err
    }

    err = w.Close()
    if err != nil {
      return err
    }

    c.Quit()

    return nil
  }

```

### References

[Задание](https://github.com/avito-tech/bx-backend-trainee-assignment)
