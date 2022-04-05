
const goWasm = new Go()

WebAssembly.instantiateStreaming(fetch("app.wasm"), goWasm.importObject)
    .then((result) => {
        console.log("webassembly Loaded");
        goWasm.run(result.instance)
    }).catch(err => {
        console.log("Error loading wen assembly ", err);
    })