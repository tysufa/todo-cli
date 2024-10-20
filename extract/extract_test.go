package extract

import "testing"

func TestReadCsv(t *testing.T){
  data := ReadCsv("./todo_data.csv")

  var res = [2][4]string{
    {"ID", "Description", "CreatedAt", "IsComplete"},
    {"1", "My new task", "2024-07-27T16:45:19-05:00", "true"},
  }

  for i, ligne := range res{
    for j, val := range ligne{
      if val != data[i][j]{
        t.Fatalf("data[%d][%d] is '%v', should be '%v' instead",i,j,data[i][j],val)
      }
    }
  }

}
