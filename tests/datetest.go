package tests

import "github.com/robfig/revel"
import "github.com/pvdvreede/dateable/app/models"
import "encoding/json"
import "log"

type DateTest struct {
	revel.TestSuite
}

func (t *DateTest) GetDatesResponse() models.Dates {
	var dates models.Dates
	err := json.Unmarshal(t.ResponseBody, &dates)
	if err != nil {
		log.Fatal(err)
	}
	return dates
}

func (t *DateTest) TestIndexContentType() {
	t.Get("/date/2013-03-04")
	t.AssertOk()
	t.AssertContentType("application/json")
}

func (t *DateTest) TestIndexWithWeekday() {
	t.Get("/date/2013-03-04")
	t.AssertOk()
	t.AssertContains("\"Weekend\": false")
}

func (t *DateTest) TestIndexWithWeekend() {
	t.Get("/date/2013-03-02")
	t.AssertOk()
	t.AssertContains("\"Weekend\": true")
}

func (t *DateTest) TestIndexWithWeekNumber() {
	t.Get("/date/2013-01-01")
	t.AssertOk()
	t.AssertContains("\"WeekNumber\": 1")
	t.Get("/date/2013-12-28")
	t.AssertOk()
	t.AssertContains("\"WeekNumber\": 52")
	t.Get("/date/2012-12-28")
	t.AssertOk()
	t.AssertContains("\"WeekNumber\": 52")
	t.Get("/date/2013-06-20")
	t.AssertOk()
	t.AssertContains("\"WeekNumber\": 25")
}

func (t *DateTest) TestIndexWithMonday() {
	t.Get("/date/2013-09-16")
	t.AssertOk()
	t.AssertContains("\"Day\": \"Monday\"")
}

func (t *DateTest) TestIndexWithSaturday() {
	t.Get("/date/2013-07-27")
	t.AssertOk()
	t.AssertContains("\"Day\": \"Saturday\"")
}
