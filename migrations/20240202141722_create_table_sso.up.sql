CREATE TABLE sso_users
(
    id              INT             NOT NULL    AUTO_INCREMENT,
    provider        VARCHAR(100)    NOT NULL,
    external_id     VARCHAR(255)    NOT NULL,
    access_token    VARCHAR(255)    NOT NULL,
    refresh_token   VARCHAR(255),
    expires_at      TIMESTAMP       NOT NULL,

    PRIMARY KEY(id)
) ENGINE = InnoDB;