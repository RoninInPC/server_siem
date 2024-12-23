package mongosubject

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"server_siem/entity/subject"
)

type MongoDB struct {
	client   *mongo.Client
	ctx      context.Context
	Address  string
	Database string
	Username string
	Password string
}

func Init(Address string, Database string, Username string, Password string) (*MongoDB, error) {
	m, err := mongo.Connect(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s", Address)).
		SetAuth(options.Credential{
			Username: Username,
			Password: Password}))
	if err != nil {
		return nil, err
	}
	return &MongoDB{client: m, ctx: context.TODO(), Address: Address, Database: Database, Username: Username, Password: Password}, nil
}

var subjectMap = map[subject.SubjectType]string{
	subject.ProcessT:    "Processes",
	subject.PortTablesT: "Ports",
	subject.UserT:       "Users",
	subject.FileT:       "Files",
}

func (m MongoDB) Add(sub subject.Subject) bool {
	switch sub.Type() {
	case subject.FileT:
		m.client.Database(m.Database).Collection("Files").InsertOne(m.ctx, sub.(subject.File))
		return true
	case subject.ProcessT:
		m.client.Database(m.Database).Collection("Processes").InsertOne(m.ctx, sub.(subject.Process))
		return true
	case subject.UserT:
		m.client.Database(m.Database).Collection("Users").InsertOne(m.ctx, sub.(subject.User))
		return true
	case subject.PortTablesT:
		m.client.Database(m.Database).Collection("Ports").InsertOne(m.ctx, sub.(subject.PortTables))
		return true
	default:
		return false

	}
}

func (m MongoDB) Update(sub subject.Subject) bool {
	switch sub.Type() {
	case subject.FileT:
		f := sub.(subject.File)
		filter := bson.D{{"filename", f.FullName}}
		m.client.Database(m.Database).Collection("Files").UpdateOne(m.ctx, filter, f)
		return true
	case subject.ProcessT:
		f := sub.(subject.Process)
		filter := bson.D{{"pid", f.PID}}
		m.client.Database(m.Database).Collection("Processes").UpdateOne(m.ctx, filter, f)
		return true
	case subject.UserT:
		f := sub.(subject.User)
		filter := bson.D{{"uid", f.Uid}}
		m.client.Database(m.Database).Collection("Users").UpdateOne(m.ctx, filter, f)
		return true
	case subject.PortTablesT:
		f := sub.(subject.PortTables)
		filter := bson.D{{"port", f.Port}}
		m.client.Database(m.Database).Collection("Ports").UpdateOne(m.ctx, filter, f)
		return true
	default:
		return false

	}
}

func (m MongoDB) Delete(sub subject.Subject) bool {
	switch sub.Type() {
	case subject.FileT:
		f := sub.(subject.File)
		filter := bson.D{{"filename", f.FullName}}
		m.client.Database(m.Database).Collection("Files").DeleteOne(m.ctx, filter)
		return true
	case subject.ProcessT:
		f := sub.(subject.Process)
		filter := bson.D{{"pid", f.PID}}
		m.client.Database(m.Database).Collection("Processes").DeleteOne(m.ctx, filter)
		return true
	case subject.UserT:
		f := sub.(subject.User)
		filter := bson.D{{"uid", f.Uid}}
		m.client.Database(m.Database).Collection("Users").DeleteOne(m.ctx, filter)
		return true
	case subject.PortTablesT:
		f := sub.(subject.PortTables)
		filter := bson.D{{"port", f.Port}}
		m.client.Database(m.Database).Collection("Ports").DeleteOne(m.ctx, filter)
		return true
	default:
		return false

	}
}

func (m MongoDB) Get(sub subject.Subject) subject.Subject {
	switch sub.Type() {
	case subject.FileT:
		f := sub.(subject.File)
		filter := bson.D{{"filename", f.FullName}}
		file := subject.File{}
		m.client.Database(m.Database).Collection("Files").FindOne(m.ctx, filter).Decode(&file)
		return file
	case subject.ProcessT:
		f := sub.(subject.Process)
		filter := bson.D{{"pid", f.PID}}
		file := subject.Process{}
		m.client.Database(m.Database).Collection("Processes").FindOne(m.ctx, filter).Decode(&file)
		return file
	case subject.UserT:
		f := sub.(subject.User)
		filter := bson.D{{"uid", f.Uid}}
		file := subject.User{}
		m.client.Database(m.Database).Collection("Users").FindOne(m.ctx, filter).Decode(&file)
		return file
	case subject.PortTablesT:
		f := sub.(subject.PortTables)
		filter := bson.D{{"port", f.Port}}
		file := subject.PortTables{}
		m.client.Database(m.Database).Collection("Ports").FindOne(m.ctx, filter).Decode(&file)
		return file
	default:
		return subject.Nope{}

	}
}
