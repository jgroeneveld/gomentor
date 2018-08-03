# losmentor

A version of mentor in go.

## Libraries used

- `github.com/jgroeneveld/schema` for JSON testing
- `github.com/jgroeneveld/trial/assert` for Lightweight Assertions
- `github.com/labstack/echo` as Lightweight web Framework

## Decisions

### trial/assert vs t.Fatal

I make use of my own library `trial/assert` for assertions. Its easy to use the build in error reporting but
mostly in tests values are compared and having a library that provides clear error messages and a super fast api
improves the speed of development.
Trial also makes it very obvious when the issue lies within the types (int, int64) when comparing. Which is a common mistake.

### echo vs standardlib

I decided to use echo instead of the standardlib to get quickly up to speed.
echo gives routing and the structure to quickly return json and errors.
Middlewares like logging are build in.
Its easy to build this things on your own but why invent the wheel again?
Of course its a dependency and if echo is not supported anymore it might be an issue, but as its open source, one could
start forking and maintaining at any point.