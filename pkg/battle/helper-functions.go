package battle

import (
	"encoding/json"
	"io/fs"
	"log"
	"path/filepath"

	"github.com/gorilla/websocket"
)

func ConvertToSlice(listInput [][]string) []Item {
	var itemsList []Item

	for _, itemInput := range listInput {
		itemsList = append(itemsList, Item{Name: itemInput[0], Score: 0})
	}

	return itemsList
}

func SendBattleOptions(conn *websocket.Conn) {
	fileNames := GetFileNames()

	jsonData, err := json.Marshal(fileNames)
	if err != nil {
		log.Println("Error marshalling json:", err)
		return
	}
	err = conn.WriteMessage(websocket.TextMessage, jsonData)
	if err != nil {
		log.Println("Error sending message:", err)
	}
}

func GetFileNames() []string {
	dirPath := "game-data"
	var files []string
	err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			files = append(files, d.Name())
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return files
}
