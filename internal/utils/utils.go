package utils

import "log"

func Check(e error, msg string) {
	if e != nil {
		if msg != "" {
			log.Println(msg)
		}
		log.Fatal(e)
	}
}
