# booking-insights- 
### This is a challenge for an event - not for a real use-case
### It is an DDD training challenge, with caveats due to the lang and my thoughts: 
  - I am not doibg a pure DDD as in golang you can skip some of the verbosisty of some aspects of that i.e Java (it does not make any sense all thte folders/subfolders here)
  - The domain and infrastructure are the main folders for the internal's app
  - infrastructure/ for all the external connections, in this case http, but it also could have DB, events, etc
  - domain/ for the model that the ubuquitous lang decided internally, in this case mine :)  

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

