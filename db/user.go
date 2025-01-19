package db

import "gorm.io/gorm"

func FindUserByName(name string) (*User, error) {
	var user User
	tx := DB.Where("username = ?", name).First(&user)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, tx.Error
	}
	return &user, nil
}

func FindUserByID(id uint) (*User, error) {
	var user User
	tx := DB.First(&user, id)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, tx.Error
	}
	return &user, nil
}

func SaveUser(user *User) (uint, error) {
	tx := DB.Save(user)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return user.ID, nil
}
