package user

import (
	"be-catatin/internal/entity"
	userRepo "be-catatin/internal/repository/user"
)

type Usecase interface {
	LoginOrCreate(phone string, username string, pin string) (*entity.User, error)
}

type usecase struct {
	repo userRepo.Repository
}

func NewUsecase(repo userRepo.Repository) Usecase {
	return &usecase{repo}
}

func (u *usecase) LoginOrCreate(phone string, username string, pin string) (*entity.User, error) {
	// 1. Cek apakah user sudah ada berdasarkan phone atau username
	user, err := u.repo.FindByPhoneOrUsername(phone, username)
	if err != nil {
		return nil, err
	}

	// 2. Jika user sudah ada, return data user tersebut (nanti bisa tambah cek validitas pin di sini kalau mau login beneran)
	if user != nil {
		// Asumsi sementara: kalau pin cocok bisa lanjut (di tahap berikutnya bisa implement hash/bcrypt)
		return user, nil
	}

	// 3. Jika belum ada, buat user baru
	newUser := &entity.User{
		Username:    username,
		PhoneNumber: phone,
		Pin:         pin, // Todo: bcrypt
	}

	err = u.repo.Create(newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
