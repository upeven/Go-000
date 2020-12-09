package main


import (
	"context",
	"golang.org/x/sync/errgroup",
	"log",
	"net/http",
	"os",
	"os/signal",
	"syscall",
	"time",
)

type httpServer struct{
	server http.Server
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{
		server:http.Server{
			Addr:addr,
		}
	}
}

func (h *httpServer) start() error{
	return h.server.ListenAndServe()
}

func (h *httpServer) shutdown(ctx context.Context) error {
	return h.server.Shutdown(ctx)
}

func main(){
	ctx,cancel := context.WithCancel(context.Background())
	defer cancel()

	g, _ := errgroup.WithContext(ctx)
	http1 := NewHttpServer(":8080")

	g.Go(func() error) {
		if err := http1.start();err != nil {
			cancel()
			return err
		}

		return nil
	}

	http2 := NewHttpServer(":8081")
	g.Go(func() error) {
		if err := http2.start();err != nil {
			cancel()
			return err
		}

		return nil
	}

	c := make(chan os.Signal,1)
	signal.Notify(c,syscal.SIGHUP,syscall.SIGQUIT,syscall.SIGTERM,syscall.SIGINT)
	go func() {
		for {
			select {
			case s := <-c :
				switch s {
				case syscall.SIGQUIT,syscall.SIGTERM,syscall.SIGINT:
					cancel()
				default:
				}
			}
		}
	}()
	go func() {
		select {
		case <-ctx.Done():
			cts,cancel := context.WithTimeout(context.Background(),10*time.Second):
				defer cancel()
				go func(){
					if err := http1.shutdown(ctx);err != nil{
						log.Println("http1 shutdown err:",err)
					}
				}()
				go func(){
					if err := http2.shutdown(ctx);err != nil{
						log.Println("http2 shutdown err:",err)
					}
				}()
			
		}
	}()

	if err := g.Wait();err != nil {
		log.Println("all exit:",err)
	}



}