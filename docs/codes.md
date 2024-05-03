# LaunchDarkly Log Codes

## Introduction 

Log codes provide a standardized way to reference different log conditions across LaunchDarkly SDKs.

## Codes

### client

Logs associated with the functionality of the client. Errors that can be covered by a more specific system, like "evaluation" should use that system.

#### 1:4:0 - missingInitTimeout

Initialization, or a function to wait for initialization, was done without a timeout.

| code | system | class |
|------|--------|-------|
| 1:4:0 | client | runtimeWarning |

##### Message

`The ${function} function was called without a timeout specified. In a future version a default timeout will be applied.`

| parameter | description |
|-----------|-------------|
| function | The function which was called without a timeout. |
### configuration

Codes associated with SDK configuration.

#### 0:2:0 - proxyTlsAuth

There is a proxy configuration, and that configuration specifies to use TLS, but it is not using HTTPS authorization. This is likely not a desired configuration.

| code | system | class |
|------|--------|-------|
| 0:2:0 | configuration | usageWarning |

##### Message

`Proxy configured with TLS options, but is not using HTTPS authentication.`

### dataSource

Messages that apply to multiple data sources. Not specific to streaming or polling.

#### 10:0:0 - invalidJsonDebug

The connection has received invalid JSON. The error should have been logged with code 10:5:0. This debug log provides additional information. This should be logged after the error message.

| code | system | class |
|------|--------|-------|
| 10:0:0 | dataSource | debug |

##### Message

`Invalid JSON follows: ${json}`

| parameter | description |
|-----------|-------------|
| json | The JSON data that was invalid. |
#### 10:5:0 - invalidJson

The LaunchDarkly data source received a payload with invalid JSON.

| code | system | class |
|------|--------|-------|
| 10:5:0 | dataSource | runtimeError |

##### Message

`${connection} connection received invalid data in "${type}" message.`

| parameter | description |
|-----------|-------------|
| type | The type of the message. For example "put", "patch", or "delete". |
| connection | The type of the connection. Should be capitalized. Examples: "Streaming" or "Polling" |
#### 10:5:1 - invalidPayload

The data source received an invalid payload. The JSON was correctly formed, but the data did not match the expected schema.

| code | system | class |
|------|--------|-------|
| 10:5:1 | dataSource | runtimeError |

##### Message

`${connection} connection received invalid data in "${type}" message.`

| parameter | description |
|-----------|-------------|
| connection | The type of the connection. Should be capitalized. Examples: "Streaming" or "Polling" |
| type | The type of the message. For example "put", "patch", or "delete" |
### evaluation

Codes associated with evaluation.

### events

Codes associated with events.

#### 8:0:0 - flushFailed

Flushing of an event batch failed.

| code | system | class |
|------|--------|-------|
| 8:0:0 | events | debug |

##### Message

`Failed to flush events. Reason: ${error}`

| parameter | description |
|-----------|-------------|
| error | The reason the flush failed. |
#### 8:0:1 - eventProcessorStarted

The event processor has started.

| code | system | class |
|------|--------|-------|
| 8:0:1 | events | debug |

##### Message

`Started event processor.`

#### 8:0:2 - flushingEvents

The event processor is flushing events.

| code | system | class |
|------|--------|-------|
| 8:0:2 | events | debug |

##### Message

`Flushing ${eventCount} events.`

| parameter | description |
|-----------|-------------|
| eventCount | The number of events in the batch being flushed. |
#### 8:0:3 - eventRetry

Event delivery failed, but a retry attempt is going to be made.

| code | system | class |
|------|--------|-------|
| 8:0:3 | events | debug |

##### Message

`Encountered a problem sending events, will retry.`

#### 8:4:0 - eventCapacityExceeded

The capacity of the event queue was exceeded. The user may want to increase the capacity, or increase the frequency of flushing. In this situation some events have already been dropped. We only want to log this warning the first time it happens.

| code | system | class |
|------|--------|-------|
| 8:4:0 | events | runtimeWarning |

##### Message

`Exceeded event queue capacity. Increase capacity, or decrease flushing interval, to avoid dropping events.`

### generalNetwork

Codes associated with network conditions and failures. When possible use "streaming", "polling", or "events" instead, or log something in one of those systems in addition to the general network error. For instance an event source may log a "generalNetwork"  condition and the streaming data source a "sreaming" condition.

### hooks

Codes associated with hooks.

### memoryStore

Codes associated with the in-memory store for flags/segments.

### persistentStore

Codes associated with a persistent store.

### polling

Codes associated with LaunchDarkly polling connections and payloads.

### streaming

Codes associated with LaunchDarkly streaming connections and payloads.

