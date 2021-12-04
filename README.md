# Backend-разработка на Go. Уровень 2

---

```bash
goose postgres 'port=8100 user=test password=test dbname=test sslmode=disable' up
goose postgres 'port=8110 user=test password=test dbname=test sslmode=disable' up
goose postgres 'port=8120 user=test password=test dbname=test sslmode=disable' up

docker exec -it go_backend_2_replica_0_1 /bin/sh
psql  --username=test -c 'select * from users'
```

