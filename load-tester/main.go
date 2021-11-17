package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	url_load := "http://squidgame.35.184.172.30.nip.io"
	args := os.Args[1:]
	fmt.Printf("Arguments without program name: %v\n", args)

	//Obtiene los nombres de los juegos por jugar
	game_names := strings.Split(args[2], "|")
	fmt.Println(game_names)

	players, _ := strconv.Atoi(args[4])
	fmt.Println(players)

	rungames, _ := strconv.Atoi(args[6])
	fmt.Println(rungames)

	concurrence := args[8]
	fmt.Println(concurrence)

	timeout, _ := strconv.Atoi(strings.Replace(args[10], "m", "", -1))
	fmt.Println(timeout)

	t1 := time.Now()
	for rungames > 0 {
		//Se verifica que no se haya alcanzado el tiempo limite para enviar peticiones
		t2 := time.Now()
		diff := t2.Sub(t1)
		if diff.Minutes() > float64(timeout) {
			break
		}
		//Se genera un numero aleatorio para el numero de jugadores
		random_players := rand.Intn(players)
		if random_players == 0 {
			random_players = 1
		}
		random_game := strings.Split(game_names[rand.Intn(len(game_names))], ";")
		url_params := "/game/" + random_game[0] + "/gamename/" + random_game[1] + "/players/" + strconv.Itoa(random_players)

		url := url_load + url_params
		fmt.Println(url)
		/*_, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
		if err != nil {
			log.Fatalln(err)
		}*/

		rungames--
	}

}
