package mem

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func TestFunc() {
	data := make([][]byte, 0)
	for i := 0; i < 1000; i++ {
		b := make([]byte, 10*1024*1024) // 10MBずつ確保
		for j := range b {
			b[j] = byte(j % 256)
		}
		data = append(data, b)
	}
	fmt.Println("Allocated", len(data)*10, "MB")

}

func Memo() {
	memFile, err := os.Create("mem.pprof")
	if err != nil {
		panic(err)
	}
	defer memFile.Close()

	// === メモリを使う処理 ===
	TestFunc()
	// GCを明示的に走らせてからスナップショットを取る
	// runtime.GC()

	// ヒーププロファイルを出力
	if err := pprof.WriteHeapProfile(memFile); err != nil {
		panic(err)
	}
	fmt.Println("Heap profile written to mem.pprof")
}
