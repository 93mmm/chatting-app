package models

type ChatAdd struct {
    CreatorId   int     `json:"creatorId"`
    Members     []int   `json:"members"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
}

type EditChat struct {
    ChatId      int     `json:"id"`
    Name        *string `json:"name"`
    Description *string `json:"description"`
}

type MoveUserInChat struct {
    ChatId      int     `json:"chatId"`
    UserId      int     `json:"userId"`
}

type UpdateUserInChat struct {
    ChatId      int     `json:"chatId"`
    UserId      int     `json:"userId"`
    Role        rune    `json:"role"`
}

type User struct {
    Id          int     `json:"id"`
    Role        rune    `json:"role"`
}

type ChatInfo struct {
    ChatId      int     `json:"chatId"`
    CreatorId   int     `json:"creatorId"`
    Members     []User  `json:"members"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
}
