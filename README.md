# Server Profiler

Send an e-mail when the process ends.



The format of the e-mail:
```text
The command name is [ls]
The start time is 2025-12-1 22:44:19
The end time is 2025-12-1 22:44:19

```

## How to use?

1. Replace `<XXX>` in the `pkg/sender.go`.
2. Run `go build`.
3. Run like `./profiler <command>`, such as `./profiler python train.py`.