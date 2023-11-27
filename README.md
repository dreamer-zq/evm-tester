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
./build/tester gentx --contract=<contract-addr> --output ~/Downloads --url http://localhost:8545 --chain-id 1333 --count 10 --gas-fee-cap 150000 --gas-tip-cap 50000 --gas-limit 200000
```
