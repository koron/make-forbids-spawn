# (GNU) make forbids to spawn after SIGINT

Logs in this files are recorded on MSYS2/Windows 10

## Try without make

```console
$ make
$ ./make-forbids-spawn
2021/10/19 23:46:11 The program is waiting SIGINT or SIGTERM. Press Ctrl-C
2021/10/19 23:46:12 recived signal (Go)
received signal (echo)
2021/10/19 23:46:12 completed
```

## Try with make

```console
$ make test
./make-forbids-spawn
2021/10/19 23:46:27 The program is waiting SIGINT or SIGTERM. Press Ctrl-C
2021/10/19 23:46:30 recived signal (Go)
make: *** [Makefile:5: test] エラー 127
```

`echo` won't be spawned. and no errors are logged.

```console
$ make test-without-spawn
./make-forbids-spawn -spawn=false
2021/10/19 23:50:20 The program is waiting SIGINT or SIGTERM. Press Ctrl-C
2021/10/19 23:50:21 recived signal (Go)
2021/10/19 23:50:21 completed
```

If skip to spawn `echo`, it completed as expectedly.

## WTF

Where does this behavior come from?
What is the cause?

This did not happen on Linux (Xubuntu 21.04).
