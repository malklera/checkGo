# Ideas to check or implement

## New flag parsing library

Why use multiple functions to parse inputs?

It is always a string that is converted to some other type.

Try something like this:

```go
// f is the command the flag is related to
// p is the variable where to save the input, take `any` then parse all std types
func (f *FlagSet) VarP(p *any, name string, shorthand string, value any, usage string)
```

```go
// f is the command the flag is related to
// p is the variable where to save the input, take `any`
// this one will deal with third party packages types, you have to pass the parsing
// function to be used to convert the input `string` to the desired type, have to
// satisfy the same interface as the ones used for the std, i do not think this
// should be difficult, every package should have the `string`->`customType` function
func (f *FlagSet) VarCustomP(p *any, parser *func(), name string, shorthand string, value any, usage string)
```

## When generating docs, list all possible errors

Is there a "simple" way to extract all possible errors returned by all functions
in a package?

Working with `github.com/lib/pq` i knew if i try to add a new row with the same
value in a unique column it will return error, but the function can return other
types of errors, would have been nice to have a list of possible errors each function
can return so i can dealt with just a general error or manage a particular one.

## go docs -http should allow to output the files without serving them in a port
