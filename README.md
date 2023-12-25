# X Protocol WebAssembly Miner

## How to use

Download or build the WebAssembly binary and `wasm_exec.js` file.

Here's a simple single-thread for you to get started:

```html

<html>
<head>
    <script src="wasm_exec.js"></script>
    <script>
        const go = new Go();
        WebAssembly.instantiateStreaming(fetch('main.wasm'), go.importObject).then((result) => {
            go.run(result.instance);
        });
    </script>
</head>
<body>
<p>Contract Address</p>
<input id="contract_address" type="text"/>
<p>Sender Address</p>
<input id="sender_address" type="text"/>
<p>Difficulty</p>
<input id="difficulty" type="number"/>
<button onclick="mine()">Mine</button>

<script>
    function mine() {
        const contractAddress = document.getElementById('contract_address').value;
        const senderAddress = document.getElementById('sender_address').value;
        const difficulty = document.getElementById('difficulty').value;
        let nonce = null;
        
        while(nonce === null) {
            nonce = calcBulk(contractAddress, senderAddress, difficulty, 100000);  // 100000 stands for the number of hashes to calculate each time
        }
        
        alert('Found a valid nonce: ' + nonce);  // nonce is a decimal string
    }
</script>
</body>
</html>
```

It's recommended to use Web Workers for multi-threading. An example will be provided soon.

## Build from source

Make sure you have installed [Golang](https://golang.org/).

The following steps assume that you are using Linux.

```shell
git clone https://github.com/X-Proto/wasm-miner

cd wasm-miner

bash build.sh

ls output  # main.wasm  wasm_exec.js
```

## License

Licensed under the [MIT License](LICENSE).
