# gomod-json

Read go.mod file contents and dump them as JSON.

## Install

```sh
make install
```

## Why

For example, tagging your Docker images based on some dependency:

```sh
# Install tools:
go install github.com/mikefarah/yq/v4@v4.34.1
# Setup build directory:
mkdir -p /tmp/jsonnet-playground && cd /tmp/jsonnet-playground
# Clone some program to build a Docker image from:
git clone https://github.com/radekg/jsonnet-playground.git .
# Find the direct jsonnet dependency, there should be only one:
JSONNET_VERSION=$(gomod-json --path=./go.mod | yq '.Require[] | select(.Mod.Path == "github.com/google/go-jsonnet" and .Indirect == false) | .Mod.Version')
# Use it to create a Docker image:
docker build -t docker.io/radekg/jsonnet-playground:${JSONNET_VERSION}
docker tag docker.io/radekg/jsonnet-playground:${JSONNET_VERSION} docker.io/radekg/jsonnet-playground:latest
# Cleanup
cd - && rm -rf /tmp/jsonnet-playground
```

## Example

Run on itself:

```sh
gomod-json --path=./go.mod | yq '.' -o json -P
```

```json
{
  "Module": {
    "Mod": {
      "Path": "github.com/radekg/gomod-json"
    },
    "Deprecated": "",
    "Syntax": {
      "Before": null,
      "Suffix": null,
      "After": null,
      "Start": {
        "Line": 1,
        "LineRune": 1,
        "Byte": 0
      },
      "Token": [
        "module",
        "github.com/radekg/gomod-json"
      ],
      "InBlock": false,
      "End": {
        "Line": 1,
        "LineRune": 36,
        "Byte": 35
      }
    }
  },
  "Go": {
    "Version": "1.20",
    "Syntax": {
      "Before": null,
      "Suffix": null,
      "After": null,
      "Start": {
        "Line": 3,
        "LineRune": 1,
        "Byte": 37
      },
      "Token": [
        "go",
        "1.20"
      ],
      "InBlock": false,
      "End": {
        "Line": 3,
        "LineRune": 8,
        "Byte": 44
      }
    }
  },
  "Toolchain": null,
  "Require": [
    {
      "Mod": {
        "Path": "github.com/mitchellh/go-homedir",
        "Version": "v1.1.0"
      },
      "Indirect": false,
      "Syntax": {
        "Before": null,
        "Suffix": null,
        "After": null,
        "Start": {
          "Line": 6,
          "LineRune": 2,
          "Byte": 57
        },
        "Token": [
          "github.com/mitchellh/go-homedir",
          "v1.1.0"
        ],
        "InBlock": true,
        "End": {
          "Line": 6,
          "LineRune": 40,
          "Byte": 95
        }
      }
    },
    {
      "Mod": {
        "Path": "github.com/pkg/errors",
        "Version": "v0.9.1"
      },
      "Indirect": false,
      "Syntax": {
        "Before": null,
        "Suffix": null,
        "After": null,
        "Start": {
          "Line": 7,
          "LineRune": 2,
          "Byte": 97
        },
        "Token": [
          "github.com/pkg/errors",
          "v0.9.1"
        ],
        "InBlock": true,
        "End": {
          "Line": 7,
          "LineRune": 30,
          "Byte": 125
        }
      }
    },
    {
      "Mod": {
        "Path": "github.com/spf13/pflag",
        "Version": "v1.0.5"
      },
      "Indirect": false,
      "Syntax": {
        "Before": null,
        "Suffix": null,
        "After": null,
        "Start": {
          "Line": 8,
          "LineRune": 2,
          "Byte": 127
        },
        "Token": [
          "github.com/spf13/pflag",
          "v1.0.5"
        ],
        "InBlock": true,
        "End": {
          "Line": 8,
          "LineRune": 31,
          "Byte": 156
        }
      }
    },
    {
      "Mod": {
        "Path": "golang.org/x/mod",
        "Version": "v0.12.0"
      },
      "Indirect": false,
      "Syntax": {
        "Before": null,
        "Suffix": null,
        "After": null,
        "Start": {
          "Line": 9,
          "LineRune": 2,
          "Byte": 158
        },
        "Token": [
          "golang.org/x/mod",
          "v0.12.0"
        ],
        "InBlock": true,
        "End": {
          "Line": 9,
          "LineRune": 26,
          "Byte": 182
        }
      }
    }
  ],
  "Exclude": null,
  "Replace": null,
  "Retract": null,
  "Syntax": {
    "Name": "/Users/radek/dev/golang/src/github.com/radekg/gomod-json/go.mod",
    "Before": null,
    "Suffix": null,
    "After": null,
    "Stmt": [
      {
        "Before": null,
        "Suffix": null,
        "After": null,
        "Start": {
          "Line": 1,
          "LineRune": 1,
          "Byte": 0
        },
        "Token": [
          "module",
          "github.com/radekg/gomod-json"
        ],
        "InBlock": false,
        "End": {
          "Line": 1,
          "LineRune": 36,
          "Byte": 35
        }
      },
      {
        "Before": null,
        "Suffix": null,
        "After": null,
        "Start": {
          "Line": 3,
          "LineRune": 1,
          "Byte": 37
        },
        "Token": [
          "go",
          "1.20"
        ],
        "InBlock": false,
        "End": {
          "Line": 3,
          "LineRune": 8,
          "Byte": 44
        }
      },
      {
        "Before": null,
        "Suffix": null,
        "After": null,
        "Start": {
          "Line": 5,
          "LineRune": 1,
          "Byte": 46
        },
        "LParen": {
          "Before": null,
          "Suffix": null,
          "After": null,
          "Pos": {
            "Line": 5,
            "LineRune": 9,
            "Byte": 54
          }
        },
        "Token": [
          "require"
        ],
        "Line": [
          {
            "Before": null,
            "Suffix": null,
            "After": null,
            "Start": {
              "Line": 6,
              "LineRune": 2,
              "Byte": 57
            },
            "Token": [
              "github.com/mitchellh/go-homedir",
              "v1.1.0"
            ],
            "InBlock": true,
            "End": {
              "Line": 6,
              "LineRune": 40,
              "Byte": 95
            }
          },
          {
            "Before": null,
            "Suffix": null,
            "After": null,
            "Start": {
              "Line": 7,
              "LineRune": 2,
              "Byte": 97
            },
            "Token": [
              "github.com/pkg/errors",
              "v0.9.1"
            ],
            "InBlock": true,
            "End": {
              "Line": 7,
              "LineRune": 30,
              "Byte": 125
            }
          },
          {
            "Before": null,
            "Suffix": null,
            "After": null,
            "Start": {
              "Line": 8,
              "LineRune": 2,
              "Byte": 127
            },
            "Token": [
              "github.com/spf13/pflag",
              "v1.0.5"
            ],
            "InBlock": true,
            "End": {
              "Line": 8,
              "LineRune": 31,
              "Byte": 156
            }
          },
          {
            "Before": null,
            "Suffix": null,
            "After": null,
            "Start": {
              "Line": 9,
              "LineRune": 2,
              "Byte": 158
            },
            "Token": [
              "golang.org/x/mod",
              "v0.12.0"
            ],
            "InBlock": true,
            "End": {
              "Line": 9,
              "LineRune": 26,
              "Byte": 182
            }
          }
        ],
        "RParen": {
          "Before": null,
          "Suffix": null,
          "After": null,
          "Pos": {
            "Line": 10,
            "LineRune": 1,
            "Byte": 183
          }
        }
      }
    ]
  }
}
```
