package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(string) (error)
	ActionInfo() (string, error)
}

//Info формирует отчет о тренировках
func Info(dataset []string, dp DataParser) {

	for _, v := range dataset{
		err := dp.Parse(v)
		if err != nil {
			log.Println(err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println(info)
		
	}
}
