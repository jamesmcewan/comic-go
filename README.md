# Comicvine list comics from a publisher - gopher's version

_This is basically the same programme as the one I wrote in Rust but for fun this one is in Go._

For some reason the comicvine website does not have a way to list all the comics from a publisher. This is a simple cli tool that does that.

Requires a comicvine api key. You can get one by creating an account at https://comicvine.gamespot.com/api/ then add the `COMICVINE_API_KEY` to a `.env` file in the root of the project.

## Usage

Get the publisher id from the comicvine website. For example, the publisher id for Titan Books is `4010-4212`.

```bash
#Remember  the quotes
go run ./cmd "4010-4212"

# To save the output to a file
go run ./cmd "4010-4212" > titan-books.txt
```

