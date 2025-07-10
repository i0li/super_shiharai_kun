## 🚀 Getting Started

### 1. リポジトリをクローン
```
git clone <this repository's URL>
cd super_shiharai_kun
```

### 2. 環境変数ファイルをコピー
```
cp .env.example .env
```

### 3. コンテナを起動
```
docker compose up -d
```

### 4. ユーザ登録APIを叩いてみる
以下のリクエストを送ってみて204が返ってくればOK
```
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "company_name": "testCorp",
    "name": "testUser",
    "email": "test@example.com",
    "password": "testPassword"
  }'
```

---

#### 📝 PostgreSQL コンテナに入る（Memo）
```
docker compose exec db psql -U myuser -d super_shiharai_kun
```