package main

var justString string //переменная объявлена глобальна(возможная причина непредвиденного поведения программы)

//func createHugeString(){}делает огромную строку

func someFunc() {
	v := createHugeString(1 << 10) //проще использовать 1024 для улучшения читаемости кода и избавиться от выполнения операции 1<<10
	// justString = v[:100]           //огромная строка v не будет освобождена, пока живект juststing
	justString := string([]byte(v[:100])) //имеет свою собсвтенную память
}
func main() {
	someFunc()
}
