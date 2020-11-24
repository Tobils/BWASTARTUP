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

## Reference
- [gorm-connect-mysql](https://gorm.io/docs/connecting_to_the_database.html)