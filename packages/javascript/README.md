@launchdarkly/sdk-logs-js
===========================

This package contains generated code for log codes and the associated methods.

## Code Generation

Code and message functions are generated via `generate.ts` in the tools directory. The generation code is not included in the final package.

The generated code includes a message and a `_code` function for each defined log condition. This method is used over namespaced static functions to ensure the generated code can be tree-shaken. The log codes include all conditions for client-side and server-side SDKs, but any given SDK may only use a subset of conditions.

## Commands

`npm run generate`: Generates schema code and then uses that generate schema in the generation of log codes and messages.

`npm run generate-schema`: Generates just the schema code. Useful when making updates to the code generator.

`npm run build`: Builds CJS and ESM modules.

`npm run build-cjs`: Builds CJS.

`npm run build-esm`: Builds ESM.


LaunchDarkly overview
-------------------------
[LaunchDarkly](https://www.launchdarkly.com) is a feature management platform that serves trillions of feature flags daily to help teams build better software, faster. [Get started](https://docs.launchdarkly.com/home/getting-started) using LaunchDarkly today!

[![Twitter Follow](https://img.shields.io/twitter/follow/launchdarkly.svg?style=social&label=Follow&maxAge=2592000)](https://twitter.com/intent/follow?screen_name=launchdarkly)
