import {client_runtimeWarning_missingInitTimeout, dataSource_runtimeError_invalidJson} from "./codes";

async function sleep(time: number) {
    return new Promise((resolve) => {setTimeout(resolve, time)})
}

async function main(){
    while(true) {
        console.warn(client_runtimeWarning_missingInitTimeout("waitForInitialization"));
        await sleep(5000)
        console.log("Some normal output.")
        await sleep(5000)
        console.warn(dataSource_runtimeError_invalidJson("streaming", "put"))
    }
}

main();