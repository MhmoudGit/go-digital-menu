package utils

import (
	"fmt"

	"github.com/MhmoudGit/go-digital-menu/database"
	h "github.com/MhmoudGit/go-digital-menu/helpers"

	"time"
)

func checkSubscriptions() {
	users, err := h.GetAllUsers(database.Db)
	if err != nil {
		fmt.Println(err)
	}
	today := time.Now()

	for _, user := range users {
		if user.EndDate.Before(today) && user.Paid {
			user.Paid = false
			user.Restaurant.IsActive = false
			h.UpdateUser(database.Db, &user, user.ID)
		}
	}
}

func RunFunctionEvery12Hours() {
	ticker := time.NewTicker(12 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Println("checking db")
		checkSubscriptions()
	}
}
