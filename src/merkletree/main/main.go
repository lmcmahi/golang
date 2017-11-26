package main
// this is release 0
// this is release 1
import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"log"
	"net/http"
	_ "os"

	"github.com/cbergoon/merkletree"
	"github.com/olebedev/config"
	"github.com/vardius/gorouter"

	"math/big"
	"os"
	"net"
	"fmt"
	"merkletree/server"
	"sync"
	"encoding/base64"
	"bytes"
)
//
// testing a http router fw...
// this function will be called when /index is called...
func index(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
	fmt.Fprint(w, "Welcome \n")

}

// handle coco
func coco(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This a coco")
}

// http serer func...

func webServer() {
	router := gorouter.New()
	router.GET("/", http.HandlerFunc(index))
	router.GET("/coco", http.HandlerFunc(coco))

	log.Fatal(http.ListenAndServe(":9192", router))
}

// structure representing the data in the tree...
// should implements Content interface and have value receiver..
type TestContent struct {
	msg string
}

// calculate the hash of the content..
func (tc TestContent) CalculateHash() []byte {
	// we are using sha256 hashing algo.
	h := sha256.New()
	h.Write([]byte(tc.msg))
	return h.Sum(nil) // adding nothing
}

// compare content ..
func (tc TestContent) Equals(other merkletree.Content) bool {
	// first we need to cast the other to TestContent
	return tc.msg == other.(TestContent).msg
}
// encoding....
func encodingFunc() {
	message := []byte("Hello world....")
	var output bytes.Buffer
	var decodedMessage bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding,&output)
	encoder.Write(message)
	log.Println("encoded message..",output)
	decodedMessage.ReadFrom(base64.NewDecoder(base64.StdEncoding,&output))
	log.Println("dencoded message..",decodedMessage.String())

}
func main() {
	addrs,errs:= net.InterfaceAddrs()
	if errs != nil {
		log.Fatal("Error reading the interfaces..",errs)
	}
	for i,v := range addrs {
		log.Println(v,i)
	}
	//encoding...
	encodingFunc()
	log.Println("This is my first merkle tree\n")
	//log.Println(os.Getwd())
	cfg, errfile := config.ParseYamlFile("./config.yaml")
	if errfile != nil {
		log.Println("Error opening file..\n", errfile)
		return
	}
	value, _ := cfg.String("development.database.host")
	log.Println("returned value..\n", value)
	// now build 2 contents

	slice := []merkletree.Content{}

	content1 := TestContent{msg: "madi1"}
	content2 := TestContent{msg: "hirab"}
	content3 := TestContent{msg: "fifi"}

	slice = append(slice, content1)
	slice = append(slice, content2)
	log.Println("Number of content...\n", len(slice))
	// created a merkle tree from the list
	mkTree, err := merkletree.NewTree(slice)

	if err != nil {
		log.Println("Error ", err)
		return
	}
	log.Println("The root of the MK Tree: ", mkTree.MerkleRoot())

	ok := mkTree.VerifyTree()
	if ok {
		log.Println("Tree is good..")
	}
	// check if the tree content an element..
	//if (mkTree.VerifyContent(mkTree.MerkleRoot(),content3)){
	//	log.Println("Content is in the tree..")
	//}
	slice = append(slice, content3)
	////content2 = TestContent{msg:"hirassb"}
	//
	mkTree.RebuildTreeWith(slice)
	if mkTree.VerifyContent(mkTree.MerkleRoot(), content3) {
		log.Println("Content is in the tree..")
	}
	log.Println(mkTree.String())
	// random number..
	randNumber()
	//log.SetFlags()
	//webServer()
	madi()
	moussa()

	//send message
	sendEncryptedMessage()
	// receive encrypted message..
	readEncryptedMessage()
	// signing a message
	signingMessage()
	// verify the message...
	verifySignature()
	// describe pKey
	//describePrivateKey()
	//hashPOW()
	//register the rpc service..
	server.RegisterService()
	startServer()

}


func startServer(){
	wc := sync.WaitGroup{}
	wc.Add(1)
	l,err := net.Listen("tcp",":8976")
	if err != nil {
		log.Fatal("Error starting the server..",err)
	}
	go http.Serve(l,nil)
	log.Println("Listen...")
	wc.Wait()
}
func randNumber() {
	// reading random number
	var flib []byte
	flib = make([]byte, 23)
	rand.Read(flib) // Read uses io.ReadFull under the hood
	log.Println("Random number...\n", flib[:])
	max := big.Int{}
	max.SetUint64(93)
	number, _ := rand.Int(rand.Reader, &max)
	log.Println("number: ..", number)
}

///
// Madi private key and public key..
var madiPrivateKey *rsa.PrivateKey
var madiPublicKey *rsa.PublicKey
var message = []byte("Hello moussa")

// Moussa private and publick key
var moussaPrivateKey *rsa.PrivateKey
var moussaPublicKey *rsa.PublicKey

//common error
var err error

// contain the encrypted message
var encryptedMesssage []byte

//hashed message
var hashedMessage []byte

// hash function used to hash the message..In this case we are using the famous sha256
var hash = sha256.New()

// label
var label = []byte("")

// madi signature..
var signature []byte

// sending an encrypted message to moussa using his publickey....
func madi() {
	madiPrivateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal("Error generating private/public keypair for madi...")
	}
	madiPublicKey = &madiPrivateKey.PublicKey
	log.Println("madi PrivateKey: ", madiPrivateKey)
	log.Println("madi PublicKey: ", madiPublicKey)
}

// moussa receiving encrypted message...

func moussa() {
	moussaPrivateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal("Error generating priate/public keypair for moussa..")
	}
	moussaPublicKey = &moussaPrivateKey.PublicKey
	log.Println("moussa PrivateKey: ", moussaPrivateKey)
	log.Println("moussa PublicKey: ", moussaPublicKey)

}

// send encrypted message..

func sendEncryptedMessage() {
	// madi encrypte the message
	encryptedMesssage, err = rsa.EncryptOAEP(hash, rand.Reader, moussaPublicKey, message, label)
	if err != nil {
		log.Fatal("Error encrypting the message..", err)
	}
	log.Println("Encryptedmessage..", encryptedMesssage)
}

// now try to decrypte the message....

func readEncryptedMessage() {
	clearmessage, err := rsa.DecryptOAEP(hash, rand.Reader, moussaPrivateKey, encryptedMesssage, label)
	if err != nil {
		log.Fatal("Error decrypting the message..")
	}
	log.Println("the decrypted message is...", string(clearmessage))
}

func signingMessage() {
	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto
	//PSSMessage := message
	hash.Write(message)
	hashedMessage = hash.Sum(nil)
	signature, err = rsa.SignPSS(rand.Reader, madiPrivateKey, crypto.SHA256, hashedMessage, &opts)
	if err != nil {
		log.Fatal("Error creating signed message...", err)
	}
	log.Println("Original message: ", message)
	log.Println("hashedMessage: ", hashedMessage)
	log.Println("Signed message..", signature)
}

func verifySignature() {
	// verify if the message was signed...
	//extra := []byte("extra")
	//hashedMessage = append(hashedMessage,extra[1])
	err = rsa.VerifyPSS(madiPublicKey, crypto.SHA256, hashedMessage, signature, nil)
	if err != nil {
		log.Fatal("You are not the one that sign the message...sorry...", err)
	}
	log.Println("Good message signed by madi...")
	//crypto.Decrypter.Decrypt()
}

func describePrivateKey() {

	D := madiPrivateKey.D //private exponent
	Primes := madiPrivateKey.Primes
	PCValues := madiPrivateKey.Precomputed

	// Note : Only used for 3rd and subsequent primes
	//CRTVal := privatekey.Precomputed.CRTValues

	log.Println("Private Key : ", madiPrivateKey)
	log.Println()
	log.Println("Private Exponent : ", D.String())
	log.Println()
	log.Printf("Primes : %s %s \n", Primes[0].String(), Primes[1].String())
	log.Println()
	log.Printf("Precomputed Values : Dp[%s] Dq[%s]\n", PCValues.Dp.String(), PCValues.Dq.String())
	log.Println()
	log.Printf("Precomputed Values : Qinv[%s]", PCValues.Qinv.String())
	log.Println()
}

func hashPOW() {
	message := []byte("coco")
	var num uint16
	for {
		hash.Write(append(message, byte(num)))
		data := hash.Sum(nil)
		if (data[0] == 0) && (data[1] == 0) && (data[2] == 0) && (data[3] == 0) {
			log.Println("message..", data, " nonce: ", num)
			os.Exit(1)
			num++
			log.Println(num, " and data: ", data)
		} else {
			num++
			//log.Println(num," and data: ",data)
		}

	}
}
