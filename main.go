
package main

import (
	"github.com/HugoJBello/go-heartbeat-service/controllers"
)

func main() {
	controllers.SetupServer().Run()
}
