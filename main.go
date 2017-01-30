package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/entropyx/tools"
)

func main() {
	data := [][]string{
		[]string{"X1", "X2", "y"},
		[]string{"21", "1", "1"},
		[]string{"35", "1", "0"},
		[]string{"19", "0", "1"},
		[]string{"20", "0", "0"},
		[]string{"33", "1", "1"},
	}
	path, _ := os.Getwd()
	tools.Save(data, path+"/data.csv")
	file := "file='" + path + "/data.csv'"
	predicter := "y"
	out, _ := exec.Command("Rscript", "test.R", file, predicter).Output()
	outsplit := strings.Split(fmt.Sprintf("%s", out), " /")

	aux := strings.Split(outsplit[0], " ")

	coef := make(map[string]float64, len(aux))
	for _, c := range aux {
		k := strings.Split(c, ":")
		coef[k[0]], _ = strconv.ParseFloat(k[1], 64)
	}

	fmt.Println(coef)
}
