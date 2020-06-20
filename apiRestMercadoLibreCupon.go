package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"github.com/gorilla/mux"
	"io/ioutil"	
)
type CuponMeLi struct{
	ValorCupon 	string `json:"valorCupon"`
	IdItem 		[]string	`json:"idItem"` 
}

type ResultadoMeLi struct{
	ResultadoOperacion 	string `json:"resultadoOperacion"`
	ListaOptima 		[]string	`json:"listaOptima"` 
}

type mytype []map[string]string

var listaValores []float64
var listakeys []string
func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/calculateService",CalculateService).Methods("POST")
	log.Println("Escuchando por el puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func CalculateService(w http.ResponseWriter, r *http.Request) {
	var cupon CuponMeLi
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&cupon)
	log.Println("Cupon: ", cupon)
	if err != nil{
		http.Error(w, "Mal Request", http.StatusBadRequest)
		return
	}	
	var items map[string]float64
	items = make(map[string]float64)
	listaValores = make([]float64, len(cupon.IdItem))
	listakeys = make([]string, len(cupon.IdItem))
	var value = 1000.0
	for i := len(cupon.IdItem) - 1; i >= 0; i-- {
		value = obtenerValorServicio(cupon.IdItem[i])
		items[cupon.IdItem[i]]=value
		listaValores[i]=value
		listakeys[i]=cupon.IdItem[i]
	}
	amount, _ := strconv.ParseFloat(cupon.ValorCupon, 2)	
	listakeys=calculate(items,amount)
	if len(listakeys)>0 {
		resultado:=ResultadoMeLi{
			ResultadoOperacion 	:"Operacion exitosa!!",
			ListaOptima : listakeys,
		}
		json.NewEncoder(w).Encode(resultado)
	}else {
		resultado:=ResultadoMeLi{
			ResultadoOperacion 	:"Operacion fallida!!",
			ListaOptima : listakeys,
		}
		json.NewEncoder(w).Encode(resultado)
	}	
}

func obtenerValorServicio(item string)float64{
	resp, err := http.Get("https://api.mercadolibre.com/items/"+item)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	jsonMap := make(map[string]float64)
	json.Unmarshal(contents, &jsonMap)
	value, _ := jsonMap["price"]	
	return value
}

func calculate(items map[string]float64, amount float64) []string {
	ordenarlista()	
	log.Println("items: ", listaValores)
	log.Println("keys: ", listakeys)
	log.Println("amount: ", amount)
	var stringOptima =""
	var listaOptima []string
	var suma=0.0
	var itemAnterior=""
	for x := 0; x < len(listaValores); x++{
		log.Println(itemAnterior+" - "+listakeys[x])
		if itemAnterior != listakeys[x]{
			itemAnterior=listakeys[x]
			suma+=listaValores[x]
			log.Println("suma: ", suma)
			if suma > amount{
				break	
			}
			stringOptima+=listakeys[x]+","
		}
	}
	if len(stringOptima) > 0{
		listaOptima = strings.Split(stringOptima[:len(stringOptima)-1], ",")
	}	
	log.Println("optima: ", stringOptima)
    return listaOptima
}

func ordenarlista(){
	tmp := 0.0
	tmp2 := ""
    for x := 0; x < len(listaValores); x++ {
        for y := 0; y < len(listaValores); y++ {
            if listaValores[x] < listaValores[y] {
                tmp = listaValores[x]
                listaValores[x] = listaValores[y]
				listaValores[y] = tmp
				tmp2 = listakeys[x]
                listakeys[x] = listakeys[y]
				listakeys[y] = tmp2
				
            }
        }
	}
}
