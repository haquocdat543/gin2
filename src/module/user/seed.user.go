package user

import (
	"gin/src/share"
)

var Seed = []*User{
    {
        Name:     "Jinzhu",
        Email:    "jinzhu@gmail.com",
        Password: "12345678", // Should be hashed in real usage
        Dob:      share.ParseDate("2006-01-02"),
    },
    {
        Name:     "xiaoqinghua",
        Email:    "xiaoqinhua@gmail.com",
        Password: "87654321",
        Dob:      share.ParseDate("2006-07-05"), // ~19 years old
    },
    {
        Name:     "lalala",
        Email:    "lalala@gmail.com",
        Password: "lalala123",
        Dob:      share.ParseDate("2004-07-05"), // ~21 years old
    },
    {
        Name:     "lihanjiaqie",
        Email:    "lihanjiaqie@gmail.com",
        Password: "lihanmeinu",
        Dob:      share.ParseDate("2002-07-05"), // ~23 years old
    },
    {
        Name:     "dengdengdengwo",
        Email:    "dengdengdengwo@gmail.com",
        Password: "woaini<3",
        Dob:      share.ParseDate("2001-07-05"), // ~24 years old
    },
}

