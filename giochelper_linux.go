//go:build linux
// +build linux

package giochelper

const (
	iocNrBits    uint = 8  // 序数的位长度
	iocMagicBits uint = 8  // 幻数的位长度
	iocSizeBits  uint = 14 // 数据大小的位长度
	iocDirBits   uint = 2  // 数据传输方向的位长度

	iocNrMask    uint = ((1 << iocNrBits) - 1)    // 序数的掩码，0x000000FF
	iocMagicMask uint = ((1 << iocMagicBits) - 1) // 幻数掩码，0x000000FF
	iocSizeMask  uint = ((1 << iocSizeBits) - 1)  // 数据大小掩码，0x00003FFF
	iocDirMask   uint = ((1 << iocDirBits) - 1)   // 数据传输方向掩码，0x00000003

	iocNrShift    uint = 0                            // 序数在命令参数中的位偏移，0
	iocMagicShift uint = iocNrShift + iocNrBits       // 幻数在命令参数中的位偏移，8
	iocSizeShift  uint = iocMagicShift + iocMagicBits // 数据大小大小在命令参数中的位偏移，16
	iocDirShift   uint = iocSizeShift + iocSizeBits   // 数据传输方向在命令参数中的位偏移，30

	iocNone  uint = 0b00 // 没有数据传输
	iocWrite uint = 0b01 // 向设备写入数据，驱动程序必须从用户空间读入数据
	iocRead  uint = 0b10 // 从设备中读取数据，驱动程序必须向用户空间写入数据
)

// 将dir，magic，nr，size四个参数组合成一个cmd参数
//	`dir` 方向
//	`magic` 幻数
//	`nr` 序数
//	`size` 数据大小（字节）
// 返回命令参数
func ioc(dir uint, magic rune, nr, size uint) uint {
	return ((dir << iocDirShift) | (uint(magic) << iocMagicShift) | (nr << iocNrShift) | (size << iocSizeShift))
}

// 构造无参数的命令参数
//	`magic` 幻数
//	`nr` 序数
// 返回命令参数
func Io(magic rune, nr uint) uint {
	return ioc(iocNone, magic, nr, 0)
}

// 构造从驱动程序中读取数据的命令参数
//	`magic` 幻数
//	`nr` 序数
//	`size` 数据大小（字节）
// 返回命令参数
func IoR(magic rune, nr, size uint) uint {
	return ioc(iocRead, magic, nr, size)
}

// 构造从驱动程序中写入数据的命令参数
//	`magic` 幻数
//	`nr` 序数
//	`size` 数据大小（字节）
// 返回命令参数
func IoW(magic rune, nr, size uint) uint {
	return ioc(iocWrite, magic, nr, size)
}

// 构造从驱动程序中写入数据后再读取数据的命令参数
//	`magic` 幻数
//	`nr` 序数
//	`size` 数据大小（字节）
// 返回命令参数
func IoWR(magic rune, nr, size uint) uint {
	return ioc(iocRead|iocWrite, magic, nr, size)
}

// 从命令参数中解析出数据方向，即写进还是读出
//	`cmd` 命令参数
// 返回方向
func IocDir(cmd uint) uint {
	return ((cmd >> iocDirShift) & iocDirMask)
}

// 从命令参数中解析出幻数magic
//	`cmd` 命令参数
// 返回幻数
func IocMagic(cmd uint) rune {
	return rune((cmd >> iocMagicShift) & iocMagicMask)
}

// 从命令参数中解析出序数number
//	`cmd` 命令参数
// 返回序数
func IocNr(cmd uint) uint {
	return ((cmd >> iocNrShift) & iocNrMask)
}

// 从命令参数中解析出用户数据大小
//	`cmd` 命令参数
// 返回数据大小（字节）
func IocSize(cmd uint) uint {
	return ((cmd >> iocSizeShift) & iocSizeMask)
}
