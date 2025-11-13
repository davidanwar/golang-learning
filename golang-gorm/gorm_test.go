package golang_gorm

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

func OpenConnection() *gorm.DB {
	dsn := "root:d4v1d4nw4r@tcp(localhost:3306)/golang_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	return db
}

var db = OpenConnection()

func TestOpenConnection(t *testing.T) {
	assert.NotNil(t, db)
}

func TestExecSQL(t *testing.T) {
	err := db.Exec("insert into sample(id, name) values(?, ?)", "12345", "David").Error
	assert.Nil(t, err)

	err = db.Exec("insert into sample(id, name) values(?, ?)", "12346", "Anwar").Error
	assert.Nil(t, err)

	err = db.Exec("insert into sample(id, name) values(?, ?)", "12347", "Fatih").Error
	assert.Nil(t, err)

	err = db.Exec("insert into sample(id, name) values(?, ?)", "12348", "Zia").Error
	assert.Nil(t, err)
}

type Sample struct {
	Id   string
	Name string
}

func TestRawSQL(t *testing.T) {
	var sample Sample

	err := db.Raw("select * from sample where id = ?", "12345").Scan(&sample).Error
	assert.Nil(t, err)
	assert.Equal(t, "12345", sample.Id)

	var samples []Sample
	err = db.Raw("select * from sample").Scan(&samples).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(samples))
}

func TestScanRow(t *testing.T) {
	var samples []Sample
	rows, err := db.Raw("select * from sample").Rows()
	assert.Nil(t, err)
	defer rows.Close()

	for rows.Next() {
		err := db.ScanRows(rows, &samples)
		assert.Nil(t, err)
	}
	assert.Nil(t, err)
	assert.Equal(t, 4, len(samples))
}

func TestCreateUser(t *testing.T) {
	user := User{
		ID:       "110011",
		Password: "123456",
		Name: Name{
			FirstName:  "Eka",
			MiddleName: "Kusuma",
			LastName:   "Wijaya",
		},
	}

	response := db.Create(&user)
	assert.Nil(t, response.Error)
	assert.Equal(t, int64(1), response.RowsAffected)
}

func TestBatchInsert(t *testing.T) {
	var users []User
	for i := 1; i < 100; i++ {
		users = append(users, User{
			ID:       strconv.Itoa(i),
			Password: "123456" + strconv.Itoa(i),
			Name: Name{
				FirstName:  "User " + strconv.Itoa(i),
				MiddleName: "Middle User " + strconv.Itoa(i),
				LastName:   "Last User " + strconv.Itoa(i),
			},
		})
	}
	//response := db.Create(&users)
	//assert.Nil(t, response.Error)
	//assert.Equal(t, int64(len(users)), response.RowsAffected)

	responseBatch := db.CreateInBatches(&users, 10)
	assert.Nil(t, responseBatch.Error)
	assert.Equal(t, int64(len(users)), responseBatch.RowsAffected)
}

func TestTransaction(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&User{
			ID:       "100",
			Password: "123456",
			Name: Name{
				FirstName:  "User",
				MiddleName: "Middle User",
				LastName:   "Last User",
			},
		}).Error
		if err != nil {
			return err
		}

		err = tx.Create(&User{
			ID:       "101",
			Password: "123456",
			Name: Name{
				FirstName:  "User",
				MiddleName: "Middle User",
				LastName:   "Last User",
			},
		}).Error
		if err != nil {
			return err
		}

		return nil
	})

	assert.Nil(t, err)
}

func TestTransactionError(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&User{
			ID:       "105",
			Password: "123456",
			Name: Name{
				FirstName:  "User",
				MiddleName: "Middle User",
				LastName:   "Last User",
			},
		}).Error
		if err != nil {
			return err
		}

		err = tx.Create(&User{
			ID:       "101",
			Password: "123456",
			Name: Name{
				FirstName:  "User",
				MiddleName: "Middle User",
				LastName:   "Last User",
			},
		}).Error
		if err != nil {
			return err
		}

		return nil
	})

	assert.NotNil(t, err)
}

func TestQuerySingleObject(t *testing.T) {
	user := User{}

	tx := db.First(&user)
	assert.Nil(t, tx.Error)
	assert.Equal(t, "1", user.ID)

	user = User{}
	tx = db.Last(&user)
	assert.Nil(t, tx.Error)
}

func TestQueryInlineSingleObject(t *testing.T) {
	user := User{}

	tx := db.First(&user, "id = ?", "100")
	assert.Nil(t, tx.Error)
	assert.Equal(t, "100", user.ID)

	user = User{}
	tx = db.Last(&user, "id = ?", "100")
	assert.Nil(t, tx.Error)
	assert.Equal(t, "100", user.ID)
}

func TestQueryMultiObject(t *testing.T) {
	var users []User
	tx := db.Find(&users, "id in ?", []string{"1", "2", "3"})
	assert.Nil(t, tx.Error)
	assert.Equal(t, 3, len(users))
}

func TestQueryCondition(t *testing.T) {
	var users []User
	tx := db.Where("id in (?)", []string{"1", "2", "3"}).Where("first_name like ?", "%User%").Find(&users)
	assert.Nil(t, tx.Error)
}

func TestQueryOrCondition(t *testing.T) {
	var users []User
	tx := db.Where("id in (?)", []string{"1", "2", "3"}).Or("first_name like ?", "%User%").Find(&users)
	assert.Nil(t, tx.Error)
}

func TestSelectedFields(t *testing.T) {
	var users []User
	tx := db.Select("id", "first_name").Find(&users)
	assert.Nil(t, tx.Error)
	for _, user := range users {
		assert.NotNil(t, user.ID)
		assert.NotEqual(t, "", user.Name.FirstName)
	}

	fmt.Println(users)
}

type UserResponse struct {
	ID         string
	FirstName  string
	MiddleName string
}

func TestQueryNonModel(t *testing.T) {
	var users []UserResponse
	tx := db.Model(&User{}).Select("id", "first_name", "middle_name").Find(&users)
	assert.Nil(t, tx.Error)
	fmt.Println(users)
}

func TestUpdate(t *testing.T) {
	user := User{}
	tx := db.Take(&user, "id = ?", "1")
	assert.Nil(t, tx.Error)

	user.Password = "secret"
	user.Name = Name{
		FirstName:  "First Name",
		MiddleName: "Middle Name",
		LastName:   "Last Name",
	}
	err := db.Save(&user).Error
	assert.Nil(t, err)
	assert.Equal(t, "secret", user.Password)
}

func TestUpdateSelectedColumns(t *testing.T) {
	err := db.Model(&User{}).Where("id = ?", "1").Updates(map[string]interface{}{
		"middle_name": "",
		"last_name":   "Morro",
	}).Error
	assert.Nil(t, err)

	err = db.Model(&User{}).Where("id = ?", "1").Update("password", "diubahlagi").Error
	assert.Nil(t, err)

	err = db.Where("id = ?", "1").Updates(User{
		Name: Name{
			FirstName: "Eko",
			LastName:  "Khannedy",
		},
	}).Error
	assert.Nil(t, err)
}

func TestDelete(t *testing.T) {

	user := User{}
	db.Take(&user, "id = ?", "90")
	tx := db.Delete(&user)
	assert.Nil(t, tx.Error)

	tx = db.Delete(&User{}, "id = ?", "91")
	assert.Nil(t, tx.Error)

	tx = db.Where("id = ?", "92").Delete(&User{})
	assert.Nil(t, tx.Error)
}

func TestLocking(t *testing.T) {

	err := db.Transaction(func(tx *gorm.DB) error {
		user := User{}
		tx.Clauses(clause.Locking{Strength: "UPDATE"}).Take(&user, "id = ?", "1")

		user.Password = "very-secret"
		err := tx.Save(&user).Error
		return err
	})
	assert.Nil(t, err)
}

func TestInsertRelation(t *testing.T) {
	wallet := Wallet{
		ID:      "10",
		UserId:  "1",
		Balance: 100,
	}

	err := db.Create(&wallet).Error
	assert.Nil(t, err)
}

func TestRetrieveRelation(t *testing.T) {
	user := User{}
	err := db.Model(&User{}).Preload("Wallet").Take(&user, "id = ?", "1").Error
	assert.Nil(t, err)
	assert.Equal(t, "1", user.ID)
	assert.Equal(t, "1", user.Wallet.ID)
	assert.Equal(t, int64(100), user.Wallet.Balance)
}

func TestRetrieveRelationJoin(t *testing.T) {
	user := User{}
	err := db.Model(&User{}).Joins("Wallet").Take(&user, "users.id = ?", "1").Error
	assert.Nil(t, err)
	assert.Equal(t, "1", user.ID)
	assert.Equal(t, "1", user.Wallet.ID)
	assert.Equal(t, int64(100), user.Wallet.Balance)
}

func TestUserAndAddresses(t *testing.T) {
	user := User{
		ID:       "201",
		Password: "rahasia",
		Name: Name{
			FirstName: "User 50",
		},
		Wallet: Wallet{
			ID:      "201",
			UserId:  "200",
			Balance: 1000000,
		},
		Addresses: []Address{
			{
				UserId:  "200",
				Address: "Jalan A",
			},
			{
				UserId:  "200",
				Address: "Jalan B",
			},
		},
	}

	createUser := User{
		ID:       user.ID,
		Password: user.Password,
		Name:     user.Name,
	}
	err := db.Create(&createUser).Error
	assert.Nil(t, err)

	// Then create the wallet referencing the created user
	wallet := user.Wallet
	err = db.Create(&wallet).Error

	addresses := user.Addresses
	err = db.Create(&addresses).Error
	assert.Nil(t, err)
}

func TestPreloadJoinOneToMany(t *testing.T) {
	var user []User
	err := db.Model(&User{}).Preload("Addresses").Joins("Wallet").Find(&user).Error
	assert.Nil(t, err)
	fmt.Println(user[0].Addresses)
}

func TestBelongsTo(t *testing.T) {
	fmt.Println("Preload")
	var addresses []Address
	err := db.Model(&Address{}).Preload("User").Find(&addresses).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(addresses))

	fmt.Println("Joins")
	addresses = []Address{}
	err = db.Model(&Address{}).Joins("User").Find(&addresses).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(addresses))
}

func TestBelongsToWallet(t *testing.T) {
	fmt.Println("Preload")
	var wallets []Wallet
	err := db.Model(&Wallet{}).Preload("User").Find(&wallets).Error
	assert.Nil(t, err)

	fmt.Println("Joins")
	wallets = []Wallet{}
	err = db.Model(&Wallet{}).Joins("User").Find(&wallets).Error
	assert.Nil(t, err)
}

func TestCreateManyToMany(t *testing.T) {
	product := Product{
		ID:    "P001",
		Name:  "Contoh Product",
		Price: 1000000,
	}
	err := db.Create(&product).Error
	assert.Nil(t, err)

	err = db.Table("user_like_product").Create(map[string]interface{}{
		"user_id":    "200",
		"product_id": "P001",
	}).Error
	assert.Nil(t, err)

	err = db.Table("user_like_product").Create(map[string]interface{}{
		"user_id":    "201",
		"product_id": "P001",
	}).Error
	assert.Nil(t, err)
}

func TestPreloadManyToMany(t *testing.T) {
	var product Product
	err := db.Preload("LikedByUsers").Take(&product, "id = ?", "P001").Error
	assert.Nil(t, err)
	assert.Equal(t, 2, len(product.LikedByUsers))
}

func TestPreloadManyToManyUser(t *testing.T) {
	var user User
	err := db.Preload("LikeProducts").Take(&user, "id = ?", "1").Error
	assert.Nil(t, err)
	assert.Equal(t, 1, len(user.LikeProducts))
}

func TestAssociationFind(t *testing.T) {
	var product Product
	err := db.Take(&product, "id = ?", "P001").Error
	assert.Nil(t, err)

	var users []User
	err = db.Model(&product).Where("users.first_name LIKE ?", "User%").Association("LikedByUsers").Find(&users)
	assert.Nil(t, err)
	//assert.Equal(t, 1, len(users))
}

func TestAssociationAppend(t *testing.T) {
	var user User
	err := db.Take(&user, "id = ?", "200").Error
	assert.Nil(t, err)

	var product Product
	err = db.Take(&product, "id = ?", "P001").Error
	assert.Nil(t, err)

	err = db.Model(&product).Association("LikedByUsers").Append(&user)
	assert.Nil(t, err)
}

func TestAssociationReplace(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		var user User
		err := tx.Take(&user, "id = ?", "1").Error
		assert.Nil(t, err)

		wallet := Wallet{
			ID:      "01",
			UserId:  user.ID,
			Balance: 1000000,
		}

		err = tx.Model(&user).Association("Wallet").Replace(&wallet)
		return err
	})
	assert.Nil(t, err)
}

func TestAssociationDelete(t *testing.T) {
	var user User
	err := db.Take(&user, "id = ?", "3").Error
	assert.Nil(t, err)

	var product Product
	err = db.Take(&product, "id = ?", "P001").Error
	assert.Nil(t, err)

	err = db.Model(&product).Association("LikedByUsers").Delete(&user)
	assert.Nil(t, err)
}

func TestAssociationClear(t *testing.T) {
	var product Product
	err := db.Take(&product, "id = ?", "P001").Error
	assert.Nil(t, err)

	err = db.Model(&product).Association("LikedByUsers").Clear()
	assert.Nil(t, err)
}

func TestPreloadingWithCondition(t *testing.T) {
	var user User
	err := db.Preload("Wallet", "balance > ?", 100).Take(&user, "id = ?", "1").Error
	assert.Nil(t, err)

	fmt.Println(user)
}

func TestPreloadingNested(t *testing.T) {
	var wallet Wallet
	err := db.Preload("User.Addresses").Take(&wallet, "id = ?", "2").Error
	assert.Nil(t, err)

	fmt.Println(wallet)
	fmt.Println(wallet.User)
	fmt.Println(wallet.User.Addresses)
}

func TestPreloadingAll(t *testing.T) {
	var user User
	err := db.Preload(clause.Associations).Take(&user, "id = ?", "1").Error
	assert.Nil(t, err)
}

func TestJoinQuery(t *testing.T) {
	var users []User
	err := db.Joins("join wallets on wallets.user_id = users.id").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(users))

	users = []User{}
	err = db.Joins("Wallet").Find(&users).Error // left join
	assert.Nil(t, err)
	assert.Equal(t, 17, len(users))
}

func TestJoinWithCondition(t *testing.T) {
	var users []User
	err := db.Joins("join wallets on wallets.user_id = users.id AND wallets.balance > ?", 500000).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(users))

	users = []User{}
	err = db.Joins("Wallet").Where("Wallet.balance > ?", 500000).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(users))
}

func TestCount(t *testing.T) {
	var count int64
	err := db.Model(&User{}).Joins("Wallet").Where("Wallet.balance > ?", 500000).Count(&count).Error
	assert.Nil(t, err)
	assert.Equal(t, int64(4), count)
}

type AggregationResult struct {
	TotalBalance int64
	MinBalance   int64
	MaxBalance   int64
	AvgBalance   float64
}

func TestAggregation(t *testing.T) {
	var result AggregationResult
	err := db.Model(&Wallet{}).Select("sum(balance) as total_balance", "min(balance) as min_balance",
		"max(balance) as max_balance", "avg(balance) as avg_balance").Take(&result).Error
	assert.Nil(t, err)

	assert.Equal(t, int64(4000000), result.TotalBalance)
	assert.Equal(t, int64(1000000), result.MinBalance)
	assert.Equal(t, int64(1000000), result.MaxBalance)
	assert.Equal(t, float64(1000000), result.AvgBalance)
}

func TestAggregationGroupByAndHaving(t *testing.T) {
	var results []AggregationResult
	err := db.Model(&Wallet{}).Select("sum(balance) as total_balance", "min(balance) as min_balance",
		"max(balance) as max_balance", "avg(balance) as avg_balance").
		Joins("User").Group("User.id").Having("sum(balance) > ?", 500000).
		Find(&results).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(results))
}

func TestContext(t *testing.T) {
	ctx := context.Background()

	var users []User
	err := db.WithContext(ctx).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 17, len(users))
}

func BrokeWalletBalance(db *gorm.DB) *gorm.DB {
	return db.Where("balance = ?", 0)
}

func SultanWalletBalance(db *gorm.DB) *gorm.DB {
	return db.Where("balance > ?", 1000000)
}

func TestScopes(t *testing.T) {
	var wallets []Wallet
	err := db.Scopes(BrokeWalletBalance).Find(&wallets).Error
	assert.Nil(t, err)

	wallets = []Wallet{}
	err = db.Scopes(SultanWalletBalance).Find(&wallets).Error
	assert.Nil(t, err)
}
