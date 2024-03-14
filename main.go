package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {

	usage := `
    _________________               _________________
     ~-.              \  |\___/|  /              .-~
         ~-.           \ / o o \ /           .-~
            >           \\  W  //           <
           /             /~---~\             \
          /_            |       |            _\
             ~-.        |       |        .-~
                ;        \     /        i
               /___      /\   /\      ___\
                    ~-. /  \_/  \ .-~
                       V         V
                    By Batuhan Çelik
                	Dirbuster GO tool
				   `
	fmt.Println(usage)

	fc := flag.Int("fc", 400, "Dönen istek filtrelemek için")
	fs := flag.Int("fs", 0, "İstek gönderme hızını ayarlamak için")
	post := flag.Bool("post", false, "Post isteği göndermek için")
	u := flag.String("u", "", "Url eklemek için")
	w := flag.String("w", "", "wordlist eklemek için")

	flag.Parse()

	if *w == "" {
		usage := `                    Dosayaya erişilemiyor veya dosya boş.`
		fmt.Println(usage)
		os.Exit(1)

	}

	speed := strconv.Itoa(*fs)

	if u != nil && *post == false {

		usage := `                      METHOD: GET
                       Speed: ` + speed + `sn
                  Status Filter: ` + fmt.Sprint(*fc) + `
												   `

		fmt.Println(usage)

		file, _ := os.OpenFile(*w, os.O_RDONLY, 0755)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {

			url := *u + scanner.Text()
			resp, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			if resp.StatusCode != *fc {
				fmt.Println("url:", *u+scanner.Text()+" statuscode: "+resp.Status)

			}
			time.Sleep(time.Duration(*fs) * time.Second)
		}

	}

	if u != nil && *post == true {
		usage := `                      METHOD: POST
		       Speed: ` + speed + `sn
		    Status Filter: ` + fmt.Sprint(*fc) + `    
					  
					  `

		fmt.Println(usage)

		file, _ := os.OpenFile(*w, os.O_RDONLY, 0755)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			url := *u + scanner.Text()
			resp, err := http.Post(url, "application/x-www-form-urlencoded", nil)
			if err != nil {
				panic(err)
			}
			if resp.StatusCode != *fc {
				fmt.Println("url:", *u+scanner.Text()+" statuscode: "+resp.Status)
			}
			time.Sleep(time.Duration(*fs) * time.Second)

		}

	}

}
