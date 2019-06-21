
# scrypta-paperwallet-go
Paper wallet implementation in GO

## Compile
### Windows
```
go get ./...
$env:GOOS = "windows"
go build
```
### Linux
```
go get ./...
$env:GOOS = "linux"
go build
```
## Run
- amount indicates the number of lyra addresses you want to generate
- genpdf indicates if you want to generate a pdf file containing the address and private key and a qr code (if the amount is not higher than 10)
```
./scrypta-paperwallet-go -amount=5 -genpdf=true
```

## Help
```
./scrypta-paperwallet-go -h

Usage of ./scrypta-paperwallet-go: 
-amount int an int (default 1) 
-genpdf true/false
```

## Dependencies
*  https://github.com/btcsuite/btcd
*  https://github.com/btcsuite/btcutil
*  https://github.com/jung-kurt/gofpdf
*  https://github.com/skip2/go-qrcode
