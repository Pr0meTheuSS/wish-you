
.PHONY: env

define ENV_SAMPLE
SMTP_EMAIL=yr.olimpiev@gmail.com
SMTP_PASSWORD=stkkftucujhicqxp
SMTP_HOST=smtp.gmail.com
SMTP_PORT=576
endef
export ENV_SAMPLE
env:
	@if [ ! -f ".env" ];\
		then echo "$$ENV_SAMPLE" > .env;\
	 fi

