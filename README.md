## Repro repo

### Steps to repro:

```
$ docker compose up -d

    [+] Running 1/2
    - Network repro_default   Created
    ✔ Container repro_edgedb  Started

$ go run main.go

    Connected in 328.3618ms
    Client didn't have to reconnect (0s)
    Connected had to reconnect (320.3006ms)
    Client didn't have to reconnect (0s)
    Connected had to reconnect (321.6555ms)
    Client didn't have to reconnect (0s)
    Connected had to reconnect (318.9663ms)
    Client didn't have to reconnect (0s)
    Connected had to reconnect (316.6645ms)
```
