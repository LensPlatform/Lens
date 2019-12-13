package tests

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-kit/kit/metrics/discard"
	"github.com/jinzhu/gorm"
	"github.com/opentracing/opentracing-go"
	"github.com/openzipkin/zipkin-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"

	"github.com/LensPlatform/Lens/src/pkg/config"
	"github.com/LensPlatform/Lens/src/pkg/endpoint"
	"github.com/LensPlatform/Lens/src/pkg/service"
	"github.com/LensPlatform/Lens/src/pkg/transport"
)

func TestHTTP(t *testing.T) {
	var err error
	// Load config file
	err = config.LoadConfig()
	if err != nil {
		return
	}

	logger := zaptest.NewLogger(t)
	zkt, _ := zipkin.NewTracer(nil, zipkin.WithNoopTracer(true))
	db, err := initDbConnection(logger)
	if err == nil {
		return
	}

	defer db.Close()
	amqpproducerconn, amqpconsumerconn := initQueues(err, logger)

	svc := service.New(logger, db, amqpproducerconn, amqpconsumerconn, discard.NewCounter(), discard.NewCounter(),
		discard.NewCounter(), discard.NewCounter(), discard.NewCounter(),
		discard.NewCounter(), discard.NewCounter(), discard.NewCounter())
	eps := endpoint.New(svc, logger, discard.NewHistogram(), opentracing.GlobalTracer(), zkt)

	mux := transport.NewHTTPHandler(svc, eps, discard.NewHistogram(),  opentracing.GlobalTracer(), zkt, logger)
	srv := httptest.NewServer(mux)
	defer srv.Close()

	for _, testcase := range []struct {
		method, url, body, want string
	}{
		{       "POST",
				srv.URL + "/v1/user/create",
				`
				{"User":{"id":"","first_name":"yvan1","last_name":"yomba1","user_name":"yvanyomba1","email":"yvanyomba1@gmail.com",
				"password":"Granada123!","password_confirmed":"Granada123!","age":24,"birth_date":"1996","phone_number":"07/12/1996",
				"location":null,"bio":"Hello All!","education":null,"interests":{"industries_of_interest":null,
				"topics_of_interest":null},"headline":"I am new here","subscriptions":null,"intent":"investors"}
				}`,
				`{"User":{"id":"","first_name":"yvan1","last_name":"yomba1","user_name":"yvanyomba1","email":"yvanyomba1@gmail.com",
				"password":"Granada123!","password_confirmed":"Granada123!","age":24,"birth_date":"1996","phone_number":"07/12/1996",
				"location":null,"bio":"Hello All!","education":null,"interests":{"industries_of_interest":null,
				"topics_of_interest":null},"headline":"I am new here","subscriptions":null,"intent":"investors"}
				}`},
	} {
		req, _ := http.NewRequest(testcase.method, testcase.url, strings.NewReader(testcase.body))
		resp, _ := http.DefaultClient.Do(req)
		body, _ := ioutil.ReadAll(resp.Body)
		if want, have := testcase.want, strings.TrimSpace(string(body)); want != have {
			t.Errorf("%s %s %s: want %q, have %q", testcase.method, testcase.url, testcase.body, want, have)
		}
	}
}

func initDbConnection(zapLogger *zap.Logger) (*gorm.DB, error) {
	connString := config.Config.GetDatabaseConnectionString()
	db, err := gorm.Open("postgres", connString)
	if err != nil {
		zapLogger.Error(err.Error())
		os.Exit(1)
	}

	zapLogger.Info("successfully connected to database", )
	return db, err
}

func initQueues(err error, zapLogger *zap.Logger) (service.Queue, service.Queue) {
	// connect to rabbitmq
	amqpConnString := "amqp://user:bitnami@stats/"
	producerQueueNames := []string{"lens_welcome_email", "lens_password_reset_email", "lens_email_reset_email"}
	consumerQueueNames := []string{"user_inactive"}
	amqpproducerconn, err := service.NewAmqpConnection(amqpConnString, producerQueueNames)
	if err != nil {
		zapLogger.Error(err.Error())
	}
	amqpconsumerconn, err := service.NewAmqpConnection(amqpConnString, consumerQueueNames)
	if err != nil {
		zapLogger.Error(err.Error())
	}
	return amqpproducerconn, amqpconsumerconn
}