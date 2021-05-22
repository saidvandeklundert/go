`string`: a read-only slice of immutable bytes. Strings are utf-8 encoded by default. The zero value is an empty string

In the context of strings, a `rune` is a unicode code point. The `rune` is a alias for `int32`, which can represent a unicode code point. When referrring to characters, use `rune` to clarify intent.

More on strings [here](https://github.com/saidvandeklundert/go/blob/main/atrings.md).