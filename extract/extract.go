package extract

import(
  "os"
  "encoding/csv"
  "strings"
	"time"
  "strconv"
)

func ReadCsv(filePath string) [][]string{
  data, err := os.ReadFile(filePath)
  if err != nil{
    panic(err)
  }

	r := csv.NewReader(strings.NewReader(string(data)))

	records, err := r.ReadAll()
	if err != nil {
    panic(err)
	}

  return records
}

func StrToTime(date string) time.Time{
  splitDate := strings.Split(date, " ") // sépare la date et l'heure
  ddate := strings.Split(splitDate[0], "-") // sépare la partie aaaa-mm-jj de la date
  // récupère individuellement le jour mois et annee
  annee, _ := strconv.Atoi(ddate[0])
  mois, _ := strconv.Atoi(ddate[1])
  jour, _ := strconv.Atoi(ddate[2])

  //idem ici pour l'heure
  t := strings.Split(splitDate[1], ":")
  hour, _ := strconv.Atoi(t[0])
  minutes, _ := strconv.Atoi(t[1])
  seconds, _ := strconv.Atoi(t[2])

  location := time.FixedZone("UTC+2", 2*60*60) // renvoie une localisation de UTC+2 (correspond à l'heure de paris)

  res := time.Date(annee, time.Month(mois), jour, hour, minutes, seconds, 0, location) // créer une date a partir des informations précédentes

  return res
}
