## Machine Test
The application is a student management application and store students information like first_name, last_name, grade, email, phone_number. All the APIs provided here is to alter or modify these datas. The database used here is MongoDB Atlas and the server is hosted in the AWS EC2 instance. 

## API Specification
You can test the API's using Postman. Use this [postman collection](https://www.postman.com/abgth/workspace/machine-test/collection/20732200-f28800df-248c-4721-9a8b-74b0164c8481?action=share&creator=20732200) to test the API's, all the documentation you needed is provided in the following.<br>
Below is the APIs used in the application and some examples along with it. 

## 👉 Insert student details
  ### Endpoint :
  ```
  http://3.110.186.160/insert
  ```  
  ### Method:
  `POST`
  
  ### Request Body:
  | Parameter | Type | Description |
  |-----------|------|-------------|
  | `fname` | string | First name of the student |
  | `lname` | string | Last name of the student |
  | `grade` | integer | Class or Standard of the student |
  | `email` | string | E-Mail ID of the student |
  | `phone` | string | Phone number of the student |
  
  ### Example Request:
  ```
  curl -X POST http://3.110.186.160/insert \
  -H "Content-Type: application/json" \
  -d '{
      "fname": "John",
      "lname": "Doe",
      "grade": 10,
      "email": "johndoe@example.com",
      "phone": "+1234567890"
  }'
  ```
  
  ### Success Response:
  HTTP Code: `200 OK`
  
  ```
  {
    "error": false,
    "message": "User has been inserted"
  }
  ```
  
## 👉 Get all student details
  ### Endpoint :
  ```
  http://3.110.186.160/get-all-students
  ```  
  ### Method:
  `GET`
  
  ### Example Request:
  ```
  curl -X GET http://3.110.186.160/get-all-students
  ```
  
  ### Success Response:
  HTTP Code: `200 OK`
  
  ```
  "error": false,
    "message": "Successfully fetched all user data",
    "data": [
        {
            "SID": "63e638d6cfb9e4b870113b20",
            "f_name": "Test",
            "l_name": "1",
            "grade": "Four",
            "email": "test@gmail.com",
            "phone": "1234567808"
        },
        {
            "SID": "63e62b88cfb9e4b870113b1f",
            "f_name": "Rahul",
            "l_name": "P K",
            "grade": "Three",
            "email": "pqrs@gmail.com",
            "phone": "1234567808"
        },
        {
            "SID": "63e62b6acfb9e4b870113b1e",
            "f_name": "Abhijith",
            "l_name": "A",
            "grade": "Two",
            "email": "abc@gmail.com",
            "phone": "1234567999"
        }
    ]
  }
  ```
  
## 👉 Get one student details
  ### Endpoint :
  ```
  http://3.110.186.160/get/{id}
  ```  
  ### Method:
  `GET`
 
  ### Example Request:
  ```
  curl -X GET http://3.110.186.160/get/63e62b6acfb9e4b870113b1e
  ```
  
  ### Success Response:
  HTTP Code: `200 OK`
  
  ```
  {
    "error": false,
    "message": "Successfully fetched user data",
    "data": {
      "SID": "63e62b6acfb9e4b870113b1e",
      "f_name": "Abhijith",
      "l_name": "A",
      "grade": "Two",
      "email": "abc@gmail.com",
      "phone": "1234567999"
     }
  }
  ```
  
## 👉 Update one student details
  ### Endpoint :
  ```
  http://3.110.186.160/update/{id}
  ```  
  ### Method:
  `PUT/PATCH`
 
  ### Request Body:
  | Parameter | Type | Description |
  |-----------|------|-------------|
  | `fname` | string | First name of the student |
  | `lname` | string | Last name of the student |
  | `grade` | integer | Class or Standard of the student |
  | `email` | string | E-Mail ID of the student |
  | `phone` | string | Phone number of the student |
  
  ### Example Request:
  ```
  curl -X PUT http://3.110.186.160/update/63e62b6acfb9e4b870113b1e \
  -H "Content-Type: application/json" \
  -d '{
      "fname": "John",
      "lname": "Doe",
      "grade": 10,
      "email": "johndoe@example.com",
      "phone": "+1234567890"
  }'
  ```
  
  ### Success Response:
  HTTP Code: `200 OK`
  
  ```
  {
    "error": false,
    "message": "Successfully updated the user data"
  } 
  ```
  
## 👉 Delete one student details
  ### Endpoint :
  ```
  http://3.110.186.160/delete/63e62b44cfb9e4b870113b1d
  ```  
  ### Method:
  `DELETE`
  
  ### Example Request:
  ```
  curl -X DELETE http://3.110.186.160/delete/63e62b44cfb9e4b870113b1d
  ```
  
  ### Success Response:
  HTTP Code: `200 OK`
  
  ```
  {
    "error": false,
    "message": "Successfully deleted the user data"
  }
  ```
## 👉 Delete entire collection or database
  ### Endpoint :
  ```
  http://3.110.186.160/drop-collection
  ```  
  ### Method:
  `DELETE`
  
  ### Example Request:
  ```
  curl -X DELETE http://3.110.186.160/drop-collection
  ```
  
  ### Success Response:
  HTTP Code: `200 OK`
  
  ```
  {
    "error": false,
    "message": "Successfully dropped the collection"
  }
  ```

