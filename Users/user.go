package Users

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID			uint32		`gorm:"primary_key;auto_increment" json:"id"`
	Username	string		`gorm:"size:255;not null;unique" json:"username"`
	Email		string		`gorm:"size:100;not null;" json:"email"`
	Password	string		`gorm:"size:100;not null;" json:"password"`
	CreatedAt	time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt	time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

