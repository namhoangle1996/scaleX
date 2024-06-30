package main

import (
	"log"
	"scaleX/internal/handlers/restHandler"
	"scaleX/internal/repository"
	"scaleX/internal/usecase"
	"sort"
)

func main() {

	userRepo := repository.NewUserRepo()
	authService := usecase.NewAuthService(userRepo)
	bookService := usecase.NewBookService(userRepo)

	handler := restHandler.NewRestHandler(authService, bookService)

	echoServer := restHandler.Echo(handler)

	log.Fatal("failed to start server", echoServer.Start(":8888"))

}

func kthSmallest(matrix [][]int, k int) int {

	var tmpSlice []int

	for i := range matrix {
		tmpSlice = append(tmpSlice, matrix[i]...)
	}

	sort.Ints(tmpSlice)

	return tmpSlice[k-1]

}

func kSmallestPairs(nums1 []int, nums2 []int, k int) (res [][]int) {

	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)

	for i := range nums1 {
		for j := i + 1; j < len(nums1); j++ {
			tmp := []int{nums1[i], nums1[j]}
			res = append(res, tmp)
			k--
			if k == 0 {
				break
			}
		}
	}

	return res
}
