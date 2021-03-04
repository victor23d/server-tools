package benchs

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	wg        sync.WaitGroup
	logger, _ = zap.NewProduction()
	log       = logger.Sugar()
)

type PrimeBench struct {
	Proc  int    `json:"proc"`
	Job   int    `json:"job"`
	Prime uint64 `json:"prime"`
}

type CPU struct {
	Num   int    `json:"num"`
	User  uint64 `json:"user"`
	Idle  uint64 `json:"idle"`
	Total uint64 `json:"total"`
}

var (
	Pb = PrimeBench{
		Proc:  0,
		Job:   1,
		Prime: 100000,
	}
)


func HandleP(c *gin.Context) {
	if err := c.ShouldBind(&Pb); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Infof("%+v\n", Pb)

	hostname, _ := os.Hostname()

	cpu := runtime.NumCPU()
	//fmt.Println(IsPrime(5212344341))
	runtime.GOMAXPROCS(Pb.Proc)

	// Calculate cpu usage
	user0, idle0, total0 := getCPUSample()

	// prime_bench
	duration := DoWork(Pb.Job, Pb.Prime)

	user1, idle1, total1 := getCPUSample()
	idleTicks := idle1 - idle0
	//fmt.Printf("%d\n",idleTicks)
	totalTicks := total1 - total0
	userTicks := user1 - user0
	//fmt.Printf("%d\n",totalTicks)
	cpuUsage := float64(100 * float64(totalTicks-idleTicks) / float64(totalTicks))

	s := fmt.Sprintf("CPU usage is %f%% [busy: %d, total: %d]\n", cpuUsage, totalTicks-idleTicks, totalTicks)
	fmt.Println(s)

	c.JSON(http.StatusOK, gin.H{
		//"status": "success",
		"time": duration.Seconds(),
		"cpu": CPU{
			Num:   cpu,
			User:  userTicks,
			Idle:  idleTicks,
			Total: totalTicks,
		},
		"Pb":       Pb,
		"hostname": hostname,
		"cpuUsage": cpuUsage,
	})
}

func IsPrime(value uint64) bool {
	for i := uint64(2); i <= uint64(math.Floor(float64(value)/2)); i++ {
		if value%uint64(i) == 0 {
			return false
		}
	}
	return true
}

func FindAllPrime(start uint64, end uint64) []uint64 {
	//slot := (end - start)/9+100
	//var primes [slot]int
	var primes []uint64
	for i := start; i <= end; i++ {
		if IsPrime(i) {
			primes = append(primes, i)
			//println(i)
		}
	}
	//print(len(primes))
	wg.Done()

	return primes
}

func DoWork(c int, endPrime uint64) time.Duration {
	start := time.Now()
	wg.Add(c)
	for i := 1; i <= c; i++ {
		go FindAllPrime(2, endPrime)
		//FindAllPrime(2,100000)
		//wg.Done()
	}
	wg.Wait()
	elapsed := time.Since(start)

	fmt.Printf("Elapsed time : ")
	fmt.Println(elapsed)
	return elapsed
}

func getCPUSample() (user, idle, total uint64) {
	contents, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		return
	}
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if fields[0] == "cpu" {
			numFields := len(fields)
			for i := 1; i < numFields; i++ {
				val, err := strconv.ParseUint(fields[i], 10, 64)
				if err != nil {
					fmt.Println("Error: ", i, fields[i], err)
				}
				total += val // tally up all the numbers to get total ticks
				if i == 4 {  // idle is the 5th field in the cpu line
					idle = val
				}
				if i == 1 {
					user = val
				}
			}
			return
		}
	}
	return
}
