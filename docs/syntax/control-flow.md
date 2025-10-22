# Control flow

Zirric offers both expression and statement forms for its `if` and `for` constructs.
Expressions produce a value, while statements are used when only side effects are
required.

## `if`

The expression form evaluates to one of its branches and can be embedded in
larger expressions:

```zirric
let label = if answer == 42 { "yes" } else { "no" }
```

Each branch of an `if` expression contains exactly one expression, and it cannot
include `return` statements.

The statement form chooses a branch for its effects and produces no value:

```zirric
if answer == 42 {
    print("yes")
} else {
    print("no")
}
```

`if` statements may hold multiple statements in their branches, including
`return`, and the branches may even be empty.

## `for`

The `for` expression iterates over a sequence and gathers the values produced by
its body:

```zirric
let doubled = for n <- [1, 2, 3] { n * 2 }
```

`for` expressions behave like a combined filter and map operation. The values
produced by the body are collected into an array. A `continue` skips a value, and
`break` ends the result early. The body may bind a single variable and consists
of exactly one expression; `return` statements are not allowed.

The statement form simply walks a collection for its effects:

```zirric
for n <- [1, 2, 3] {
    print(n)
}
```

`for` statements may bind multiple variables, contain multiple statements
including `return`, or be completely empty.

An empty loop written as `for { }` runs indefinitely and is useful for
processes that wait for external events. It only terminates when a `break`
statement is executed.

Both variants share the same syntax for conditions and iterators. Expression
forms yield values, whereas statements have no result and are used purely for
side effects.
