package test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestDecodeJson(t *testing.T) {
	jsonString := `{"FirstName":"Iqbal","MiddleName":"Muhammad","LastName":"Fauzan","Gender":"Male","Age":27,"Hobbies":["Running","Coding","Eating","Hangout"]}`
	jsonBytes := []byte(jsonString)
	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.Gender)
	fmt.Println(customer.Age)
	fmt.Println(customer.Hobbies)
}

func logjson(data interface{}) {
	bytes, err := json.Marshal(data) // untuk encode data dari golang -> json
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
func TestEncode(t *testing.T) {
	logjson("Iqbal")
	logjson(1)
	logjson(false)
	logjson([]string{"iqbal", "fauzan"})
}

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Gender     string
	Age        int
	Hobbies    []string
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Iqbal",
		LastName:   "Fauzan",
		MiddleName: "Muhammad",
		Gender:     "Male",
		Age:        27,
	}
	marshal, _ := json.Marshal(customer)
	fmt.Println(string(marshal))
}

func TestJsonArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Iqbal",
		LastName:   "Fauzan",
		MiddleName: "Muhammad",
		Gender:     "Male",
		Age:        27,
		Hobbies:    []string{"Running", "Coding", "Eating", "Hangout"},
	}
	marshalBytes, _ := json.Marshal(customer)
	fmt.Println(string(marshalBytes))
}

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func TestJsonTag(t *testing.T) {
	product := Product{
		Id:       "P001",
		Name:     "product 1",
		ImageUrl: "http://example.com/imageurl",
	}
	marshal, _ := json.Marshal(product)
	fmt.Println(string(marshal))
}

func TestJsonTagDecode(t *testing.T) {
	jsonString := `{"id":"P001","name":"product 1","image_url":"http://example.com/imageurl"}`
	jsonByte := []byte(jsonString)
	product := &Product{}
	json.Unmarshal(jsonByte, product)
	fmt.Println(product)
	fmt.Println(product.Id)
	fmt.Println(product.Name)
	fmt.Println(product.ImageUrl)
}

func TestMap(t *testing.T) {
	jsonString := `{"id":"P001","name":"product 1","image_url":"http://example.com/imageurl","price":200000}`
	jsonByte := []byte(jsonString)

	var result map[string]interface{}
	err := json.Unmarshal(jsonByte, &result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["image_url"])
	fmt.Println(result["price"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":        "P00001",
		"name":      "Macbook Pro M3 16GB RAM",
		"image_url": "http://example.com/image_url.png",
		"price":     25000000,
	}
	marshal, _ := json.Marshal(product)
	fmt.Println(string(marshal))
}

func TestStreamDecoder(t *testing.T) {
	reader, err := os.Open("customer.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(reader)
	customer := &Customer{}
	decoder.Decode(customer)
	fmt.Println(customer)
}

func TestStreamEncoder(t *testing.T) {
	writter, err := os.Create("customer_out.json")
	if err != nil {
		panic(err)
	}
	encoder := json.NewEncoder(writter)
	customer := Customer{
		FirstName:  "FirstName nichhh",
		MiddleName: "Middle name nichh",
		LastName:   "lastname nichhh",
	}
	encoder.Encode(customer)
}
