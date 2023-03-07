package repo

import (
	"log"
	"reflect"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/hfleury/henrybillogram/pkg/discount"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func TestMain(m *testing.M) {
	var err error
	dsn := "host=localhost user=rootuser password=nosecret dbname=billodb port=5432 sslmode=disable"
	dbConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connect to db: %v", err)
	}
	m.Run()
}

func TestDiscountBrandRepo_InsertBatchDiscount(t *testing.T) {
	discounts := createDiscount(10, 1, "brand1")
	type fields struct {
		dbConn *gorm.DB
	}
	type args struct {
		d []discount.Discount
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []discount.Discount
		wantErr bool
	}{
		{
			name: "Success insert 10 discount codes",
			fields: fields{
				dbConn: dbConn,
			},
			args: args{
				d: discounts,
			},
			want:    discounts,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbr := &DiscountBrandRepo{
				dbConn: tt.fields.dbConn,
			}
			got, err := dbr.InsertBatchDiscount(tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("DiscountBrandRepo.InsertBatchDiscount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiscountBrandRepo.InsertBatchDiscount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createDiscount(n, brandId int, bname string) []discount.Discount {
	var discounts []discount.Discount

	for i := 0; i < n; i++ {
		discounts = append(discounts, discount.Discount{
			BrandID: brandId,
			Code:    bname + gofakeit.Word(),
		})
	}

	return discounts
}
