package main

import (
	"os"

	service "github.com/nukesz/pipeline/service"

	"github.com/olekukonko/tablewriter"
)

func main() {

	data := [][]string{
		[]string{"Alfred", "15", "10/20", "(10.32, 56.21, 30.25)"},
		[]string{"Beelzebub", "30", "30/50", "(1,1,1)"},
		[]string{"Hortense", "21", "80/80", "(1,1,1)"},
		[]string{"Pokey", "8", "30/40", "(1,1,1)"},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"NPC", "Speed", "Power", "Location"})
	table.AppendBulk(data)
	table.Render()

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	s := service.NewServer()
	s.Run(":" + port)
}
