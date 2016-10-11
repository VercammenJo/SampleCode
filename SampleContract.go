package main

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)
 // "strings"	"strconv"  "encoding/json"

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

//var marbleIndexStr = "_marbleindex"				//name for the key/value that will store a list of all known marbles
//var openTradesStr = "_opentrades"				//name for the key/value that will store all open trades

type Customer struct{
	Name string `json:"name"`					//the fieldtags are needed to keep case from bouncing around
	Surname string `json:"surname"`
	Address string `json:"address"`
	City string `json:"city"`
	Country string `json:"country"`
}

// ============================================================================================================================
// Main
// ============================================================================================================================

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Example Chaincode: %s", err)
	}
}

// ============================================================================================================================
// Init - reset all the things
// ============================================================================================================================

func (t *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	return nil, nil
}

// ============================================================================================================================
// Invoke - Our entry point for Invocations
// ============================================================================================================================

func (t *SimpleChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {													//initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	} else if function == "create_customer" {							   	//Creates a customer
		return t.create_customer(stub, args)
	//} else if function == "update_customer" {											//writes a value to the chaincode state
	//	return t.update_customer(stub, args)
	//} else if function == "init_Document" {									//create a new Document
	//	return t.init_Document(stub, args)
	//} else if function == "set_user" {										//change owner of a marble
	//	return t.set_user(stub, args)
	}
	fmt.Println("invoke did not find func: " + function)					//error

	return nil, errors.New("Received unknown function invocation")
}

// ============================================================================================================================
// Query - Our entry point for Queries
// ============================================================================================================================

func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read_customer" {													//read a variable
		return t.read_customer(stub, args)
	}
	fmt.Println("query did not find func: " + function)						//error

	return nil, errors.New("Received unknown function query")
}

// ============================================================================================================================
// Read Customer - read a customer from state 
// ============================================================================================================================

func (t *SimpleChaincode) read_customer(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var name, jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the var to query")
	}

	name = args[0]
	valAsbytes, err := stub.GetState(name)									//get the var from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + name + "\"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil													//send it onward
}

// ============================================================================================================================
// Create Customer - create a customer in state 
// ============================================================================================================================

func (t *SimpleChaincode) create_customer(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var err error
	//   0       1       2     3
	// "asdf", "blue", "35", "bob"
	if len(args) != 5 {
		return nil, errors.New("Incorrect number of arguments. Expecting 5")
	}

	fmt.Println("- start creation of customer")
	if len(args[0]) <= 0 {
		return nil, errors.New("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return nil, errors.New("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return nil, errors.New("3rd argument must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return nil, errors.New("4th argument must be a non-empty string")
	}
	if len(args[4]) <= 0 {
		return nil, errors.New("5th argument must be a non-empty string")
	}
	

	str := `{"name": "` + args[0] + `", "surname": "` + args[1] + `", "address": "` + args[2]  + `, "city": "` + args[3]+  `, "country": "` + args[4]+ `"}`
	err = stub.PutState(args[0], []byte(str))								//store marble with id as key
	if err != nil {
		return nil, err
	}

	fmt.Println("- end create customer")
	return nil, nil
}
