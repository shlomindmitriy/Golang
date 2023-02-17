package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
	добавить 3 хендлера + - *
	/v1/caclulation/minus
	/v1/caclulation/plus
	/v1/caclulation/умн
	описать бизнес логику этих хендлеров
	проверить в инсомнии работу этих методов
*/

type request struct {
	Number1 int    `json:"number1"`
	Number2 int    `json:"number2"`
	Sumbol  string `json:"sumbol"`
}

type Respose struct {
	Sum int `json:"sum"`
}
type requestDivision struct {
	Number1 float64 `json"number1"`
	Number2 float64 `json"number2"`
	Sumbol  string  `json"sumbol"`
}
type ResposeDivision struct {
	Sum float64 `json:"sum"`
}

func main() {
	// http.HandleFunc("/v1/caclulation/minus", caclulationMinus) // each request calls handler
	// // http.HandleFunc("/", handler) // each request calls handler

	// http.HandleFunc("/v1/caclulation/plus", caclulationPlus)

	// http.HandleFunc("/v1/caclulation/multiplication", calculationMultiplication)

	// http.HandleFunc("/v1/caclulation/division", calculationDivision)
	http.HandleFunc("/v1/calculation/", calculation)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func calculation(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	req := request{}

	err = json.Unmarshal(b, &req)
	if err != nil {
		panic(err)
	}
	fmt.Println(req)

	resp := Respose{}

	switch req.Sumbol {
	case "+":
		resp.Sum = req.Number1 + req.Number2

	case "-":
		resp.Sum = req.Number1 - req.Number2

	case "*":
		resp.Sum = req.Number1 * req.Number2
	}

	buf, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}

	w.Write(buf)
}

// func caclulationMinus(w http.ResponseWriter, r *http.Request) {

// 	b, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		panic(err)
// 	}

// 	req := request{}

// 	err = json.Unmarshal(b, &req)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(req)

// 	resp := Respose{}

// 	resp.Sum = req.Number1 - req.Number2

// 	buf, err := json.Marshal(resp)
// 	if err != nil {
// 		panic(err)
// 	}
// 	w.Write(buf)

// }
// func caclulationPlus(w http.ResponseWriter, r *http.Request) {

// 	b, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// 	req := request{}

// 	err = json.Unmarshal(b, &req)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(req)

// 	resp := Respose{}

// 	resp.Sum = req.Number1 + req.Number2

// 	buf, err := json.Marshal(resp)
// 	if err != nil {
// 		panic(err)
// 	}
// 	w.Write(buf)
// }

// func calculationMultiplication(w http.ResponseWriter, r *http.Request) {

// 	b, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// 	req := request{}

// 	err = json.Unmarshal(b, &req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(req)

// 	resp := Respose{}

// 	resp.Sum = req.Number1 * req.Number2

// 	buf, err := json.Marshal(resp)

// 	if err != nil {
// 		panic(err)
// 	}

// 	w.Write(buf)

// }
// func calculationDivision(w http.ResponseWriter, r *http.Request) {

// 	b, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		panic(err)
// 	}

// 	req := requestDivision{}

// 	err = json.Unmarshal(b, &req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(req)
// 	resp := ResposeDivision{}

// 	resp.Sum = req.Number1 / req.Number2

// 	buf, err := json.Marshal(resp)
// 	if err != nil {
// 		panic(err)
// 	}
// 	w.Write(buf)
// }
