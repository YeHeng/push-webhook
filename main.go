package main

import (
	. "github.com/YeHeng/qy-wexin-webhook/util"
)

func main() {
	r := App()
	configRoute(r)

	err := r.Run(":9091")
	if err != nil {
		Logger.Fatalf("Gin start fail. %v", err)
	}
}
