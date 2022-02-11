package main

import ( 
	"fmt" 
	"math/rand"
	"strings"
	"time"
	"math"
	"strconv"
	"encoding/json"
	"os"
	"sort"
	"reflect"
	"errors"
	"log"
	"io/ioutil"
	"net/http"
)

func main() {
	copy_or_link()

	lesson3_text()
	lesson3_task()
	lesson4_text()
	lesson4_text()
	lesson6_text()
	lesson7_task()
	lesson8_text()
	lesson9_text()
	lesson11_text()
	lesson11_task()
	lesson12_text()
	lesson12_task()
	lesson14_text()
	lesson14_task()
	lesson15_text()
	lesson15_task()
	lesson16_text()
	lesson16_task()
	lesson17_task()
	lesson18_text()
	lesson19_text()
	lesson19_task()
	lesson20_text()
	lesson20_task()
	lesson21_text()
	lesson21_task()
	//lesson22_task()
	lesson23_text()
	lesson24_text()
	lesson24_task()
	lesson25_text()
	lesson26_text()
	lesson26_task()
	lesson27_task()
	lesson28_text()
	//lesson28_task()
	lesson29_text()
	lesson30_text()
	lesson35_text()

	copy_or_link()
}

func lesson3_text() {
	fmt.Printf("\n Lesson_3 \n \n")
	var speed = 90
	const spend_time = 1

	fmt.Printf("Time: %v hours, Speed: %v mps \n", speed, spend_time)
	fmt.Printf("Distance: %v miles \n", speed * spend_time)

	var v0, v1 = 12, 13
	fmt.Printf("V1: %v, V2: %v \n", v0, v1)

	var rand_value = rand.Intn(10) + 1
	fmt.Printf("Random value: %v \n", rand_value)
}

func lesson3_task() {
	const distance = 56000000
	var spend_time = 28 * 24

	fmt.Printf("Should have speed: %v \n", distance / spend_time)
}

func lesson4_text() {
	fmt.Printf("\n Lesson_4 \n \n")

	var bool_val = true
	var word = "hello my friend"

	bool_val = strings.Contains(word, "friend")

	fmt.Printf("Should contain friend: %v \n", bool_val)

	var number = 20

	if number % 3 == 0 {
		fmt.Printf("Devide by 3 with 0: \n")
	} else if number % 3 == 1 {
		fmt.Printf("Devide by 3 with 1: \n")
	} else {
		fmt.Printf("Devide by 3 with 2: \n")
	}

	var kind = "table"

	switch kind {
	case "table":
		fmt.Printf("It is Table \n")
		fallthrough
	case "sun":
		fmt.Printf("It is Sun \n")
	default: 
		fmt.Printf("Default \n")
	}

	var counter = 0

	for counter < 10 {
		fmt.Printf("%3v \n", counter)
		counter++
	}
}

func lesson4_task() {
	const count_of_nums = 10
	var s = rand.NewSource(time.Now().Unix())
	var r = rand.New(s) // initialize local pseudorandom generator 

	var my_number = r.Intn(count_of_nums)

	for {
		var computer_number = r.Intn(count_of_nums)

		if my_number > computer_number {
			fmt.Printf("My number greater then computer number %v \n", computer_number)
		} else if my_number < computer_number {
			fmt.Printf("My number less then computer number %v \n", computer_number)
		} else {
			fmt.Printf("Success! My number is: %v \n", my_number)
			break
		}
		
	}
}

var outside = "Outside value"

func lesson6_text()  {
	fmt.Printf("Lesson_6 \n \n")
	// Область видимости в {}

	// var num = 1
	// num := 1

	for counter := 10; counter > 0; counter-- {
		fmt.Printf("Number %v \n", counter)
	}

	// counter -> undefined: counter

	counter1 := 10

	for counter1 := 10; counter1 > 0; counter1-- {
		fmt.Printf("Conuner again %v \n", counter1)
	}

	fmt.Printf("Final number1 %v \n", counter1)

	if num := 1; num == 1 {
		fmt.Printf("true \n")
	}

	if num := 2; num == 2 {
		fmt.Printf("true \n")
	}

	// fmt.Printf("Catch error %v \n", num) -> undefined: num

	num := 1

	if num == 1 {
		fmt.Printf("true \n")
	}

	if num == 2 {
		fmt.Printf("true \n")
	}

	fmt.Printf("No catch error %v \n", num)

	// Outside value
	fmt.Printf(outside)
}

func lesson7_task() {
	fmt.Printf("\n Lesson_7 \n \n")

	fmt.Printf("%20v %20v %20v %20v \n \n", "Spaceline", "Days", "Trip Type", "Price")

	for num := 0; num < 10; num ++ {

		speed := rand.Intn(14) + 16;

		s := ""
		d := rand.Intn(10) + 30 - speed;
		t := ""
		p := rand.Intn(30) + 30 + speed;

		switch num := rand.Intn(3); num {
		case 0:
			s = "Space Adventures"
		case 1:
			s = "SpaceX"
		case 2:
			s = "Virgin Galactic"
		}

		switch num := rand.Intn(2); num {
		case 0:
			t = "One-way"
		case 1:
			t = "Round-trip"
			p = p * 2
		}

		fmt.Printf("%20v %20v %20v %20v \n", s, d, t, p)
	}
}

func lesson8_text() {
	fmt.Printf("\n Lesson_8 \n \n")

	// Текст "Go" и число 28487 на компьютере с архитектурой x86 представлены одним и тем же набором нулей и единиц — 0110111101000111. 
	// Тип устанавливает, что данные биты и байты означают.

	var pi64 = math.Pi
	var pi32 float32 = math.Pi
	
	fmt.Println(pi64) // Выводит: 3.141592653589793              
	fmt.Println(pi32) // Выводит: 3.1415927

	var p string

	fmt.Println(p) // Выводит: 3.1415927

}

func lesson9_text() {
	fmt.Printf("\n Lesson_9 \n \n")

	// var num int = 1
	// var unum uint = 1

	var blue uint8 = 255
	fmt.Printf("%08b\n", blue) // Выводит: 11111111  
}

func lesson11_text() {
	fmt.Printf("\n Lesson_11 \n \n")

	var word string = "some word"

	// fmt.Printf("Letter 6 is %c \n", word[6])

	for counter := 0; counter < len(word); counter++ {
		fmt.Printf("Letter %c is %v \n", word[counter], counter)
	}

	c := 'a'
	c = c + 3
	fmt.Printf("%c\n", c) // Выводит: d

	message := "uv vagreangvbany fcnpr fgngvba"
	
	for i := 0; i < len(message); i++ { // Итерирует каждый символ ASCII     
		c := message[i]
		if c >= 'a' && c <= 'z' { // Оставляет оригинальную пунктуацию и пробелы           
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}
	fmt.Printf(" \n")
}

func lesson11_task() {
	var enc_text string = "l fdph, l vdz, l frqtxhuhg"

	for _, c := range enc_text {
		n_c := c
		if c >= 'a' && c <= 'z' {
			n_c = c - 3

			if c < 'd' {
				n_c += 26
			}

		}
		fmt.Printf("%c", n_c)
	}
	fmt.Printf("\n")
}

func lesson12_text() {
	fmt.Printf("\n Lesson_12 \n \n")

	n_int := 12
	n_flt := 10.01

	// res := n_int * n_flt // (mismatched types int and float64)
	// res := n_int > n_flt // (mismatched types int and float64)

	res := float64(n_int) * n_flt

	fmt.Printf("%v \n", res)
	fmt.Printf("%v \n", int(res))

	// var pi float64 = 3.14 // cannot convert pi (type float64) to type string
	var n1_int int = 100
	// rune и byte являются просто другими названиями для int32 и uint8
	var n2 rune = 940
	var n3 byte = 12

	fmt.Print("Converted byte (uint8) and rune(int32): ")
	fmt.Print(string(n1_int), string(n2), string(n3))
	fmt.Printf("\n")

	fmt.Print("Converted intreger: ")
	fmt.Println(strconv.Itoa(n1_int))

	countdown, err := strconv.Atoi("10")
	if err != nil {
		// о нет, что-то пошло не так
	}
	fmt.Println(countdown) // Выводит: 10

}

func lesson12_task() {
	v := 39999

	fmt.Printf("MaxUint8 %v \n", math.MaxInt8) // -127 .. 127 (2^8)
	fmt.Printf("MaxUint8 %v \n", math.MaxUint8) // 0 .. 257 (2^8)

	if v < math.MaxUint8 && v > 0 {
		fmt.Printf("Ok \n")
	}

	value := "no"
	var res int

	switch value {
	case "yes", "1", "true":
		res = 1
	case "no", "0", "false":
		res = 0
	default: 
		fmt.Printf("Error \n")
	}

	fmt.Printf("Res is %v \n", res)
}

func lesson14_text() {
	fmt.Printf("\n Lesson_14 \n \n")

	num := 33

	fmt.Printf(strconv.Itoa(num))
	fmt.Printf(string(num))
	fmt.Printf("\n")

}

func lesson14_task() {
	fmt.Printf("Result kelvin to celsius = %v \n", kelvinToCelsius(200))
	fmt.Printf("Result celsius to farenheit = %v \n", celsiusToFahrenheit(200))
	fmt.Printf("Result kelvin to farenheit = %v \n", kelvinToFahrenheit(0.0))
}

func kelvinToCelsius(num float64) float64 {
	num -= 273.3
	return num
}

func celsiusToFahrenheit(num float64) float64 {
	num = (num * 9.0 / 5.0) + 32.0
	return num
}

func kelvinToFahrenheit(num float64) float64 {
	celsius := kelvinToCelsius(num)
	farenheit := celsiusToFahrenheit(celsius)
	return farenheit
}

func lesson15_text() {
	fmt.Printf("\n Lesson_15 \n \n")
	// type celsius float64
	// type fahrenheit float64
	// type kelvin float64

	var temp_c celsius = 20
	var temp_f fahrenheit = 20
	var warmUp float64 = 32

	// temp += warmUp // temp += warmUp (mismatched types celsius and float64)
	temp_c += celsius(warmUp)

	// same_values := fahrenheit == celsius // fahrenheit == celsius (mismatched types fahrenheit and celsius)

	fmt.Printf("Celsius %v \n", temp_c)
	fmt.Printf("Fahrenheit %v \n", temp_f)

	// Здесь нет классов или даже объектов, но есть методы.
	var temp_k kelvin = 400

	c := temp_k.celsius() // return type celsius
	fmt.Printf("Kelvin -> Celsius %v \n", c)
}

type celsius float64
type fahrenheit float64
type kelvin float64

func lesson15_task() {
	var cur_temp_c celsius = 127
	cur_temp_k := celsiusToKelvinV1(cur_temp_c)

	fmt.Printf("Task: Kelvin %v \n", cur_temp_k)

	var temp_c celsius = 100
	var temp_k kelvin = 100
	var temp_f fahrenheit = 100

	c_to_k := temp_c.kelvin()
	c_to_f := temp_c.fahrenheit()

	k_to_c := temp_k.celsius()
	k_to_f := temp_k.fahrenheit()

	f_to_c := temp_f.celsius()
	t_to_k := temp_f.kelvin()

	fmt.Printf("Task: c_to_k %v \n", c_to_k)
	fmt.Printf("Task: c_to_f %v \n", c_to_f)
	fmt.Printf("Task: k_to_c %v \n", k_to_c)
	fmt.Printf("Task: k_to_f %v \n", k_to_f)
	fmt.Printf("Task: f_to_c %v \n", f_to_c)
	fmt.Printf("Task: t_to_k %v \n", t_to_k)

	fmt.Printf("Task: c - %v to c_to_f - %v to f_to_c - %v \n", temp_c, c_to_f, c_to_f.celsius())

}

// celsiusToKelvinV1(cur_temp_c)
func kelvinToCelsiusV1(num kelvin) celsius { // функция kelvinToCelsiusV1 (к: похож на просто метод)
	num -= 273.3
	return celsius(num)
}

func celsiusToKelvinV1(num celsius) kelvin { 
	num += 273.3
	return kelvin(num)
}

// c = k.celsius()
func (num kelvin) celsius() celsius { // метод celsius для типа kelvin (к: почти как метод инстанса, но только это метод инстанса типа)
	num -= 273.3
	return celsius(num)
}

func (num kelvin) fahrenheit() fahrenheit {
	num = ((num - 273.3) * 9.0 / 5.0) + 32.0
	return fahrenheit(num)
}

func (num celsius) kelvin() kelvin {
	num += 273.3
	return kelvin(num)
}

func (num celsius) fahrenheit() fahrenheit {
	num = (num * 9.0 / 5.0) + 32.0
	return fahrenheit(num)
}

func (num fahrenheit) kelvin() kelvin {
	num = ((num - 32.0) * 5.0) / 9.0 + 273.3
	return kelvin(num)
}

func (num fahrenheit) celsius() celsius {
	num = ((num - 32.0) * 5.0) / 9.0
	return celsius(num)
}

func lesson16_text() {
	fmt.Printf("\n Lesson_16 \n \n")
	
	sensor := fakeSensor

	fmt.Printf("Sensor %v \n", sensor())

	var adapter func() kelvin
	
	sensor = adapter
	sensor = fakeSensor

	res := measureTemperature(sensor)
	fmt.Printf("Measure res %v \n", res)

	res_v1 := measureTemperatureV1(sensor)
	fmt.Printf("Measure res_v1 %v \n", res_v1)

	var f = func ()  { // Присваивает анонимную функцию переменной
		fmt.Printf("Inside anon func \n")
	}

	f()

	func () {
		fmt.Printf("Inside anon func in func \n")
	}()

	// В отличие от обычных функций, функциональные литералы являются замыканиями, потому что они сохраняют отсылки к переменным в окружающей области видимости.

	sensor = realSensor
	res_s := calibrate(sensor, kelvin(5))
	fmt.Printf("Calibrate %v \n", res_s())

		// var k kelvin = 294.0
		//
		// sensor := func() kelvin {
		//	return k
		//}
		//fmt.Println(sensor()) // Выводит: 294
		//
		//k++
		//fmt.Println(sensor()) // Выводит: 295
}

func lesson16_task() {
	// Перепишите следующую сигнатуру функции для использования типа функции:
	// func drawTable(rows int, getRow func(row int) (string, string))
	// type getRowFn func(row int) (string, string)
	// func drawTable(rows int, getRow getRowFn)
	var tmp kelvin = 6
	sensor := realSensor
	res_s := calibrate(sensor, tmp)
	fmt.Printf("Task: Calibrate %v \n", res_s())
	tmp = 10
	fmt.Printf("Task: Calibrate %v \n", res_s())
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(100) + 100)
}

func measureTemperature(sensor func() kelvin) kelvin{
	return kelvin(sensor())
}

type sensor func() kelvin // Объявление типов функции в Golang

func measureTemperatureV1(s sensor) kelvin{
	return kelvin(s())
}

type getRowFn func(row int) (string, string)

// Анонимная функция в предыдущем листинге использует замыкания. Это отсылка к переменным s и offset, что функция calibrate принимает в качестве параметров.
// Из-за того, что замыкание сохраняет ссылку на ближайшие переменные, а не копирует их значения, изменения с этими переменными отражаются в вызовах к анонимной функции
func calibrate(s sensor, offset kelvin) sensor {
	// Замыкания сохраняют отсылки к переменным в окружающей области видимости.
	return func() kelvin { 
		return s() + offset // offset это ссылка, а не значение, аналогично с s()
	}
} 

func realSensor() kelvin {
	return kelvin(0)
}

func lesson17_task() {
	var data_gen dataForRow

	data_gen = celsiusToKelvinV2

	drawHeader()
	drawTable(40, 100, 5, data_gen)
	drawBottom()
}

type dataForRow func(value float64) (string, string)

func celsiusToKelvinV2(value float64) (string, string) {
	return fmt.Sprintf("%.1f", celsius(value)), fmt.Sprintf("%.1f", celsius(value).kelvin())
}

func drawTable (start_value int, finish_value int, step int, generate dataForRow) {
	for value := start_value; value <= finish_value; value += step {
		col_1_value, col_2_value := generate(float64(value))
		drawTwoValues(col_1_value, col_2_value)
	}
}

func drawLine()  {
	fmt.Printf("=======================\n")
}

func drawEmptyLine()  {
	fmt.Printf("|          |          |\n")
}

func drawHeaderValue()  {
	drawEmptyLine()
	drawTwoValues("C", "K")
	drawEmptyLine()
}

func drawHeader()  {
	drawLine()
	drawHeaderValue()
	drawLine()
}

func drawBottom()  {
	drawLine()
}

func drawTwoValues(value_1 string, value_2 string) {
	fmt.Printf("| %8s | %8s |\n", value_1, value_2)
}

func lesson18_text() {
	fmt.Printf("\n Lesson_18 \n \n")

	var planets [8]string
	planets[0] = "Меркурий"
	fmt.Printf("Planet 1 %v \n", planets[0])
	fmt.Printf("Length planets %v \n", len(planets))

	planets = [...]string{ // Компилятор Go подсчитывает элементы (Такая краткая запись называется композитный литерал)
		"Меркурий",
		"Венера",
		"Земля",
		"Марс",
		"Юпитер",
		"Сатурн",
		"Уран",
		"Нептун", // Запятая в самом конце является обязательной
	}
	fmt.Printf("Length planets %v \n", len(planets))

	for i := 0; i < len(planets); i++ {
		if i + 1 == len(planets) {
			fmt.Printf(" %v \n", planets[i])
		} else {
			fmt.Printf(" %v |", planets[i])
		}
	}

	for i, n := range planets {
		if i + 1 == len(planets) {
			fmt.Printf(" %v \n", n)
		} else {
			fmt.Printf(" %v |", n)
		}
	}

	planets_copy := planets
	planets_copy[0] = "Oops"

	output_array(planets_copy)

	output_array(planets)

	a := 3
	b := a // := передача значения
	a += 3
	fmt.Printf("a %v, b %v \n", a, b)

	changable(planets, b)

	fmt.Printf("a %v, b %v \n", a, b)

	output_array(planets)

	func () {
		planets[0] = "Ooooopsee"
	}()
	fmt.Printf("Anonymous: ")
	output_array(planets)

}

func changable(a [8]string, b int) {
	a[1] = "Oops"
	b += 4
}

func output_array(planets [8]string) {
	for i, n := range planets {
		if i + 1 == len(planets) {
			fmt.Printf(" %v \n", n)
		} else {
			fmt.Printf(" %v |", n)
		}
	}
}

func output_slcie(planets []string) {
	for i, n := range planets {
		if i + 1 == len(planets) {
			fmt.Printf(" %v \n", n)
		} else {
			fmt.Printf(" %v |", n)
		}
	}
}

func lesson19_text() {
	fmt.Printf("\n Lesson_19 \n \n")

	planets := [...]string{
		"Меркурий",
		"Венера",
		"Земля",
		"Марс",
		"Юпитер",
		"Сатурн",
		"Уран",
		"Нептун",
	}
	 
	terrestrial := planets[0:4] // Срез - ссылается на основной массив
	gasGiants := planets[4:6]
	iceGiants := planets[6:8]

	fmt.Println(terrestrial, gasGiants, iceGiants)

	terrestrial[0] = "Opps!"

	output_array(planets)

	neptune := "Neptune"
	tune := neptune[3:]
	fmt.Println(tune) // Выводит: tune

	// Имейте в виду, что индексы учитывают количество байтов, но не рун

	// Инициализации среза

	dwarfArray := [...]string{"Церера", "Плутон", "Хаумеа", "Макемаке", "Эрида"}
	dwarfSlice := dwarfArray[:]

	// Кратко
	dwarfs := []string{"Церера", "Плутон", "Хаумеа", "Макемаке", "Эрида"}

	fmt.Printf("array %T\n", dwarfArray) // Выводит: array [5]string
	fmt.Printf("slice %T\n", dwarfSlice) // Выводит: slice []string
	fmt.Printf("slice %T\n", dwarfs) // Выводит: slice []string

	// Срез это ссылки на базовый массив.
	fmt.Printf("Before: ")
	output_array(planets)
	slice_planets := planets[:]
	unchangable_planets := planets[:]
	change_slice(slice_planets)
	fmt.Printf("After: ")
	output_array(planets)

	// Однако hyperspace может достигнуть базового массива, на который указывает worlds, и изменить его элементы. Данные изменения доступны другим срезам (видам) массива.
	fmt.Printf("Slice of planets: ")
	output_slcie(unchangable_planets)
	// Массивы редко используются напрямую. Разработчики Go предпочитают срезы ввиду их гибкости, особенно при передачи аргументов функции.

}

func change_slice(sliced []string) {
	for i, n := range sliced {
		sliced[i] = strconv.Itoa(i) + " " + n
	}
}

type Planets string

func lesson19_task() {
	planets := []Planets{
		"Меркурий",
		"Венера",
		"Земля",
		"Марс",
		"Юпитер",
		"Сатурн",
		"Уран",
		"Нептун",
	}

	output_planets(planets)

	for i, p := range planets {
		planets[i] = p.terraform()
	}

	output_planets(planets)
}

func (p Planets) terraform() Planets {
	p = "Новый " + p
	return Planets(p)
}

func output_planets(planets []Planets) {
	for i, n := range planets {
		if i + 1 == len(planets) {
			fmt.Printf(" %v \n", n)
		} else {
			fmt.Printf(" %v |", n)
		}
	}
}

func lesson20_text() {
	fmt.Printf("\n Lesson_20 \n \n")

	// Число элементов, что видны через срез, определяется его длиной. Если у слайса есть базовый массив, что большего размера, у среза остается вместимость для роста.

	dwarfs := []string{"Церера", "Плутон", "Хаумеа", "Макемаке", "Эрида"} 
    dump("dwarfs", dwarfs) // Выводит: dwarfs: длина 5, вместимость 5 [Церера Плутон Хаумеа Макемаке Эрида]                                             
    dump("dwarfs[1:2]", dwarfs[1:2]) // Выводит: dwarfs[1:2]: длина 1, вместимость 4 [Плутон]  

	dwarfs1 := []string{"Церера", "Плутон", "Хаумеа", "Макемаке", "Эрида"}  // Длина 5, вместимость 5
	dwarfs2 := append(dwarfs1, "Оркус") // Длина 6, вместимость 10
	dwarfs3 := append(dwarfs2, "Салация", "Квавар", "Седна") // Длина 9, вместимость 10
	dump("dwarfs1", dwarfs1)
	dump("dwarfs2", dwarfs2)
	dump("dwarfs3", dwarfs3)
	
	// С появлением версии Go 1.2 был введен трех-индексный срез, что нужен для ограничения вместимости итогового среза.
	planets := []string{
		"Меркурий", "Венера", "Земля", "Марс",
		"Юпитер", "Сатурн", "Уран", "Нептун",
	}
	 
	terrestrial := planets[0:4:4] // Длина 4, вместимость 4
	worlds := append(terrestrial, "Церера")
	 
	fmt.Println(planets) // Выводит: [Меркурий Венера Земля Марс Юпитер Сатурн Уран Нептун]

	// Если третий индекс не уточняется, вместимость terrestrial равна 8. Добавление элемента "Церера" не перемещает новый массив, а вместо этого переписывает "Юпитер"

	terrestrial = planets[0:4] // Длина 4, вместимость 8
	worlds = append(terrestrial, "Церера")
	dump("worlds", worlds)
 
    fmt.Println(planets) // Выводит: [Меркурий Венера Земля Марс Церера Сатурн Уран Нептун]

	// Если для append недостаточно вместимости, Go должен выделить новый массив и скопировать туда содержимое старого массива.

	slice_1 := make([]string, 0, 10)
	slice_2 := []string {"1", "2"}
	fmt.Printf("slice_1 %T\n", slice_1) // Выводит: slice []string
	dump("slice_1", slice_1)
	fmt.Printf("slice_2 %T\n", slice_2) // Выводит: slice []string
	dump("slice_2", slice_2)

	dwarfs_slice := make([]string, 0, 10)
	dwarfs_slice = append(dwarfs_slice, "Церера", "Плутно", "Хаумеа", "Макемаке", "Эрида")
	dump("dwarfs_slice", dwarfs_slice)

	dwarfs_array := [5]string{"Церера", "Плутон", "Хаумеа", "Макемаке", "Эрида"} 
	dwarfs_slice_by_array := dwarfs_array[:]
	dump("dwarfs_slice_by_array", dwarfs_slice_by_array)

	// Приемущество СРЕЗА по сравнению с МАССИВОМ, что СРЕЗ указывает на МАССИВ и при передачи среза в функцию, передается ссылка, а не копируется значение (в отличии от массива).

	word_counter("G", "G")
	words := make([]string, 0, 3)
	words = append(words, "G", "G", "G")
	word_counter(words...)
}

func dump(label string, slice []string) {
    fmt.Printf("%v: длина %v, вместимость %v %v\n", label, len(slice), cap(slice), slice)
}

func word_counter(words ...string) {
	fmt.Printf("words %T\n", words)
	fmt.Printf("Передано %v слова \n", len(words))
}

func lesson20_task() {
	fmt.Printf("\n")
 	s := []string{}
 	lastCap := cap(s)
  
 	for i := 0; i < 10000; i++ {
	 	s = append(s, "An element")
	 	if cap(s) != lastCap {
		 	fmt.Println(cap(s))
		 	lastCap = cap(s)
	 	}
 	}
}

func lesson21_text() {
	fmt.Printf("\n Lesson_21 \n \n")

	hash := map[string]int{
		"Земля": 1,
		"Венера": 2,
	}

	temperature := map[string]int{
		"Земля": 1,
		"Венера": 2,
	}

	hash["Земля"] = 3

	fmt.Printf("Hash %v \n", hash["Земля"])

	// Переменная moon должна содержать значение ключа "Луна" или же нулевое значение. При наличии ключа значение дополнительной переменной ok будет равно true,  в противном случае — false.
	if moon, ok := temperature["Луна"]; ok { // Синтаксис comma, ok
		fmt.Printf("Средняя температура на поверхности Луны составляет %v° C.\n", moon)
   	} else {
	   fmt.Println("Где Луна?") // Выводит: Где Луна?
   	}

   	// temp, found := temperature["Венера"]

	// arrays
	array_v1 := [2]string {"1", "2"}
	array_v2 := array_v1 // := копирует
	array_v2[0] = "777"

	fmt.Printf("Begining array_v1: %v\n", array_v1[0])
	fmt.Printf("After array_v1: %v\n", array_v2[0])

	// slices
	slice_v1 := make([]string, 0, 10)
	slice_v1 = append(slice_v1, "1", "2", "3")
	slice_v2 := slice_v1 // := не копирует, а ссылает
	slice_v2[1] = "4"

	fmt.Println("Begining slice")
	output_slcie(slice_v1)
	fmt.Println("Slice after")
	output_slcie(slice_v2)

	// string
	string_v1 := "word"
	string_v2 := string_v1 // := копирует
	string_v2 = "change"
	fmt.Printf("Begining string %v \n", string_v1)
	fmt.Printf("After string %v \n", string_v2)

	// maps
	hash_v1 := hash // := не копирует, а ссылает
	hash_v1["Земля"] = 4

	fmt.Printf("Begining hash %v \n", hash["Земля"])
	fmt.Printf("After hash %v \n", hash_v1["Земля"])

	// struct
	struct_v1 := planet {100.0} 
	struct_v2 := struct_v1 // := копирует

	struct_v2.radius = 200.0

	fmt.Printf("Begining struct %v \n", struct_v1.radius)
	fmt.Printf("After struct %v \n", struct_v2.radius)

	// maps
	hash_2 := make(map[string]int, 2)
	hash_2["1"] = 1
	hash_2["2"] = 2

	hash_3 := hash_2

	hash_3["1"] = 100

	fmt.Printf("Begining hash %v \n", hash_2["1"])
	fmt.Printf("After hash %v \n", hash_3["1"])

	for t, num := range hash_2 { // Итерирует через карту (ключ, значение)
		fmt.Printf("%+.2f встречается %d раз(а) \n", t, num)
   	}
}

func lesson21_task() {
	var text string = "As far as eye could reach he saw nothing but the stems of the great plants about him receding in the violet shade, and far overhead the multiple transparency of huge leaves filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever he felt able he ran again; the ground continued soft and springy, covered with the same resilient weed which was the first thing his hands had touched in Malacandra. Once or twice a small red creature scuttled across his path, but otherwise there seemed to be no life stirring in the wood; nothing to fear—except the fact of wandering unprovisioned and alone in a forest of unknown vegetation thousands or millions of miles beyond the reach or knowledge of man."

	prepared_text := strings.Fields(strings.ToLower(text))
	fmt.Println(prepared_text)

	set := make(map[string]int, len(prepared_text))

	for _, value := range prepared_text {
		value = strings.Trim(value, `.,"-`)
		set[value]++
	}

	for t, num := range set { // Итерирует через карту (ключ, значение)
		fmt.Printf("%v встречается %d раз(а) \n", t, num)
   	}
}

func lesson22_task() {
	const (
		width = 20
		height = 20
		lifetime = 10
	)
	u := newUniverse(width, height)
	next_u := newUniverse(width, height)
	u.Seed()
	u.Show()
	for i := 0; i < lifetime; i++ {
		next_u = nextUniverse(u, next_u)
		next_u.Show()
		u, next_u = next_u, u
		time.Sleep(time.Second)
		fmt.Print("\033[H")
	}
}

type Universe [][]bool

func newUniverse(width int, height int) Universe {
	slice := make([][]bool, width)
	for i, _ := range slice {
		slice[i] = make([]bool, height)
	}
	return Universe(slice)
}

func (u Universe) Show() {
	fmt.Printf("######################\n")
	var item_s string
	for _, row := range u {
		fmt.Printf("#")
		for _, item := range row {
			if item {
				item_s = "*"
			} else {
				item_s = " "
			}
			fmt.Printf("%v", item_s)
		}
		fmt.Printf("#\n")
	}
	fmt.Printf("######################\n")
}

func (u Universe) Seed() {
	for j, row := range u {
		for i, _ := range row {
			u[i][j] = rand.Intn(4) > 2
		}
	}
}

func (u Universe) Alive(x, y int) bool {
	width := len(u)
	height := len(u[0])

	index_x := x
	index_y := y 

	if x < 0 {
		index_x = x + width
	} else if x >= width {
		index_x = x - width
	} else {}

	if y < 0 {
		index_y = y + height
	} else if y >= height {
		index_y = y - height
	}
	//fmt.Printf("| %v %v %v %v |", x, index_x, y, index_y)
	return u[index_x][index_y]
}

func (u Universe) CountNeighbors(x, y int) int {
	var counter int;
	for i := x - 1; i <= x + 1; i++ {
		for j := y - 1; j <= y + 1; j++ {
			if i != x && j != y {
				if u.Alive(i, j) {
					counter++
				}
			}
		}
	}
	return counter
}

func (u Universe) Next(x, y int) bool {
	var is_alive bool;
	if u.Alive(x, y) {
		//fmt.Printf("| CountNeighbors %v |", u.CountNeighbors(x, y))
		if u.CountNeighbors(x, y) < 2 || u.CountNeighbors(x, y) > 3 {
			is_alive = false
		} else {
			is_alive = true
		}
	} else {
		if u.CountNeighbors(x, y) == 3 {
			is_alive = true
		}
	}
	return is_alive
}

func nextUniverse(old_u Universe, new_u Universe) Universe {
	for i, row := range old_u {
		for j, _ := range row {
			//fmt.Printf("| %v %v |", i, j)
			new_u[i][j] = old_u.Next(i, j)
		}
	}
	new_u.Show()
	return new_u
}

func lesson23_text() {
	fmt.Printf("\n Lesson_23 \n \n")

	var curiosity struct {
		lat  float64
		long float64
	}

	// Присваивание значений полям структуры
	curiosity.lat = -4.5895 
	curiosity.long = 137.4417 
	
	fmt.Println(curiosity.lat, curiosity.long) // Выводит: -4.5895 137.4417
	fmt.Println(curiosity) // Выводит: {-4.5895 137.4417}

	type location struct {
		lat  float64
		long float64
	}

	var a location
	a.lat = 200.0
	a.long = 400.0

	fmt.Println(a)

	b := location{ lat: 100.0, long: 200.0 }

	fmt.Println(b)

	c := location{ 100.0, 200.0 }

	fmt.Println(c)

	d := c // := копирует значение

	d.lat = c.lat + 100

	fmt.Println("c and d:")
	fmt.Println(c)
	fmt.Println(d)

	type location_v1 struct {
		Name string `json:"name"`
		Lat  float64 `json:"lat"`
		Long float64 `json:"long"`
	}
	 
	locations := []location_v1{
		{Name: "Bradbury Landing", Lat: -4.5895, Long: 137.4417},
		{Name: "Columbia Memorial Station", Lat: -14.5684, Long: 175.472636},
		{Name: "Challenger Memorial Station", Lat: -1.9462, Long: 354.4734},
	}

	fmt.Printf("%+v\n", locations)

	// JSON
	// bytes, err := json.Marshal(locations)
	bytes, err := json.MarshalIndent(locations, "", "  ")
    exitOnError(err)

	fmt.Println(string(bytes)) 
}

// exitOnError выводит любые ошибки и выходит.
func exitOnError(err error) {
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func lesson24_text() {
	fmt.Printf("\n Lesson_24 \n \n")

	// координаты в градусах, минутах, секундах для сферы N/S/E/W 


	lat := coordinate{4, 35, 22.2, 'S'}
	long := coordinate{137, 26, 30.12, 'E'}

	fmt.Printf("%v\n", lat.decimal())
	fmt.Printf("%v\n", long.decimal())

	lat_v1 := newCoordinate(4, 35, 22.2, 'S')
	long_v1 := newCoordinate(137, 26, 30.12, 'E')

	fmt.Printf("%v\n", lat_v1.decimal())
	fmt.Printf("%v\n", long_v1.decimal())

	earth := newPlanet(300000.0)

	fmt.Printf("Earth = %v\n", earth.s(100.0, 200.0))
}

type coordinate struct {
	d, m, s float64
	h       rune
}

func (c coordinate) decimal() float64 {
    sign := 1.0
    switch c.h {
    case 'S', 'W', 's', 'w':
        sign = -1
    }
    return sign * (c.d + c.m/60 + c.s/3600)
}

// функция конструктора new <- Правило именования конструктора
// В Go функции, переменные и другие идентификаторы, что начинаются с большой буквы, экспортируются и становятся доступными для других пакетов.
func newCoordinate(d, m, s float64, h rune) coordinate {
	return coordinate{d, m, s, h}
}

type planet struct {
	radius float64
}

func newPlanet(r float64) planet{
	return planet{radius: r}
}

func (p planet) s(a float64, b float64) float64 {
	return (a + b / p.radius)
}

func lesson24_task() {

	place1 := newPlace("Spirit", "Moscow", newCoordinate(4, 35, 22.2, 'S'), newCoordinate(137, 26, 30.12, 'E'))
	place2 := newPlace("Orgs", "Voronezhs", newCoordinate(4, 35, 22.2, 'S'), newCoordinate(137, 26, 30.12, 'E'))

	fmt.Printf("Values: \n %+v\n %+v\n", place1, place2)
	fmt.Printf("Decimals: \n %+v\n %+v\n", place1.lon.decimal(), place2.lang.decimal())
}

type place struct {
	module string
	name string
	lon coordinate
	lang coordinate
}

func newPlace(m string, l string, lon coordinate, lang coordinate) place {
	return place{m, l, lon, lang}
}

func lesson25_text() {
	fmt.Printf("\n Lesson_25 \n \n")

	report := report {1, temperature{-10, 10}, location{20, 50}}
	fmt.Println(report)
	fmt.Printf("Average report: %v \n", report.temperature.average())
	fmt.Printf("Average report: %v \n", report.average())

	report_v1 := report_v1 { 1, temperature{-10, 10}, location{20, 50} }
	fmt.Printf("Average report_v1: %v \n", report_v1.average())
	fmt.Printf("Average report_v1: %v \n", report_v1.high)
}

type report struct {
	sol int
	temperature temperature
	location location
}

type location struct {
	lat, long float64
}

type temperature struct {
    high, low celsius
}

func (t temperature) average() celsius {
	return celsius((t.high + t.low) / 2)
}

func (r report) average() celsius {
	return r.temperature.average()
}

type report_v1 struct {
	sol int
	temperature
	location
}

func lesson26_text() {
	fmt.Printf("\n Lesson_26 \n \n")

	var talker_inst interface {
		talk() string
	}
	 
	talker_inst = martian{}
	fmt.Println(talker_inst.talk()) // Выводит: nack nack
	 
	talker_inst = laser(3)
	fmt.Println(talker_inst.talk()) // Выводит: pew pew pew

	shout(martian{}) // Выводит: NACK NACK
	shout(laser(2)) // Выводит: PEW PEW

	type starship struct {
		laser
	}
	 
	s := starship{laser(3)} // космический корабль starship также удовлетворяет требования интерфейса talker, позволяя ему использоваться вместе с shout
	 
	fmt.Println(s.talk()) // Выводит: pew pew pew
	shout(s) // Выводит: PEW PEW PEW
}

type martian struct{}
 
func (m martian) talk() string {
    return "nack nack"
}
 
type laser int
 
func (l laser) talk() string {
    return strings.Repeat("pew ", int(l))
}

func shout(t talker) {
    louder := strings.ToUpper(t.talk())
    fmt.Println(louder)
}

// Обычно интерфейсы объявляются как именованные типы
type talker interface {
	talk() string
}

func lesson26_task() {
	elysium := location_v2{
        Name: "Elysium Planitia",
        Lat:  coordinate{4, 30, 0.0, 'N'},
        Long: coordinate{135, 54, 0.0, 'E'},
    }

    bytes, err := json.MarshalIndent(elysium, "", "  ")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Println(string(bytes))
}

// type coordinate struct {
// 	d, m, s float64
// 	h       rune
// }

// func (c coordinate) decimal() float64 {
//     sign := 1.0
//     switch c.h {
//     case 'S', 'W', 's', 'w':
//         sign = -1
//     }
//     return sign * (c.d + c.m/60 + c.s/3600)
// }

type location_v2 struct {
    Name string     `json:"name"`
    Lat  coordinate `json:"latitude"`
    Long coordinate `json:"longitude"`
}

func (c coordinate) toString() string {
	return fmt.Sprintf("%v°%v'%.1f\" %c", c.d, c.m, c.s, c.h)
}

func (c coordinate) MarshalJSON() ([]byte, error) {
	// var struct_var struct{
	// 	DD float64 `json:"decimal"`
	// 	DMS string  `json:"dms"`
	// 	D   float64 `json:"degrees"`
	// 	M   float64 `json:"minutes"`
	// 	S   float64 `json:"seconds"`
	// 	H   string  `json:"hemisphere"`
	// }

	// struct_var.DD = c.decimal()
	// struct_var.DMS = c.toString()
	// struct_var.D = c.d
	// struct_var.M = c.m
	// struct_var.S = c.s
	// struct_var.H = string(c.h)

	return json.Marshal(
	struct{
		DD float64 `json:"decimal"`
		DMS string  `json:"dms"`
		D   float64 `json:"degrees"`
		M   float64 `json:"minutes"`
		S   float64 `json:"seconds"`
		H   string  `json:"hemisphere"`
	}{
		DD:  c.decimal(),
        DMS: c.toString(),
        D:   c.d,
        M:   c.m,
        S:   c.s,
        H:   string(c.h),
	})
}

func lesson27_task() {
	// TODO:
}

func lesson28_text() {
	fmt.Printf("\n Lesson_28 \n \n")

	value := 12

	fmt.Println(&value) // 0xc00001c9f0

	value_address := &value

	fmt.Println(*value_address) // 12
	fmt.Printf("Address это %T\n", value_address) // *int - указателем типа

	var linker *int

	linker = &value
	fmt.Printf("Linker: %v \n", *linker) // 12

	// Звездочка перед типом обозначает тип указателя, а звездочка перед названием переменной нужна для указания на значение, к которому отсылается указатель.

	var link_to_president *string 

	old_president := "Obama"
	link_to_president = &old_president
	fmt.Printf("Current president: %v \n", *link_to_president)
	
	new_president := "Baiden" // new_president занимает ячейку памяти, link_to_president = &new_president ссылка на эту ячейку
	link_to_president = &new_president
	fmt.Printf("Next president: %v \n", *link_to_president)

	new_president = "Tramp"
	fmt.Printf("Changed president: %v \n", *link_to_president)

	*link_to_president = "Baiden"
	fmt.Printf("Returned president: %v \n", new_president)

	premier := link_to_president
	*premier = "Putin?"
	fmt.Printf("Premer became president: %v \n", new_president)

	fmt.Println(premier == link_to_president)
	fmt.Println(premier == &new_president) // Указатели premier и link_to_president оба содержат один и тот же адрес памяти, следовательно, они равны

	ufo := "UFO Delegation"

	link_to_president = &ufo
	fmt.Println(premier == link_to_president)

	fake_ufo := *link_to_president // Копия значения!
	*link_to_president = "UFO second flight"

	fmt.Printf("Ufo: %v \n", ufo)
	fmt.Printf("Fake ufo: %v \n", fake_ufo) // не поменялася!

	a := "123"
	b := a
	c := *&a

	fmt.Printf("Addresess: %v %v %v \n", &a, &b, &c)
	fmt.Printf("Values: %v %v %v \n", a, b, c)

	rebecca := person{
		name:       "Rebecca",
		superpower: "imagination",
		age:        14,
	}
	 
	birthday(&rebecca)
	 
	fmt.Printf("%+v\n", rebecca) // Выводит: {name:Rebecca superpower:imagination age:15}

	birthday_value(rebecca)

	fmt.Printf("%+v\n", rebecca) // Выводит: {name:Rebecca superpower:imagination age:15}

	link_to_person := &rebecca

	link_to_person.birthday()

	fmt.Printf("%+v\n", *link_to_person) // Выводит: {name:Rebecca superpower:imagination age:15}

	unstaged_person := *link_to_person
	unstaged_person.birthday()
	// Альтернативно, вызов метода в следующем листинге не использует указатель, однако по-прежнему работает. 
	// Go автоматически определяет адрес переменной (&) при вызове метода через пояснение с точкой, поэтому вам не нужно писать (&unstaged_person).birthday()

	fmt.Printf("%+v\n", unstaged_person) // Выводит: {name:Rebecca superpower:imagination age:15}

	// !Структуры регулярно передаются с указателями.!
}

type person struct {
	name, superpower string
	age              int
}

func birthday(p *person) {
	p.age++ // Синтаксис следующего листинга предпочтительнее (*p).age++.
}

func birthday_value(p person) {
	p.age++ // Синтаксис следующего листинга предпочтительнее (*p).age++.
}
 
func (p *person) birthday() {
    p.age++
}

func copy_or_link() {
	// string, int, float
	string1 := "word"
	string2 := string1 // Копирует значение.

	string2 = "owl"

	fmt.Printf("String before: %v \n", string1)
	fmt.Printf("String after: %v \n", string2)

	// array
	array1 := [2]string {"word", "owl"}
	array2 := array1 // Копирует значение.

	array2[0] = "bird"

	fmt.Printf("Array before: %v \n", array1[0])
	fmt.Printf("Array after: %v \n", array2[0])

	// struct
	struct1 := &struct{ radius float64 }{ radius: 100.0 }
	struct2 := struct1 // Копирует значение.

	struct2.radius = 200.0

	fmt.Printf("Struct before: %v \n", struct1.radius)
	fmt.Printf("Struct after: %v \n", struct2.radius)

	// map
	map1 := map[string]int { "word": 2, "owl": 1 }
	map2 := map1 // Передает ссылку.

	map2["word"] = 100

	fmt.Printf("Map before: %v \n", map1["word"])
	fmt.Printf("Map after: %v \n", map2["word"])

	// slice
	slice1 := make([]string, 2)
	slice1 = append(slice1, "word", "owl")

	slice2 := slice1  // Передает ссылку.
	slice2[0] = "bird"

	fmt.Printf("Slice before: %v \n", slice1[0])
	fmt.Printf("Map after: %v \n", slice2[0])

}

func lesson28_task() {
	const(
		start_x = 0
		start_y = 0
		lifetime = 10
	)
	turtle := newTurtle(start_x, start_y)
	
	for i:=0; i < lifetime; i++ {
		Step(&turtle)
		time.Sleep(time.Second)
		turtle.show()
	}
}

type turtle struct {
	x int
	y int
}

func newTurtle(x int, y int) turtle {
	return turtle {x, y}
}

func (t *turtle) up() {
	t.y += 1
}

func (t *turtle) down() {
	t.y -= 1
}

func (t *turtle) left() {
	t.x -= 1
}

func (t *turtle) right() {
	t.x += 1
}

func (t *turtle) show() {
	fmt.Printf("turtle is on x: %v y: %v \n", t.x, t.y)
}

func Step(t *turtle) {
	switch rand.Intn(4) {
	case 0:
		t.up()
	case 1:
		t.down()
	case 2:
		t.right()
	case 3:
		t.left()
	}
}

func lesson29_text() {
	fmt.Printf("\n Lesson_29 \n \n")
	var nowhere *int
	fmt.Println(nowhere) // Выводит: <nil>

	if nowhere != nil {
		fmt.Println(*nowhere)
	}

	var fn func(a, b int) int
	fmt.Println(fn == nil) // Выводит: true

	food := []string{"onion", "carrot", "celery"}
    sortStrings(food, nil)
    fmt.Println(food) // Выводит: [carrot celery onion]

	// slice
	var soup []string
	fmt.Println(soup == nil) // Выводит: true
	
	for _, ingredient := range soup {
		fmt.Println(ingredient)
	}
	
	fmt.Println(len(soup)) // Выводит: 0
	soup = append(soup, "onion", "carrot", "celery")
	fmt.Println(soup) // Выводит: [onion carrot celery]

	// map
	var soup1 map[string]int
	fmt.Println(soup == nil) // Выводит: true
	
	measurement, ok := soup1["onion"]
	if ok {
		fmt.Println(measurement)
	}

	// interface
	var v interface{}
	fmt.Printf("%T %v %v\n", v, v, v == nil)
}

func sortStrings(s []string, less func(i, j int) bool) {
    if less == nil {
        less = func(i, j int) bool { return s[i] < s[j] }
    }
    sort.Slice(s, less)
}

func lesson30_text() {
	fmt.Printf("\n Lesson_30 \n \n")

	// files, err := ioutil.ReadDir(".")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	
	// for _, file := range files {
	// 	fmt.Println(file.Name())
	// }
	// Старый метод
	// f, err := os.Create(name)
    // if err != nil {
    //     return err
    // }
	// Новый
	// type safeWriter struct {
	// 	w   io.Writer
	// 	err error // Место для хранения первой ошибки
	// }
	 
	// func (sw *safeWriter) writeln(s string) {
	// 	if sw.err != nil {
	// 		return // Пропускает запись, если раньше была ошибка
	// 	}
	// 	_, sw.err = fmt.Fprintln(sw.w, s) // Записывает строку и затем хранить любую ошибку
	// }
	// f, err := os.Create(name)
    // if err != nil {
    //     return err
    // }
    // defer f.Close()
 
    // sw := safeWriter{w: f}
    // sw.writeln("Errors are values.")
    // sw.writeln("Don't just check errors, handle them gracefully.")
    // sw.writeln("Don't panic.")
	// return sw.err // Возвращает ошибку в случае ее возникновения

	// func (g *Grid) Set(row, column int, digit int8) error {
	// 	if !inBounds(row, column) {
	// 		return errors.New("out of bounds")
	// 	}
	 
	// 	g[row][column] = digit
	// 	return nil
	// }

	// func main() {
	// 	var g Grid
	// 	err := g.Set(10, 0, 5)
	// 	if err != nil {
	// 		fmt.Printf("An error occurred: %v.\n", err)
	// 		os.Exit(1)
	// 	}
	// }
	// var (
	// 	ErrBounds = errors.New("out of bounds")
	// 	ErrDigit  = errors.New("invalid digit")
	// )

	// if !inBounds(row, column) {
	// 	return ErrBounds
	// }

	err := something(1, 4)

	if err != nil {
		fmt.Println("Can not do something.")
		os.Exit(1)
	}

	dog := dog{ 12 }

	err = dog.Gaw()

	if err != nil {
		if errs, ok := err.(ManyErrors); ok {
			fmt.Printf("%d error(s) occurred:\n", len(errs))
			for _, e := range errs {
				fmt.Printf("- %v\n", e)
			}
		}
	}

	fmt.Printf("- %v\n", NotBirth.Error())
	fmt.Printf("- %v\n", NotBirth)
}

var (
	OutOfRange = errors.New("Out of range")
	NotIntValue = errors.New("Not in value")
	ToOld = errors.New("To old")
	ToYoung = errors.New("To young")
	NotBirth = errors.New("Not birth")
)

func something(x int, y int) error {
	if x < 0 || y < 0 {
		return OutOfRange
	}

	if reflect.TypeOf(x).Kind() != reflect.Int {
		return NotIntValue
	}
	fmt.Println("Everything is good.")
	return nil
} 

type dog struct {
	age int
}

// error < итерфейс. Срез []error удалетворяет инферфейсу error.
type ManyErrors []error

func (me ManyErrors) Error() string {
	var s []string
	for _, err := range me {
        s = append(s, err.Error()) // Конвертирует ошибки в строки
    }
    return strings.Join(s, ", ")
}

func (d dog) Gaw() error {
	var errs ManyErrors

	if d.age > 10 {
		errs = append(errs, ToOld)
	}

	if d.age < 2 && d.age > 0 {
		errs = append(errs, ToYoung)
	}

	if d.age == 0 {
		errs = append(errs, NotBirth)
		return errs
	}

	return errs
}

func lesson35_text() {
	fmt.Printf("\n Lesson_35 \n \n")

	user := User {1, "Стенли", "Кубрик"}

	json_data, err := json.MarshalIndent(user, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json_data))

	users := []User {
		{1, "Стенли", "Кубрик"},
		{2, "Девид", "Линч"},
	}

	json_data1, err := json.MarshalIndent(users, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json_data1))

	var u1 User

	json_data2 := []byte(`{
		"Id": 1,
		"Name": "Стенли",
		"Surname": "Кубрик"
	  }`)

	err1 := json.Unmarshal(json_data2, &u1)

	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Println(u1)

	filename, err := os.Open("data.json")

	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(filename)
 
    if err != nil {
		filename.Close()
        log.Fatal(err)
    }

	var result []ParsedUser

	jsonErr := json.Unmarshal(data, &result)
 
    if jsonErr != nil {
		filename.Close()
        log.Fatal(jsonErr)
    }

	fmt.Println(result)

	url := "http://api.open-notify.org/astros.json"
 
    var netClient = http.Client{
        Timeout: time.Second * 10,
    }

	res, err := netClient.Get(url)

	if err != nil {
        log.Fatal(err)
    }
 
    defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
 
    fmt.Println(body)
 
    if err != nil {
        log.Fatal(err)
    }

	astros := people{}
 
    jsonErr1 := json.Unmarshal(body, &astros)
 
    if jsonErr1 != nil {
        log.Fatal(jsonErr1)
    }
 
    fmt.Println(astros)

}

type User struct {
	Id int
	Name string
	Surname string
}

type ParsedUser struct {
    Name       string
    Occupation string
    Born       string
}

type Astronaut struct {
    Name  string
    Craft string
}
 
type people struct {
    Number  int
    People  []Astronaut
    Message string
}

