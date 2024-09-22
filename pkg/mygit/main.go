package mygit

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
)

func CatFile(args []string) error {
	// fmt.Println("args recieved: ", args)
	
	if args[0] != "-p" {
		return fmt.Errorf("unknown flag")
	}

	objDir := ".git/objects"
	objSHA := args[1]

	entries, err := os.ReadDir(objDir + "/" + objSHA[:2])

	if err != nil {
		return fmt.Errorf("could not find git object with SHA-1: %s", objSHA)
	}

	for _, e := range entries {
		if e.Name() != objSHA[2:] {
			continue
		}

		file, err := os.Open(objDir + "/" + objSHA[:2] + "/" + objSHA[2:])

		if err != nil {
			return fmt.Errorf("error opening file %s", err.Error())
		}

		rc, err := zlib.NewReader(file)

		if err != nil {
			return fmt.Errorf("error decompressing %s", err.Error())
		}

		content, err := io.ReadAll(rc)

		if err != nil {
			return fmt.Errorf("error reading content: %s", err.Error())
		}

		parts := bytes.SplitN(content, []byte{0}, 2)

		if len(parts) != 2 {
			return fmt.Errorf("invalid blob format")
		}

		_, err = os.Stdout.Write(parts[1])

		if err != nil {
			return fmt.Errorf("error writing content: %s", err)

		}

		rc.Close()
	}

	return nil
}

func HashFile(args []string) (string, error) {

	return "", nil
}
