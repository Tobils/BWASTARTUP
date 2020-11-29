# BWASTARTUP

Course Build with angga membangun web crowds funding menggunakan Golang dan Nuxtjs

---

## Analisis Entity

- User
- Campaigns
- Campaign Images
- Transactions

---

## Entity Relationship Diagram (ERD)

kita bisa menggunaka erdplus.com

1. Diagram

<img src="./ERD-BWASTARTUP.png" style="align:denter;">

2. Details
   - Users
     - bisa membuat banyak campaign dan bersifat opsional
     - bisa memiliki banyak transaksi dan bersifat opsional
   - Campaign
     - wajib dimiliki oleh 1 user
     - wajib memiliki beberapa gambar campaign
     - bisa memiliki banyak transaksi
   - Campaign Image
     - wajib dimiliki oleh 1 campaign
   - Transactions
     - wajib dimiliki oleh 1 user
     - wajib memiliki oleh 1 campaign

---

## Entity Fields / Columns

1. User

   - id : int
   - name : varchar
   - occupations : varchar
   - email : varchar
   - password_hash : varchar
   - avatar_file_name : varchar
   - role : varchar
   - token : varchar
   - created_at : datetime
   - updated_at : datetime

2. Campaigns

   - id : int
   - user_id : int
   - name : varchar
   - short_description : varchar
   - goal_amount : int
   - current_amount : int
   - description : text
   - perks : text
   - backer_count : int
   - slug : varchar
   - created_at : datetime
   - updated_at : datetime

3. Campaign Images

   - id : int
   - campaign_id : int
   - file_name : varchar
   - is_primary : boolean (tinyint)
   - created_at : datetime
   - updated_at : datetime

4. Transaction
   - id : int
   - campaign_id : int
   - user_id : int
   - amount : int
   - status : varchar
   - code : varchar
   - created_at : datetime
   - updated_at : datetime

## Table init

1. Table User
   ```sql
   CREATE TABLE users (
   id int(6) NOT NULL primary key AUTO_INCREMENT,
   name VARCHAR(50),
   occupation VARCHAR(50),
   email VARCHAR(50),
   password_hash VARCHAR(255),
   avatar_file_name VARCHAR(50),
   role VARCHAR(50),
   token VARCHAR(50),
   created_at DATETIME,
   updated_at DATETIME
   )
   ```

## Init Project

1. Config

   ```bash
   mkdir database
   mkdir bwastartup
   cd bwastartup
   go mod init bwastartup

   # auto reload
   ```

2. Install Gin dan Gorm

   ```bash
   # install GIN
   go get -u github.com/gin-gonic/gin

   # install GORM
   go get -u gorm.io/gorm

   # driver gorm mysql
   go get -u gorm.io/driver/mysql
   ```

3. Run Mysql Docker

   ```bash
   cd database
   docker-compose up -d
   ```

4. Struktur Direktory

   - handler
   - service
   - repository -> db

5. Register

   - input
   - handler -> mapping input ke struct
   - service -> mapping ke struct input ke struct user
   - repository -> save struct user ke db

6. Login
   - input
   - handler -> mapping input ke trsuct input
   - service -> mapping struct input ke struct user, bcrypt password
   - repository -> find logged in by email

## JWT Auth / Json Web Tokwn Auth

jwt dimanfaatkan untuk authentikasi API berdasarkan token user. [cek details jwt](jwt.io)

```bash
# get jwt module
go get github.com/dgrijalva/jwt-go
```

1. create auth service /bawastartup/auth/service.go

   ```go

   // service.go
   package auth
   import "github.com/dgrijalva/jwt-go"

   /**
   1. generate token
   2. validasi token
   */

   type Service interface {
      GenerateToken(userID int) (string, error)
   }

   type jwtService struct {
   }

   func NewService() *jwtService {
      return &jwtService{}
   }

   var SECRET_KEY = []byte("BWASTARTUP_s3cr3t_k3y")

   func (s *jwtService) GenerateToken(userID int) (string, error) {
      claim := jwt.MapClaims{}
      claim["user_id"] = userID

      // generate token
      token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

      // tanda tangani token
      signedToken, err := token.SignedString(SECRET_KEY)
      if err != nil {
         return signedToken, err
      }

      return signedToken, nil
   }

   ```

2. API response with token
   ```json
   {
     "meta": {
       "message": "Suuccessfully loggedin",
       "code": 200,
       "status": "success"
     },
     "data": {
       "id": 1,
       "name": "test simpan dari service",
       "occupation": "petanikode",
       "email": "contoh@gmail.com",
       "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxfQ.pR5mUyz1tm_Ni6-mCi-ankpmIwVifpJ0k_tNjbyp6p8",
       "ImageUrl": ""
     }
   }
   ```

## Reference

- [gorm-connect-mysql](https://gorm.io/docs/connecting_to_the_database.html)
- [validation error](https://github.com/go-playground/validator)
