CREATE TABLE Users (
    id           BIGSERIAL   PRIMARY KEY,
    passwordHash TEXT        NOT NULL,
    username     VARCHAR(63) UNIQUE,
    email        VARCHAR(63) UNIQUE NOT NULL,
    regAt        TIMESTAMP   DEFAULT NOW()
);

-- some-id + 20 000 000 000 => group chat id
-- some-id                  => chat with user some-id
CREATE TABLE Chats (
    id          bigint       PRIMARY KEY,
    chatName    VARCHAR(63),
    FOREIGN KEY (id) REFERENCES Users(id)
);

CREATE TABLE ChatParticipants (
    userId          bigint    NOT NULL,
    chatId          bigint    NOT NULL,
    joinedAt        TIMESTAMP DEFAULT NOW(),
    role            char      DEFAULT 'm',
    FOREIGN KEY (userId) REFERENCES Users(id),
    FOREIGN KEY (chatId) REFERENCES Chats(id)
);

CREATE TABLE Messages (
    messageId        BIGSERIAL PRIMARY KEY,
    userId           bigint    NOT NULL,
    chatId           bigint    NOT NULL,
    messageIdForUser SERIAL    UNIQUE NOT NULL,
    sent             TIMESTAMP DEFAULT NOW(),
    text             TEXT
);

-- TODO: test these databases
