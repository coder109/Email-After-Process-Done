# Server Profiler

Send an e-mail when the process ends.



The format of the e-mail:
```text
The command name is [ls]
The start time is 2025-12-1 22:44:19
The end time is 2025-12-1 22:44:19

```

## How to use?

1. Run `go mod init profiler` and `go mod tidy`.
2. Remove `TEST` in `macro_test.go` and change the necessary information.
3. Run `go build`.
4. Run like `./profiler <command>`, such as `./profiler python train.py`.