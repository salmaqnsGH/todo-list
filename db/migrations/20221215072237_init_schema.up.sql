CREATE TABLE activities (
    id                  INT             NOT NULL AUTO_INCREMENT,
    email               VARCHAR(255),
    title               VARCHAR(255),
    created_at          datetime,
    updated_at          datetime,
    deleted_at          datetime,
    PRIMARY KEY (id)
)ENGINE = InnoDB;

CREATE TABLE todos (
    id                  INT             NOT NULL AUTO_INCREMENT,
    activity_group_id   INT NOT NULL,
    title               VARCHAR(255),
    is_active           VARCHAR(255),
    priority            VARCHAR(255),
    created_at          datetime,
    updated_at          datetime,
    deleted_at          datetime,
    PRIMARY KEY (id)
)ENGINE = InnoDB;

ALTER TABLE todo ADD FOREIGN KEY (activity_group_id) REFERENCES activity(id);