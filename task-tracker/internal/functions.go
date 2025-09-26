package internal

import (
	"encoding/json"
	"log"
	"os"
	"path"
)

func getFilePath() string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Println("Cannot access current working directory")
		return ""
	}

	return path.Join(cwd, "tasks.json")
}

func ReadTasksFromFile() ([]Task, error) {
	filepath := getFilePath()
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		file, err := os.Create(filepath)
		os.WriteFile(filepath, []byte("[]"), 0644)

		if err != nil {
			log.Println("Cannot create a file")
			return nil, err
		}

		defer func() {
			if err := file.Close(); err != nil {
				log.Println("Cannot close the file")
			}
		}()

		return []Task{}, nil
	}

	file, err := os.Open(filepath)
	if err != nil {
		log.Println("Cannot open the file")
		return nil, err
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Println("Cannot close the file")
		}
	}()

	tasks := []Task{}
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		log.Println("Cannot decode the file")
		return nil, err
	}

	return tasks, nil
}

func WriteTasksToFile(tasks []Task) error {
	filepath := getFilePath()
	file, err := os.Create(filepath)
	if err != nil {
		log.Println("Cannot recreate a file")
		return err
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Println("Cannot close the file")
		}
	}()

	err = json.NewEncoder(file).Encode(tasks)
	if err != nil {
		log.Println("Cannot encode the file")
		return err
	}

	return nil
}
