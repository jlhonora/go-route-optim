Route optimization in Go
========================

A quick solver for the route optimization problem. The goal is to get a decently optimized route in less than a second.

**Warning:** This is a work in progress, don't expect it to stay consistent.

## Input

Here's an input example:

```
{
    "points": [
        {
            "id": 123,
            "waypoint": {
                "latitude": -12.082021,
                "longitude": -77.058512
            }
        },
        {
            "id": 124,
            "waypoint": {
                "latitude": -12.071229,
                "longitude": -77.039371
            }
        },
        {
            "id": 125,
            "waypoint": {
                "latitude": -12.085998,
                "longitude": -77.027968
            },
        }
    ],
    "start": {
        "waypoint": {
            "latitude": -33.1,
            "longitude": -71.0
        }
    }
}
```

The point structure could be extended to include time windows, priorities and many others.

## Output

```
{
    "points": [
        {
            "id": 123,
            "slot": 0,
            "waypoint": {
                "latitude": -12.082021,
                "longitude": -77.058512
            }
        },
        {
            "id": 124,
            "slot": 1,
            "waypoint": {
                "latitude": -12.071229,
                "longitude": -77.039371
            }
        },
        {
            "id": 125,
            "slot": 2,
            "waypoint": {
                "latitude": -12.085998,
                "longitude": -77.027968
            }
        }
    ]
}
```
