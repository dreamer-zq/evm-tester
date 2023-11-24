# turbo-tester

1. Compile the contract

```bash
npx hardhat compile
```

2. Generator code

```bash
abigen  --abi ./abi/contracts/TicketGame.sol/TicketGame.json --pkg main --type TicketGame --out TicketGame.go
```

3. Build

```bash
make build
```

4. Generate

```bash
./build/tester gentx --contract=0x547bD9C389686441d9a56Db1DaffF505bC216073 --output ~/Downloads --url http://localhost:8545
```
