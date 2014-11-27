package util

import(
    "time"
    //"fmt"
)

const DateFormat = "2006-01-02"

var DefaultDate time.Time = time.Time{}
var Days = [13]int{-1, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

//the @label is like yyyyMM
func ParseDate(label string) time.Time {
    t := DefaultDate
    if len(label) == 0 {
        return t
    }

    ystr := label[:4]
    mstr := label[4:]

    year := ToInt(ystr)
    month := ToInt(mstr)
    day := LastDay(year, month)
    
    if day < 1 {
        NewLog().Error("Cannot parse the date: ", label)
    }
    
    //fmt.Println("Before:", t, year, month, day)
    t = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
    //fmt.Println("After: ", t)

    return t
}

func IsLeapYear(year int) bool {
    return (year % 400) == 0 || ( (year % 100) != 0 && (year % 4) == 0) 
}

func LastDay(year, month int) int {
    day := Days[month]
    if IsLeapYear(year) && month == 2 {
        day += 1
    }

    return day
}

func FormatDate(date time.Time) string {
	return date.Format(DateFormat)
}
