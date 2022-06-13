# srategy
ztrade examples

strategy is written by go language, you can simply run "go build" "go test" to test if there has errors

## backtest

```
./ztrade backtest --fee 0.0002 --balance 5000 --start "2022-01-01 00:00:00" --end "2022-06-01 00:00:00" --exchange binance --symbol ETHUSDT --script macd.go --param '{"bin": "30m", "fast": 7, "slow": 20, "dea": 9, "amount": 2}'
```
