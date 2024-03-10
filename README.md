В это проекте вы можете найти почти все пункты в вашем опроснике, а именно:
1) Проект, в котором я ранее использовали module, package, exported names и структурировал проект.
2) Проект, в котором я ранее использовал errors, panic, recover ( panic, recover мне не приходилось использовать, но я понимаю для чего они нужны)
3) Проект, где вы ранее  работали с JSON (Marshalling & Unmarshalling) - в проекте я в принципе с фронтом общаюсь по средством Json файлов,  а так же с jwt токенами.
4) Создайте API для того, чтобы можно было реализовать регистрацию и авторизацию: эндпоинты /signup, /login реализуют эти функции
{
    "first_name": "Erkanat",
    "last_name" : "Azamtov",
    "password": "Erk223",
    "avatar_link": "link",
    "gender": "true",
    "age":28,
    "phone_number":"564654123",
    "city_of_residence":"Astana",
    "description": "test of login post request"
    "email": "rakhat.oshakbayev@gmail.com"
} - это пример JSON файла для регистрации

{
    "email": "rakhat.oshakbayev@gmail.com"
    "password": "Kanat"
} - это пример JSON файла логина
5) На основе задачи 5 создайте API запросы, чтобы пользователи могли менять свои данные - данную функцию реализует эндпойнт /updateProfile
{
    "jwtToken": {
        "Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJha2hhdC5vc2hha2JheWV2QGdtYWlsLmNvbSIsImxldmVsIjoidXNlciIsImV4cCI6MTcxMDA2NTk4MX0.ss-Xj-2uUeXAuSoNVUVGW7QKvxQQd6fZwyHusFZKsA8"
    },
    "userInfo":{
        "first_name": "Erkanat",
        "last_name" : "Malhayev",
        "password": "Kanat",
        "avatar_link": "AZAZAZ",
        "gender": "false",
        "age":228,
        "phone_number":"87756400316",
        "city_of_residence":"Astana",
        "description": "test of login post request",
        }
} - это пример JSON файла для смены данных
