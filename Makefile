.PHONY: up-db
up-db: 
	sudo docker container restart postgres

.PHONY: env

define ENV_SAMPLE
SMTP_EMAIL=yr.olimpiev@gmail.com
SMTP_PASSWORD=stkkftucujhicqxp
SMTP_HOST=smtp.gmail.com
SMTP_PORT=576
DB_DSN=postgres://postgres:password@localhost:5431/postgres?sslmode=disable
endef
export ENV_SAMPLE

env:
	@if [ ! -f ".env" ];\
		then echo "$$ENV_SAMPLE" > .env;\
	 fi

