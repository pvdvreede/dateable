package tests

import "github.com/robfig/revel"
import "github.com/pvdvreede/dateable/app/models"
import "encoding/json"
import "log"
import "fmt"

type DatesTest struct {
	revel.TestSuite
}

func (t *DatesTest) GetDatesResponse() models.Dates {
	var dates models.Dates
	err := json.Unmarshal(t.ResponseBody, &dates)
	if err != nil {
		log.Fatal(err)
	}
	return dates
}

type DateRangeTest struct {
	From        string
	To          string
	Length      int
	Weekdays    int
	WeekendDays int
}

func (t *DatesTest) TestIndexContentType() {
	t.Get("/dates/2013-03-04/2013-05-06")
	t.AssertOk()
	t.AssertContentType("application/json")
}

func (t *DatesTest) TestIndexCorrectDates() {
	toRun := []DateRangeTest{
		DateRangeTest{"2013-05-01", "2013-05-01", 1, 1, 0},
		DateRangeTest{"2013-05-01", "2013-05-02", 2, 2, 0},
		DateRangeTest{"2013-05-01", "2013-05-10", 10, 8, 2},
		DateRangeTest{"2013-05-01", "2013-06-02", 33, 23, 10},
	}

	for _, v := range toRun {
		t.Get(fmt.Sprintf("/dates/%v/%v", v.From, v.To))
		t.AssertOk()
		t.AssertEqual(v.Length, len(t.GetDatesResponse().Results))
		t.AssertEqual(v.Length, t.GetDatesResponse().TotalDays)
		t.AssertEqual(v.Weekdays, t.GetDatesResponse().Weekdays)
		t.AssertEqual(v.WeekendDays, t.GetDatesResponse().WeekendDays)
	}
}
