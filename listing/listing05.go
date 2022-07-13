package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

// интерфейс хранит в себе тип интерфейса и тип самого значения
// значение любого интерфейса является nil, когда и значение, и тип являются nil
// функция возвращает nil типа *customError, результат мы сравниваем с nil типа nil, откуда и следует неравенство
func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
