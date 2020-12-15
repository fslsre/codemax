1. Set your api keys in .env file

2. Run below command in project directory:
go run app.go

2. To Send email run below command from terminal:
curl --location --request GET 'localhost:8888' \
--header 'Authorization: Basic cm9vdDpyb290' \
--header 'Content-Type: application/json' \
--data-raw '{
    "from_email": "faisal.iete@gmail.com",
	"from_name": "Faisal",
	"to_email": "faisal.iete@gmail.com",
	"to_name": "Faisal",
	"subject": "Test Subject",
	"text_part": "My first Mailjet email",
	"html_part": "<h3>Dear passenger 1, welcome to <a href='\''https://www.mailjet.com/'\''>Mailjet</a>!</h3><br />May the delivery force be with you!",
    "custom_id": "AppGettingStartedTest"
}'
