package add

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"time"
)

func changeId(filename string) string{
  curID, err := os.ReadFile(filename)
  if err != nil {
    panic(err)
  }

  nextID, err := strconv.Atoi(string(curID[:len(curID)-1])) // on récupère l'identifiant sous forme de nb on retire le dernier élément car c'est le byte \n
  if err != nil {
    panic(err)
  }
  nextID+=1

  idFile, err := os.OpenFile("currentID.txt", os.O_WRONLY, 0644)
  defer idFile.Close()
  if err != nil {
    panic(err)
  }

  idFile.WriteString(strconv.Itoa(nextID)) // on incremente l'id dans le fichier

  return strconv.Itoa(nextID)
} 

func writeTask(task string){

  nextID := changeId("currentID.txt")

  currentTime := time.Now()

  var tasks [][]string

  var f *os.File // on déclare un type file pour pouvoir l'utilser sans redéfinir f dans chaque if
  defer f.Close()


  if _, err := os.Stat("test.csv"); err == nil { // si le fichier existe on l'ouvre en mode append
    f, err = os.OpenFile("test.csv", os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
      panic(err)
    }
  } else if errors.Is(err, os.ErrNotExist) {
    f, err = os.OpenFile("test.csv", os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
      panic(err)
    }

    tasks = append(tasks, []string{"Id", "Description", "Created", "Completed"}) // si on créer le fichier on ajoute l'en tête
  } else {
    panic(err)
  }

  tasks = append(tasks, []string{nextID, task, currentTime.Format("2006-01-02 15:04:05"), "false"})

  w := csv.NewWriter(f)
  w.WriteAll(tasks)
}

func Main(args []string){
  writeTask(args[0])
}
