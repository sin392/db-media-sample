##################################################
### lint 系
##################################################

.PHONY: lint
lint: lint/cspell lint/go ## lint系のtargetを一括実行します

.PHONY: lint/go
lint/go: ## goのlintを実行します
	go list -f '{{.Dir}}/...' -m | xargs golangci-lint run

.PHONY: lint/cspell
lint/cspell: ## cspell によるスペルチェックを実行します
	npx cspell lint --no-progress --show-context .

##################################################
### 依存関係のupdate系
##################################################

.PHONY: update/cspell
update/cspell: ## cspellで使っているカスタム辞書を更新します。実行後に差分を見てtypoが含まれていないか確認してください。
	npx cspell --words-only --unique . | LC_ALL=C sort -u --ignore-case | LC_ALL=C tr '[:upper:]' '[:lower:]' > .cspell/project-words.txt
