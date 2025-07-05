package user

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"uniqueIndex" json:"name"`
	Password  string    `json:"password"`
	Email     string    `gorm:"uniqueIndex" json:"email"`
	Dob       *time.Time `json:"dob"`
	Role      *string    `json:"role"`
	Address   *string    `json:"address"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Hook: Generate UUID before inserting
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()

	hashedBytes, err := bcrypt.GenerateFromPassword(
		[]byte(
			u.Password,
		),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	u.Password = string(hashedBytes)

	return
}

// Hook: Generate UUID before inserting
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {

	hashedBytes, err := bcrypt.GenerateFromPassword(
		[]byte(
			u.Password,
		),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	u.Password = string(hashedBytes)

	return
}
