# memorydb
In memory database implementation for Golang.

- ORM docs: [https://gorm.io/docs/](https://gorm.io/docs/)
- SQLite docs: [https://www.sqlite.org/inmemorydb.html](https://www.sqlite.org/inmemorydb.html)

## Build
```bash
bash build.sh
```

## Start memorydb
```
./memorydb --ip 127.0.0.1 --port 8080 --tz "Asia/Baku"
```

## Test on Windows
```bash
telnet 127.0.0.1 8080 
```

## Result
```bash
[{33f4de6b-aa1d-4e88-8f84-cdc42dd0b9ea Hello World
}]
  Who are you?
[{33f4de6b-aa1d-4e88-8f84-cdc42dd0b9ea Hello World
} {218be113-d88e-476d-8307-6809aba8886f Who are you?
}]
  Maybe
[{33f4de6b-aa1d-4e88-8f84-cdc42dd0b9ea Hello World
} {218be113-d88e-476d-8307-6809aba8886f Who are you?
} {b4d2391b-9a07-4d87-adf8-83b60511c255 Maybe
}]
```