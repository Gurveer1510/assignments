package main

func main() {
	// fileName := flag.String("f", "problems.csv", "specify the filename to use for the quiz. ")
	// limit := flag.Int("l", 10, "time limit in seconds for the quiz.")
	// flag.Parse()

	// dir, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal("could not get current directory")
	// }
	// path := dir + "/" + *fileName
	// file, err := os.Open(path)
	// if err != nil {
	// 	log.Println(err)
	// 	log.Fatalf("Could not open %v ", path)
	// }
	// defer file.Close()

	// scanner := bufio.NewScanner(file)
	// reader := bufio.NewReader(os.Stdin)
	// correct := 0
	// incorrect := 0

	// _ = time.AfterFunc(time.Duration(*limit)*time.Second, func() {

	// 	fmt.Println("\nTimer Over. :(")
	// 	fmt.Printf("You answered %v questions correctly.", correct)
	// 	fmt.Printf("\nYou answered %v questions incorrectly.\n", incorrect)
	// 	os.Exit(0)
	// })

	// fmt.Printf("You have %v seconds to answer the questions.\n", *limit)
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	r := csv.NewReader(strings.NewReader(line))
	// 	record, err := r.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	ques := record[0]
	// 	ans := record[1]

	// 	fmt.Printf("%v ?  - > ", ques)

	// 	txt, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		log.Println("Could not read the user's answer")
	// 		log.Fatal(err)
	// 	}
	// 	txt = strings.TrimSpace(txt)

	// 	if ans == txt {
	// 		correct += 1
	// 	} else {
	// 		incorrect += 1
	// 	}
	// }

	// fmt.Printf("You answered %v questions correctly.", correct)
	// fmt.Printf("\nYou answered %v questions incorrectly.\n", incorrect)
	Solve()
}
