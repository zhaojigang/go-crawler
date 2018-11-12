package persist

import "log"

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		for {
			item := <-out
			log.Printf("ItemSaver getItems, items: %v", item)
		}
	}()
	return out
}
