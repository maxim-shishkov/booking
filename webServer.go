package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func startServer() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/user/", userPage)
	http.HandleFunc("/bookingDetails/", bookingDetailsPage)
	http.HandleFunc("/booking/", bookingPage)
	http.HandleFunc("/addRoom/", RoomPage)
	http.HandleFunc("/save", addRoomPage)
	http.HandleFunc("/add", addBookingPage)
	http.HandleFunc("/delete/", deletePage)
	port := ":8080"
	println("Server listen on port:",port)
	err := http.ListenAndServe(port,nil)
	if err != nil {
		log.Fatal("ListenAndServe",err)
	}
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	tmpl,err := template.ParseFiles("static/index.html", "static/header.html", "static/footer.html")
	if err != nil {http.Error(w, err.Error(),400)
		return
	}
	err = tmpl.Execute(w,rooms)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func userPage(w http.ResponseWriter, r *http.Request) {
	tmpl,err := template.ParseFiles("static/user.html", "static/header.html", "static/footer.html")
	if err != nil {http.Error(w, err.Error(),400)
		return
	}
	if err := tmpl.Execute(w,rooms);	err != nil {http.Error(w, err.Error(),400)
		return
	}
}

func bookingDetailsPage(w http.ResponseWriter, r *http.Request) {
	tmpl,err := template.ParseFiles("static/bookingDetails.html", "static/header.html", "static/footer.html")
	if err != nil {http.Error(w, err.Error(),400)
		return
	}
	var id, _ =   strconv.Atoi(strings.Replace(r.URL.Path,"/bookingDetails/","",1))
	if err := tmpl.Execute(w, getBooking(id) );	err != nil {http.Error(w, err.Error(),400)
		return
	}
}

func bookingPage(w http.ResponseWriter, r *http.Request) {
	tmpl,err := template.ParseFiles("static/booking.html", "static/header.html", "static/footer.html")
	if err != nil {http.Error(w, err.Error(),400)
		return
	}
	id, _ :=  strconv.Atoi(strings.Replace(r.URL.Path,"/booking/","",1))

	err = tmpl.Execute(w, (id) )
	if err != nil {
		log.Fatal(err)
		return
	}
}

func deletePage(w http.ResponseWriter, r *http.Request) {
		id, _ :=  strconv.Atoi(strings.Replace(r.URL.Path,"/delete/","",1))
		deleteRoom(id)
		http.Redirect(w, r, "/", 302)
}

func RoomPage(w http.ResponseWriter, r *http.Request) {
	tmpl,err := template.ParseFiles("static/addRoom.html", "static/header.html", "static/footer.html")
	if err != nil {http.Error(w, err.Error(),400)
		return
	}
	err = tmpl.Execute(w,nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func addRoomPage(w http.ResponseWriter, r *http.Request)  {
	name := r.FormValue("name")
	if len(name)==0 { {}
		errorPage(w,r, []string{"empty field"} )
	} else {
		addRoom(name)
		http.Redirect(w, r, "/", 302)
	}
}

func errorPage(w http.ResponseWriter, r *http.Request, errorMessage []string)  {
	tmpl,err := template.ParseFiles("static/error.html", "static/header.html", "static/footer.html")
	if err != nil {http.Error(w, err.Error(),400)
		return
	}
	err = tmpl.Execute(w,errorMessage)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func addBookingPage(w http.ResponseWriter, r *http.Request)  {
	id, _ := strconv.Atoi( r.FormValue("id") )
	name := r.FormValue("name")
	date := r.FormValue("date")
	hour := r.FormValue("hour")
	minut := r.FormValue("minut")
	b,ar := addBooking(id,name,date,hour,minut)
	if b {
		http.Redirect(w, r, "/", 302)
	} else {
		errorPage(w,r, ar )
	}
}