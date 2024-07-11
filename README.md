### Application run on port 8888
I'd be delighted to discuss this exam with you. Please do not hesitate to contact me via Telegram (https://t.me/naamleee) or WhatsApp (84886680912). Thank you very much.

### Go mod tidy
    make tidy

### Local development usage:
    make run

### Login Api
- You can use user_name = "**admin**" for admin role 
or user_name = "**regular**" for regular role


    curl --location 'localhost:8888/api/v1/login' \
    --header 'Content-Type: application/json' \
    --data '{
        "user_name" : "admin"
    }'

![img.png](sampleFile/basicLogin.png)

### Home Api
    curl --location 'localhost:8888/api/v1/home' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiYWRtaW5Vc2VySWQiLCJyb2xlIjoiYWRtaW4iLCJ1c2VyX25hbWUiOiJhZG1pbiIsImV4cCI6MTcxNDE5NTEyOX0._d8JYg12DthfEUsJ9IPNWGupQpYM0dhAiDILgMxNy5U'

![img_1.png](sampleFile/homeApi.png)

### AddBook Api
    curl --location 'localhost:8888/api/v1/addBook' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiYWRtaW5Vc2VySWQiLCJyb2xlIjoiYWRtaW4iLCJ1c2VyX25hbWUiOiJhZG1pbiIsImV4cCI6MTcxNDE5MzEwMn0.wRj5ni7n6JS2wEDoaZNrKMSxudZyyCv2dExe7FXH_kQ' \
    --header 'Content-Type: application/json' \
    --data '{
        "name" : "TEST2023",
        "author" : "Nam Lee" ,
        "publication_year" : 2023
    }'

![img_2.png](sampleFile/addBookApi.png)
-- Added TEST2023 record

### DeleteBook Api
    curl --location --request DELETE 'localhost:8888/api/v1/deleteBook' \
        --header 'Content-Type: application/json' \
        --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiYWRtaW5Vc2VySWQiLCJyb2xlIjoiYWRtaW4iLCJ1c2VyX25hbWUiOiJhZG1pbiIsImV4cCI6MTcxNDE5MzEwMn0.wRj5ni7n6JS2wEDoaZNrKMSxudZyyCv2dExe7FXH_kQ' \
        --data '{
            "name" : "TEST2023"
    }'

![img_3.png](sampleFile/deleteBookAPi.png)
        -- No more TEST2023 record


### Get summary content of each chapter
    curl --location --request GET 'localhost:8888/api/v1/chapterContent'

![img.png](sampleFile/getChapterContent.png)

##### Create  worker 