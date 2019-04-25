version_tag = `scm-status | jq --raw-output '.tags[0]' | cut -c 2-`

default:
	docker build -t jongschneider/wait-for-elasticsearch .
	docker tag jongschneider/wait-for-elasticsearch:latest jongschneider/wait-for-elasticsearch:$(version_tag)
	docker push jongschneider/wait-for-elasticsearch:latest
	docker push jongschneider/wait-for-elasticsearch:$(version_tag)
