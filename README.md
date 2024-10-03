
# MyWallet backend API

A simple e-wallet app backend API for tracking your income and outcome. This is my first project for practicing with go  
programming language, gorm, and gofiber

## List of endpoint
Base URL = http://localhost:3000/api this is default base url and default port is 3000, you can replace with your own domain

## Login
`POST /auth/login`
Body :
| name | type | required| description |
|--|--|--|--|
| username | string |yes|-|
| password | string |yes|-|

Responses
>200

    {
	  "data": {
	    "token": "exampletoken"
	  },
	  "message": "success",
	  "status": "0"
	}

## Register
`POST /auth/register`
Body :
| name | type | required| description |
|--|--|--|--|
| username | string |yes|max 255 character|
| password | string |yes|-|

Responses
>200

    {
	  "data": {
	        "id": 2,
		    "username": "joko"
	  },
	  "message": "success",
	  "status": "0"
	}
## Get Balance
`GET /user/balance`
Headers
| name|type| required | description |
|--|--|--|--|
|  Authorization|string|yes  |jwt token, get jwt token after login|


Responses
>200

    {
	  "data": {
	        "current_balance": 1100000
	  },
	  "message": "success",
	  "status": "0"
	}
## Create Transaction
`POST /transaction`
Headers
| name|type| required | description |
|--|--|--|--|
|  Authorization|string|yes  |jwt token, get jwt token after login|
Body :
| name | type | required| description |
|--|--|--|--|
| amount | number / double |yes|-|
| desc| string |no|simple description max 255 char|

Responses
>200

    {
	  "data": {
	        "amount": 100000,
		    "desc": "",
		    "id": 1,
		    "created_at": "2024-10-03T23:04:10.793+07:00"
	  },
	  "message": "success",
	  "status": "0"
	}
## Update Transaction
`PUT /transaction/:id`
Headers
| name|type| required | description |
|--|--|--|--|
|  Authorization|string|yes  |jwt token, get jwt token after login|
Body :
| name | type | required| description |
|--|--|--|--|
| amount | number / double |yes|fill with new data or use old data|
| desc| string |no|simple description max 255 char, fill with new data or use old data|

Responses
>200

    {
	  "data": {
	        "amount": 100000,
		    "desc": "",
		    "id": 1,
		    "created_at": "2024-10-03T23:04:10.793+07:00"
	  },
	  "message": "success",
	  "status": "0"
	}
## Get All Transaction
`GET /transaction`
Headers
| name|type| required | description |
|--|--|--|--|
|  Authorization|string|yes  |jwt token, get jwt token after login|

Responses
>200

    {
	  "data": [
    {
      "amount": 100000,
      "desc": "Jangan makan malam",
      "id": 1,
      "created_at": "2024-10-03T14:44:49.38Z"
    },
    {
      "amount": 100000,
      "desc": "Jangan makan malam",
      "id": 2,
      "created_at": "2024-10-03T14:44:54.838Z"
    }
	  ],
	  "message": "success",
	  "status": "0"
	}
## Get Transaction by Id
`GET /transaction/:id`
Headers
| name|type| required | description |
|--|--|--|--|
|  Authorization|string|yes  |jwt token, get jwt token after login|

Responses
>200

    {
	  "data":     {
      "amount": 100000,
      "desc": "Jangan makan malam",
      "id": 1,
      "created_at": "2024-10-03T14:44:54.838Z"
    },
	  "message": "success",
	  "status": "0"
	}
