# gocli-demo

## Directory structure
- cmd: main applications, directories inside should match executable
- internal: private package code
- pkg: public package code
- vendor: managed by dep tool
- api: openapi specs, schema and proto files
- web: web app components, templates etc
- config: configuration files
- build: package -> build configs i.e. docker, rpm, ci -> for travis circle etc
- test: external tests and data
- examples: examples for applicaton and public packages

## Inspired by:
- [golang-standards/project-layout](https://github.com/golang-standards/project-layout)

## Requirements

- Export the API key with `export API_KEY=<key>`

### Running the Project
- To run the project
  ```
  make build
  ./bin/cli --help
  ```

## About the CLI
- Welcome Screen for API
  ```
  A simple CLI application for fetching data from a REST API

  Usage:
    gocli-demo [flags]
    gocli-demo [command]

  Available Commands:
    completion  Generate the autocompletion script for the specified shell
    explain     Explain word meaning from Meriam Webster dictionary REST API
    help        Help about any command

  Flags:
    -h, --help   help for gocli-demo
  ```

- Get specific word or list of words separated by spaces
  ```
  Explain word meaning from Meriam Webster dictionary REST API

  Usage:
    gocli-demo explain [flags]

  Flags:
    -h, --help              help for explain
        --wordForm string   Specify 'noun/verb/adjective' form for the word meaning (not implemented)
  ```

- Example Output:
  ```
  $ ./bin/cli explain exercise program

  ˈek-sər-ˌsīz noun {bc}the act of bringing into play or realizing in action {bc}{sx|use||} 
  ˈprō-ˌgram noun {bc}a public notice
  ```

## Scope of Improvement

- Dockerise the entire Dev setup to not have hard depencies on system with VSCode devcontainer extension
- Have more specific commands where the user can actually signify their intent, example, return only meaning, or a specific context if it exists from the json response and only return something default as fallback if nothing is found
- We can cache the API calls, so as to not keep searching for the same queries over and again locally into a SQLlite database with a TTL for these entries
- Add Tests
- Add a Package for parsing the results, and marshalling the JSON Response into a `struct` for better operational capabilities. Atm, these are all contained in the `main` pkg


