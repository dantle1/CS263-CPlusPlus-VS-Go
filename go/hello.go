// this is a test go fileworkspace
package main

import (
	"flag"
	"log"
	"fmt"
	"sync"
	"time"
	//import "os"

	"net/http"
	_ "net/http/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
    flag.Parse()
    if *cpuprofile != "" {
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal(err)
        }
        pprof.StartCPUProfile(f)
        defer pprof.StopCPUProfile()
    }
}

// func main() {
//     // we need a webserver to get the pprof webserver
//     go func() {
//         log.Println(http.ListenAndServe("localhost:6060", nil))
//     }()
//     fmt.Println("hello world")
//     var wg sync.WaitGroup
//     wg.Add(1)
//     go leakyFunction(wg)
//     wg.Wait()
// }

// func leakyFunction(wg sync.WaitGroup) {
//     defer wg.Done()
//     s := make([]string, 3)
//     for i:= 0; i < 10000000; i++{
//         s = append(s, "magical pandas")
//         if (i % 100000) == 0 {
//             time.Sleep(500 * time.Millisecond)
//         }
//     }
// }
