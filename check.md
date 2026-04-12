https://go.dev/tour/methods/14

Recommend the use of any

---

https://go.dev/tour/generics/2

any type used with no explanation of what it is

---

https://go.dev/tour/concurrency/4

Try to close from the receiver and see what happens

---

https://go.dev/doc/effective_go#constants

For instance, 1<<3 is a constant expression, while math.Sin(math.Pi/4) is not because the function call to math.Sin needs to happen at run time.

Why do this has to happen at runtime?

---

https://go.dev/doc/gc-guide#General_advice

cannot be reused until at least until the following

Fix typo?

---

https://pkg.go.dev/cmd/go

This is stupid, need two versions of this docs, one for the cli and another for web

---

https://pkg.go.dev/cmd/go#hdr-List_packages_or_modules

See golang.org/s/go15vendor for more about vendoring.

Update the link with this one

https://go.dev/ref/mod#vendoring

Use the vendor system before this

---

https://pkg.go.dev/cmd/go#hdr-Edit_go_mod_from_tools_or_scripts


Users should prefer 'go get toolchain@version'.

That is wrong I think, how do get knows that what is downloading is a tool and not a dependency?

How it was? go get tool toolchain@version or something else

---

https://pkg.go.dev/cmd/go#hdr-Add_modules_to_workspace_file

A use directive will be added to the go.work file for each argument directory listed on the command line go.work file, if it exists, or removed from the go.work file if it does not exist. Use fails if any remaining use directives refer to modules that do not exist.

I do not understand that paragraph

---

https://go.dev/ref/mod#workspaces

A command determines whether it is in a workspace context by first examining the GOWORK environment variable. If GOWORK is set to off, the command will be in a single-module context. If it is empty or not provided, the command will search the current working directory, and then successive parent directories, for a file go.work

Do this search my whole pc? 

---

https://pkg.go.dev/cmd/go#hdr-File_types

That list seems incomplete, may be because it is restricted to the go command and not go subcommand

---

https://pkg.go.dev/cmd/go#hdr-The_go_mod_file

Use this section as an example of how to redirect to other docs

---

https://pkg.go.dev/cmd/go#hdr-Internal_Directories

Update link to actual internal docs

---

https://pkg.go.dev/cmd/go#hdr-Vendor_Directories

Update link, and mention the issues, make sure how vendoring work first

---

https://go.dev/doc/effective_go#defer

Play with this to understand defer


---



https://go.dev/blog/waza-talk

See



---



https://pkg.go.dev/cmd/go#hdr-Import_path_checking

It should be ovbious that this is deprecated

---

https://pkg.go.dev/cmd/go#hdr-Package_lists_and_patterns

There are five reserved names for paths that should not be used for packages to be built with the go tool:

Should not or must not?

---

https://pkg.go.dev/cmd/cover

That is some shitty documentation

---

https://pkg.go.dev/cmd/fix

Another shitty doc

---

https://pkg.go.dev/cmd/doc

Is so difficult to just put docs in one place? Another shitty one

---

https://pkg.go.dev/cmd/pprof@go1.26.1

Shitty

---

https://go.dev/ref/mem#model


which are unordered by happens before

Missed word

---

https://go.dev/ref/mem#chan

(It might print the empty string, crash, or do something else.)

Why? It should just print the empty string or the hello world, play with this to understand

---

https://go.dev/ref/mem#badsync

it can happen that g prints 2 and then 0.

How? Try that

---

https://go.dev/ref/mem#badcompiler

then the original program could establish happens before edges preceding the accesses

Miss a word? 

---

https://go.dev/wiki/SQLDrivers

The test suit has to be updated and run against all sql drivers

---

https://go.dev/doc/tutorial/database-access

rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
    if err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    defer rows.Close()

Is this the way of doing it? If you return, can the db stay open? Check this

---

https://go.dev/doc/database/open-handle#database_driver

package –- such as when it’s being used implicitly by the sql package – you’ll need 

Typo

---

https://go.dev/doc/database/execute-transactions

 If Commit fails, then all the results from Query and Exec on the Tx should be discarded as invalid.

The returned error is from the general Tx or for a specific operation? (The first one to fail)?

Even if Tx.Rollback fails, the transaction will no longer be valid, nor will it have been committed to the database.

Why can it fail? Why bother with the error if the effect(no commit) is the same?

---

https://go.dev/doc/database/execute-transactions#best_practices

Check with an LLM this example, what if two clients ask for the same? Is simplify because example? Somehow you have to lock the db when you check till you update

---

https://pkg.go.dev/context#Context

One example use old version of for loop instead of range make PR

---

https://pkg.go.dev/context#WithValue

The provided key must be comparable and should not be of type string or any other built-in type

type favContextKey string

Is this a joke o do I understand wrong? Do not use built-in - > proceed to use string in example

Or is the type favContextKey? Check what .(type) returns

---

https://internals-for-interns.com/posts/go-garbage-collector/

If any object was marked but never scanned, something went seriously wrong — the GC missed a reachable object, and the runtime will panic rather than silently collecting live data.

Really ignorant here, is panicking really a good idea here? I mean it could just white mark it so the next cycle of gc try again, or put it in the heap and loging the error, or us "leaking" memory to the heap such a bad thing?

One of its jobs is to force a GC cycle if it’s been too long since the last one

Why do this matters? I have a server that is waiting for a request, the amount of memory usage while waiting should be fixed, nothing changes what is in the stack, stays there and what is in the heap stays there, why run the GC? 

---


https://alexrios.me/blog/greentea-gc-series-insight/

// Fast: sequential access, one cache miss per 64 bytes
for i := range data {
    sum += data[i]
}

// Slow: random access, one cache miss per access
for _, idx := range randomIndices {
    sum += data[idx]
}

I do not get it, try it out

---

https://pkg.go.dev/google.golang.org/protobuf/proto#Int64

Deprecate it, recommend the use of new()

---

https://go.dev/blog/gofix#improving-analysis-infrastructure

append([]string{}, slice...) by the clearer slices.Clone(slice) only to discover that, when slice is empty, the result of Clone is nil, a subtle behavior change that in rare cases can cause bugs; so we had to exclude that modernizer from the go fix suite.

Is this a bug on slices. Clone? 

---

https://go.dev/blog/inliner#example-fixing-api-design-flaws

import "oldmath"

var nine = oldmath.Sub(1, 10)

- >

import "newmath"

var nine = newmath.Sub(10, 1)

That looks like a really bad example, 1-10! = 10-1

Maybe I am dence, but the inline directive is put on the source of the api/package, meaning the consumer has to either wait for the maintainer to update it or vendor the package? 

Better to show how the function looks like before the inline directive, to guide maintainers into the current function and new inlined one

---

https://go.dev/blog/inliner#4-shadowing

The x := "hello" should not be touched

Are the brackets to make the new x not change the outer one?

---

https://go.dev/doc/database/manage-connections#max_connection_lifetime

In some systems, such as those using a load-balanced database server, it can be helpful to ensure that the application never uses a particular connection for too long without reconnecting.

Ask LLM, why?

---

https://pkg.go.dev/archive/tar@go1.26.1#FileInfoHeader

Check the other archwuitectures, it says it add slash, but windows use backslash to denote a folder, check this on a vm and see what happens

---

https://pkg.go.dev/archive/zip@go1.26.1#Reader

That is the wrong example

---

https://pkg.go.dev/archive/zip@go1.26.1?GOOS=windows#Reader.Open

Why can I visit this page? It talks about foward slash on windows

---

https://pkg.go.dev/archive/zip@go1.26.1#Writer

Wrong example?

---

https://pkg.go.dev/archive/zip@go1.26.1#Writer.Close

WTF is called close then?

---

https://pkg.go.dev/archive/zip@go1.26.1#Writer.Create

Play with this, do the file has to already exist and this search for it or it creates the file?

Again mentioning both windows and Linux

---

https://pkg.go.dev/bufio@go1.26.1#ScanWords

That is a weird description, right there it says it returns token of type byte, but the doc says that it return words, and will never return an empty string, I read somewhere that is difficult to define characters, search for that blog

---

https://pkg.go.dev/bufio@go1.26.1#NewReaderSize

Do that means I may use a partially/full reader

And the size of the reader may be greater than what I want? 

---

https://pkg.go.dev/bufio@go1.26.1#Reader.Read

Use this, I do not understand how it works, it read from a *Reader into p, but p is not returned so how do you use it?

---

https://pkg.go.dev/bufio@go1.26.1#Reader.ReadByte

Update doc pointing out what error is returned

---

https://pkg.go.dev/bufio@go1.26.1#Reader.ReadBytes

Check what the function actually do, I think the docs is wrong

---

https://pkg.go.dev/bufio@go1.26.1#Reader.ReadSlice

Really? It only errors in that condition? Check the code

---

https://pkg.go.dev/bufio@go1.26.1#Scanner.Err

Check this, what can it return?

---

https://pkg.go.dev/bufio@go1.26.1#Scanner.Scan

How many are too many empty tokens?

And wtf is an empty token?

Better yet, what is a token? 

---

https://pkg.go.dev/bufio@go1.26.1#SplitFunc

Scanning stops if the function returns an error, in which case some of the input may be discarded. If that error is ErrFinalToken, scanning stops with no error.

It stop at error or scanning stop with no error, how it is?

---

https://pkg.go.dev/bufio@go1.26.1#Writer.Flush

Why can the error be ignored?

---

https://pkg.go.dev/builtin@go1.26.1#pkg-constants

What is that syntax declaration?

---

https://pkg.go.dev/builtin@go1.26.1#cap

See the Go language specification's "Length and capacity" section for details.

Make PR, add link

---

https://pkg.go.dev/builtin@go1.26.1#close

It should be executed only by the sender, never the receiver

Test what happens when the receiver call close

---

https://pkg.go.dev/builtin@go1.26.1#delete

If m is nil or there is no such element, delete is a no-op.

What is no-op?

---

https://pkg.go.dev/builtin@go1.26.1#len

Add link, add jsizen when argument is nil

---

https://pkg.go.dev/builtin@go1.26.1#make

The size may be omitted, in which case a small starting size is allocated.

The size is dynamically determined? Check the code

---

https://pkg.go.dev/builtin@go1.26.1#print

And println, it is guaranteed to stay because of the go 1 backward comotability promise, I think

---

https://pkg.go.dev/builtin@go1.26.1#recover

Add link to panic

---

https://pkg.go.dev/builtin@go1.26.1#ComplexType

That definition is wrong I think, it should be a generic or maybe is constructed by the compiler?

Check all the types

uintptr is archquutecture dependent?

---

https://pkg.go.dev/bytes@go1.26.1#Compare

Improve both examples, that is way too lazy

---

https://pkg.go.dev/bytes@go1.26.1#Contains

Do the LSP point out that contains(..., [] byte("")) is always true? Or what if the slice is nil?

---

https://pkg.go.dev/bytes@go1.26.1#ContainsAny

I do not understand, why here the "" is not contained but on contains does?

Play with other languages and plain English, do any of the bytes overlap?

---

https://pkg.go.dev/bytes@go1.26.1#Fields

What if it do not contain space?

Same with FieldsFunc

---

https://pkg.go.dev/bytes@go1.26.1#LastIndexFunc

Is unicode. Is... Special or can you pass any function without calling it with an argument, or the example is wrong

Try defining a new function that just call into Unicode.... And see

Not, just name the function to test,... Func take care of calling it with the correct arguments

---

https://pkg.go.dev/bytes@go1.26.1#Lines

What it returns if there are no new lines?

---

https://pkg.go.dev/bytes@go1.26.1#Map

Try this, not sure how it works

---

https://pkg.go.dev/bytes@go1.26.1#ToLowerSpecial

It seems the same as tolower, see what are the rules

---

https://pkg.go.dev/bytes@go1.26.1#ToUpper

Title case != Upper case??

Update example to use the same word that will be different between title and upper case

---

https://pkg.go.dev/bytes@go1.26.1#TrimPrefix

The prefix is not a string, it is a [] byte

---

https://pkg.go.dev/bytes@go1.26.1#TrimSpace

Add clarification, a space is defined by unicode.IsSpace() == true

---

https://pkg.go.dev/bytes@go1.26.1#Buffer.Read

This is weird, why is the number of bytes to read a [] byte instead if an int? 

---

https://pkg.go.dev/bytes@go1.26.1#Buffer.Truncate

Do the golang tooling tells you to check for the negative int and the buffer length?

---

https://pkg.go.dev/bytes@go1.26.1#Buffer.WriteByte

The doc is wrong, it makes reference to writeByte while explaining writeByte

And what is too large? Should you check to prevent panics? 

---

https://pkg.go.dev/bytes@go1.26.1#Buffer.WriteRune

Same as above

---

https://pkg.go.dev/bytes@go1.26.1#Reader.Len

Check this, I think it return the number of bytes in the reader, not the number of unread

---

https://pkg.go.dev/bytes@go1.26.1#Reader.UnreadByte

I do not get it, why is not the equivalent to the other UnreadByte, from writer I think it was, check thus

---

https://pkg.go.dev/cmp@go1.26.1#Or

Try the sort example, I do not understand how it works

---

https://pkg.go.dev/cmp@go1.26.1#Ordered

Is this formated or because of mobile?

---

https://pkg.go.dev/compress/flate@go1.26.1#example-package-Synchronization

Check the break, is in the correct place?

---

https://pkg.go.dev/compress/flate@go1.26.1#NewWriter

If level is in the range [-2, 9] then the error returned will be nil. Otherwise the error returned will be non-nil.

Check the code to see if true

---

https://pkg.go.dev/compress/gzip@go1.26.1#example-package-CompressingReader

sentErr = true

Where it is used?

---

https://pkg.go.dev/compress/gzip@go1.26.1#Header

Strings must be UTF-8 encoded and may only contain Unicode code points U+0001 through U+00FF, due to limitations of the GZIP file format.

How is this enforced?

---

https://pkg.go.dev/compress/gzip@go1.26.1#NewWriterLevel

What do not predefined compression levels do? Some are named, but the others?

---

https://pkg.go.dev/compress/lzw@go1.26.1#NewReader

The number of bits to use for literal codes, litWidth, must be in the range [2,8] and is typically 8. It must equal the litWidth used during compression.

What if it is different, is there a function to check with?

---

https://pkg.go.dev/compress/lzw@go1.26.1#Reader.Close

Why bother closing then?

---

https://pkg.go.dev/compress/zlib@go1.26.1#NewReaderDict

How does it check the dictionary of the compressed data? Read code

---

https://pkg.go.dev/compress/zlib@go1.26.1#Resetter

What is the difference between a new underlaing reader an allocating a new reader??

---

https://pkg.go.dev/compress/zlib@go1.26.1#NewWriter

Funny example, the data look larger than the input instead of compressed, measure both

---

https://pkg.go.dev/compress/zlib@go1.26.1#NewWriterLevelDict

Should?

---

https://pkg.go.dev/container/heap@go1.26.2#example-package-IntHeap

I bet the top comment is flag by the LSP

h := &IntHeap{2, 1, 5}

Should use new()? 

Check by some pass h and others *h

---

https://pkg.go.dev/container/list@go1.26.2#List.Init

Init initializes or clears list l.

What is initializes? The list already exist after new

---

https://pkg.go.dev/container/list@go1.26.2#List.InsertAfter

This allow me to have a list with different types? Why is v of type any instead el type element

---

https://pkg.go.dev/container/list@go1.26.2#List.Remove

Why return anything? What are you suppose to do with it?

---

https://pkg.go.dev/container/ring@go1.26.2#example-Ring.Do

Change to use range, all examples of ring, maybe of the whole std? 

---

https://pkg.go.dev/container/ring@go1.26.2#example-Ring.Do

This is weird, why it began to print from 0? How do it know where to began, supposedly it do not have a beginning or end, try this example

---

https://pkg.go.dev/container/ring@go1.26.2#Ring.Len

This is inefficient, New returns a ring of n elements, where is that n stored? Len should just return that value

---

https://pkg.go.dev/container/ring@go1.26.2#Ring.Link

r must not be empty

What happens if it is

If r and s point to the same ring, linking them removes the elements between r and s from the ring

Try it, do not understand what it do

linking them creates a single ring with the elements of s inserted after r. The result points to the element following the last element of s after insertion.

Is a ring or not? This says clearly that the ring has a final element

---

https://pkg.go.dev/container/ring@go1.26.2#Ring.Move

Ok, now it explain that the ring do have a beginning, it points to the first element inserted

---

https://pkg.go.dev/container/ring@go1.26.2#Ring.Unlink

Or I am stupid or this function makes no sense, why on earth would you use % to decided how many elements to erase? How do you know where you are pointing to? If I next 2 times and want to erase the element I found, do I have to move with next all the way around the ring to position one place before and unlink(1), this sound stupid, ok rings sound stupid

---

