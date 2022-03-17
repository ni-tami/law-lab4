package main

import (
	"context"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	e := echo.New()

	client := db()
	mongo := &Database{}

	mongo.coll = client.Database("mongo").Collection("cookies")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello!")
	})
	e.GET("/cookies", mongo.getAllCookie)
	e.POST("/cookies", mongo.createCookie)
	e.GET("/cookies/:tag", mongo.getCookie)
	e.PATCH("/cookies/:tag", mongo.updateUser)
	e.DELETE("/cookies/:tag", mongo.deleteCookie)

	e.Logger.Fatal(e.Start("127.0.0.1:8080"))

}

type Database struct {
	coll *mongo.Collection
}

func (db *Database) getCookie(c echo.Context) error {
	opts := options.Find()
	tag := c.Param("tag")
	filter := bson.D{{"tag", bson.D{{"$eq", tag}}}}
	cursor, err := db.coll.Find(context.TODO(), filter, opts)

	if err != nil {
		log.Fatal(err)
	}
	var result []bson.M
	if err = cursor.All(context.TODO(), &result); err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, result)

}

func (db *Database) getAllCookie(c echo.Context) error {
	opts := options.Find()
	cursor, err := db.coll.Find(context.TODO(), opts)

	if err != nil {
		log.Fatal(err)
	}

	var result []bson.M
	if err = cursor.All(context.TODO(), &result); err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, result)

}

func (db *Database) createCookie(c echo.Context) error {
	var cookie Cookie
	if err := c.Bind(&cookie); err != nil {
		return err
	}
	result, err := db.coll.InsertOne(context.TODO(), cookie)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, result)
}

func (db *Database) updateUser(c echo.Context) error {
	var cookie Cookie
	if err := c.Bind(&cookie); err != nil {
		return err
	}

	toUpdate := getNonEmptyBson(c, cookie)

	tag := c.Param("tag")
	filter := bson.M{"tag": tag}
	update := bson.M{"$set": toUpdate}
	result, err := db.coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, result)
}

func (db *Database) deleteCookie(c echo.Context) error {
	tag := c.Param("tag")
	filter := bson.M{"tag": tag}
	result, err := db.coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, result.DeletedCount)
}

func getNonEmptyBson(c echo.Context, obj interface{}) bson.M {
	f := reflect.ValueOf(obj)
	fType := f.Type()
	update := bson.M{}
	for i := 0; i < f.NumField(); i++ {
		key := fType.Field(i).Name
		value := f.Field(i).Interface()
		switch value.(type) {
		case string:
			if value.(string) != "" {
				update[strings.ToLower(key)] = value
			}
		case float64:
			if value.(float64) > 0 {
				update[strings.ToLower(key)] = value
			}
		case []string:
			if len(value.([]string)) > 0 {
				update[strings.ToLower(key)] = value
			}
		}
	}
	return update
}
