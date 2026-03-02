<p align="center">
  <img src="doc/logo.png" width="120" alt="Logos">
</p>

<h1 align="center">Logos</h1>

Logos (from Greek, meaning word, reason, or logic) is a scripting language designed to be read and understood, not deciphered.

It's what you reach for when Bash becomes unreadable after line three. A CLI-first language built for the kind of work Bash handles poorly: readable logic, proper error handling, and code you can come back to a week later and still understand.

```logos
let res = httpGet("https://api.example.com/data")

if res.ok {
    let data = parseJson(res.value.body)
    print(data["message"])
} else {
    print("Error: " + res.error)
}
```

## Highlights

- Readable, C-like syntax that favors clarity
- Result-based error handling: functions return `{ok, value, error}` instead of throwing
- Built-in HTTP, JSON, and file I/O primitives
- Lightweight concurrency with `spawn` blocks
- Embeddable in Go via `github.com/codetesla51/logos/logos`
- Scripts use the `.lgs` extension
- Can be compiled into a standalone binary using `lgs build`

## Install

Recommended (installer hosted in this repo):

```sh
curl -fsSL https://raw.githubusercontent.com/codetesla51/logos/main/install.sh | sh
```

This downloads the installer from the repository and picks the correct release asset for your OS/arch. Works on Linux and macOS.

## Quick Start

```sh
lgs script.lgs        # run a script
lgs                   # start REPL
lgs fmt script.lgs    # format code
lgs build script.lgs  # compile to a standalone binary
```

## Docs

Full documentation and the landing page are available at https://logos-lang.vercel.app/ (docs at https://logos-lang.vercel.app/docs).

## Contributing

Contributions are welcome. If you'd like to add examples, standard library utilities, or help with releases, open an issue or a PR on GitHub: https://github.com/codetesla51/logos

## License

MIT
