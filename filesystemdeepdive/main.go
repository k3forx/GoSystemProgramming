package main

import (
	"log"

	"gopkg.in/fsnotify/fsnotify.v1"
)

func main() {
	var counter int
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event: ", event)
				if event.Op&fsnotify.Create == fsnotify.Create {
					log.Println("created file: ", event.Name)
					counter++
				} else if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file: ", event.Name)
					counter++
				}
			case err := <-watcher.Errors:
				log.Panicln("error: ", err)
			}
			if counter > 3 {
				done <- true
			}
		}
	}()

	if err := watcher.Add("."); err != nil {
		panic(err)
	}
	<-done
}
