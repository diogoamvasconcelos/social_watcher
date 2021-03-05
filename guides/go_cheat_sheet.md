- slice half-open indexing:

  - s[1:len(s)-1]
  - s[1:]
  - s[:2]

- zero/implicit initialization of vars (`var s string`)

  - string: ""
  - int/float64: 0
  - bool: false

- forloop

  - normal: `for initialization; condition; post {`
  - while: `for condition {`
  - while true: `for {`
  - range: `for i, val := range s {`

- strings

  - join: `strings.Join(os.Args[1:], ",")`

- Stopwatch

  ```
  start := time.Now()
  fmt.Printf("\nTook: %.2fs", time.Since(start).Seconds())
  ```

- inline `if err`

  - `if err := DoSomething(); err != nil {}`

- tuple assignemnt

  ```
  i, j := 0, 1 // declaration
  i, j = j, i // tuple assignement (swap i, j)
  ```

- Printf:

  - %d: decimal integer
  - %x, %o, %b: integer in hex, oct and binary
  - %f, %g, %e: floating point, floatint point more precision, scientific representation (exp)
  - %t: boolean
  - %s: string
  - %q: quoted string
  - %v: any value in a natural format (struct, etc)
  - %T: type of any type (useful for dynamic interfaces: `interface{}`)
  - %%: literal percent sign (no operand)
  - %[1]{some letter}: use the first operand (e.g: `fmt.Printf("%d, %[1]x", 15) -> "15 F"`)
  - %*s: the asterisc prints a string padded with number of spaces( e.g: `fmt.Printf("%*sL", depth, " ")`)

- type assertion

```
  var i interface{} = "hello"
  s, ok := i.(string) // and discard the value: `_, ok := i.(string)`
```

- array initialization
  ` a := []int{0,1,2}`

- operator precedence (decreasing order):

  - `* / % << >> & &^` ( &^ -> AND NOT, bit clear)
  - `+ - | ^` ( ^ -> bitwise XOR)
  - `== != < <= > >=`
  - `&&`
  - `||`

- use raw string literals (``, back quote) as don't have to escape any char fir:

  - regex
  - JSON
  - HTML templates

- for effecient string manupulation (like appending many times, which makes a copy every time as strings are immutable), use `bytes.Buffer` (pg.74 of gobook)

- using `iota` for const/enum initialization

  ```
  const (
    a = 0,
    b,
    c = 1,
    d
  )
  ```

  initializes a=0,b=0,c=1,d=1

  with `iota`, it will be incremented by one

  ```
  const(
    a = iota,
    b,
    c,
    d
  )
  ```

  a=0,b=1,c=2,d=3

  for bitflags, this is pretty handy

  ```
  const(
    on = (1 << iota),
    even,
    blinking,
    fast
  )
  ```

- composite types:

  - array: fixed size, rarely used (use slices instead, arrays is for lowlevel things)
    - initialization: `var arr [3]int = [3]int{1,2,3}` ,or `var arr = [...]int{1,2}`
      - `r := [...]int{99:-1}` is an array with 100 elemetns, all 0, but the last one which is -1
  - slice: dynamic size
    - initialization: `arr := []int{1,2,3}`
  - struct: fixed size

    - use `struct embedding` and `anonymous fields`

      ```
      type Point struct {
        X, Y, int
      }

      type Circle struct {
        Point
        Radius int
      }

      c := Circle{Point{0,0}, 1}
      ```

      > Note, the anonymous field can be a pointer `&Point` for example

  - map: dynamic size
    - can be used to emulate a `set` (unique keys)

- all arguments are passed by value (copied) except the following that are by reference:

  - pointers
  - slices
  - maps
  - functions
  - channels

- error handling

  - use `fmt.ErrorF` to format errors, example
    `return nil, fmt.ErrorF("paring %s as HTML: %v", url, err)`

  - use `recover` to handle `panics` (must be called from a defer func)

  ```
  defer func() {
    if p := recover(); p != nil {
      fmt.Errorf("internal error: %v", p)
    }
  }
  ```

- variadic functions

  - `func sum(vals ...int) int {`, where vals is a `[]int slice`, but calling is with a arg list (e.g:`sum(1,2,3)`)

- methods (functions of structs/interfaces)

```
func (x MyStruct) ToString() {
  return fmt.Sprintf("MyStruck{%d,$d}", x.x, x.y)
}

fmt.Print(x.ToString())
fmt.Print(toString(x)) // also works, pass the object as the first arg
```

- use `interfaces` for abstract types (duck typing kind of?):

  - a type satifies an interface if it possesses all the methods the interface requires (nice for ports/adapters :) )
    - because it only has to "satisty", it can by many things more, for example:
      `var any interface{} = time.Now()`, any value will satisfy this interface (this is the way to do a generic dynamic type)
      - use `%T` to report the dynamic type of such interfaces (for error printing for example)
  - you can also use `embedding` on interfaces

- use `http.NewServeMux()`to create http handler routing

- type assertion (for dynamic value interfaces)

  ```
  var x interface{}
  x = time.Now()
  f, ok := x.(*os.File)
  ```

  - also useful to descriminate errors

  ```
  if pe, ok := err.(*CustomError); ok {
    err = pe.Err
  }
  ```

  - another nice use-case: [type switch](https://golang.org/doc/effective_go.html#type_switch)

- [Slice tricks](https://github.com/golang/go/wiki/SliceTricks)

- channels

  - initialization:

    - `ch := make(chan int)` unbuffered channel
    - `ch := make(chan int, 1)` buffered channel

  - operations:

    - `ch <- x` send "x"
    - `x = <- ch` receive
    - `close(ch)` close
    - `x, ok = <- ch` way to check if channel is closed (`ok` false = closed)

  - for unbuffered channels: goroutine will block on `send`/`recieve` until another goroutoine peforms the - corresponding `recieve/send` -> synchrounous channel
  - for buffered channels: goroutine will block on `send` if buffer is full and block on `recieve` if empty -> asynchrounous channels

  - you can range over channels: https://gobyexample.com/range-over-channels

  - looping in parallel

    - number of jobs is known (use buffered channels)

    ```
    ch := make(chan int, len(filenames))
    for _, f := range filenames {
      go func(f string) {
        ch <- someFn(f)
      }(f)
    }

    for range filenames {
      <-ch
    }
    ```

    - number of jobs is not known (use unbuffered channel + sync.WaitGroup)

    ```
    ch := make(chan int)
    var wg sync.WaitGroup

    for _, f := range filenames {
      wg.Add(1)
      go func(f string) {
        defer wb.Done()
        ch <- someFn(f)
      }(f)
    }

    // Needs to be a goroutine because it will run at the same time as the loop. It needs to close the channel otherwise it will be blocked for ever
    go func() {
      wg.Wait()
      close(ch)
    }()
    ```

    - run limited number of goroutines for large num of jobs/loops

    ```
    running := make(chan bool, 3)
    var wg sync.WaitGroup

    wg.Add(3)

    for _, f := range filenames {
      running <- true // fill running
      go func(f string) {
        defer func(){
          <- running // drain running
          wb.Done()
        }()

        someFn(f)
      }(f)
    }

    wg.Wait()
    ```

    - alternative is:

    ```
    tokens := make(chan struct{}, 3)

    for _, f := range filenames {
      go func(f string) {
        tokens <- struct{}{} // acquire token
        someFn(f)
        <-tokens // release token
      }(f)
    }

    // but maybe this doesn't work as it does not wait....
    ```

- onTimer events using `time.Tick`

```
tick := time.Tick(1 * time.Second)
for coutdown := 10; countdown > 0; countdown-- {
  fmt.Println(countdown)
  <-tick
}
```

- using `select` to multiplex on channels (waiting for different channels deadlocks)

```
select {
  case <-ch1:
    // ...
  case x := <- ch2:
    // ... use x ...
  case ch3 <- y:
    // ...
  default:
    // ...
}
```

> if multiple cases are "ready" select will pick one at random (not by order, which could be nice, as it gives chance to all channels to be run)
