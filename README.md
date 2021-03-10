[![codecov](https://codecov.io/gh/tufin/oasdiff/branch/master/graph/badge.svg?token=Y8BM6X77JY)](https://codecov.io/gh/tufin/oasdiff)
[![CircleCI](https://circleci.com/gh/Tufin/oasdiff.svg?style=svg)](https://circleci.com/gh/Tufin/oasdiff)
[![Go Report Card](https://goreportcard.com/badge/github.com/tufin/oasdiff)](https://goreportcard.com/report/github.com/tufin/oasdiff)
[![GoDoc](https://godoc.org/github.com/tufin/oasdiff?status.svg)](https://godoc.org/github.com/tufin/oasdiff)

# OpenAPI Diff Go Module
This [Go](https://golang.org) module provides a diff utility for [OpenAPI Spec 3](https://swagger.io/specification/).

## Unique features vs. other diff tools
- go module
- deep diff into paths, parameters, request bodies, responses, schemas, enums, callbacks etc.

## Build
```
git clone https://github.com/Tufin/oasdiff.git
cd oasdiff
go build
```

## Running from the command-line
```
./oasdiff -base data/openapi-test1.yaml -revision data/openapi-test2.yaml
```

## Help
```
./oasdiff --help
```

## Output example

```yaml
spec:
  paths:
    deleted:
      - /api/{domain}/{project}/install-command
      - /register
      - /subscribe
    modified:
      /api/{domain}/{project}/badges/security-score:
        operations:
          added:
            - POST
          modified:
            GET:
              tags:
                deleted:
                  - security
              parameters:
                deleted:
                  cookie:
                    - test
                  header:
                    - user
                    - X-Auth-Name
                modified:
                  path:
                    domain:
                      schema:
                        type:
                          from: string
                          to: integer
                        format:
                          from: hyphen-separated list
                          to: non-negative integer
                        description:
                          from: Hyphen-separated list of lowercase string
                          to: Non-negative integers (including zero)
                        min:
                          from: null
                          to: 7
                        pattern:
                          from: ^(?:([a-z]+-)*([a-z]+)?)$
                          to: ^(?:\d+)$
                  query:
                    filter:
                      content:
                        schema:
                          properties:
                            modified:
                              color:
                                type:
                                  from: string
                                  to: number
                    image:
                      explode:
                        from: null
                        to: true
                      schema:
                        description:
                          from: alphanumeric
                          to: alphanumeric with underscore, dash, period, slash and colon
                    token:
                      schema:
                        anyOf: true
                        type:
                          from: string
                          to: ""
                        format:
                          from: uuid
                          to: ""
                        description:
                          from: RFC 4122 UUID
                          to: ""
                        pattern:
                          from: ^(?:[0-9a-f]{8}-[0-9a-f]{4}-[0-5][0-9a-f]{3}-[089ab][0-9a-f]{3}-[0-9a-f]{12})$
                          to: ""
              responses:
                added:
                  - default
                deleted:
                  - "200"
                  - "201"
        parameters:
          deleted:
            path:
              - domain
  security:
    deleted:
      - bearerAuth
  servers:
    deleted:
      - tufin.com
  tags:
    deleted:
      - security
      - reuven
  componentsdiff:
    schemas:
      deleted:
        - network-policies
        - rules
    parameters:
      deleted:
        header:
          - network-policies
    headers:
      deleted:
        - new
        - test
        - testc
    requestBodies:
      deleted:
        - reuven
    responses:
      deleted:
        - OK
    securitySchemes:
      deleted:
        - OAuth
        - bearerAuth
summary:
  diff: true
  components:
    headers:
      deleted: 3
    parameters:
      deleted: 1
    paths:
      deleted: 3
      modified: 1
    requestBodies:
      deleted: 1
    responses:
      deleted: 1
    schemas:
      deleted: 2
    security:
      deleted: 1
    securitySchemes:
      deleted: 2
    servers:
      deleted: 1
    tags:
      deleted: 2
```

## Embedding into your Go program
```
diff.Get(&diff.Config{}, spec1, spec2)
```
See full example: [main.go](main.go)

## Notes
1. oasdiff expects [OpenAPI References](https://swagger.io/docs/specification/using-ref/) to be resolved.  
You can resolve refs using [this function](https://pkg.go.dev/github.com/getkin/kin-openapi/openapi3#SwaggerLoader.ResolveRefsIn) from the openapi3 package.

2. oasdiff ignores changes to [Examples](https://swagger.io/specification/#example-object) and [Extensions](https://swagger.io/specification/#specification-extensions) by default. You can change this behavior through [configuration](diff/config.go).

## Work in progress
While most objects of OpenAPI Spec are already supported by this diff tool, some are still missing, notably: Examples, ExternalDocs, Links, Variables and a couple more.  
Pull requests are welcome!

