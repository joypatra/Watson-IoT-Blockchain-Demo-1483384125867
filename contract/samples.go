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
                "behaviorArray": [
                    "NO TYPE PROPERTY"
                ],
                "licenceNumber": "carpe noctem",
                "startLatitude": "carpe noctem",
                "startLongitude": "carpe noctem",
                "startTime": "carpe noctem",
                "stopLatitude": "carpe noctem",
                "stopLongitude": "carpe noctem",
                "stopTime": "carpe noctem",
                "tripID": "carpe noctem"
            }
        },
        "lastKnownLocation": {
            "latitude": 123.456,
            "longitude": 123.456
        },
        "timestamp": "2017-01-24T08:14:45.040395277-06:00",
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
                "behaviorArray": [
                    "NO TYPE PROPERTY"
                ],
                "licenceNumber": "carpe noctem",
                "startLatitude": "carpe noctem",
                "startLongitude": "carpe noctem",
                "startTime": "carpe noctem",
                "stopLatitude": "carpe noctem",
                "stopLongitude": "carpe noctem",
                "stopTime": "carpe noctem",
                "tripID": "carpe noctem"
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
        "timestamp": "2017-01-24T08:14:45.040508409-06:00",
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