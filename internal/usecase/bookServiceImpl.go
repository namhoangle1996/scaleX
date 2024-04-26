package usecase

import (
	"context"
	"encoding/csv"
	"github.com/labstack/gommon/log"
	"os"
	"scaleX/internal/dto"
	"scaleX/internal/repository"
	"sync"
)

type bookService struct {
	userRepo repository.UserRepo
}

func (b bookService) FetchBook(ctx context.Context, userId string) (res dto.FetchBookResp, err error) {
	userInfo, err := b.userRepo.GetUserById(ctx, userId)
	if err != nil {
		return res, err
	}

	var bookNamesResp []string

	switch userInfo.Role {
	case "regular":
		regularBookNames, err := readBooksInfoFromFile("regularUser.csv")
		if err != nil {
			return res, err
		}
		bookNamesResp = append(bookNamesResp, regularBookNames...)
	case "admin":

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			regularBookNames, err := readBooksInfoFromFile("regularUser.csv")
			if err != nil {
				return
			}
			bookNamesResp = append(bookNamesResp, regularBookNames...)
		}()

		go func() {
			defer wg.Done()
			adminBookNames, err := readBooksInfoFromFile("adminUser.csv")
			if err != nil {
				return
			}
			bookNamesResp = append(bookNamesResp, adminBookNames...)
		}()

		wg.Wait()

	}

	res.BookNames = bookNamesResp

	return res, err

}

func readBooksInfoFromFile(fileName string) (bookNames []string, err error) {
	filePath := "./sampleFile/"

	file, err := os.Open(filePath + fileName)
	if err != nil {
		log.Errorf("Error opening file: ", err)
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read all records from the CSV
	records, err := reader.ReadAll()
	if err != nil {
		log.Errorf("Error reading file: ", err)
		return nil, err
	}

	for _, record := range records {
		bookNames = append(bookNames, record[0])
	}

	return bookNames, err
}

func NewBookService(userRepo repository.UserRepo) BookService {
	return &bookService{
		userRepo: userRepo,
	}
}
