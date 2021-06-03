# Linters

## go vet ./...

```
# gopkg.in/olebedev/go-duktape.v3
duk_logging.c: In function ‘duk__logger_prototype_log_shared’:
duk_logging.c:184:64: warning: ‘Z’ directive writing 1 byte into a region of size between 0 and 9 [-Wformat-overflow=]
  184 |  sprintf((char *) date_buf, "%04d-%02d-%02dT%02d:%02d:%02d.%03dZ",
      |                                                                ^
In file included from /usr/include/stdio.h:867,
                 from duk_logging.c:5:
/usr/include/x86_64-linux-gnu/bits/stdio2.h:36:10: note: ‘__builtin___sprintf_chk’ output between 25 and 85 bytes into a destination of size 32
   36 |   return __builtin___sprintf_chk (__s, __USE_FORTIFY_LEVEL - 1,
      |          ^~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
   37 |       __bos (__s), __fmt, __va_arg_pack ());
      |       ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
# github.com/ethereum/go-ethereum/tests/fuzzers/trie
tests/fuzzers/trie/trie-fuzzer.go:72:23: method ReadByte() byte should have signature ReadByte() (byte, error)
```

## golangci-lint run

```
consensus/ethash/algorithm.go:154:12: unsafeptr: possible misuse of reflect.SliceHeader (govet)
        header := *(*reflect.SliceHeader)(unsafe.Pointer(&dest))
                  ^
consensus/ethash/algorithm.go:157:37: unsafeptr: possible misuse of reflect.SliceHeader (govet)
        cache := *(*[]byte)(unsafe.Pointer(&header))
                                           ^
consensus/ethash/algorithm.go:286:12: unsafeptr: possible misuse of reflect.SliceHeader (govet)
        header := *(*reflect.SliceHeader)(unsafe.Pointer(&dest))
```

## make test

```
# gopkg.in/olebedev/go-duktape.v3
duk_logging.c: In function ‘duk__logger_prototype_log_shared’:
duk_logging.c:184:64: warning: ‘Z’ directive writing 1 byte into a region of size between 0 and 9 [-Wformat-overflow=]
  184 |  sprintf((char *) date_buf, "%04d-%02d-%02dT%02d:%02d:%02d.%03dZ",
      |                                                                ^
In file included from /usr/include/stdio.h:867,
                 from duk_logging.c:5:
/usr/include/x86_64-linux-gnu/bits/stdio2.h:36:10: note: ‘__builtin___sprintf_chk’ output between 25 and 85 bytes into a destination of size 32
   36 |   return __builtin___sprintf_chk (__s, __USE_FORTIFY_LEVEL - 1,
      |          ^~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
   37 |       __bos (__s), __fmt, __va_arg_pack ());
      |       ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
```

```
--- FAIL: TestParseNode (0.00s)
    urlv4_test.go:191: test "enode://1dd9d65c4552b5eb43d5ad55a2ee3f56c6cbc1c64a5c8d659f51fcd51bace24351232b8d7821617d2b29b54b81cdefb9b3e9c37d7fd5f63270bcc9e1a6f6a439@invalid.:3":
          got error `lookup invalid.: Temporary failure in name resolution`, expected `no such host`
FAIL
FAIL    github.com/ethereum/go-ethereum/p2p/enode       1.512s
```
