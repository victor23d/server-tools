package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

var (
	wg sync.WaitGroup
	logger, _ = zap.NewProduction()
	log = logger.Sugar()
	pb	PrimeBench
)

type PrimeBench struct {
	Proc	int `json:"proc"`
	Job		int `json:"job"`
	Prime 	int64 `json:"prime"`
}



func main() {

	port := 80

	pb = PrimeBench{
		Proc: 0,
		Job: 1,
		Prime: 100000,
	}

	r := gin.Default()
	r.GET("/prime", func(c *gin.Context) {
		//c.JSON(200, gin.H{
		//	"value": pb,
		//)}
		c.String(http.StatusOK, "%v",pb)
	})
	r.POST("/prime", handleV)
	r.Run(":" + strconv.Itoa(port))
}

func handleV(c *gin.Context) {
	var pb PrimeBench
	if err := c.ShouldBind(&pb); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Infof("%v",pb)

	//default zero pb
	if pb.Prime == 0{
		pb.Prime = 100000
	}
	if pb.Job == 0{
		pb.Job = 1
	}

	cpu := runtime.NumCPU()
	//fmt.Println(IsPrime(5212344341))
	runtime.GOMAXPROCS(pb.Proc)
	duration := DoWork(pb.Job, pb.Prime)

	c.JSON(http.StatusOK, gin.H{
		//"status": "success",
		"time": duration.Seconds(),
		"cpu": cpu,
		"pb": pb,
	})
}


func IsPrime(value int64) bool{
	for i:=int64(2); i <= int64(math.Floor(float64(value)/2)); i++{
		if value % int64(i) == 0{
			return false
		}
	}
	return true
}

func FindAllPrime(start int64, end int64) []int64 {
	//slot := (end - start)/9+100
	//var primes [slot]int
	var primes []int64
	for i := start; i <= end; i++ {
		if IsPrime(i) {
			primes = append(primes, i)
			//println(i)
		}
	}
	print(len(primes))
	wg.Done()

	return primes
}

func DoWork(c int, endPrime int64) time.Duration {
	start := time.Now()
	wg.Add(c)
	for i := 1; i <= c; i++ {
		go FindAllPrime(2,endPrime)
		//FindAllPrime(2,100000)
		//wg.Done()
	}
	wg.Wait()
	elapsed := time.Since(start)

	fmt.Printf("Elapsed time : ")
	fmt.Println(elapsed)
	return elapsed
}

