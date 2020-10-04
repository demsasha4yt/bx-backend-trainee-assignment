CREATE TABLE ads_emails(
	id BIGSERIAL NOT NULL PRIMARY KEY,
	ad_id BIGINT,
	email_id BIGINT,
	CONSTRAINT fk_ads
		FOREIGN KEY(ad_id)
			REFERENCES ads(id)
				ON DELETE CASCADE,
	CONSTRAINT fk_emails
			FOREIGN KEY(email_id)
				REFERENCES emails(id)
					ON DELETE CASCADE,
	UNIQUE(ad_id, email_id)
);
