/*
 * Created Date: Friday, July 8th 2022, 2:10:57 pm
 * Author: ZhanG
 */
package database

import (
	"context"
	"fmt"
	"time"

	"github.com/ElieenAndBella/MStarExamSystem/constant"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Client

func init() {
	uri := fmt.Sprintf("mongodb://%s:%s", constant.MC.Addr, constant.MC.Port)
	if constant.MC.User != "" || constant.MC.Pass != "" {
		uri = fmt.Sprintf("mongodb://%s:%s@%s:%s", constant.MC.User, constant.MC.Pass, constant.MC.Addr, constant.MC.Port)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	MongoDB, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	err = MongoDB.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
}
