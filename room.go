package main

import (
	"sort"
	"strconv"
	"time"
)

type Room struct {
	Name string
	Booking []Bookings
}
type Bookings struct {
	Name string
	Times time.Time
}
var cns int
var rooms = make(map[int]Room)

func addRoom(name string)  {
	rooms[cns] = Room{name, make([]Bookings,0 )  }
	cns++
}
func deleteRoom(id int)  {
	delete(rooms,id)
}

func addBooking(id int,name string,date string,hour string,minute string) (bool,[]string) {
	err := make([]string,0)
	addZero := func(str string) string {
		if len(str)==1 {
			str="0" + str
		}
		return str
	}
	check := func(str string,min int,max int,multiplicity int) bool{
		i, _ := strconv.Atoi(str)
		return len(str)==0||i<min||i>max||i%multiplicity!=0
	}
	t,_ := time.Parse("2006-01-02 15:04:05",addZero(date)+" "+addZero(hour)+":"+addZero(minute)+":00")

	if check(minute,0,60,30) {
		err = append(err,"Minutes input error")
	}
	if check(hour,0,24,1) {
		err = append(err,"Hours input error")
	}
	if (len(date)==0)||t.Before( time.Now() )  {
		err = append(err,"Invalid date! (It should be equal to or greater than the current date)")
	}
	if checkTime(id,t) {
		err = append(err,"Invalid date! (is busy)")
	}
	if len(err) == 0{
		rooms[id] = Room{getRoomName(id),append(rooms[id].Booking,Bookings{name,t}) }
		sortTime(id)
	}
	return len(err) == 0, err
}

func getRoomName(id int) string {
	return rooms[id].Name
}

type print struct {
	Name string
	Date string
	Time string
	Expired bool
}

func getBooking(id int) ([]print) {

	ar2 := make([]print, len(rooms[id].Booking) )
	ar := make([]string, len(rooms[id].Booking) )
	for i, v := range rooms[id].Booking{
		ar[i] = v.Name +" " + v.Times.Format("02.01.2006 15:04")
		if v.Times.Before(time.Now()){ ar[i] += " expired"}

		ar2[i].Name = v.Name
		ar2[i].Time = v.Times.Format("15:04")
		ar2[i].Date = v.Times.Format("02.01.2006")
		ar2[i].Expired = v.Times.Before(time.Now())
	}
	return ar2
}

func sortTime(id int)  {
	ar := rooms[id].Booking
	sort.Slice(ar, func(i, j int) bool {
		return ar[i].Times.Before( ar[j].Times )
	})
}

func checkTime(id int, newTime time.Time) bool {
	for _, v := range rooms[id].Booking {
		if newTime.Equal(v.Times) {
		return true }
	}
	return false
}