[[String types](https://go.dev/ref/spec#String_types)


Strings are immutable: once created, it is impossible to change the contents of 
a string.

String types](https://go.dev/ref/spec#String_types)


Strings are immutable: once created, it is impossible to change the contents of 
a string.

`strType.go`

When you change a variable of type string you are creating a new string and pointing
to it, normally strings go into the Stack(read only) but there are situations where
you can put it on the Heap(read/write) then you can modify the underlying bytes
(not recommended).

---

Why did I put this here?

[Rune literals](https://go.dev/ref/spec#Rune_literals)

A rune literal represents a rune constant, an integer value identifying a Unicode 
code point. A rune literal is expressed as one or more characters enclosed in single 
quotes, as in 'x' or '\n'. Within the quotes, any character may appear except newline 
and unescaped single quote. A single quoted character represents the Unicode value 
of the character itself, while multi-character sequences beginning with a backslash 
encode values in various formats.

The simplest form represents the single character within the quotes; since Go source 
text is Unicode characters encoded in UTF-8, multiple UTF-8-encoded bytes may represent 
a single integer value. For instance, the literal 'a' holds a single byte representing 
a literal a, Unicode U+0061, value 0x61, while 'ä' holds two bytes (0xc3 0xa4) 
representing a literal a-dieresis, U+00E4, value 0xe4.


[String literals](https://go.dev/ref/spec#String_literals)

If the source code represents a character as two code points, such as a combining
form involving an accent and a letter, the result will be an error if placed in 
a rune literal (it is not a single code point), and will appear as two code points 
if placed in a string literal.

Explain this two things

---

[Constants](https://go.dev/ref/spec#Constants)

Implementation restriction: Although numeric constants have arbitrary precision 
in the language, a compiler may implement them using an internal representation 
with limited precision. That said, every implementation must:

Round to the nearest representable constant if unable to represent a floating-point 
or complex constant due to limits on precision.

When the "nearest" is both upper and lower number, what di the implementers do? 

And on "common" cases do they round up or down conditionally or always one way?

check this for the answer

https://go.dev/ref/spec#Representability

---

[Variables](https://go.dev/ref/spec#Variables)

var x interface{}  // x is nil and has static type interface{}

https://leapcell.medium.com/why-gos-any-and-interface-are-the-same-but-not-quite-ba6d0a54aade

interface{}: Represents "unconstrained dynamic types" (think JSON parsing, reflection). Emphasizes dynamism, requires type assertions.
any: Designed specifically for generics. Marks unconstrained type parameters, emphasizing type-safe generality—once a type is set, it stays consistent.

Is static or not? Should the example use any instead of interface{}?

---

[Numeric types](https://go.dev/ref/spec#Numeric_types)

Explicit conversions are required when different numeric types are mixed in an 
expression or assignment.

Can I operate uint8 and byte, int32 and rune? 

---

[String types](https://go.dev/ref/spec#String_types)

Strings are immutable: once created, it is impossible to change the contents of a string.

The length is a compile-time constant if the string is a constant.

First it says strings are constants? Then add a conditional if they are constants, 
strings are not constant if declared as var, actually test this out, what part 
of a string is immutable?

---

[Array type](https://go.dev/ref/spec#Array_types)

An array type T may not have an element of type T, or of a type containing T as 
a component, directly or indirectly, if those containing types are only array or struct types.

I do not get this, try to understand.

Check what happens with the any type

---

[Slice types](https://go.dev/ref/spec#Slice_types)

The value of an uninitialized slice is nil.

Check this

---

[https://go.dev/ref/spec#Representability](https://go.dev/ref/spec#Representability)

Rounding uses IEEE 754 round-to-even rules but with an IEEE negative zero further 
simplified to an unsigned zero. Note that constant values never result in an IEEE 
negative zero, NaN, or infinity.

----

[https://go.dev/ref/spec#Declarations_and_scope](https://go.dev/ref/spec#Declarations_and_scope)

An identifier declared in a block may be redeclared in an inner block. While the 
identifier of the inner declaration is in scope, it denotes the entity declared by the inner declaration.

How do the compiler know which are you referring to? Check this

---

[Predeclared identifiers](https://go.dev/ref/spec#Predeclared_identifiers)

Check the identifiers, there are some I do not recognize, and check print, can I use it without fmt?

---

[Satisfying a type constraint](https://go.dev/ref/spec#Satisfying_a_type_constraint)

Reread this I do not understand

---

[Variable declarations](https://go.dev/ref/spec#Variable_declarations)

Implementation restriction: A compiler may make it illegal to declare a variable inside a function body if the variable is never used.

Can I check this on the playground? Why the "may", why make it optional?

If the function declaration specifies type parameters, the function name denotes a generic function. A generic function must be instantiated before it can be called or used as a value.

Can I declare and not use a global variable? Why?

---

[Short variable declarations](https://go.dev/ref/spec#Short_variable_declarations)

In some contexts such as the initializers for "if", "for", or "switch" statements, they can be used to declare local temporary variables.

I thought all variables/const where temporary respective to the scope, once out of scope it get cleaned by the GC
Check this out

---

[Function declarations](https://go.dev/ref/spec#Function_declarations)

A function declaration without type parameters may omit the body. Such a declaration provides the signature for a function implemented outside Go, such as an assembly routine.

I do not understand, check how this work, can you call upon bynaries of other language as if they are functions? 

But how do I indicate to the compiler where to get the function? Can I call binaries like this?

---

[](https://go.dev/ref/spec#Composite_literals)

A parsing ambiguity arises when a composite literal using the TypeName form of the LiteralType appears as an operand between the keyword and the opening brace of the block of an "if", "for", or "switch" statement, and the composite literal is not enclosed in parentheses, square brackets, or curly braces. In this rare case, the opening brace of the literal is erroneously parsed as the one introducing the block of statements. To resolve the ambiguity, the composite literal must appear within parentheses.

Try the given example to really understand what the ambiguity is

---

[](https://go.dev/ref/spec#Selectors)

If x is of pointer type and has the value nil and x.f denotes a struct field, assigning to or evaluating x.f causes a run-time panic.

Do the lsp cach this?

---

[](https://go.dev/ref/spec#Index_expressions)

if x is out of range at run time, a run-time panic occurs

Do LSP catch this?

if the map is nil or does not contain such an entry, a[x] is the zero value for the element type of M

That sounds bad, how do I know that the key x do not exist or exist and is of zero value?

---

[](https://go.dev/ref/spec#Calls)


Calling a nil function value causes a run-time panic.

Lsp? 

---

[](https://go.dev/ref/spec#Integer_overflow)

Why is this ok? Should i as user of the language dealth with this case if it happens?

---

[](https://go.dev/ref/spec#Comparison_operators)

---

https://go.dev/ref/spec#Constants

Not map or byte, slice constants?
Because they are dynamic length? 

---

https://go.dev/ref/spec#Comparison_operators

Interface types that are not type parameters are comparable. 

I do not understand

---

https://go.dev/ref/spec#Receive_operator

<-strobe  // wait until clock pulse and discard received value

Where is strobe defined?

---

https://go.dev/ref/spec#Conversions

https://go.dev/ref/spec#Conversions_from_slice_to_array_or_array_pointer

Check array to slice 

, if the length of the slice is less than the length of the array, a run-time panic occurs.

Check this with different lenght array/slice combinations

---

https://go.dev/ref/spec#Assignment_statements

Each left-hand side operand must be addressable, a map index expression, or (for = assignments only) the blank identifier. Operands may be parenthesized

What do the parenthesis statement means? 

---

The number of Ps is determined by GOMAXPROCS, which defaults to the number of CPU cores detected earlier. So on an 8-core machine, you get 8 Ps—meaning up to 8 goroutines can run truly in parallel at any given moment.

What are threaths then? And what are light threaths? 

---

https://go.dev/ref/spec#For_range

If the range expression is a channel or integer, at most one iteration variable is permitted; otherwise there may be up to two. If the last iteration variable is the blank identifier, the range clause is equivalent to the same clause without that identifier.

Check if i can do for i, n := range(10)

---

https://go.dev/ref/spec#Defer_statements

Check what happens with defer after the return

---

https://go.dev/ref/spec#Clear

After manipulating a secret you apply clear to completely erase the sensitive data from memory?

---

Check typed vs untyped on golang

---

https://go.dev/ref/spec#Deletion_of_map_elements

If the map m is nil or the element m[k] does not exist, delete is a no-op.

What is no-op?

---

https://go.dev/ref/spec#Handling_panics

 To ensure this, calling panic with a nil interface value (or an untyped nil) causes a run-time panic.

Calling panic(nil) causes a panic? How do this work

I do not think I understand how recover works is just caching the panic to handle it?

Read again

---

https://go.dev/ref/spec#Import_declarations

Consider a compiled a package containing the package clause package math,

Typo "a" twice 

Should the compiled word be there? I do not import compiled packages I import source code, check

---

https://go.dev/ref/spec#The_zero_value

Do all variables/constants get into my binary? Or only the used ones?

---

https://go.dev/ref/spec#Run_time_panics

Read handling panics, this is kind of pointless section

---

https://go.dev/ref/spec#Package_unsafe

Since Go strings are immutable, the bytes passed to String must not be modified afterwards. [Go 1.20]

Must not != cannot? What happen if I do it? Test

