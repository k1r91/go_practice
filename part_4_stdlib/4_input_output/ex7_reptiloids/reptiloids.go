package main

import (
	"bufio"
	"fmt"
	mathrand "math/rand"
	"os"
	"path/filepath"
)

// алфавит планеты Нибиру
const alphabet = "aeiourtnsl"

// Census реализует перепись населения.
// Записи о рептилоидах хранятся в каталоге census, в отдельных файлах,
// по одному файлу на каждую букву алфавита.
// В каждом файле перечислены рептилоиды, чьи имена начинаются
// на соответствующую букву, по одному рептилоиду на строку.
type Census struct{
	total int
}

// Count возвращает общее количество переписанных рептилоидов.
func (c *Census) Count() int {
	return c.total
}

// Add записывает сведения о рептилоиде.
func (c *Census) Add(name string) {
	c.total += 1
	fileName := filepath.Join("census", fmt.Sprintf("%s.txt", string(name[0])))
	file, err := os.OpenFile(fileName, os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(file)
	writer.WriteString(fmt.Sprintf("%s\n", name))
	err = writer.Flush()
	if err != nil {
		panic(err)
	}
	err = file.Close()
	if err != nil {
		panic(err)
	}
}

// Close закрывает файлы, использованные переписью.
func (c *Census) Close() {
}

// NewCensus создает новую перепись и пустые файлы
// для будущих записей о населении.
func NewCensus() *Census {
	for _, c := range alphabet {
		fileName := filepath.Join("census", fmt.Sprintf("%s.txt", string(c)))
		os.Remove(fileName)
		f, err := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0666)
		if err != nil {
			panic(err)
		}
		err = f.Close()
		if err != nil {
			panic(err)
		}
	}
	return &Census{}
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

var rand = mathrand.New(mathrand.NewSource(0))

// randomName возвращает имя очередного рептилоида.
func randomName(n int) string {
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(chars)
}

func main() {
	census := NewCensus()
	defer census.Close()
	for i := 0; i < 1024; i++ {
		reptoid := randomName(5)
		census.Add(reptoid)
	}
	fmt.Println(census.Count())
}
