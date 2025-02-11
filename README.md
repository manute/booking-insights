# booking-insights

Represents the booking's insights for an apartnmnet.
The goal is to make better decisions for renting the apartament.

## API

[API specification](https://app.swaggerhub.com/apis-docs/BlackfireSFL/BackendChallenge/1.0.1)

## Running
For running the app and the http server use the target `run`
```sh
make run 
```
And ths envs are accepted:

- *HTTP_PORT* default:"8080"
- *HTTP_READ_TIMEOUT* ,default:"10s"
- *HTTP_WRITE_TIMEOUT* default:"10s"
- *HTTP_MAX_HEADER_BYTES* default:"1048576"

## Testing
For running all the tests you can use the target`test`

```sh
make test
```

## Building/Developing/Extras
The makefile includes more targets to help the development process, such as: `image`, `build`, `fmt`.
You can see all of them with the `help` target.

## TODO

- [] maximize endpoint
- [] e2e test
- [] performant
- [] posibility of running sep. unit and integr. tests
- [] better logs
