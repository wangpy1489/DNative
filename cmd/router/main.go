package main 
import (
	"log"
	"os"

	"github.com/wangpy1489/DNative/pkg/router"
)

var logger = log.New(os.Stdout,"", log.LstdFlags|log.Llongfile)

func main ()  {
	rou, err := router.MakeRouter(logger)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Println("server start")
	rou.Serve(8000)
	logger.Fatal("server done")
}