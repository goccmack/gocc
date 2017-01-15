regenerate:
	make -C t1 regenerate
	# TODO: Remove once https://github.com/goccmack/gocc/issues/36 gets resolved.
	find . -type f -name '*.go' | xargs -I '{}' goimports -w '{}'
