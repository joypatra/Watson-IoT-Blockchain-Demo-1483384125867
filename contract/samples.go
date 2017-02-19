package main

var samples = `
{
    "contractState": {
        "activeAssets": [
            "The ID of a managed asset. The resource focal point for a smart contract."
        ],
        "nickname": "IOT_BLOCKCHAIN_DEMO",
        "version": "The version number of the current contract"
    },
    "event": {
        "assetID": "The ID of a managed asset. The resource focal point for a smart contract.",
        "extension": {
            "authorityRating": {
                "currentRating": 123.456,
                "lastIncrement": 123.456,
                "lastIncrementTime": "carpe noctem"
            },
            "insurerRating": {
                "currentRating": 123.456,
                "lastIncrement": 123.456,
                "lastIncrementTime": "carpe noctem"
            },
            "lastTrip": {
                "ctx_sub_trips": [
                    "NO TYPE PROPERTY"
                ],
                "driver_id": "carpe noctem",
                "end_altitude": 123.456,
                "end_latitude": 123.456,
                "end_longitude": 123.456,
                "end_time": 123.456,
                "generated_time": 123.456,
                "id": {
                    "tenant_id": "carpe noctem",
                    "trip_uuid": "carpe noctem"
                },
                "mo_id": "carpe noctem",
                "start_altitude": 123.456,
                "start_latitude": 123.456,
                "start_longitude": 123.456,
                "start_time": 123.456,
                "trip_features": [
                    "NO TYPE PROPERTY"
                ],
                "trip_id": "carpe noctem"
            }
        },
        "lastKnownLocation": {
            "latitude": 123.456,
            "longitude": 123.456
        },
        "timestamp": "2017-02-19T20:08:36.789786674Z",
        "vehicleID": {
            "driverID": [
                "NO TYPE PROPERTY"
            ],
            "vin": "carpe noctem"
        }
    },
    "initEvent": {
        "nickname": "IOT_BLOCKCHAIN_DEMO",
        "version": "The ID of a managed asset. The resource focal point for a smart contract."
    },
    "state": {
        "alerts": {
            "active": [
                "TRIP_GOOD",
                "TRIP_DANGEROUS",
                "RATING_DRIVER_INCREASE",
                "RATING_DRIVER_DECREASE",
                "RATING_INSURER_INCREASE",
                "RATING_INSURER_DECREASE"
            ],
            "cleared": [
                "TRIP_GOOD",
                "TRIP_DANGEROUS",
                "RATING_DRIVER_INCREASE",
                "RATING_DRIVER_DECREASE",
                "RATING_INSURER_INCREASE",
                "RATING_INSURER_DECREASE"
            ],
            "raised": [
                "TRIP_GOOD",
                "TRIP_DANGEROUS",
                "RATING_DRIVER_INCREASE",
                "RATING_DRIVER_DECREASE",
                "RATING_INSURER_INCREASE",
                "RATING_INSURER_DECREASE"
            ]
        },
        "assetID": "The ID of a managed asset. The resource focal point for a smart contract.",
        "compliant": true,
        "extension": {
            "authorityRating": {
                "currentRating": 123.456,
                "lastIncrement": 123.456,
                "lastIncrementTime": "carpe noctem"
            },
            "insurerRating": {
                "currentRating": 123.456,
                "lastIncrement": 123.456,
                "lastIncrementTime": "carpe noctem"
            },
            "lastTrip": {
                "ctx_sub_trips": [
                    "NO TYPE PROPERTY"
                ],
                "driver_id": "carpe noctem",
                "end_altitude": 123.456,
                "end_latitude": 123.456,
                "end_longitude": 123.456,
                "end_time": 123.456,
                "generated_time": 123.456,
                "id": {
                    "tenant_id": "carpe noctem",
                    "trip_uuid": "carpe noctem"
                },
                "mo_id": "carpe noctem",
                "start_altitude": 123.456,
                "start_latitude": 123.456,
                "start_longitude": 123.456,
                "start_time": 123.456,
                "trip_features": [
                    "NO TYPE PROPERTY"
                ],
                "trip_id": "carpe noctem"
            }
        },
        "lastEvent": {
            "args": [
                "parameters to the function, usually args[0] is populated with a JSON encoded event object"
            ],
            "function": "function that created this state object",
            "redirectedFromFunction": "function that originally received the event"
        },
        "lastKnownLocation": {
            "latitude": 123.456,
            "longitude": 123.456
        },
        "timestamp": "2017-02-19T20:08:36.790030416Z",
        "txntimestamp": "Transaction timestamp matching that in the blockchain.",
        "txnuuid": "Transaction UUID matching that in the blockchain.",
        "vehicleID": {
            "driverID": [
                "NO TYPE PROPERTY"
            ],
            "vin": "carpe noctem"
        }
    }
}`