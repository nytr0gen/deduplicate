# Deduplicate

Parsing through long lists of endpoints, full of params, became a hassle? You want to get out? Say no more, Deduplicate is the solution. With less than 200 lines of code, written in Go, it will remove any duplicates from your lists and make visual parsing a bliss.

I made this so that I can go faster through [wayback urls](https://github.com/tomnomnom/waybackurls) output.

## How does it work?

Glad you asked. It formats an url to `%host%path%query_names`, checks if the url was already found, if not, adds it into the map.

## Install

```bash
go get github.com/nytr0gen/deduplicate
```

## Usage

```bash
echo domain | waybackurls | deduplicate --hide-useless --sort
```
