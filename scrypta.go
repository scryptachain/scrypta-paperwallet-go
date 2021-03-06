package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/jung-kurt/gofpdf"
	"github.com/skip2/go-qrcode"
)

//Network Struct
type Network struct {
	name        string
	symbol      string
	xpubkey     byte
	xprivatekey byte
}

//Wallet Struct
type Wallet struct {
	ADDRESS string `json:"address"`
	PRIVKEY string `json:"privkey"`
}

//Scrypta Network Config
var network = Network{name: "scrypta", symbol: "lyra", xpubkey: 0x30, xprivatekey: 0xae}

//GetNetworkParams --> Set the enviroment with scrypta config
func (network Network) GetNetworkParams() *chaincfg.Params {
	networkParams := &chaincfg.MainNetParams
	networkParams.PubKeyHashAddrID = network.xpubkey
	networkParams.PrivateKeyID = network.xprivatekey
	return networkParams
}

//CreatePrivateKey --> Create scrypta privkey
func (network Network) CreatePrivateKey() (*btcutil.WIF, error) {
	secret, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return nil, err
	}
	return btcutil.NewWIF(secret, network.GetNetworkParams(), true)
}

//GetAddress --> Generate scrypta address from privkey
func (network Network) GetAddress(wif *btcutil.WIF) (*btcutil.AddressPubKey, error) {
	return btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeCompressed(), network.GetNetworkParams())
}

//CreateAddress --> accept number of addresses --> return json
func CreateAddress(amount int) (string, []Wallet) {

	wallets := []Wallet{}
	for i := 0; i < amount; i++ {
		wif, _ := network.CreatePrivateKey()
		address, _ := network.GetAddress(wif)
		var wallet = Wallet{ADDRESS: address.EncodeAddress(), PRIVKEY: wif.String()}
		wallets = append(wallets, wallet)
	}

	json := ConvertToJSON(&wallets)

	log.Println("Generated", amount, "addresses")

	return json, wallets

}

func main() {

	initLog()

	amount := flag.Int("amount", 1, "amount of lyra addresses to generate")
	genpdf := flag.Bool("genpdf", false, "true/false")

	flag.Parse()

	json, wallets := CreateAddress(*amount)

	if *genpdf == true && *amount <= 10 {
		GenPDF(wallets)
	} else if *genpdf == true && *amount > 10 {
		log.Println("PDFs Not generated (max addresses amount = 10)")
	}

	fmt.Println(json)
}

//GenPDF --> Create PDF File
func GenPDF(wallets []Wallet) {

	for _, wallet := range wallets {

		pdf := gofpdf.New("P", "mm", "A4", "")
		pdf.AddPage()
		pdf.SetFont("Arial", "B", 8)
		pdf.Cell(15, 15, wallet.ADDRESS+":"+wallet.PRIVKEY)

		var png []byte
		png, err2 := qrcode.Encode(wallet.ADDRESS+":"+wallet.PRIVKEY, qrcode.Medium, 256)
		r := bytes.NewReader(png)
		pdf.RegisterImageOptionsReader("qr-code.png", gofpdf.ImageOptions{ImageType: "PNG"}, r)
		pdf.Image("qr-code.png", 7, 20, 50, 50, false, "", 0, "")

		err := pdf.OutputFileAndClose(wallet.ADDRESS + ".pdf")
		if err != nil && err2 != nil {
			log.Fatalln(err)
		}

	}

}

//ConvertToJSON Interface to JSON
func ConvertToJSON(input interface{}) string {
	btResult, _ := json.MarshalIndent(&input, "", "  ")
	return string(btResult)
}

func initLog() {
	var err error
	//Creating logs folder if not exists
	_ = os.Mkdir("logs", os.ModeDir)
	//Set LogFile name and path
	logFile, err := os.Create("logs/" + time.Now().Format("2006-01-02 15-04-05") + ".txt")
	mw := io.MultiWriter(os.Stdout, logFile)
	if err != nil {
		fmt.Println(err)
	}
	//Update logger
	log.SetOutput(mw)
}
