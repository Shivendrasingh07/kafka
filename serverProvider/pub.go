package serverProvider

import (
	"bufio"
	"encoding/json"
	"example.com/m/models"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

func (srv *Server) Pub() {

	fmt.Print("Enter text: ")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	// remove the delimeter from the string
	//input = strings.TrimSuffix(input, "\n")
	//fmt.Println(input)

	outboundMessageBytes, err := json.Marshal(&models.Message{
		Message: input,
	})

	if err != nil {
		logrus.Errorf("toggleBlock: error marshal outbound message data  %v", err)
		return
	}
	srv.Messenger.Publish(outboundMessageBytes)

}
