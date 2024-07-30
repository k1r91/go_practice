package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// начало решения

// ConvertEmployees преобразует XML-документ с информацией об организации
// в плоский CSV-документ с информацией о сотрудниках
func ConvertEmployees(outCSV io.Writer, inXML io.Reader) error {
	reader := xml.NewDecoder(inXML)
	writer := csv.NewWriter(outCSV)
	for {
		token, err := reader.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		switch element := token.(type) {
		case xml.StartElement:
			if element.Name.Local == "department" {
				err := startWriteDepartment(reader, writer)
				if err !=  nil {
					return err
				}
			}
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return err
	}
	return nil
}

type Code string

type Employee struct {
	Id int `xml:"id,attr,omitempty"`
	Name string `xml:"name"`
	City string `xml:"city"`
	Salary int `xml:"salary"`
	Code Code
}

func (emp *Employee) Slice() []string {
	return []string{strconv.Itoa(emp.Id), emp.Name, emp.City, string(emp.Code), strconv.Itoa(emp.Salary)}
}

func (emp *Employee) String() string {
	return fmt.Sprintf("%d,%s,%s,%s,%d\n", emp.Id, emp.Name, emp.City, emp.Code, emp.Salary)
}

func startWriteDepartment(reader *xml.Decoder, writer *csv.Writer) error {
	var code Code
	var employees []Employee
	department:
	for {
		token, err := reader.Token()
		if err != nil {
			return err
		}
		switch element := token.(type) {
		case xml.StartElement:
			switch element.Name.Local {
			case "code":
				err = reader.DecodeElement(&code, &element)
				if err != nil {
					return err
				}
			case "employee":
				var employee Employee
				err = reader.DecodeElement(&employee, &element)
				if err != nil {
					return err
				}
				employees = append(employees, employee)
			}
		case xml.EndElement:
			if element.Name.Local == "department" {
				break department
			}
		}
	}
	for i := range employees {
		employees[i].Code = code
		err := writer.Write(employees[i].Slice())
		if err != nil {
			return err
		}
	}
	return nil
}

// конец решения

func main() {
	fmt.Println("start")
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
	src = `hello`

	in := strings.NewReader(src)
	out := os.Stdout
	ConvertEmployees(out, in)
	/*
		id,name,city,department,salary
		11,Дарья,Самара,hr,70
		12,Борис,Самара,hr,78
		21,Елена,Самара,it,84
	*/
}
