package complete

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/todo_app/extract"
)

func Main(id string){
  tasks := extract.ReadCsv("test.csv")

  f, err := os.OpenFile("test.csv", os.O_WRONLY, 0644)
  defer f.Close()

  if err != nil {
    panic(err)
  }

  idExists := false

  for i, line := range tasks{
    if line[0] == id{
      if tasks[i][3] == "true"{
        fmt.Printf("Attention, la tache d'id %s est déjà validé\n", id)
      }
      tasks[i][3] = "true"
      idExists = true
    }
  }
  
  if !idExists{
    fmt.Printf("Attention, la tache d'id %s n'existe pas\n", id)
  }

  w := csv.NewWriter(f)
  w.WriteAll(tasks)
  w.Flush()

}
