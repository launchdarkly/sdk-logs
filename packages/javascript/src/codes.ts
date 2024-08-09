/* eslint-disable prettier/prettier */
/* eslint-disable @typescript-eslint/no-unused-vars */
/* eslint-disable @typescript-eslint/naming-convention */
// This code is automatically generated and should not be manually edited.

export const Logs = {
  /**
   * Logs associated with the functionality of the client. Errors that can be covered by a more specific system, like "evaluation" should use that system.
  */
  Client: {
    /**
     * Codes associated with debugging.
    */
    Debug: {
    },
    /**
     * An error that should not happen in correctly implemented code. For instance missing a condition in a switch statement.
    */
    ImplementationError: {
    },
    /**
     * Codes for informative messages logged during normal operations.
    */
    Informative: {
    },
    /**
     * A non-usage error which interferes with operation and likely requires user intervention.
    */
    RuntimeError: {
    },
    /**
     * An unexpected, but recoverable, runtime issue not associated with usage.
    */
    RuntimeWarning: {
      /**
       * Initialization, or a function to wait for initialization, was done without a timeout.
      */
      MissingInitTimeout: {
        /**
         * Generate a log string for this code.
         *
         * This function will automatically include the log code.
         * @param _function The function which was called without a timeout.
        */
        message:(_function: string) => {
          return `1:4:0 The ${_function} function was called without a timeout specified. In a future version a default timeout will be applied.`;
        },
        /**
         * Get the code for this condition.
        */
        code:() => {
          return '1:4:0';
        },
      },
    },
    /**
     * An error which represents a mis-use of an API and impedes correct functionality.
    */
    UsageError: {
    },
    /**
     * A warning about the usage of an API or configuration. The usage or configuration does not interfere with operation, but is not recommended or may result in unexpected behavior.
    */
    UsageWarning: {
    },
  },
  /**
   * Codes associated with SDK configuration.
  */
  Configuration: {
    /**
     * Codes associated with debugging.
    */
    Debug: {
    },
    /**
     * An error that should not happen in correctly implemented code. For instance missing a condition in a switch statement.
    */
    ImplementationError: {
    },
    /**
     * Codes for informative messages logged during normal operations.
    */
    Informative: {
    },
    /**
     * A non-usage error which interferes with operation and likely requires user intervention.
    */
    RuntimeError: {
    },
    /**
     * An unexpected, but recoverable, runtime issue not associated with usage.
    */
    RuntimeWarning: {
    },
    /**
     * An error which represents a mis-use of an API and impedes correct functionality.
    */
    UsageError: {
    },
    /**
     * A warning about the usage of an API or configuration. The usage or configuration does not interfere with operation, but is not recommended or may result in unexpected behavior.
    */
    UsageWarning: {
      /**
       * There is a proxy configuration, and that configuration specifies to use TLS, but it is not using HTTPS authorization. This is likely not a desired configuration.
      */
      ProxyTlsAuth: {
        /**
         * Generate a log string for this code.
         *
         * This function will automatically include the log code.
        */
        message:() => {
          return `0:2:0 Proxy configured with TLS options, but is not using HTTPS authentication.`;
        },
        /**
         * Get the code for this condition.
        */
        code:() => {
          return '0:2:0';
        },
      },
    },
  },
  /**
   * Messages that apply to multiple data sources. Not specific to streaming or polling.
  */
  DataSource: {
    /**
     * Codes associated with debugging.
    */
    Debug: {
      /**
       * The connection has received invalid JSON. The error should have been logged with code 10:5:0. This debug log provides additional information. This should be logged after the error message.
      */
      InvalidJsonDebug: {
        /**
         * Generate a log string for this code.
         *
         * This function will automatically include the log code.
         * @param json The JSON data that was invalid.
        */
        message:(json: string) => {
          return `10:0:0 Invalid JSON follows: ${json}`;
        },
        /**
         * Get the code for this condition.
        */
        code:() => {
          return '10:0:0';
        },
      },
    },
    /**
     * An error that should not happen in correctly implemented code. For instance missing a condition in a switch statement.
    */
    ImplementationError: {
    },
    /**
     * Codes for informative messages logged during normal operations.
    */
    Informative: {
    },
    /**
     * A non-usage error which interferes with operation and likely requires user intervention.
    */
    RuntimeError: {
      /**
       * The LaunchDarkly data source received a payload with invalid JSON.
      */
      InvalidJson: {
        /**
         * Generate a log string for this code.
         *
         * This function will automatically include the log code.
         * @param connection The type of the connection. Should be capitalized. Examples: "Streaming" or "Polling"
         * @param _type The type of the message. For example "put", "patch", or "delete".
        */
        message:(connection: string, _type: string) => {
          return `10:5:0 ${connection} connection received invalid data in "${_type}" message.`;
        },
        /**
         * Get the code for this condition.
        */
        code:() => {
          return '10:5:0';
        },
      },
      /**
       * The data source received an invalid payload. The JSON was correctly formed, but the data did not match the expected schema.
      */
      InvalidPayload: {
        /**
         * Generate a log string for this code.
         *
         * This function will automatically include the log code.
         * @param connection The type of the connection. Should be capitalized. Examples: "Streaming" or "Polling"
         * @param _type The type of the message. For example "put", "patch", or "delete"
        */
        message:(connection: string, _type: string) => {
          return `10:5:1 ${connection} connection received invalid data in "${_type}" message.`;
        },
        /**
         * Get the code for this condition.
        */
        code:() => {
          return '10:5:1';
        },
      },
    },
    /**
     * An unexpected, but recoverable, runtime issue not associated with usage.
    */
    RuntimeWarning: {
    },
    /**
     * An error which represents a mis-use of an API and impedes correct functionality.
    */
    UsageError: {
    },
    /**
     * A warning about the usage of an API or configuration. The usage or configuration does not interfere with operation, but is not recommended or may result in unexpected behavior.
    */
    UsageWarning: {
    },
  },
  /**
   * Codes associated with evaluation.
  */
  Evaluation: {
    /**
     * Codes associated with debugging.
    */
    Debug: {
    },
    /**
     * An error that should not happen in correctly implemented code. For instance missing a condition in a switch statement.
    */
    ImplementationError: {
    },
    /**
     * Codes for informative messages logged during normal operations.
    */
    Informative: {
    },
    /**
     * A non-usage error which interferes with operation and likely requires user intervention.
    */
    RuntimeError: {
    },
    /**
     * An unexpected, but recoverable, runtime issue not associated with usage.
    */
    RuntimeWarning: {
    },
    /**
     * An error which represents a mis-use of an API and impedes correct functionality.
    */
    UsageError: {
    },
    /**
     * A warning about the usage of an API or configuration. The usage or configuration does not interfere with operation, but is not recommended or may result in unexpected behavior.
    */
    UsageWarning: {
    },
  },
  /**
   * Codes associated with events.
  */
  Events: {
    /**
     * Codes associated with debugging.
    */
    Debug: {
      /**
       * Flushing of an event batch failed.
      */
      FlushFailed: {
        /**
         * Generate a log string for this code.
         *
         * This function will automatically include the log code.
         * @param error The reason the flush failed.
        */
        message:(error: string) => {
          return `8:0:0 Failed to flush events. Reason: ${error}`;
        },
        /**
         * Get the code for this condition.
        */
        code:() => {
          return '8:0:0';
        },
      },
      /**
       * The event processor has started.
      */
      EventProcessorStarted: {
        /**
         * Generate a log string for this code.
         *
         * This function will automatically include the log code.
        */
        message:() => {
          return `8:0:1 Started event processor.`;
        },
        /**
         * Get the code for this condition.
        */
        code:() => {
          return '8:0:1';
        },
      },
      /**
       * The event processor is flushing events.
      */
      FlushingEvents: {
        /**
         * Generate a log string for this code.
         *
         * This function will automatically include the log code.
         * @param eventCount The number of events in the batch being flushed.
        */
        message:(eventCount: string) => {
          return `8:0:2 Flushing ${eventCount} events.`;
        },
        /**
         * Get the code for this condition.
        */
        code:() => {
          return '8:0:2';
        },
      },
      /**
       * Event delivery failed, but a retry attempt is going to be made.
      */
      EventRetry: {
        /**
         * Generate a log string for this code.
         *
         * This function will automatically include the log code.
        */
        message:() => {
          return `8:0:3 Encountered a problem sending events, will retry.`;
        },
        /**
         * Get the code for this condition.
        */
        code:() => {
          return '8:0:3';
        },
      },
    },
    /**
     * An error that should not happen in correctly implemented code. For instance missing a condition in a switch statement.
    */
    ImplementationError: {
    },
    /**
     * Codes for informative messages logged during normal operations.
    */
    Informative: {
    },
    /**
     * A non-usage error which interferes with operation and likely requires user intervention.
    */
    RuntimeError: {
    },
    /**
     * An unexpected, but recoverable, runtime issue not associated with usage.
    */
    RuntimeWarning: {
      /**
       * The capacity of the event queue was exceeded. The user may want to increase the capacity, or increase the frequency of flushing. In this situation some events have already been dropped. We only want to log this warning the first time it happens.
      */
      EventCapacityExceeded: {
        /**
         * Generate a log string for this code.
         *
         * This function will automatically include the log code.
        */
        message:() => {
          return `8:4:0 Exceeded event queue capacity. Increase capacity, or decrease flushing interval, to avoid dropping events.`;
        },
        /**
         * Get the code for this condition.
        */
        code:() => {
          return '8:4:0';
        },
      },
    },
    /**
     * An error which represents a mis-use of an API and impedes correct functionality.
    */
    UsageError: {
    },
    /**
     * A warning about the usage of an API or configuration. The usage or configuration does not interfere with operation, but is not recommended or may result in unexpected behavior.
    */
    UsageWarning: {
    },
  },
  /**
   * Codes associated with network conditions and failures. When possible use "streaming", "polling", or "events" instead, or log something in one of those systems in addition to the general network error. For instance an event source may log a "generalNetwork"  condition and the streaming data source a "sreaming" condition.
  */
  GeneralNetwork: {
    /**
     * Codes associated with debugging.
    */
    Debug: {
    },
    /**
     * An error that should not happen in correctly implemented code. For instance missing a condition in a switch statement.
    */
    ImplementationError: {
    },
    /**
     * Codes for informative messages logged during normal operations.
    */
    Informative: {
    },
    /**
     * A non-usage error which interferes with operation and likely requires user intervention.
    */
    RuntimeError: {
    },
    /**
     * An unexpected, but recoverable, runtime issue not associated with usage.
    */
    RuntimeWarning: {
    },
    /**
     * An error which represents a mis-use of an API and impedes correct functionality.
    */
    UsageError: {
    },
    /**
     * A warning about the usage of an API or configuration. The usage or configuration does not interfere with operation, but is not recommended or may result in unexpected behavior.
    */
    UsageWarning: {
    },
  },
  /**
   * Codes associated with hooks.
  */
  Hooks: {
    /**
     * Codes associated with debugging.
    */
    Debug: {
    },
    /**
     * An error that should not happen in correctly implemented code. For instance missing a condition in a switch statement.
    */
    ImplementationError: {
    },
    /**
     * Codes for informative messages logged during normal operations.
    */
    Informative: {
    },
    /**
     * A non-usage error which interferes with operation and likely requires user intervention.
    */
    RuntimeError: {
    },
    /**
     * An unexpected, but recoverable, runtime issue not associated with usage.
    */
    RuntimeWarning: {
    },
    /**
     * An error which represents a mis-use of an API and impedes correct functionality.
    */
    UsageError: {
    },
    /**
     * A warning about the usage of an API or configuration. The usage or configuration does not interfere with operation, but is not recommended or may result in unexpected behavior.
    */
    UsageWarning: {
    },
  },
  /**
   * Codes associated with the in-memory store for flags/segments.
  */
  MemoryStore: {
    /**
     * Codes associated with debugging.
    */
    Debug: {
    },
    /**
     * An error that should not happen in correctly implemented code. For instance missing a condition in a switch statement.
    */
    ImplementationError: {
    },
    /**
     * Codes for informative messages logged during normal operations.
    */
    Informative: {
    },
    /**
     * A non-usage error which interferes with operation and likely requires user intervention.
    */
    RuntimeError: {
    },
    /**
     * An unexpected, but recoverable, runtime issue not associated with usage.
    */
    RuntimeWarning: {
    },
    /**
     * An error which represents a mis-use of an API and impedes correct functionality.
    */
    UsageError: {
    },
    /**
     * A warning about the usage of an API or configuration. The usage or configuration does not interfere with operation, but is not recommended or may result in unexpected behavior.
    */
    UsageWarning: {
    },
  },
  /**
   * Codes associated with a persistent store.
  */
  PersistentStore: {
    /**
     * Codes associated with debugging.
    */
    Debug: {
    },
    /**
     * An error that should not happen in correctly implemented code. For instance missing a condition in a switch statement.
    */
    ImplementationError: {
    },
    /**
     * Codes for informative messages logged during normal operations.
    */
    Informative: {
    },
    /**
     * A non-usage error which interferes with operation and likely requires user intervention.
    */
    RuntimeError: {
    },
    /**
     * An unexpected, but recoverable, runtime issue not associated with usage.
    */
    RuntimeWarning: {
    },
    /**
     * An error which represents a mis-use of an API and impedes correct functionality.
    */
    UsageError: {
    },
    /**
     * A warning about the usage of an API or configuration. The usage or configuration does not interfere with operation, but is not recommended or may result in unexpected behavior.
    */
    UsageWarning: {
    },
  },
  /**
   * Codes associated with LaunchDarkly polling connections and payloads.
  */
  Polling: {
    /**
     * Codes associated with debugging.
    */
    Debug: {
    },
    /**
     * An error that should not happen in correctly implemented code. For instance missing a condition in a switch statement.
    */
    ImplementationError: {
    },
    /**
     * Codes for informative messages logged during normal operations.
    */
    Informative: {
    },
    /**
     * A non-usage error which interferes with operation and likely requires user intervention.
    */
    RuntimeError: {
    },
    /**
     * An unexpected, but recoverable, runtime issue not associated with usage.
    */
    RuntimeWarning: {
    },
    /**
     * An error which represents a mis-use of an API and impedes correct functionality.
    */
    UsageError: {
    },
    /**
     * A warning about the usage of an API or configuration. The usage or configuration does not interfere with operation, but is not recommended or may result in unexpected behavior.
    */
    UsageWarning: {
    },
  },
  /**
   * Codes associated with LaunchDarkly streaming connections and payloads.
  */
  Streaming: {
    /**
     * Codes associated with debugging.
    */
    Debug: {
    },
    /**
     * An error that should not happen in correctly implemented code. For instance missing a condition in a switch statement.
    */
    ImplementationError: {
    },
    /**
     * Codes for informative messages logged during normal operations.
    */
    Informative: {
    },
    /**
     * A non-usage error which interferes with operation and likely requires user intervention.
    */
    RuntimeError: {
    },
    /**
     * An unexpected, but recoverable, runtime issue not associated with usage.
    */
    RuntimeWarning: {
    },
    /**
     * An error which represents a mis-use of an API and impedes correct functionality.
    */
    UsageError: {
    },
    /**
     * A warning about the usage of an API or configuration. The usage or configuration does not interfere with operation, but is not recommended or may result in unexpected behavior.
    */
    UsageWarning: {
    },
  },
}
