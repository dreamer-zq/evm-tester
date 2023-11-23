# turbo-tester

1. Compile the contract

```bash
npx hardhat compile
```

2. Generator code

```bash
abigen  --abi ./abi/contracts/TicketGame.sol/TicketGame.json --pkg main --type TicketGame --out TicketGame.go
```
