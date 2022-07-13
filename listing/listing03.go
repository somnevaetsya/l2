package main

import (
	"fmt"
	"os"
)

// интерфейс хранит в себе тип интерфейса и тип самого значения
// значение любого интерфейса является nil, когда и значение, и тип являются nil
// функция возвращает nil типа *os.PathError, результат мы сравниваем с nil типа nil, откуда и следует неравенство

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
