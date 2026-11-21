package file

import "os"

func WriteFile(path string, content string) error {
	data := []byte(content) // Convert the string content to a byte slice

	err := os.WriteFile(path, data, 0644) //WriteFile function writes the byte slice to the specified file path with the given permissions//

	//0644 means read and write permissions for the owner, and read-only permissions for the group and others (0644)//

	if err != nil {
		return err // If there is an error while writing the file, return the error//
	}
	return nil // If the file is written successfully, return nil indicating no error//
}

