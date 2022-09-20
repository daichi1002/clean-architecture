package domain

import (
	"clean-architecture/pb"
	"time"
)

type User struct {
	UserId    string
	Name      string
	Email     string
	Birthday  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Users []User

// ToProto is convert user to proto.
func (u *User) ToProto() *pb.User {
	return &pb.User{
		UserId: u.UserId,
		Name:   u.Name,
		Email:  u.Email,
	}
}

// ToProto is convert users model to proto.
func (u Users) ToProto() []*pb.User {
	var users []*pb.User
	for _, user := range u {
		users = append(users, user.ToProto())
	}
	return users
}
