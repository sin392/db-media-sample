##################################################
### generate 系
##################################################
.PHONY: generate
generate: generate/proto ## generate系のtargetを一括で実行します

.PHONY: generate/proto
generate/proto: ## protoファイルのコード生成を実行します
	cd proto && buf generate

##################################################
### docs 系
##################################################

.PHONY: docs
docs: docs/deps ## ドキュメント生成系のtargetを一括で実行します

.PHONY: docs/deps
docs/deps: ## パッケージの依存関係のドキュメント生成をします
	gomod graph --packages --annotate 'github.com/sin392/db-media-sample/sample/internal/**/*' | \
	sed -e 's|github.com/sin392/db-media-sample/sample/internal/||g' | \
	grep -v -e 'golang.org' -e 'app' -e 'config' -e 'errors' | \
	python3 ../scripts/color_dot_graph.py | \
	dot -Tpng -o docs/graph/deps.png

##################################################
### lint 系
##################################################

.PHONY: lint
lint: lint/proto lint/cspell ## lint系のtargetを一括実行します

.PHONY: lint/proto
lint/proto: ## protoファイルのlintを実行します
	cd proto && buf lint

.PHONY: lint/cspell
lint/cspell: ## cspell によるスペルチェックを実行します
	npx cspell lint --no-progress --show-context .

##################################################
### 依存関係のupdate系
##################################################

.PHONY: update
update: update/proto update/link update/cspell ## 依存関係のupdate系のtargetを一括実行します

.PHONY: update/proto
update/proto: ## bufで管理しているprotoの依存関係を更新します
	cd proto && buf mod update

.PHONY: update/link
update/link: ## 外部protoのsymlinkを作成
	ln -sf ../../../../vendor/grpc-graphql-gateway/include/graphql/graphql.proto proto/vendor/grpc-graphql-gateway/graphql.proto

.PHONY: update/cspell
update/cspell: ## cspellで使っているカスタム辞書を更新します。実行後に差分を見てtypoが含まれていないか確認してください。
	npx cspell --words-only --unique . | LC_ALL=C sort -u --ignore-case | LC_ALL=C tr '[:upper:]' '[:lower:]' > ../.cspell/project-words.txt
