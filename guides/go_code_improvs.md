## Based on [Effective GO](https://golang.org/doc/effective_go.html)

- use the `if` intialization method for error checking

```
if err := file.Chmod(0664); err != nil {
  log.Print(err)
  return err
}
```

- use `switch` instead of [if-else-if-else chains](https://golang.org/doc/effective_go.html#switch)

  - the `switch` expression doesn't need to be a constant: it is evaluated at run time

  - Note: there is no automatic fall through, but you can use `comma` for the same behaviour:

    ```
    func shouldEscape(c byte) bool {
      switch c {
      case ' ', '?', '&', '=', '#', '+', '%':
        return true
      }
      return false
    }
    ```

  - Less need to use `break` because of the no-fall through. Also, `break` and `continue` can take an extra "label" for nested for/switch cases, e.g: `break Loop`
    xx

- Dynamic type (i.e: `interface {}`) run-time type [check this](https://golang.org/doc/effective_go.html#interface_conversions)

  - Use a [type switch](https://golang.org/doc/effective_go.html#type_switch)

  - Cast using the `type assertion`
    - e.g: `str, ok := value.(string)`

- Use [named return paramters](https://golang.org/doc/effective_go.html#named-results) more often (can make code more readable)

  - Remember that if not assigned, will have their default values

- use [defer](https://golang.org/doc/effective_go.html#defer) to schedule execution to run right before functioin returns

  - used to release mutex's or closing files
  - deferred functions are executed in LIFO order (stack)
  - arguments to the functio called in `defer` are evaluated when read, so you can exploit []()

- use [slices](https://golang.org/doc/effective_go.html#slices) and take advantage of the range operators (e.g: buf[1:])

- use the second `comma ok` from a map get values to check if map constains:

  ```
  _, present := timeZone[tz]
  ```

- use append almost like the js spread

  - https://golang.org/doc/effective_go.html#append

- Use `const` to create enumlike data

  - https://golang.org/doc/effective_go.html#constants

- use the `init()` function of a package, as it can be quite convient (it runs after all package initialization)

  - https://golang.org/doc/effective_go.html#init

- user nested structs: https://medium.com/@xcoulon/nested-structs-in-golang-2c750403a007

- write functions on named-types (https://golang.org/doc/effective_go.html#pointers_vs_values)

  - ex. myType.ToString()

- panic/recover pattern (throw/catch kind of)

  - only works for goroutines

  ```
    func server(workChan <-chan *Work) {
        for work := range workChan {
            go safelyDo(work)
        }
    }

    func safelyDo(work *Work) {
        defer func() {
            if err := recover(); err != nil {
                log.Println("work failed:", err)
            }
        }()
        do(work)
  }
  ```

- user `defer` (so powerful!!)

  - deferred functions can modify named return values (crazy!!)

- the [flag pckage](https://golang.org/pkg/flag/) is very nice to do CLIs :)

- use [interfaces](https://gobyexample.com/interfaces) for ports/adapters
