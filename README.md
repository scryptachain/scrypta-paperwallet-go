# Scrypta Paperwallet GO
<p><a href="https://camo.githubusercontent.com/4e892209b4b1e2d1a773ec97e544a92f068a6f0b/68747470733a2f2f6d69726f2e6d656469756d2e636f6d2f6d61782f333136382f312a31674778414b57714b5135577a635170755f766932412e6a706567" target="_blank" rel="noopener noreferrer"><img style="display: block; margin-left: auto; margin-right: auto;" src="https://camo.githubusercontent.com/4e892209b4b1e2d1a773ec97e544a92f068a6f0b/68747470733a2f2f6d69726f2e6d656469756d2e636f6d2f6d61782f333136382f312a31674778414b57714b5135577a635170755f766932412e6a706567" alt="" data-canonical-src="https://miro.medium.com/max/3168/1*1gGxAKWqKQ5WzcQpu_vi2A.jpeg" /></a></p>
<p style="text-align: center;">&nbsp;<a title="English &mdash; Scrypta Wiki" href="https://en.scrypta.wiki" target="_blank" rel="nofollow noopener"><strong>Wiki English</strong></a>&nbsp;&middot; &middot; &middot;&nbsp;<a title="Italiano &mdash; Scrypta Wiki" href="https://it.scrypta.wiki" target="_blank" rel="nofollow noopener"><strong>Wiki italiano</strong></a></p>

## Scrypta Paper wallet implementation in GO

### Compile
#### Windows
```
go get ./...
$env:GOOS = "windows"
go build
```
#### Linux
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
