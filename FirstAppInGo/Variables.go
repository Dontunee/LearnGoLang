//Declaring Variables
var i int
var s string
i = 20
s = "Some string"

//Omitting the type when initializing the variable
var k int = 35
var score = 10

//Using the shot variable declarator :=
j := 50
str := "Some string"

//Declaring multiple Variables
firstName, lastName := "tunde", "afolabi"

//Variable declaration block
var (
	name = "Tunde Afolabi"
	afe  = 50
)

//Declaring a constant // name of constant is written in uppercase
const PI float64 = 3.14159265359
const VALUE = 1000

//Converting datatypes
//string to float
s := "3.1428383939393"
f, err := strconv.ParseFloat(s, 8)

if err != null {
	fmt.Println("Error convert string to integer", err)
}

fmt.Println(reflect.TypeOf(f))

//Float to string
flo := 3.1415926535
strfLO := strconv.FormatFloat(flo, 'E', -1, 32)
fmt.Println(reflect.TypeOf(strfLO))
fmt.Println(strfLO)

//Float to int 
f := 3.14533394894
iInt := int32(f)
fmt.Println(iInt)

//Working with boolean 
var isTRUE bool = true;
fmt.Printf("%v, %T\n", n,n)

//CONSTANTS
//Typed Constants
const myConst int = 42
fmt	.Printf("%v,%T", n,n)

//Arrays 

grades := [3]int{1,2,3}

//Slices

grades := []{1,2,3,4}
make([]int, 3)
make([]int, 3,100)
append(a,10)

//Maps 

statePopulations := make(map[string]int)

statePopulations = map[string]int{
	"California": 122323,
	"Ohio": 239029320
}

//Get Ohio value
statePopulations["Ohio"]

//Delete a key 
delete(statePopulations, "California")

//Test to confirm it has been removed
//ok is false if it doesnt exist
pop,ok := statePopulations["California"]
fmt.Println(pop,ok)


//Structs

//Declaring a struct 

type Person struct {
	name string,
	age int,
	occupation string,
	Friends []string 
}

//Anonymous struct 
aDoctor := struct{name string}{name: "John Tunde"}

//Composition
type Animal struct{
	Name string
	Origin string
}

type Bird struct{
	Animal
	SpeedKph float32
	CanFly bool
}

smallBird := Bird{}
smallBird.Name = "Emu"
smallBird.Origin = "Australia"


bigBird := Bird{
	Animal : Animal{Name: "Emu", Origin: "Australia"},
	SpeedKph : 20923,
	CanFly: false,
}

//Using TAGS

tags Car struct{
	Name string `required max: "100"`
	Origin string
}

t := reflect.TypeOf(Animal{})
field, _ := t.FieldByName("Name")
fmt.Println(field.Tag)

//If statement
number := 30
if number > 1{

}
else if number > 10 {

}

//Switch statement 

switch number{
	case 1,5,10:
		fmt.Println("one five or ten")
	case 2,4 :
		fmt.Println("two or four")
	default:
		fmt.Println("änother number")
}

//Type Switch

var i interface{} = 1
switch i.(type) {
case int:
	fmt.Println("This is an integer")
case float64:
	fmt.Println("This is a floating number")
case string:
	fmt.Println("This is a string")
default:
	fmt.Println("Type doesnt exist")
}

//Looping 

//Simple loops
for i := 0; i < 5; i++{
	fmt.Println(i)
}

//Looping through a collection

s := []int{1,2,3}

for k,v := range s{
	fmt.Println(k,v)
}

//Defer 

defer fmt.Println("middle")


//Interfaces

type Writer interface{
	Write([] byte) (int.error)
}

type ConsoleWriter struct{
	func (cw ConsoleWriter) Write(data []byte) (int ,error){
		n, err := fmt.Println(string(data))
		return n,err 
	}
}