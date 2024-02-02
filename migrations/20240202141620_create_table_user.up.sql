CREATE TABLE users
(
    id              INT             NOT NULL    AUTO_INCREMENT,
    first_name      VARCHAR(255)    NOT NULL,
    last_name       VARCHAR(255),
    role_id         INT             NOT NULL,
    sso_id          INT,
    unique_id       VARCHAR(20)     NOT NULL,
    email           VARCHAR(255)    UNIQUE NOT NULL,
    password        VARCHAR(255),
    phone_number    VARCHAR(32)     UNIQUE,
    created_at      TIMESTAMP       NOT NULL,
    updated_at      TIMESTAMP,

    PRIMARY KEY(id)
) ENGINE = InnoDB;