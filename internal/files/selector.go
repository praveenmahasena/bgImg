package files

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func selectFunc(fs []os.DirEntry) (int, error) {
	fmt.Println(fs)

	sig := make(chan os.Signal,1)

	signal.Notify(sig, syscall.SIGHUP)



	fmt.Println("we go up")


	return 0, nil
}
