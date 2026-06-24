package sqlite3

/*
// AES-NI 硬件加速路径（aes_hardware.c）仅在 x86（amd64/386）下编译，
// 需要同时启用 SSE4.1 与 AES 指令集；其它架构走软件实现，不要传这些 x86 专有选项。
#cgo amd64 386 CFLAGS: -msse4.1 -maes
#cgo CFLAGS: -DUSE_LIBSQLITE3 -DCODEC_TYPE=CODEC_TYPE_AES128 -DSQLITE_HAS_CODEC=1
*/
import "C"
