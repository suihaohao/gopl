//Learning URL https://books.studygolang.com/gopl-zh/
package chat

type client chan <- string

var (
	entering = make(chan client)
	leaving = make(chan client)
	message = make(chan string)
)

func broadcaster()  {
	clients := make(map[client]bool)
	for  {
		select {
		case msg:=<-message:
			for cli := range clients {
				cli <- msg
			}
		case cli:=<-entering:
			clients[cli]=true
		case cli:=<-leaving:
			delete(clients,cli)
			close(cli)
		}
	}
}
