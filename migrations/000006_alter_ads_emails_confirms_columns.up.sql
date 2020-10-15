ALTER TABLE ads_emails ADD confirm BOOLEAN NOT NULL DEFAULT false;
ALTER TABLE ads_emails ADD confirm_token VARCHAR NOT NULL;
ALTER TABLE ads_emails ADD unsubscribe_token VARCHAR NOT NULL;
