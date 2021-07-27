package lib

import "log"

func WriteLogError(folder string, funcName string, err error) {
	log.Printf("[%s][%s]Error deatail: %s \n", folder, funcName, err.Error())
}

func WriteLogSuccess(folder string, funcName string) {
	log.Printf("[%s][%s] Handler successful! \n", folder, funcName)
}
