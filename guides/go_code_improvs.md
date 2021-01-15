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

- Dynamic type (i.e: `interface {}`) run-time type check

  - Use a [type switch](https://golang.org/doc/effective_go.html#type_switch)

- Use [named return paramters](https://golang.org/doc/effective_go.html#named-results) more often (can make code more readable)

  - Remember that if not assigned, will have their default values

- use [defer](https://golang.org/doc/effective_go.html#defer) to schedule execution to run right before functioin returns

  - used to release mutex's or closing files
  - deferred functions are executed in LIFO order (stack)
  - arguments to the functio called in `defer` are evaluated when read, so you can exploit []()

- TODO: continue from [data](https://golang.org/doc/effective_go.html#data)
