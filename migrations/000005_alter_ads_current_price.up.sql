ALTER TABLE ads ADD current_price INT NOT NULL DEFAULT 0 CHECK (current_price >= 0);
