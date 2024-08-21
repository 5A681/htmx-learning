package filedata

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

func (f fileData) CheckNewFileRealTime() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Directory to watch
	dirToWatch := "./"

	// Start a goroutine to handle events
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					log.Println("error")
				}
				// Check if a new file was created
				if event.Op&fsnotify.Create == fsnotify.Create {
					fmt.Println("New file detected:", event.Name)
					// Read the new file
					f.readFile(event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("Error:", err)
			}
		}
	}()

	// Add the directory to the watcher
	err = watcher.Add(dirToWatch)
	if err != nil {
		log.Fatal(err)
	}

	// Block forever
	select {}
}

func (f fileData) readFile(filePath string) {
	// Wait briefly to ensure the file is fully written
	var wg sync.WaitGroup
	wg.Add(10)
	time.Sleep(1 * time.Second)
	go ReadFileXls(filePath, 3, &wg, f.dataTempRepo)
	go ReadFileXls(filePath, 4, &wg, f.dataTempRepo)
	go ReadFileXls(filePath, 5, &wg, f.dataTempRepo)
	go ReadFileXls(filePath, 6, &wg, f.dataTempRepo)
	go ReadFileXls(filePath, 7, &wg, f.dataTempRepo)
	go ReadFileXls(filePath, 8, &wg, f.dataTempRepo)
	go ReadFileXls(filePath, 9, &wg, f.dataTempRepo)
	go ReadFileXls(filePath, 10, &wg, f.dataTempRepo)
	go ReadFileXls(filePath, 11, &wg, f.dataTempRepo)
	go ReadFileXls(filePath, 12, &wg, f.dataTempRepo)
	wg.Wait()
	log.Println("Read ", filePath, "finish")
}
