package cpu

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func Cpu() {
	f, err := os.Create("cpu.pprof")
	if err != nil {
		panic(err)
	}
	pprof.StartCPUProfile(f)
	defer func() {
		pprof.StopCPUProfile()
		f.Close()
	}()

	// === ここから計測したい処理 ===
	sum := 0
	for i := 0; i < 1_000_000_000; i++ {
		sum += i % 7
	}
	fmt.Println(sum)
	// === ここまで計測 ===
}
