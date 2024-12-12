package battle

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/demartinom/list-ranker-web/pkg/global"
	"github.com/gorilla/websocket"
)

func SendBattleOptions(conn *websocket.Conn) {
	fileNames := GetFileNames()
	jsonData, err := json.Marshal(PremadeLists{"List Options", fileNames})
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
			fileName := d.Name()
			files = append(files, fileName[:len(fileName)-4])
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return files
}

// Readies CSV file to be used in game
func ReadCSV(fileName string) []*global.Item {
	filePath := fmt.Sprintf("game-data/%s.csv", fileName)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	reader := csv.NewReader(file)
	//Skip over header line
	reader.Read()

	listItems, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var itemsList []*global.Item

	for _, itemInput := range listItems {
		itemsList = append(itemsList, &global.Item{Name: itemInput[0], Score: 0})
	}

	return itemsList
}

// Readies User lisst to be used in game
func ReadCustom(customList []string) []*global.Item {
	var itemsList []*global.Item

	for _, itemInput := range customList {
		itemsList = append(itemsList, &global.Item{Name: itemInput, Score: 0})
	}

	return itemsList
}

func sendCombatants(battlers []*global.Item, conn *websocket.Conn) {
	jsonData, err := json.Marshal(CombatantsList{"Combatants", battlers})
	if err != nil {
		log.Println("Error marshalling json:", err)
		return
	}
	err = conn.WriteMessage(websocket.TextMessage, jsonData)
	if err != nil {
		log.Println("Error sending message:", err)
	}
}
