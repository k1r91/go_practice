package main

import (
	"encoding/xml"
	"encoding/csv"
	"strconv"
	"fmt"
	"io"
	"os"
	"strings"
	"errors"
)

// начало решения

// ConvertEmployees преобразует XML-документ с информацией об организации
// в плоский CSV-документ с информацией о сотрудниках
type Organization struct {
	Departments []*Department `xml:"department"`
}

type Department struct {
	Code string `xml:"code"`
	Employees []*Employee `xml:"employees>employee"`
}

type Employee struct {
	Id int `xml:"id,attr,omitempty"`
	Name string `xml:"name"`
	City string `xml:"city"`
	Salary int `xml:"salary"`
	Code string
}

func (emp *Employee) String() string {
	return fmt.Sprintf("%d,%s,%s,%s,%d\n", emp.Id, emp.Name, emp.City, emp.Code, emp.Salary)
}

func (emp *Employee) Slice() []string {
	return []string{strconv.Itoa(emp.Id), emp.Name, emp.City, string(emp.Code), strconv.Itoa(emp.Salary)}
}

func ConvertEmployees(outCSV io.Writer, inXML io.Reader) error {
	var org Organization
	dec := xml.NewDecoder(inXML)
	token, err := dec.Token()
	if err != nil {
		return err
	}
	switch element := token.(type) {
	case xml.StartElement:
		err := dec.DecodeElement(&org, &element)
		if err != nil {
			return err
		}
	default:
		return errors.New("invalid xml format")
	}
	writer := csv.NewWriter(outCSV)
	writer.Write([]string{"id", "name", "city", "department", "salary"})
	if err := writer.Error(); err != nil {
		return err
	}
	for _, d := range org.Departments {
		for _, e := range d.Employees {
			employee := Employee{e.Id, e.Name, e.City, e.Salary, d.Code}
			writer.Write(employee.Slice())
			if err := writer.Error(); err != nil {
				return err
			}
		}
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		return err
	}
	return nil
}

// конец решения

func main() {
	src := `<organization>
    <department>
        <code>hr</code>
        <employees>
            <employee id="11">
                <name>Дарья</name>
                <city>Самара</city>
                <salary>70</salary>
            </employee>
            <employee id="12">
                <name>Борис</name>
                <city>Самара</city>
                <salary>78</salary>
            </employee>
        </employees>
    </department>
    <department>
        <code>it</code>
        <employees>
            <employee id="21">
                <name>Елена</name>
                <city>Самара</city>
                <salary>84</salary>
            </employee>
        </employees>
    </department>
</organization>`
	in := strings.NewReader(src)
	out := os.Stdout
	err := ConvertEmployees(out, in)
	fmt.Println(err)
	/*
		id,name,city,department,salary
		11,Дарья,Самара,hr,70
		12,Борис,Самара,hr,78
		21,Елена,Самара,it,84
	*/
}
