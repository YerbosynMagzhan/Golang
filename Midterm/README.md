# SteamGames

## Project description 

All of us are interested in what games our friends have. This project allows you to finally find out what games your friends have.

## Info about me

- 22B031173
- Yerbosyn Magzhan
- 2nd year graduate 

## Games REST API

### HealthCheck

```
 - GET /health-check
```

### Games

```
 - POST /games
 - GET /games/:id
 - PUT /games/:id
 - DELETE /menus/:id
```

## DB Structure

```
 // Use DBML to define your database structure
 // Docs: https://dbml.dbdiagram.io/docs

 Table "games" {
  "id" SERIAL [pk, increment]
  "name" VARCHAR(100)
  "genre" VARCHAR(50)
  "platform" VARCHAR(50)
  "created_at" timestamp(0) [not null, default: `NOW()`]
  "updated_at" timestamp(0) [not null, default: `NOW()`]
 }

 Table "users" {
  "id" SERIAL [pk, increment]
  "username" VARCHAR(50)
  "email" VARCHAR(100)
  "created_at" timestamp(0) [not null, default: `NOW()`]
  "updated_at" timestamp(0) [not null, default: `NOW()`]
 }

 Table "user_games" {
  "id" SERIAL [pk, increment]
  "user_id" INT
  "game_id" INT
  "created_at" timestamp(0) [not null, default: `NOW()`]
  "updated_at" timestamp(0) [not null, default: `NOW()`]
  }

 Ref:"users"."id" < "user_games"."user_id" 
 Ref:"games"."id" < "user_games"."game_id"
```
