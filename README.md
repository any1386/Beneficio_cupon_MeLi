# Beneficio_Cupon_MeLi
Instalacion-> se monta el proyecto en 
API que dado una lista de item_id y el monto total, retorne la lista de items que maximice el total gastado sin excederlo.
Este proyecto implementa un API que dado una lista de item_id y el monto total, retorne la lista de items que maximice el total gastado sin excederlo.

# Se sube el proyecto a github
# Se crea proyecto Google App Engine y se Hostea la API
Se ingresa a Google App Engine:
Se crea el proyecto en Google App Engine 
desde consola se da el comando:
git clone \
     https://github.com/any1386/Beneficio_cupon_MeLi

y paso seguido se instala la libreria go get github.com/gorilla/mux

y se prueba con el comando go run .
Paso seguido se crea e implementa en App Engine la aplicación, para implementar la app, se debe crear una región con el comando:
gcloud app create

Se ingresa el comando para implementar la app:
gcloud app deploy

LUEGO SE HABILITA LA APP EN LA URL
https://console.developers.google.com/apis/library/cloudbuild.googleapis.com?project=beneficiocuponmeli

Y se da el siguiente comando para obtener la url
 gcloud app browse

https://beneficiocuponmeli.rj.r.appspot.com

# Realizar pruebas Usando la Herramienta Soap UI
la siguiente es la url obtenida en el host de Google App Engine
https://beneficiocuponmeli.rj.r.appspot.com
Por metodo Post con el EndPoint https://beneficiocuponmeli.rj.r.appspot.com
con el Request:
{
	"valorCupon":"100000",
	"idItem":["MLA816019440","MLA811601010","MLA810645375","MLA805330648"]
}
Se obtiene como resultado:
{
   "resultadoOperacion": "Operacion exitosa!!",
   "listaOptima":    [
      "MLA805330648",
      "MLA810645375",
      "MLA811601010"
   ]
}
Se adjunta en github BeneficioCuponMeli-soapui.xml
con request y con TestSuite
