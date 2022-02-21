# memorydb
In memory database implementation for Golang.

- ORM docs: [https://gorm.io/docs/](https://gorm.io/docs/)
- SQLite docs: [https://www.sqlite.org/inmemorydb.html](https://www.sqlite.org/inmemorydb.html)

## Result
```json
[
    {
        "id": "05358e7e-07c9-4285-af0f-5bd53a363616",
        "name": "John Doe"
    },
    {
        "id": "b042fcef-754a-404e-b53f-28c489e7c044",
        "name": "John Smith"
    }
]
```

```bash
$ go run ./main.go

2022/02/21 16:48:38 ←[32m(...)/go/pkg/mod/gorm.io/driver/sqlite@v1.3.1/migrator.go:33
←[0m←[33m[0.504ms] ←[34;1m[rows:-]←[0m SELECT count(*) FROM sqlite_master WHERE type='table' AND name="users"

2022/02/21 16:48:38 ←[32m(...)/memorydb/db/migration.go:10
←[0m←[33m[0.000ms] ←[34;1m[rows:0]←[0m CREATE TABLE `users` (`id` uuid,`created_at` datetime,`updated_at` datetime,`deleted_at` datetime,`name` varchar(128) NOT NULL,PRIMARY KEY (`id`))

2022/02/21 16:48:38 ←[32m(...)/go/pkg/mod/gorm.io/driver/sqlite@v1.3.1/migrator.go:288
←[0m←[33m[0.000ms] ←[34;1m[rows:0]←[0m CREATE INDEX `idx_users_deleted_at` ON `users`(`deleted_at`)

2022/02/21 16:48:38 ←[32m(...)/go/pkg/mod/gorm.io/driver/sqlite@v1.3.1/migrator.go:33
←[0m←[33m[0.000ms] ←[34;1m[rows:-]←[0m SELECT count(*) FROM sqlite_master WHERE type='table' AND name="users"

2022/02/21 16:48:38 ←[32m(...)/go/pkg/mod/gorm.io/driver/sqlite@v1.3.1/migrator.go:33
←[0m←[33m[0.000ms] ←[34;1m[rows:-]←[0m SELECT count(*) FROM sqlite_master WHERE type='table' AND name="users"

2022/02/21 16:48:38 ←[32m(...)/memorydb/db/migration.go:10
←[0m←[33m[0.505ms] ←[34;1m[rows:0]←[0m CREATE TABLE `users` (`id` uuid,`created_at` datetime,`updated_at` datetime,`deleted_at` datetime,`name` varchar(128) NOT NULL,PRIMARY KEY (`id`))

2022/02/21 16:48:38 ←[32m(...)/memorydb/db/migration.go:10
←[0m←[33m[0.505ms] ←[34;1m[rows:0]←[0m CREATE TABLE `users` (`id` uuid,`created_at` datetime,`updated_at` datetime,`deleted_at` datetime,`name` varchar(128) NOT NULL,PRIMARY KEY (`id`))

2022/02/21 16:48:38 ←[32m(...)/go/pkg/mod/gorm.io/driver/sqlite@v1.3.1/migrator.go:288
←[0m←[33m[0.000ms] ←[34;1m[rows:0]←[0m CREATE INDEX `idx_users_deleted_at` ON `users`(`deleted_at`)

2022/02/21 16:48:38 ←[32m(...)/go/pkg/mod/gorm.io/driver/sqlite@v1.3.1/migrator.go:288
←[0m←[33m[0.506ms] ←[34;1m[rows:0]←[0m CREATE INDEX `idx_users_deleted_at` ON `users`(`deleted_at`)

2022/02/21 16:48:38 ←[32m(...)/go/pkg/mod/gorm.io/driver/sqlite@v1.3.1/migrator.go:33
←[0m←[33m[0.000ms] ←[34;1m[rows:-]←[0m SELECT count(*) FROM sqlite_master WHERE type='table' AND name="users"

2022/02/21 16:48:38 ←[32m(...)/memorydb/db/migration.go:10
←[0m←[33m[0.597ms] ←[34;1m[rows:0]←[0m CREATE TABLE `users` (`id` uuid,`created_at` datetime,`updated_at` datetime,`deleted_at` datetime,`name` varchar(128) NOT NULL,PRIMARY KEY (`id`))

2022/02/21 16:48:38 ←[32m(...)/go/pkg/mod/gorm.io/driver/sqlite@v1.3.1/migrator.go:288
←[0m←[33m[1.004ms] ←[34;1m[rows:0]←[0m CREATE INDEX `idx_users_deleted_at` ON `users`(`deleted_at`)

2022/02/21 16:48:38 ←[32m(...)/go/pkg/mod/gorm.io/driver/sqlite@v1.3.1/migrator.go:33
←[0m←[33m[1.013ms] ←[34;1m[rows:-]←[0m SELECT count(*) FROM sqlite_master WHERE type='table' AND name="users"

2022/02/21 16:48:38 ←[32m(...)/memorydb/db/migration.go:10
←[0m←[33m[1.367ms] ←[34;1m[rows:0]←[0m CREATE TABLE `users` (`id` uuid,`created_at` datetime,`updated_at` datetime,`deleted_at` datetime,`name` varchar(128) NOT NULL,PRIMARY KEY (`id`))

2022/02/21 16:48:38 ←[32m(...)/go/pkg/mod/gorm.io/driver/sqlite@v1.3.1/migrator.go:288
←[0m←[33m[0.000ms] ←[34;1m[rows:0]←[0m CREATE INDEX `idx_users_deleted_at` ON `users`(`deleted_at`)

2022/02/21 16:48:38 ←[32m(...)/memorydb/api/view.go:25
←[0m←[33m[1.076ms] ←[34;1m[rows:1]←[0m INSERT INTO `users` (`id`,`created_at`,`updated_at`,`deleted_at`,`name`) VALUES ("05358e7e-07c9-4285-af0f-5bd53a363616","2022-02-21 16:48:38.859","2022-02-21 16:48:38.859",NULL,"John Doe") RETURNING `id`

2022/02/21 16:48:38 ←[32m(...)/memorydb/api/view.go:30
←[0m←[33m[0.529ms] ←[34;1m[rows:1]←[0m SELECT * FROM `users` WHERE name = "John Doe" AND `users`.`deleted_at` IS NULL AND `users`.`id` = "05358e7e-07c9-4285-af0f-5bd53a363616" ORDER BY `users`.`id` DESC LIMIT 1
[2022-02-21 16:48:38]   0ea20c9e-59a4-4144-aa3a-d07e0d45906a    127.0.0.1       200     POST    /users  3.2ms

2022/02/21 16:48:45 ←[32m(...)/memorydb/api/view.go:25
←[0m←[33m[1.472ms] ←[34;1m[rows:1]←[0m INSERT INTO `users` (`id`,`created_at`,`updated_at`,`deleted_at`,`name`) VALUES ("b042fcef-754a-404e-b53f-28c489e7c044","2022-02-21 16:48:45.014","2022-02-21 16:48:45.014",NULL,"John Smith") RETURNING `id`

2022/02/21 16:48:45 ←[32m(...)/memorydb/api/view.go:30
←[0m←[33m[0.000ms] ←[34;1m[rows:1]←[0m SELECT * FROM `users` WHERE name = "John Smith" AND `users`.`deleted_at` IS NULL AND `users`.`id` = "b042fcef-754a-404e-b53f-28c489e7c044" ORDER BY `users`.`id` DESC LIMIT 1
[2022-02-21 16:48:44]   0fa20c9e-59a4-4144-aa3a-d07e0d45906a    127.0.0.1       200     POST    /users  3.0532ms

2022/02/21 16:48:49 ←[32m(...)/memorydb/api/view.go:44
←[0m←[33m[0.790ms] ←[34;1m[rows:2]←[0m SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL
[2022-02-21 16:48:49]   10a20c9e-59a4-4144-aa3a-d07e0d45906a    127.0.0.1       200     GET     /users  1.2944ms
```