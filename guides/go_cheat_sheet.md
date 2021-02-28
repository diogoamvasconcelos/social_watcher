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
  - %T: type of any type
  - %%: literal percent sign (no operand)
  - %[1]{some letter}: use the first operand (e.g: `fmt.Printf("%d, %[1]x", 15) -> "15 F"`)

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

  - map: dynamic size
    - can be used to emulate a `set` (unique keys)
