package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mediocregopher/radix/v3"
	"gocloud.dev/pubsub"
	_ "gocloud.dev/pubsub/kafkapubsub"
)

func main() {
	router := mux.NewRouter()

	router.
		HandleFunc("/rate", PostRateHandler).
		Methods(http.MethodPost)
	router.
		HandleFunc("/total", GetTotalHandler).
		Methods(http.MethodGet)

	if err := http.ListenAndServe(":8081", router); err != http.ErrServerClosed {
		panic(fmt.Errorf("error on listen and serve: %v", err))
	}
}

func GetTotalHandler(w http.ResponseWriter, r *http.Request) {
	var rates []string
	err := storage().Do(radix.Cmd(&rates, "LRANGE", "result", "0", "10"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(rates) == 0 {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	var sum int
	for _, rate := range rates {
		v, err := strconv.Atoi(rate)
		if err != nil {
			continue
		}
		sum += v
	}
	result := float64(sum) / float64(len(rates))
	_, _ = w.Write([]byte(fmt.Sprintf("%.2f", result)))
}

func PostRateHandler(w http.ResponseWriter, r *http.Request) {
	rate := r.FormValue("rate")
	if _, err := strconv.Atoi(rate); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("Send message: %s\n", rate)

	err := topic().Send(context.Background(), &pubsub.Message{
		Body: []byte(rate),
	})
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

var t *pubsub.Topic

func topic() *pubsub.Topic {
	if t != nil {
		return t
	}
	var err error
	t, err = pubsub.OpenTopic(context.Background(), "kafka://rates")
	if err != nil {
		panic(err)
	}
	return t
}

var (
	s        *radix.Pool
	connFunc = func(network, addr string) (radix.Conn, error) {
		return radix.Dial(network, addr,
			radix.DialTimeout(1*time.Minute),
		)
	}
)

func storage() *radix.Pool {
	if s != nil {
		return s
	}
	var err error
	s, err = radix.NewPool("tcp", "redis:6379", 1, radix.PoolConnFunc(connFunc))
	if err != nil {
		panic(err)
	}
	return s
}
