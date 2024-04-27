package usecase

import (
	"context"
	"encoding/csv"
	"github.com/labstack/gommon/log"
	"os"
	"scaleX/internal/constants"
	"scaleX/internal/dto"
	"scaleX/internal/repository"
	"strconv"
	"strings"
	"sync"
)

type bookService struct {
	userRepo repository.UserRepo
}

const (
	filePath        = "./sampleFile/"
	regularFileName = "regularUser.csv"
	adminFileName   = "adminUser.csv"
)

func (b bookService) AddBook(ctx context.Context, request dto.AddBookRequest) error {
	return insertBookInfoToFile(request)
}

func (b bookService) DeleteBook(ctx context.Context, request dto.DeleteBookRequest) error {
	return deleteBookFromFile(request.Name, regularFileName)
}

func (b bookService) FetchBook(ctx context.Context, userId string) (res dto.FetchBookResp, err error) {
	userInfo, err := b.userRepo.GetUserById(ctx, userId)
	if err != nil {
		return res, err
	}

	var bookNamesResp []string

	switch userInfo.Role {
	case constants.REGULAR_ROLE:
		regularBookNames, err := readBooksInfoFromFile(regularFileName)
		if err != nil {
			return res, err
		}
		bookNamesResp = append(bookNamesResp, regularBookNames...)
	case constants.ADMIN_ROLE:

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			regularBookNames, err := readBooksInfoFromFile(regularFileName)
			if err != nil {
				return
			}
			bookNamesResp = append(bookNamesResp, regularBookNames...)
		}()

		go func() {
			defer wg.Done()
			adminBookNames, err := readBooksInfoFromFile(adminFileName)
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

func insertBookInfoToFile(book dto.AddBookRequest) error {
	file, err := os.OpenFile(filePath+regularFileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	newData := [][]string{
		{book.Name, book.Author, strconv.Itoa(book.PublicationYear)},
	}

	for _, row := range newData {
		err := writer.Write(row)
		if err != nil {
			return err
		}
	}

	return err
}

func deleteBookFromFile(bookName, fileName string) error {
	file, err := os.Open(filePath + fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	//rewrite new Data except input book name to the file
	var newData [][]string
	for _, record := range records {
		bookNameRow := record[0]
		if strings.ToUpper(bookNameRow) != strings.ToUpper(bookName) {
			newData = append(newData, record)
		}
	}

	file, err = os.Create(filePath + fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(newData)
	if err != nil {
		return err
	}

	return err
}

func readBooksInfoFromFile(fileName string) (bookNames []string, err error) {
	file, err := os.Open(filePath + fileName)
	if err != nil {
		log.Errorf("Error opening file: ", err)
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Errorf("Error reading file: ", err)
		return nil, err
	}

	for i, record := range records {
		if i > 0 {
			bookNames = append(bookNames, record[0])
		}
	}

	return bookNames, err
}

func NewBookService(userRepo repository.UserRepo) BookService {
	return &bookService{
		userRepo: userRepo,
	}
}
