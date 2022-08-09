package main

//func main() {
//	r := chi.NewRouter()
//	//r.Use(middleware.Logger)
//	logrus.Println("hi, im publisher")
//	//r.Get("/", func(w http.ResponseWriter, r *http.Request) {
//	//	w.Write([]byte("welcome"))
//	//})
//	//r.Post("/pub", func(w http.ResponseWriter, r *http.Request) {
//	//	w.Write([]byte("welcome"))
//	//})
//
//	//srv.PublisherMessage()
//}

//func PublisherMessage() {
//
//	fmt.Print("Enter text: ")
//	reader := bufio.NewReader(os.Stdin)
//	// ReadString will block until the delimiter is entered
//	input, err := reader.ReadString('\n')
//	if err != nil {
//		fmt.Println("An error occured while reading input. Please try again", err)
//		return
//	}
//
//	// remove the delimeter from the string
//	input = strings.TrimSuffix(input, "\n")
//	//fmt.Println(input)
//
//	outboundMessageBytes, err := json.Marshal(&models.Message{
//		Message: input,
//	})
//
//	if err != nil {
//		logrus.Errorf("toggleBlock: error marshal outbound message data  %v", err)
//		return
//	}
//	srv.Messenger.Publish(outboundMessageBytes)
//}
