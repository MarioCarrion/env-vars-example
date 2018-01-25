# Environment Variables Example

[![Build Status](https://travis-ci.org/MarioCarrion/env-vars-example.svg?branch=master)](https://travis-ci.org/MarioCarrion/env-vars-example.svg?branch=master)

Really tiny and simple example about using both [envconfig](https://github.com/kelseyhightower/envconfig) and [godotenv](https://github.com/joho/godotenv) as a prerequisite for a 12 factor application.

## Usage

Without indicating any environment variable it uses the default values:

```
$ ./env-vars-example
root:@localhost:3306/dbname
```

but if you set the environment variables, those take preference:

```
$ DSN_USERNAME=mario DSN_HOST=fancy.com ./env-vars-example
mario:@fancy.com:3306/dbname
```
