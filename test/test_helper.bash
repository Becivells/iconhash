setup() {
  export ICONHASH_DEV_DIRNAME="${BATS_TEST_DIRNAME}/.."
  export ICONHASH_BIN="${ICONHASH_DEV_DIRNAME}/bin/iconhash"
  export TEST_DATA="${BATS_TEST_DIRNAME}/data"
  export BATS_TMP_DIRNAME="${BATS_TEST_DIRNAME}/tmp"
  export BATS_FIXTURE_DIRNAME="${BATS_TEST_DIRNAME}/fixture"
  export LC_ALL=C # Linux macOS 下 sort 排序问题
  mkdir -p "${BATS_TMP_DIRNAME}"
}

# golden_diff like gofmt golden file check method, use this function check output different with template
golden_diff() {
  diff "${BATS_TMP_DIRNAME}/${BATS_TEST_NAME}.golden" "${BATS_FIXTURE_DIRNAME}/${BATS_TEST_NAME}.golden"
}