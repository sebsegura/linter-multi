install-hooks: hooks-check-deps
	@husky init
	@husky add pre-push ""
	yes | cp -r githooks/. .husky/hooks

hooks-check-deps:
	@if [ -z `which husky` ]; then \
		echo "installing husky";\
		go install github.com/automation-co/husky@latest;\
	fi