package discount

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hfleury/henrybillogram/internal"
	"github.com/hfleury/henrybillogram/internal/app"
	"github.com/hfleury/henrybillogram/pkg/discount/repo"
	"github.com/hfleury/henrybillogram/pkg/discount/service"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ginEngine *gin.Engine
var dbConn *gorm.DB

func TestMain(m *testing.M) {
	os.Setenv("BILLO_ENVIRONMENT", "TEST")
	appBlg, errapp := app.NewAppBillogram()
	if errapp != nil {
		log.Fatalf("error initialing the App billogram %v", errapp)
	}

	var err error
	dsn := "host=localhost user=rootuser password=nosecret dbname=billodb port=5432 sslmode=disable"
	dbConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connect to db: %v", err)
	}

	blgRouter := internal.NewBlgRouter(appBlg)
	ginEngine = blgRouter.ConfigRouter()
	dscRepo := repo.NewDiscountBrandRepo(dbConn)
	dscUserRepo := repo.NewDiscountUserRepo(dbConn)
	dscService := service.NewDiscountBrandService(*dscRepo)
	dscUserService := service.NewDiscountUserService(*dscUserRepo)
	dscHandler := NewDiscountHandler(dscService, dscUserService)
	dscRouter := NewDiscountRouter(appBlg, dscHandler)
	dscRouter.SetDiscountRouters()

	m.Run()
}

func TestDiscountRouter(t *testing.T) {

	t.Run("Success Creating 10 discounts", func(t *testing.T) {
		w := httptest.NewRecorder()
		requestBodys, err := json.Marshal(map[string]int{
			"amountDiscount": 10,
		})
		assert.NoError(t, err)

		req, err := http.NewRequest("POST", "/brand/discount", bytes.NewBuffer(requestBodys))
		req.Header = http.Header{
			"Blg-Brand-Id":   {"1"},
			"Blg-Brand-Name": {"brand1"},
		}
		ginEngine.ServeHTTP(w, req)

		assert.NoError(t, err)
		assert.Equal(t, 200, w.Code)
	})

	t.Run("Fail: Brand Id required", func(t *testing.T) {
		w := httptest.NewRecorder()
		requestBody, err := json.Marshal(map[string]string{
			"amountDiscount": "10",
		})
		assert.NoError(t, err)

		req, err := http.NewRequest("POST", "/brand/discount", bytes.NewBuffer(requestBody))
		ginEngine.ServeHTTP(w, req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusPreconditionRequired, w.Code)
		assert.Equal(t, "{\"error\":\"Brand ID and Name required\"}", w.Body.String())
	})

	t.Run("Fail: Invalid brand ID", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, err := http.NewRequest("POST", "/brand/discount", nil)
		req.Header = http.Header{
			"Blg-Brand-Id": {"987"},
		}

		ginEngine.ServeHTTP(w, req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusPreconditionRequired, w.Code)
		assert.Equal(t, "{\"error\":\"Brand ID and Name required\"}", w.Body.String())
	})
}
