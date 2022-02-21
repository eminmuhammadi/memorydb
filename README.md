# memorydb
In memory database implementation for Golang.

## Result
```bash
$ go run main.go

(...) ←[32m(...)/go/pkg/mod/gorm.io/driver/sqlite@v1.3.1/migrator.go:33
←[0m←[33m[0.000ms] ←[34;1m[rows:-]←[0m SELECT count(*) FROM sqlite_master WHERE type='table' AND name="users"

(...) ←[32m(...)/(...)/memorydb/db/migration.go:10
←[0m←[33m[0.522ms] ←[34;1m[rows:0]←[0m CREATE TABLE `users` (`id` integer,`created_at` datetime,`updated_at` datetime,`deleted_at` datetime,`name` varchar(128) NOT NULL,PRIMARY KEY (`id`))

(...) ←[32m(...)/go/pkg/mod/gorm.io/driver/sqlite@v1.3.1/migrator.go:288
←[0m←[33m[0.095ms] ←[34;1m[rows:0]←[0m CREATE INDEX `idx_users_deleted_at` ON `users`(`deleted_at`,`deleted_at`)

(...) ←[32m(...)/(...)/memorydb/main.go:25
←[0m←[33m[1.086ms] ←[34;1m[rows:1]←[0m INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`) VALUES ("2022-02-21 01:22:37.737","2022-02-21 01:22:37.737",NULL,"Emin Muhammadi") RETURNING `id`

(...) ←[32m(...)/(...)/memorydb/main.go:30
←[0m←[33m[0.000ms] ←[34;1m[rows:1]←[0m SELECT * FROM `users` WHERE name = "Emin Muhammadi" AND `users`.`deleted_at` IS NULL AND `users`.`id` = 1 ORDER BY `users`.`id` LIMIT 1
26c63611-0608-471f-aac6-f298ed1fd557 2022-02-21 01:22:37.7366238 +0100 CET
```