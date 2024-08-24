# compose-astro-go-nginx-demo

### Command to prepare env vars
```
ln .env frontend/.env
ln .env backend/.env
```

### Commands to run (locally)
```
# make postgres db
docker run postgres...

# run backend
go run backend/main.go

# run frontend
cd frontend && pnpm i && pnpm start
```

### Commands to run (on deployment)
```
docker compose up --build -d
```
