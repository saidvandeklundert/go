The slice is the most important go to data type in Go.

The slice is a reference type.

It is designed to stay on your stack. The slice is 24 bytes in size and has only 3 fields:
- pointer to a backing array
- length
- capacity

```go
// Various ways to create a slice:

    // Zero value slice(nil):
var aSlice []string                 // aSlice == nil: true.
    // Create zeroed array with 10 elements and return the slice for it:
eSlice := make([]int, 10)           // len(eSlice): 10
    // Slice literal:
bSlice := []bool{true, true, false}    


// Slice of integers:
slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

len(slice)          // return lenght of the slice: 11
cap(slice)          // return the size of the slice backing array

// copy will copy the backing array as well.
// Copy a slice into newS, make sure the slice is big enough:
newS := make([]int, 11); 
copy(newS, slice)       // [0 1 2 3 4 5 6 7 8 9 10]            
newS := make([]int, 2)
copy(newS, slice)       // [0 1]
newS := make([]int, 20)         
copy(newS, slice)       // [0 1 2 3 4 5 6 7 8 9 10 0 0 0 0 0 0 0 0 0]  
// copy(to, from), the 'to' value will be a new instance with a new pointer: 


// Slicing: slice[start(including):stop(NOT including)]
slice[1]            // 1
slice[2:6]          // [2 3 4 5]
slice[3:]           // [3 4 5 6 7 8 9 10]
slice[:7]           // [0 1 2 3 4 5 6]

// For the second value, consider thinking about how long of a slice you would like.
// Then add that to the start value.
// Exampe, starting at index 2, wanting a slice with a length of 4:
// 2 + 4 = 6, slice[2:6]

// You can supply a third parameter to a slice to indicate the capacity:
slice[2:6:4]

// This limits the size of the new slice. 
// If things are appended, a new backing-array is created and the 'old' slice is left intact.

// For loop a slice:
for _, nr := range slice {
	fmt.Println(nr)
}

// Appending:
Slice := []string{"a", "b"}
Slice = append(Slice, "c", "d")             // [a b c d]
SliceAddition := []string{"e", "f"}     
Slice = append(Slice, SliceAddition...)     // [a b c d e f]

// Joining (strings):
Joined := strings.Join(Slice, ", ")         // a, b, c, d, e, f (string)


// Write slice to disk:
Slice := []string{"a", "b", "c"}
func saveFile(fileName string, slice []string) {
	stringSlice := strings.Join(slice, ",")
	ioutil.WriteFile(fileName, []byte(stringSlice), 0666)
}
saveFile("test.txt", Slice)     // cat test.txt: a,b,c

// Reading the same file from disk:
byteSlice, err := ioutil.ReadFile("test.txt")
if err != nil {
	os.Exit(1)
}
text := string(byteSlice)
fmt.Println("File content:", text)
s := strings.Split(text, ",")        // strings.Split returns slice. s: []string

```

