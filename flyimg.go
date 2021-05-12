package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	html,err:=template.ParseFiles("index.html")
	check(err)

	err=html.Execute(writer,nil)
	check(err)
}

func dataHandler(writer http.ResponseWriter, request *http.Request) {
	add:=os.Getenv("DOMAIN")
	link:=request.FormValue("url")
	quality:=request.FormValue("quality")
	format:=request.FormValue("format")
	width:=request.FormValue("width")
	height:=request.FormValue("height")
	http.Redirect(writer,request,"http://"+add+"/upload/"+"w_"+width+","+"h_"+height+","+"q_"+quality+","+"o_"+format+","+"rf_1"+"/"+link,http.StatusFound)
}


func main() {

	http.HandleFunc("/index",viewHandler)
	http.HandleFunc("/index/data",dataHandler)

	err:=http.ListenAndServe(":2222",nil)
	log.Fatal(err)

}
