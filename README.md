# booking-insights

Represens the booking's insights for an apartnmnet.

The source of data to calculate the insigths is being retirverd from tje booking request api.
The goal is to make better decisions for renting the apartament.

## API

### HTTPServer
This is an http server and it has ttwo endpoints

    -     o in the folder `cmd/n` and you can run it with the makefile target `run`
/* TODO: Pass Envs */
```sh
make run 
```

## Testing
For running the tests you can use the target`test`
/* TODO: differentiate amd implement unit, integration, e2e */

```sh
make test
```

## Building/Developing/Extras
In the makefile are included more targets to help thew devlopemtn, such as: `image`, `build, ``fmt`.
You can see all the targets with the `help` target.
```

/*TODO: Perforrmamt */
