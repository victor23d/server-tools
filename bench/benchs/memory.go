package benchs

import (
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

type Memory struct {
	Alloc      uint64 `json:"alloc"`
	TotalAlloc uint64 `json:"total_alloc"`
	Sys        uint64 `json:"sys"`
	NumGC      uint64 `json:"num_gc"`
}

var Mem = Memory{
	Alloc:      0,
	TotalAlloc: 0,
	Sys:        0,
	NumGC:      0,
}

var blocks [][]int

func GetMemUsage() Memory {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	Mem = Memory{
		Alloc:      bToMb(m.Alloc),
		TotalAlloc: bToMb(m.TotalAlloc),
		Sys:        bToMb(m.Sys),
		NumGC:      bToMb(uint64(m.NumGC)),
	}
	//s, _ := json.MarshalIndent(Mem,"","  ")
	//s, _ := json.Marshal(Mem)
	//return string(s)
	return Mem
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func MakeBlock(mb int) [][]int {

	//clean up Memory
	blocks = nil
	runtime.GC()
	time.Sleep(2 * time.Second)

	// [130000]int ~= 1mb
	//               length, capacity
	a := make([]int, 130000, 130000)
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(999)
	for i := 0; i < len(a); i++ {
		a[i] = r
	}
	//make new blocks
	for i := 0; i < mb; i++ {
		blocks = append(blocks, a)
	}
	log.Infof("%+v\n", GetMemUsage())
	return blocks
}

func HandleM(c *gin.Context) {
	type memPost struct {
		MemSize int `json:"Mem"`
	}
	var m memPost
	if err := c.ShouldBind(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Infof("%+v\n", m)

	MakeBlock(m.MemSize)

	hostname, _ := os.Hostname()

	c.JSON(http.StatusOK, gin.H{
		//"status": "success",
		"hostname": hostname,
		"memUsage": GetMemUsage(),
	})
}

func CleanUpMem() {
	log.Info("Mem Clean Up ...")
	blocks = nil
	runtime.GC()
	time.Sleep(2 * time.Second)
	log.Infof("%+v\n", GetMemUsage())
}
