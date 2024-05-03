/* eslint-disable prettier/prettier */
/* eslint-disable @typescript-eslint/no-unused-vars */
/* eslint-disable @typescript-eslint/naming-convention */
// This code is automatically generated and should not be manually edited.

/**
 * There is a proxy configuration, and that configuration specifies to use TLS, but it is not using HTTPS authorization. This is likely not a desired configuration.
 *
 * Generate a log string for this code.
 *
 * This function will automatically include the log code.
*/
export function configuration_usageWarning_proxyTlsAuth(): string {
  return `0:2:0 Proxy configured with TLS options, but is not using HTTPS authentication.`;
}
/**
 * There is a proxy configuration, and that configuration specifies to use TLS, but it is not using HTTPS authorization. This is likely not a desired configuration.
 *
 * Get the code for this condition.
*/
export function configuration_usageWarning_proxyTlsAuth_code(): string {
  return '0:2:0';
}
/**
 * The connection has received invalid JSON. The error should have been logged with code 10:5:0. This debug log provides additional information. This should be logged after the error message.
 *
 * Generate a log string for this code.
 *
 * This function will automatically include the log code.
 * @param json The JSON data that was invalid.
*/
export function dataSource_debug_invalidJsonDebug(json: string): string {
  return `10:0:0 Invalid JSON follows: ${json}`;
}
/**
 * The connection has received invalid JSON. The error should have been logged with code 10:5:0. This debug log provides additional information. This should be logged after the error message.
 *
 * Get the code for this condition.
*/
export function dataSource_debug_invalidJsonDebug_code(): string {
  return '10:0:0';
}
/**
 * The LaunchDarkly data source received a payload with invalid JSON.
 *
 * Generate a log string for this code.
 *
 * This function will automatically include the log code.
 * @param connection The type of the connection. Should be capitalized. Examples: "Streaming" or "Polling"
 * @param type The type of the message. For example "put", "patch", or "delete".
*/
export function dataSource_runtimeError_invalidJson(connection: string, type: string): string {
  return `10:5:0 ${connection} connection received invalid data in "${type}" message.`;
}
/**
 * The LaunchDarkly data source received a payload with invalid JSON.
 *
 * Get the code for this condition.
*/
export function dataSource_runtimeError_invalidJson_code(): string {
  return '10:5:0';
}
/**
 * The data source received an invalid payload. The JSON was correctly formed, but the data did not match the expected schema.
 *
 * Generate a log string for this code.
 *
 * This function will automatically include the log code.
 * @param connection The type of the connection. Should be capitalized. Examples: "Streaming" or "Polling"
 * @param type The type of the message. For example "put", "patch", or "delete"
*/
export function dataSource_runtimeError_invalidPayload(connection: string, type: string): string {
  return `10:5:1 ${connection} connection received invalid data in "${type}" message.`;
}
/**
 * The data source received an invalid payload. The JSON was correctly formed, but the data did not match the expected schema.
 *
 * Get the code for this condition.
*/
export function dataSource_runtimeError_invalidPayload_code(): string {
  return '10:5:1';
}
/**
 * Initialization, or a function to wait for initialization, was done without a timeout.
 *
 * Generate a log string for this code.
 *
 * This function will automatically include the log code.
 * @param function The function which was called without a timeout.
*/
export function client_runtimeWarning_missingInitTimeout(function: string): string {
  return `1:4:0 The ${function} function was called without a timeout specified. In a future version a default timeout will be applied.`;
}
/**
 * Initialization, or a function to wait for initialization, was done without a timeout.
 *
 * Get the code for this condition.
*/
export function client_runtimeWarning_missingInitTimeout_code(): string {
  return '1:4:0';
}
/**
 * Flushing of an event batch failed.
 *
 * Generate a log string for this code.
 *
 * This function will automatically include the log code.
 * @param error The reason the flush failed.
*/
export function events_debug_flushFailed(error: string): string {
  return `8:0:0 Failed to flush events. Reason: ${error}`;
}
/**
 * Flushing of an event batch failed.
 *
 * Get the code for this condition.
*/
export function events_debug_flushFailed_code(): string {
  return '8:0:0';
}
/**
 * The event processor has started.
 *
 * Generate a log string for this code.
 *
 * This function will automatically include the log code.
*/
export function events_debug_eventProcessorStarted(): string {
  return `8:0:1 Started event processor.`;
}
/**
 * The event processor has started.
 *
 * Get the code for this condition.
*/
export function events_debug_eventProcessorStarted_code(): string {
  return '8:0:1';
}
/**
 * The event processor is flushing events.
 *
 * Generate a log string for this code.
 *
 * This function will automatically include the log code.
 * @param eventCount The number of events in the batch being flushed.
*/
export function events_debug_flushingEvents(eventCount: string): string {
  return `8:0:2 Flushing ${eventCount} events.`;
}
/**
 * The event processor is flushing events.
 *
 * Get the code for this condition.
*/
export function events_debug_flushingEvents_code(): string {
  return '8:0:2';
}
/**
 * Event delivery failed, but a retry attempt is going to be made.
 *
 * Generate a log string for this code.
 *
 * This function will automatically include the log code.
*/
export function events_debug_eventRetry(): string {
  return `8:0:3 Encountered a problem sending events, will retry.`;
}
/**
 * Event delivery failed, but a retry attempt is going to be made.
 *
 * Get the code for this condition.
*/
export function events_debug_eventRetry_code(): string {
  return '8:0:3';
}
/**
 * The capacity of the event queue was exceeded. The user may want to increase the capacity, or increase the frequency of flushing. In this situation some events have already been dropped. We only want to log this warning the first time it happens.
 *
 * Generate a log string for this code.
 *
 * This function will automatically include the log code.
*/
export function events_runtimeWarning_eventCapacityExceeded(): string {
  return `8:4:0 Exceeded event queue capacity. Increase capacity, or decrease flushing interval, to avoid dropping events.`;
}
/**
 * The capacity of the event queue was exceeded. The user may want to increase the capacity, or increase the frequency of flushing. In this situation some events have already been dropped. We only want to log this warning the first time it happens.
 *
 * Get the code for this condition.
*/
export function events_runtimeWarning_eventCapacityExceeded_code(): string {
  return '8:4:0';
}
