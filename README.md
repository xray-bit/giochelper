# golang ioc helper

linux ioctl 命令参数构造辅助，用于补充 [golang.org/x/sys](https://pkg.go.dev/golang.org/x/sys) unix 中的缺失。

## 用例

```golang
// ...
arg := 0x12345678
cmd := giochelper.IoW('x', 0, 4)
unix.IoctlSetPointerInt(fd, cmd, arg)
// ...
```

## 参考

1. [IOCTLs](https://docs.kernel.org/userspace-api/ioctl/index.html)
2. ioctl.h

