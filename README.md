# GO-TEMPL-HTMX

This project is a simple web app made in go using [echo](https://github.com/labstack/echo), [templ](https://github.com/a-h/templ), [htmx](https://htmx.org/) and [tailwind](https://tailwindcss.com/).

## Running

First you need to generate the templ go files and tailwind styles using the command

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