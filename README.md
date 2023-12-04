# turbo-tester

1. Install `solcjs`

```bash
npm install -g solc@0.8.19
```

2. Compile the contract

```bash
rm -rf ./compiled
solcjs --bin --abi ./contracts/TicketGame.sol --base-path ./ --include-path ./node_modules/ --output-dir ./compiled/
```

3. Generator code

```bash
abigen  --abi ./compiled/contracts_TicketGame_sol_TicketGame.abi --bin ./compiled/contracts_TicketGame_sol_TicketGame.bin --pkg gen --type TicketGame --out ./simple/gen/TicketGame.go
```

4. Build

```bash
make build
```

5. Generate

```bash
./build/tester gentx --contract=0x476F62693e194C50141c62D818D6112a9a70826a --output ~/Downloads --url http://localhost:8545 --chain-id 1223 --batch-size 1000 --gas-fee-cap 150000 --gas-tip-cap 50000 --gas-limit 200000  --contract-method-params 0x476F62693e194C50141c62D818D6112a9a70826a
```

6. Start

```bash
./build/tester start --contract=0x476F62693e194C50141c62D818D6112a9a70826a --url http://localhost:8545 --chain-id 1223 --batch-size 1000 --gas-fee-cap 150000 --gas-tip-cap 50000 --gas-limit 200000  --run-period 5s --run-user-num 10  --contract-method-params 0x476F62693e194C50141c62D818D6112a9a70826a
```
