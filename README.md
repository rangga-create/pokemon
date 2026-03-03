# PokemonBE API

Base URL: `http://localhost:8080`

All responses use the same shape:
```json
{
  "status": "success|error",
  "message": "string",
  "code": 200,
  "data": {}
}
```

## GET /supabase
Returns Supabase URL and public API key (from environment variables).

Success (200):
```json
{
  "status": "success",
  "message": "supabase config",
  "code": 200,
  "data": {
    "supabase_url": "https://gdcbtiaeuepztqrcssvj.supabase.co",
    "supabase_api_key": "sb_publishable_..."
  }
}
```

Error (500):
```json
{
  "status": "error",
  "message": "SUPABASE_URL atau SUPABASE_API_KEY belum diset",
  "code": 500
}
```

## POST /register
Body:
```json
{
  "username": "ash",
  "password": "pikachu123"
}
```

Success (201):
```json
{
  "status": "success",
  "message": "register berhasil",
  "code": 201,
  "data": {
    "user": {
      "id": 1,
      "username": "ash",
      "saldo_uang": 0
    }
  }
}
```

Error (400):
```json
{
  "status": "error",
  "message": "input tidak valid",
  "code": 400
}
```

## POST /login
Body:
```json
{
  "username": "ash",
  "password": "pikachu123"
}
```

Success (200):
```json
{
  "status": "success",
  "message": "login berhasil",
  "code": 200,
  "data": {
    "token": "jwt_token_here",
    "user": {
      "id": 1,
      "username": "ash",
      "saldo_uang": 0
    }
  }
}
```

Error (401):
```json
{
  "status": "error",
  "message": "username atau password salah",
  "code": 401
}
```

## GET /profile
Protected with JWT.

Success (200):
```json
{
  "status": "success",
  "message": "ini endpoint yang diproteksi JWT",
  "code": 200,
  "data": {
    "user_id": 1
  }
}
```
