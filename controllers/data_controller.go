package controllers

import (
	"encoding/json"
	"log"
	"time"

	"github.com/AliceTrinta/GO-RMA/config"
	"github.com/AliceTrinta/GO-RMA/entity"
	"github.com/AliceTrinta/GO-RMA/infrastructure/repository"
	"github.com/AliceTrinta/GO-RMA/usecase/Data"
	"github.com/nats-io/nats.go"
)

func RmlData() {
	nc, err := nats.Connect("0.0.0.0:4222")
	if err != nil {
		log.Fatal(err)
	}
	sub, err := nc.SubscribeSync("data")
	if err != nil {
		log.Fatal(err)
	}

	for {
		msg, err := sub.NextMsg(10 * time.Second)
		if err != nil {
			continue
		}
		err = SaveData(msg.Data)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func SaveData(data []byte) error {
	db, err := config.MongoConnect()
	if err != nil {
		return err
	}
	defer db.Session.Close()
	repo := repository.NewDataMongo(db)
	service := Data.NewService(repo)

	var d entity.Data
	log.Println(string(data))
	err = json.Unmarshal(data, &d)
	if err != nil {
		return err
	}

	log.Println("registering a data...")
	return service.CreateData(d)
}
