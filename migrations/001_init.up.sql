CREATE TABLE Users (
    id          BIGSERIAL   PRIMARY KEY,
    email       VARCHAR(63) UNIQUE NOT NULL,
    username    VARCHAR(63) UNIQUE,
    pwdHash     TEXT        NOT NULL,
    regAt       TIMESTAMP   DEFAULT NOW()
);

-- some-id + 20 000 000 000 => group chat id
-- some-id                  => chat with user some-id
CREATE TABLE Chats (
    id          BIGSERIAL   PRIMARY KEY,
    creatorId   bigint,
    name        VARCHAR(63) NOT NULL,
    description VARCHAR(255),
    FOREIGN KEY (creatorId) REFERENCES Users(id)
);

CREATE TABLE ChatParticipants (
    userId      bigint      NOT NULL,
    chatId      bigint      NOT NULL,
    joinedAt    TIMESTAMP   DEFAULT NOW(),
    role        char        DEFAULT 'm',
    PRIMARY KEY (userId, chatId),
    FOREIGN KEY (userId) REFERENCES Users(id),
    FOREIGN KEY (chatId) REFERENCES Chats(id)
);

CREATE TABLE Messages (
    messageId   BIGSERIAL   PRIMARY KEY,
    userId      bigint      NOT NULL,
    chatId      bigint      NOT NULL,
    localMsgId  SERIAL      UNIQUE NOT NULL,
    sent        TIMESTAMP   DEFAULT NOW(),
    text        TEXT
);
