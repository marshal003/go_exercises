package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

//DataFile DS for DataFile which provides various helper methods to read
type DataFile struct {
	Filepath  string
	Separator string
}

// AsyncDataFileReader Reads file line by line and then pass it into the channel.
// Channel will be closed after the completion at EOF.
func (df *DataFile) AsyncDataFileReader(channel chan []string) {
	file, err := os.Open(df.Filepath)
	if err != nil {
		log.Println(err)
		panic("Unable to open file")
	}
	defer func() {
		file.Close()
		close(channel)
	}()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		channel <- strings.Split(scanner.Text(), df.Separator)
	}
}
