# FamChat
A simple, secure chat app designed to keep your family connected over your home Wi-Fi network. Whether you're in the next room or across the house, stay in touch with those who matter most, all within your local network.

## Stack
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white) 

## Features
- **Emoji Reactions**: Let family members react to messages with a variety of emojis.
- **Message Bubbles with Themes**: Customize chat bubbles with fun themes.

## Entity Relationship Diagram
```mermaid
erDiagram
    USER {
        int id
        string username
        string password_hash
        string email
        datetime created_at
        datetime last_login
        boolean is_active
    }
    PROFILE {
        int user_id
        string bio
        string profile_picture
        string location
        date birthdate
        string status_message
    }
    CHAT {
        int id
        string name
        datetime created_at
        boolean is_group
    }
    MESSAGE {
        int id
        int chat_id
        int user_id
        string content
        datetime sent_at
        boolean is_read
    }
    NOTIFICATION {
        int id
        int user_id
        string type
        string message
        datetime created_at
        boolean is_read
    }
    
    USER ||--o{ PROFILE : has
    USER ||--o{ MESSAGE : sends
    USER ||--o{ NOTIFICATION : receives
    CHAT ||--o{ MESSAGE : contains
    CHAT ||--o{ USER : participants
```

## Project Setup
FamChat uses Docker for deployment and project creation.

```
git clone git@github.com:saltnepperson/FamChat.git

cd FamChat

docker-compose build
docker-compose up
```

## Authors
- [@saltnepperson](https://www.github.com/saltnepperson)
