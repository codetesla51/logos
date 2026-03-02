<p align="center">
  <img src="doc/logo.png" width="120" alt="Logos">
</p>

<h1 align="center">Logos</h1>

<p align="center">
  Bash, but you can actually read it.
</p>

<p align="center">
  <a href="https://logos-lang.dev/docs">Docs</a> ·
  <a href="https://logos-lang.dev/docs/install">Install</a> ·
  <a href="https://logos-lang.dev/docs/examples">Examples</a>
</p>

---

Logos is a scripting language for people who got tired of writing bash. It has readable syntax, proper error handling, and compiles to a single binary.

```logos
let res = httpGet("https://api.example.com/data")

if res.ok {
    let data = parseJson(res.value.body)
    print(data["message"])
} else {
    print("Error: " + res.error)
}
```

## Features

- Readable C-like syntax
- Built-in HTTP, JSON, and file I/O
- Result-based error handling (no exceptions)
- Concurrency with `spawn`
- Compiles to standalone binary with `lgs build`
- Embeddable in Go

## Install

```sh
curl -fsSL https://install.logos-lang.dev | sh
```

## Quick Start

```sh
lgs script.lgs        # run a script
lgs                   # start REPL
lgs fmt script.lgs    # format code
lgs build script.lgs  # compile to binary
```

## Docs

See [logos-lang.dev/docs](https://logos-lang.dev/docs) for full documentation.

## License

MIT
