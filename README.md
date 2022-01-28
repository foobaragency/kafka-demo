# Kafka demo

This repository contains the demo apps using during the presentation at our knowledge-friday event. (See: [this issue](https://github.com/foobaragency/knowledge-friday/issues/27))

### Structure
```bash
    apps
        |- clients # mock client apps consuming both APIs
        |- products # Products API written in Go
        |- orders # Order APIs written in TypeScript
    scripts # useful scripts ro run all apps at once
```

### How to run this

1. Create the `.env` files based on `.env.template` files provided in each app.
2. Install the dependencies as following:

    + Products API: `go mod tidy`
    + Orders API: `yarn install`

3. You can run each API individually or all at once using the command `make run`
4. Have fun!