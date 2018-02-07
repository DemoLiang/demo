
/*package main

import "fmt"

func max(a, b int) int {
	if a > b{
		return a
	}
	return b
}

func main() {
	x := 3
	y := 4
	z := 5

	max_xy := max(x,y)
	max_xz := max(x,z)

	fmt.Printf("max(%d,%d) = %d \n",x,y,max_xy)
	fmt.Printf("max(%d,%d) = %d \n",x,z,max_xz)
	fmt.Printf("max(%d,%d) = %d \n",y,z,max(y,z))


	fmt.Printf("hello word!\n")
}
*/

/*
package main

import "fmt"

type person struct{
	name string
	age int
}

func Older(p1,p2 person) (person,int) {
	if p1.age > p2.age{
		return p1,p1.age-p2.age
	}
	return p2,p2.age-p1.age
}

func main() {
	var tom person

	tom.name,tom.age = "Tom",18
	bob := person{age:25,name:"Bob"}
	paul :=person{"paul",43}

	tb_Older,tb_diff := Older(tom,bob)
	tp_Older,tp_diff := Older(tom,paul)
	bp_Older,bp_diff := Older(bob,paul)

	fmt.Printf("Of %s and %s,%s is older by %d years\n",
		tom.name,bob.name,tb_Older.name,tb_diff)
	fmt.Printf("Of %s and %s,%s is older by %d years\n",
		tom.name,paul.name,tp_Older.name,tp_diff)
	fmt.Printf("Of %s and %s,%s is older by %d years\n",
		bob.name,paul.name,bp_Older.name,bp_diff)
}
*/
/*
package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v:=range r.Form {
		fmt.Println("key:",k)
		fmt.Println("val:",strings.Join(v,""))
	}
	fmt.Fprint(w,"hello world!")
}

func main() {
	http.HandleFunc("/",sayhelloName)
	err:=http.ListenAndServe(":9090",nil)
	if err != nil{
		log.Fatal("ListenAndServer:",err)
	}
}
*/

/*
package main

import "fmt"
import "math/rand"
import "time"

func main() {
	for i := 0;i<10;i++{
		a:=rand.Int()
		fmt.Printf("%d\n",a)
	}
	fmt.Println()
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f\n",100*rand.Float32())
	}
}
*/

/*
package main

import "fmt"
import "strconv"

func main() {
	var orig string = "666"
	var an int
	var newS string

	fmt.Printf("the size of ints is:%d \n",strconv.IntSize)

	an,_ = strconv.Atoi(orig)
	fmt.Printf("the integer is %d\n",an)
	an =an +5
	newS = strconv.Itoa(an)
	fmt.Printf("the new string is :%s\n",newS)
}
*/

/*
package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d is even: is %t \n",16,even(16))
	fmt.Printf("%d is odd:is %t \n",17,odd(17))
	fmt.Printf("%d is odd:is %t \n",18,odd(18))
}

func even(nr int) bool{
	if nr == 0{
		return true
	}
	return odd(RevSign(nr)-1)
}
func odd(nr int) bool {
	if nr == 0{
		return false
	}
	return even(RevSign(nr)-1)
}

func RevSign(nr int) int {
	if nr < 0{
		return -nr
	}
	return nr
}
*/

/*
package main

import "fmt"
import "time"

const LIM  = 41

var fibs [LIM]uint64

func main() {
	var result uint64 = 0
	start := time.Now()
	for i:=0;i<LIM;i++{
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is : %d\n",i,result)
	}
	end := time.Now();
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time:%s \n",delta)
}

func fibonacci(n int) (res uint64) {
	if fibs[n] != 0{
		res = fibs[n]
		return
	}
	if n <= 1{
		res = 1
	}else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fibs[n] = res
	return
}

*/

/*
package main

import "fmt"

func main() {
	array := [3]float64{7.0,8.5,9.1}
	x := Sum(&array)
	fmt.Printf("the sum of the array is %f\n",x)
}
func Sum(a *[3]float64) (sum float64){
	for b,v := range a{
		sum +=v
		fmt.Println(b)
		fmt.Println(v)
	}
	return
}
*/


/*
package main

import "fmt"

func main() {
	var mapLit map[string]int
	var mapAssigned map[string]int

	mapLit = map[string]int{"one":1,"two":2}
	mapCreated :=make(map[string]float32)
	mapAssigned = mapLit

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3

	fmt.Printf("Map literal at \"one\" is : %d\n",mapLit["one"])
	fmt.Printf("Map create at \"key2\" is : %f\n",mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is :%d\n",mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is : %d\n",mapLit["ten"])
}
*/

/*
package main

import "fmt"

func main() {
	var value int
	var isPresent bool

	map1 :=make(map[string]int)
	map1["New Delhi"] = 55
	map1["Beijing"] = 20
	map1["Washington"] = 25
	value,isPresent = map1["Beijing"]
	if isPresent{
		fmt.Printf("the value of \"Beijing\" in map1 is :%d\n",value)
	}else {
		fmt.Printf("map1 does not contain Beijing\n")
	}

	value,isPresent = map1["Paris"]
	fmt.Printf("is \"Paris\" in map1 ?: %t\n",isPresent)
	fmt.Printf("value is:%d\n",value)

	delete(map1,"Washington")
	value,isPresent = map1["Washington"]
	if isPresent{
		fmt.Printf("the value of \"Washington\" in map1 is :%d\n",value)
	}else {
		fmt.Printf("map1 does not contain Washington\n")
	}
}

*/

/*
package main

import "fmt"

func main() {
	items := make([]map[int]int,5)
	for i:=range items{
		items[i] = make(map[int]int,1)
		items[i][1] = 2
	}
	fmt.Printf("Version A:value of items:%v\n",items)

	items2 := make([]map[int]int,5)
	for _,item := range items2{
		item = make(map[int]int ,1)
		item[1] = 2
	}
	fmt.Printf("Version b:value of items:%v\n",items2)
}

*/

/*
package main

import "fmt"
import "regexp"
import "strconv"

func main() {
	searchIn :="John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+"

	f := func(s string) string{
		v,_ := strconv.ParseFloat(s,32)
		fmt.Printf("%f \n",v)
		return strconv.FormatFloat(v*2,'f',2,32)
	}
	if ok,_ :=regexp.Match(pat,[]byte(searchIn));ok{
		fmt.Println("match found")
	}
	re,_ := regexp.Compile(pat)
	str := re.ReplaceAllString(searchIn,"##.#")
	fmt.Println(str)

	str2 := re.ReplaceAllStringFunc(searchIn,f)
	fmt.Println(str2)
}

*/

/*
package main

import "fmt"
import "math"
import "math/big"

func main() {
	im := big.NewInt(math.MaxInt64)
	in := im
	io := big.NewInt(1956)
	ip := big.NewInt(1)
	ip.Mul(im,in).Add(ip,im).Div(ip,io)
	fmt.Printf("big int:%v\n",ip)

	rm := big.NewRat(math.MaxInt64,1956)
	fmt.Printf("%v\n",rm)
	rn := big.NewRat(-1956,math.MaxInt64)
	ro := big.NewRat(19,56)
	rp := big.NewRat(1111,2222)
	rq := big.NewRat(1,1)
	rq.Mul(rm,rn).Add(rq,ro).Mul(rq,rp)
	fmt.Printf("Big Rat:%v\n",rq)
}

*/

/*
package main

import "fmt"
import "reflect"

type TagType struct {
	field1 bool	"an important answer"
	field2 string	"The name of the thing"
	field3 int	"How much there are"
}

func main() {
	tt := TagType{true,"Barak Obama",1}
	for i:=0;i<3;i++{
		refTag(tt,i)
	}
}

func refTag(tt TagType,ix int)  {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n",ixField.Tag)
}
*/

/*
package main

import "fmt"

type Shaper interface {
	Area() float32
}
type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func main() {
	sq1 := new(Square)
	sq1.side = 5

	areaIntf := sq1
	fmt.Printf("the square has area:%f\n",areaIntf.Area())
}
*/

/*
package main

import "fmt"

type stockPosition struct {
	ticker string
	sharePrice float32
	count float32
}

func (s stockPosition) getValue() float32{
	return s.sharePrice * s.count
}

type car struct {
	make string
	model string
	price float32
}

func (c car) getValue() float32 {
	return c.price
}

type valuable interface {
	getValue() float32
}

func showValue(asset valuable) {
	fmt.Printf("value of asset is %f\n",asset.getValue())
}

func main() {
	var o valuable = stockPosition{"GOOG",577.20,4}
	showValue(o)
	o = car{"BMW","M3",66500}
	showValue(o)
}

*/

/*
package main

import "fmt"

type List []int

func (l List) Len() int{
	return len(l)
}

func (l *List) Append(val int) {
	*l = append(*l,val)
}

type Appender interface {
	Append(int)
}

func CountInto(a Appender, start, end int) {
	for i:=start;i<=end;i++{
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

func main() {
	var lst List
	if LongEnough(lst){
		fmt.Printf("- lst is long enough\n")
	}
	plst := new(List)
	CountInto(plst,1,10)
	if LongEnough(plst){
		fmt.Printf("- plst is long enough\n")
	}
}
*/

package main

import "./sort"
import 	"fmt"

func ints() {
	data := []int{74,59,238,-784,9845,959,905,0,0,42,7586,-5467984,7586}
	a := sort.IntArray(data)
	sort.Sort(a)
	if !sort.IsSorted(a){
		panic("fails")
	}
	fmt.Printf("the sorted array is:%v\n",a)
}
func strings() {
	data :=[]string{"monday","friday","tuesday","wednesday","sunday","thursday","","saturday"}
	a := sort.StringArray(data)
	sort.Sort(a)
	if !sort.IsSorted(a){
		panic("fail")
	}
	fmt.Printf("the sorted array is:%v\n",a)
}

type day struct {
	num int
	shortName string
	longName string
}

type dayArray struct {
	data []*day
}

func (p *dayArray) Len() int {
	return len(p.data)
}
func (p *dayArray) Less(i,j int) bool{
	return p.data[i].num < p.data[j].num
}
func (p *dayArray) Swap(i, j int) {
	p.data[i],p.data[j] = p.data[j],p.data[i]
}

func days() {
	Sunday := day{0,"SUN","Sunday"}
	Monday := day{1,"MON","Monday"}
	Tuesday :=day{2,"TUE","Tuesday"}
	Wednesday := day{3,"WED","Wednesday"}
	Thursday := day{4,"THU","Thursday"}
	Friday := day{5,"FRI","Friday"}
	Saturday := day{6,"SAT","Saturday"}

	data :=[]*day{&Tuesday,&Thursday,&Wednesday,&Sunday,&Monday,&Friday,&Saturday}
	a := dayArray{data}
	sort.Sort(&a)
	if !sort.IsSorted(&a){
		panic("fail")
	}
	for _, d := range data{
		fmt.Printf("%s ",d.longName)
	}
	fmt.Printf("\n")
}

func main() {
	ints()
	strings()
	days()
}