package main

import "fmt"

func main() {

	fmt.Println("Slices")

	// fixed size is something like a string to be substringed with slices - reused elements
	months := [...]string{1: "January", 2: "February", 3: "March",
		4: "April", 5: "May", 6: "June",
		7: "July", 8: "August", 9: "September",
		10: "October", 11: "November", 12: "December"}
	// cmp the conversion from strings to bytes -> where immutability was important
	// here **immutability is not** as important...

	fmt.Println(months)
	// slices are something like substrings -> do not copy, but reuse underlying array
	// **slice operator** s[i:j] 0 ≤ i ≤ j ≤ cap(s) where s is array variable/pointer to array/slice
	// default i = 0, default j= len(s) **NOT cap(s)**
	Q2 := months[4:7]                       // **capacity is until the end size of the underlying array or capacity of slice**
	fmt.Println(Q2, cap(Q2), len(months)-4) // ["April" "May" "Imaginary Month"] -> capacity = len(months) - 4
	summer := months[6:9]
	fmt.Println(summer, cap(summer), len(months)-6) // ["Imaginary Month" "July" "August"]
	Q2[2] = "\"Imaginary Month\""                   // affects summer month of June
	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}
	// ** capacity ** is a hook to the underlying array - what can the slice SAFELY use

	// fmt.Println(summer[:20])    // panic: out of range -> beyond capacity
	endlessSummer := summer[:5]                                 // extend a slice (within capacity)!!!
	fmt.Println(endlessSummer, cap(endlessSummer), cap(summer)) // "[June July August September October]" -> same capacity

	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])  // creating a slice
	fmt.Println(a) // "[5 4 3 2 1 0]" underlying array is reversed

	s := []int{0, 1, 2, 3, 4, 5}
	// main.go:43      0x4a50af        488d05ca980000                  lea rax, ptr [rip+0x98ca] -> ??
	// main.go:43      0x4a50b6        e8c57af6ff                      call $runtime.newobject // returns address of array in rax
	// main.go:43      0x4a50bb        48898424c8000000                mov qword ptr [rsp+0xc8], rax -> rax returns address of array
	// main.go:43      0x4a50c3        48c70000000000                  mov qword ptr [rax], 0x0 // first value
	// main.go:43      0x4a50ca        488b9424c8000000                mov rdx, qword ptr [rsp+0xc8] // for offsetting
	// main.go:43      0x4a50d4        48c7420801000000                mov qword ptr [rdx+0x8], 0x1 //second element of array
	// main.go:43      0x4a50e6        48c7421002000000                mov qword ptr [rdx+0x10], 0x2
	// main.go:43      0x4a50f8        48c7421803000000                mov qword ptr [rdx+0x18], 0x3
	// main.go:43      0x4a510a        48c7422004000000                mov qword ptr [rdx+0x20], 0x4
	// main.go:43      0x4a511c        48c7422805000000                mov qword ptr [rdx+0x28], 0x5
	// main.go:43      0x4a5124        488b8424c8000000                mov rax, qword ptr [rsp+0xc8]
	// main.go:43      0x4a512e        eb00                            jmp 0x4a5130
	// main.go:43      0x4a5130        48898424a8010000                mov qword ptr [rsp+0x1a8], rax -> offset of array
	// main.go:43      0x4a5138        48c78424b001000006000000        mov qword ptr [rsp+0x1b0], 0x6 -> len
	// main.go:43      0x4a5144        48c78424b801000006000000        mov qword ptr [rsp+0x1b8], 0x6 -> cap
	// Rotate s left by two positions. -> Slicing modifies the UNDERLYING ARRAY!
	reverse(s[:2]) // {0, 1} -> {1, 0}
	reverse(s[2:]) // {2, 3 , 4 ,5 } -> {5, 4, 3, 2}

	reverse(s)     // {1, 0, 5, 4, 3, 2}
	fmt.Println(s) // "[2 3 4 5 0 1]"

	// the only comparison of slics is to nil -> as in if s == nil -> cannot compare slices
	{
		// nil - no underlying array
		var s []int    // len(s) == 0, s == nil
		s = nil        // len(s) == 0, s == nil
		s = []int(nil) // len(s) == 0, s == nil -> conversion expression
		// main.go:69      0x4a52cc        48c784243002000000000000        mov qword ptr [rsp+0x230], 0x0
		// main.go:69      0x4a52d8        440f11bc2438020000              movups xmmword ptr [rsp+0x238], xmm15
		// main.go:70      0x4a52e1        48c784243002000000000000        mov qword ptr [rsp+0x230], 0x0
		// main.go:70      0x4a52ed        440f11bc2438020000              movups xmmword ptr [rsp+0x238], xmm15
		// main.go:71      0x4a52f6        48c784243002000000000000        mov qword ptr [rsp+0x230], 0x0
		// main.go:71      0x4a5302        440f11bc2438020000              movups xmmword ptr [rsp+0x238], xmm15
		if s == nil {
			fmt.Println(s)
		}
		// ** empty slite literal -> NOT NIL
		s = []int{} // len(s) == 0, s != nil ** here s IS NOT NIL but zero-initialized members
		// main.go:82      0x4a53b8        488d15b9ce0f00                  lea rdx, ptr [runtime.zerobase] -> get a non-nil pointer
		// main.go:82      0x4a53bf        4889942418010000                mov qword ptr [rsp+0x118], rdx -> non-nil pointer
		// main.go:82      0x4a53cb        4889942430020000                mov qword ptr [rsp+0x230], rdx -> non-nil pointer
		// main.go:82      0x4a53d3        440f11bc2438020000              movups xmmword ptr [rsp+0x238], xmm15
		if s != nil {
			fmt.Println(s)
			fmt.Println("Capacity of non-nil empty literal slice:", cap(s)) // 0
		}

		// test for emptiness
		if len(s) == 0 {
			fmt.Println("Slice is empty")
		}

		// passing nil -> perfectly safe
		reverse(nil) // will be compiled as nop -> xor eax, eax e.g.

	}

	// make built-in function
	// make([]T, len) // capacity = len -> underlying array length = capacity
	// make([]T, len, cap) // same as make([]T, cap)[:len]

	// 4.2.1 The append Function
	fmt.Println("\n4.2.1 The append Function")

	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes) // "['H' 'e' 'l' 'l' 'o' ',' '世', '界']"
	// alternatively
	runes = []rune("Hello, 世界")
	fmt.Printf("%q\n", runes) // "['H' 'e' 'l' 'l' 'o' ',' '世', '界']"

	// append internals
	var ints []int
	printCopyDemo = true
	ints = appendInt(ints, 3)
	fmt.Println(ints, cap(ints)) // [3] 1

	fmt.Println("\nCapacity growth")
	printCopyDemo = false
	var x, y []int //two slics to showcase what append returns
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d \t%v\n", i, cap(y), y)
		x = y
	}

	{
		fmt.Println("\nBuilt-in append")
		var x []int
		x = append(x, 1)
		x = append(x, 2, 3) // multiple values
		x = append(x, 4, 5, 6)
		x = append(x, x...) // append the slice x -> ... supplies a list of arguments
		fmt.Println(x)      // "[1 2 3 4 5 6 1 2 3 4 5 6]"
	}

	// 4.2.2 In-Place Slice Techniques
	fmt.Println("\n4.2.2 In-Place Slice Techniques")
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data)) // `["one" "three"]`
	// to avoid the possible confusion from the overwriting
	// data = nonempty(data)
	fmt.Printf("%q\n", data) // `["one" "three" "three"]` -> the second three is "" overriden

	data = []string{"one", "", "three"}
	data = nonemptyWithAppend(data)
	fmt.Printf("%q\n", data) // `["one" "three"]`

	// stack implementation

	stack := []int{}
	stack = append(stack, 3) // push
	stack = append(stack, 4)
	top := stack[len(stack)-1]
	fmt.Println(top)             // 4
	stack = stack[:len(stack)-1] // pop
	fmt.Println(stack)           // 3

	// remove in place
	{
		s := []int{5, 6, 7, 8, 9}
		fmt.Println(remove(s, 2)) // "[5 6 8 9]"
		s = []int{5, 6, 7, 8, 9}
		fmt.Println(removeNoOrder(s, 2)) // "[5 6 9 8]""
	}

}

var printCopyDemo bool

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// bytes.Equal exists, but for other slices this should be the function w/o generics
//lint:ignore U1000 ...
func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func appendInt(x []int, y int) []int {
	if len(x)+1 <= cap(x) {
		// There is room to grow. Extend the slice.
		x = x[:len(x)+1]
		x[len(x)-1] = y
		return x
	}

	// There is insufficient space. Allocate a new array.
	// Grow by doubling, for amortized linear complexity.
	var z []int
	// **len(x) == cap(x)** -> mine from len cannot be larger than cap (see first condition above)
	zlen := cap(x) + 1 // the needed new len
	var zcap int
	// double capacity of existing unless 0 which initializes to 1 capacity
	if cap(x) == 0 {
		zcap = 1
	} else {
		zcap = 2 * cap(x)
	}

	z = make([]int, zlen, zcap)
	// a built-in function; (destination, source) instead of loop
	// can be smaller into bigger
	copy(z, x)
	if printCopyDemo {
		fmt.Println("\nCopy built-in")
		bigger := []int{3, 3, 4}
		smaller := []int{1, 2}
		copied := copy(bigger, smaller) // destination > source
		fmt.Println(bigger, copied)     // [1,2,4] 2
		copied = copy(smaller, bigger)  // source > destination
		fmt.Println(smaller, copied)    // [1,2] 2
	}

	z[len(x)] = y
	return z
}

// what slices resemble -> not pure reference types (not just a pointer)
type IntSlice struct {
	//lint:ignore U1000 ...
	ptr *int // pointer to address to a member of the underlying array
	//lint:ignore U1000 ...
	len, cap int
}

// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonemptyWithAppend(strings []string) []string {
	// interesting technique -> getting zero-length of original same underying array
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func remove(slice []int, index int) []int {
	copy(slice[index:], slice[index+1:]) // copying left effectively
	// need to shrink the slice by 1
	return slice[:len(slice)-1]
}

func removeNoOrder(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
