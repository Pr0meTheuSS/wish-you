package service

/*
 * Project: I-wish-you
 * Created Date: Sunday, July 16th 2023, 7:12:46 pm
 * Author: Olimpiev Y. Y.
 * -----
 * Last Modified:  yr.olimpiev@gmail.com
 * Modified By: Olimpiev Y. Y.
 * -----
 * Copyright (c) 2023 NSU
 *
 * -----
 */

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strconv"

	"github.com/ztrue/tracerr"
)

func SendConfirmMail(recipientEmail string, body string) error {
	// Учетные данные для доступа к SMTP-серверу Gmail
	email := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")
	// Настройки SMTP-сервера Gmail
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return tracerr.Errorf("Cannot parse int from variable %s", "SMTP_PORT")
	}

	// Формирование аутентификационных данных
	auth := smtp.PlainAuth("", email, password, smtpHost)

	// Установка соединения с сервером SMTP
	client, err := smtp.Dial(fmt.Sprintf("%s:%d", smtpHost, smtpPort))
	if err != nil {
		log.Println("Ошибка при подключении к серверу SMTP:", err)
		return err
	}
	defer client.Close()

	// Инициируем шифрование TLS
	tlsConfig := &tls.Config{
		ServerName:         smtpHost,
		InsecureSkipVerify: false,
	}
	if err = client.StartTLS(tlsConfig); err != nil {
		log.Println("Ошибка при запуске TLS:", err)
		return err
	}

	// Устанавливаем аутентификацию
	if err = client.Auth(auth); err != nil {
		log.Println("Ошибка аутентификации:", err)
		return err
	}

	// Формирование сообщения
	msg := []byte("To:" + recipientEmail + "\r\n" +
		"Subject: Подтверждение регистрации\r\n" +
		"\r\n" +
		body)

	// Отправка письма через SMTP
	if err = client.Mail(email); err != nil {
		log.Println("Ошибка при указании адреса отправителя:", err)
		return err
	}
	if err = client.Rcpt(recipientEmail); err != nil {
		log.Println("Ошибка при указании адреса получателя:", err)
		return err
	}
	w, err := client.Data()
	if err != nil {
		log.Println("Ошибка при получении записи:", err)
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		log.Println("Ошибка при записи сообщения:", err)
		return err
	}
	err = w.Close()
	if err != nil {
		log.Println("Ошибка при закрытии записи:", err)
		return err
	}

	log.Println("Письмо успешно отправлено!")
	return nil
}
