package server

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"os"
	"yookassa-payment-proxy/internal/models"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Post("/payment", s.YooKassaPaymentHandler)
	s.App.Get("/", s.HelloWorldHandler)

}

func (s *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World!",
	}

	return c.JSON(resp)
}

// YooKassaPaymentHandler создает платеж через API Yookassa.
func (s *FiberServer) YooKassaPaymentHandler(c *fiber.Ctx) error {
	log.Println("run yookassa handler")
	// Считываем данные платежа из запроса
	var paymentData models.YooKassaPayment
	if err := c.BodyParser(&paymentData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Ошибка при разборе JSON"})
	}

	// Преобразуем данные платежа в JSON
	jsonData, err := json.Marshal(paymentData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Ошибка при конвертации данных в JSON"})
	}
	log.Println("json input: ", paymentData)

	log.Println("byte input: ", jsonData)

	// Подготавливаем HTTP запрос к YooKassa
	req, err := http.NewRequest("POST", "https://api.yookassa.ru/v3/payments", bytes.NewBuffer(jsonData))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Ошибка при создании запроса"})
	}

	// Устанавливаем необходимые заголовки
	req.Header.Set("Idempotence-Key", uuid.New().String())
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(os.Getenv("MERCHANT_ID"), os.Getenv("SECRET_KEY"))

	// Выполняем запрос к YooKassa
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Ошибка при отправке запроса к YooKassa"})
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)

	// Читаем ответ от YooKassa
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Ошибка при чтении ответа от YooKassa"})
	}

	// Десериализуем ответ в структуру
	var paymentResponse models.YooKassaPaymentResponse
	if err := json.Unmarshal(respBody, &paymentResponse); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Ошибка при десериализации ответа от YooKassa"})
	}

	log.Println("yookassa output: ", paymentResponse)

	// Возвращаем ответ от YooKassa обратно на фронтенд
	return c.Status(resp.StatusCode).JSON(paymentResponse)
}
