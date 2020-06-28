package quietly

import "os"

func ExampleClose() {
	file, err := os.Open("/path/to/file")
	if err != nil {
		panic(err)
	}
	defer Close(file)
	// TODO: Read the file.
}
