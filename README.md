# GO-TEMPL-HTMX

This project is a simple web app made in go using [templ](https://github.com/a-h/templ), [sqlc](https://sqlc.dev/), [htmx](https://htmx.org/) and [tailwind](https://tailwindcss.com/).

## Running

First you need to generate the templ and sqlc go files and tailwind styles using the command (note that you need to install the [templ](https://templ.guide/quick-start/installation) and [sqlc](https://docs.sqlc.dev/en/stable/overview/install.html) commands to generate these files)

```bash
make generate
```
After generating, you can build the project with the following command

```bash
make build
```

Lastly, you can run it with

```bash
make run
```

## Dev mode

You can also use [air](https://github.com/cosmtrek/air) to run the project in dev mode (it will generate the required files and build it automatically)

```bash
make dev
```

## Database

This project uses a PostgreSQL database and you can configure the connection in the PG_CONNECTION_STRING environment variable. To apply the migrations, run the following command:

```bash
make migrate
```
