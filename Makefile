migrateup:
	migrate -path internal/interfaces/postgres/migrations -database 'postgres://postgres:123456@localhost:5432/time?sslmode=disable' -verbose up

migratedown:
	migrate -path internal/interfaces/postgres/migrations -database 'postgres://postgres:123456@localhost:5432/time?sslmode=disable' -verbose down

.PHONY:  migrateup migratedow

#генерате миграции с скл файл то соурс коде(гошный)