## Templo

Templo is a basic template i've built to develop simple yet powerful Go web applications, this template is also used by my team at [Wawandco](https://github.com/wawandco) as a starting point for most of the web apps we build with Go.

#### Templo uses

- [Negroni](https://github.com/codegangsta/negroni) to chain our middlewares.
- [Pat](https://github.com/bmizerany/pat) as our mux.
- [Jet](https://github.com/CloudyKit/jet) to render our html templates.
- [Godotenv](https://github.com/joho/godotenv) to load env configuration.

#### App Configuration



------------------

#### Development Tools

There is a set of tools i use to develop in Go that may be helpful in combination with Tempo, here is a list i'll keep updated:

- [Gin](https://github.com/codegangsta/gin)
- [Godep](https://github.com/tools/godep)
- [Transporter](https://github.com/wawandco/transporter)
- [Testify/Assert](github.com/stretchr/testify/assert)

#### Running on development

Templo comes with a `.env_sample` file, it has some environment variables needed for the app, please use this as a reference to create your own `.env` file, which will be ignored and should not be pushed to the git repository.

The `.env` file will be loaded by the app every time the app runs, feel free to put there things that are environment specific.

#### Running Tests

In order to run the tests please use
```go
$ go test $(go list ./... | grep -v /vendor/)
```
