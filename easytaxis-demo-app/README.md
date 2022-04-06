# EasyTaxis Ltd. Fleet Simulator

An interactive CLI app simulating a fleet of smart taxis sending observability data to a Dynatrace Environment.
Useful for demonstrating Dynatrace open ingestion, generic topology, and unified analysis.

## Pre-requisites

Successfully running this simulation requires access to a Dynatrace SaaS or Managed environment and a Dynatrace API Token with permissions `metrics.ingest`, `logs.ingest`, `events.ingest`.

## Running the app

Choose the appropriate executable for your platform and run it.
You'll be greeted by an introductory banner and can type `help` to find out more.

```
Welcome to EasyTaxis Fleet Simulator 0.0.2
This is an interactive app. Type 'help' for usage information and examples.
> help

        EasyTaxis Fleet Simulator
        ---------------------------
        This tool is used to simulate generic entities, their relationships, and observability data such as
        metrics, events, and logs. The scenario simulates a fleet of smart taxies which communicate their
        data out to a Dynatrace Environment via APIs.

        Commands:
                start - starts the simulation, with given parameters
                stop  - stops any running simulation
                help  - prints this help message
                exit  - stops any running simulation and exits the app

        Additional 'help' is available for each command.
```

## Starting the simulation

You can use the start command to start the simulation. Usage is explained when typing `start help`:

```
> start help

        Command starts up the fleet simulation. Format is "start [arguments] [flags]"

        Arguments:
                --environment, -e       - Dynatrace SaaS or Managed tenant and domain. You don't need https:// or the ending slash
                --token, -t             - Dynatrace API Token with Metrics (V2) permission or the name of an environment variable (if -ev is used)
                --fleets, -f            - Number of fleets to simulate (max. 10) (default: 2)
                --taxisPerFleet, -tpf   - Number of taxis per fleet to simulate. Ranges supported using the 'min-max' format for more variety (default: 5)
        Flags:
            --env-vars, -ev             - Token is taken from an environment variable specified with -t
                --verbose, -v           - Prints out every single metric line sent

        Example:
                start -e abc123.live.dynatrace.com -t abcdefg1234567 -f 3 -tpf 2-5
```
## Data produced

The simulation ingests Metrics, Logs, and Events, to your Dynatrace environment.

**Example metric lines:**

* for a fleet (once every 2 minutes)
```
custom.easytaxis.fleet.cars.available,fleet.id="100001",fleet.location="London" 5
custom.easytaxis.fleet.cars.busy,fleet.id="100001",fleet.location="London" 0
custom.easytaxis.fleet.cars.total,fleet.id="100001",fleet.location="London" 5
custom.easytaxis.fleet.queue,fleet.id="100001",fleet.location="London" 0
```
* for a taxi (once per minute)
```
custom.easytaxis.taxi.speed,taxi.id="28434893",taxi.class="limo",taxi.registration="FX23 MDV",fleet.id="100000" 21.808139
custom.easytaxis.taxi.engine.temperature,taxi.id="28434893",taxi.class="limo",taxi.registration="FX23 MDV",fleet.id="100000" 91.808139
custom.easytaxis.taxi.engine.daystorevision,taxi.id="28434893",taxi.class="limo",taxi.registration="FX23 MDV",fleet.id="100000" 365
```
**Events are sent as follows:**
* Every 10 minutes each Fleet receives traffic updates
* Every 5 minutes each Fleet creates a customer request (booking) and one Taxi accepts the request

**Logs are sent as follows:**
* Every minute, each Taxi logs that it is sending performance metrics
* Every 2 minutes, each Fleet logs that it is sending performance metrics


## License consumption

Depending on how many Fleets and Taxis you choose to run your simulation, the total data ingestion will consume different
amounts of Davis Data Units.

You can calculate this as:

`DDUs = numberOfTaxis x 0.0037 + numberOfFleets x (0.00235 + 0.0002 x taxisPerFleet)`

For example, with 2 Fleets of 3 Taxis each I will ingest:

`6 x 0.0037 + 2 x (0.00235 + 0.0002 x 3) = 0.0281 DDUs per minute`
