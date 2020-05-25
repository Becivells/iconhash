#!/usr/bin/env bats

load test_helper

# 1. 检查版本是否为当天编译的
@test "Test iconhash version" {
  run ${ICONHASH_BIN} -v
  [ "$status" -eq 0 ]
  [ "${lines[0]%% *}" == "Tag:" ]
  [ "${lines[1]%% *}" == "Version:" ]
  [ "${lines[2]%% *}" == "Compile:" ]
  [ "${lines[3]%% *}" == "Branch:" ]
  [ $(expr "${lines[2]}" : "Compile: $(date +'%Y-%m-%d').*") -ne 0 ]   # 检查当前版本是否为今日编译的
}


# 2. 执行 iconhash -h 是否正常
@test "Check iconhash help" {
  cd ${ICONHASH_DEV_DIRNAME} && bin/iconhash 2> ${BATS_TMP_DIRNAME}/${BATS_TEST_NAME}.golden
  run golden_diff
  echo "${output}"
  [ $status -eq 0 ]
}

# 3. 检查执行 iconhash url 是否正常
@test "Check iconhash url" {
  ${ICONHASH_BIN} https://raw.githubusercontent.com/Becivells/iconhash/master/test/data/dede.ico > ${BATS_TMP_DIRNAME}/${BATS_TEST_NAME}.golden
  run golden_diff
  echo "${output}"
  [ $status -eq 0 ]
}
# 4. 检查执行 iconhash file 是否正常
@test "Check iconhash file" {
  ${ICONHASH_BIN} ${TEST_DATA}/favicon.ico > ${BATS_TMP_DIRNAME}/${BATS_TEST_NAME}.golden
  run golden_diff
  echo "${output}"
  [ $status -eq 0 ]
}


# 5. 检查执行 iconhash file not exist 是否正常
@test "Check iconhash file not exist" {
  run ${ICONHASH_BIN} ${TEST_DATA}/favicon.ico1 > ${BATS_TMP_DIRNAME}/${BATS_TEST_NAME}.golden
  run golden_diff
  echo "${output}"
  [ $status -eq 1 ]
}

# 6. 检查执行 iconhash args url 是否正常
@test "Check iconhash args url" {
  ${ICONHASH_BIN} -url https://raw.githubusercontent.com/Becivells/iconhash/master/test/data/dede.ico > ${BATS_TMP_DIRNAME}/${BATS_TEST_NAME}.golden
  run golden_diff
  echo "${output}"
  [ $status -eq 0 ]
}

# 7. 检查执行 iconhash args url shodan 是否正常
@test "Check iconhash args url shodan" {
  ${ICONHASH_BIN} -url https://raw.githubusercontent.com/Becivells/iconhash/master/test/data/dede.ico -shodan > ${BATS_TMP_DIRNAME}/${BATS_TEST_NAME}.golden
  run golden_diff
  echo "${output}"
  [ $status -eq 0 ]
}

# 8. 检查执行 iconhash args file 是否正常
@test "Check iconhash args file" {
  ${ICONHASH_BIN} -file ${TEST_DATA}/favicon.ico  > ${BATS_TMP_DIRNAME}/${BATS_TEST_NAME}.golden
  run golden_diff
  echo "${output}"
  [ $status -eq 0 ]
}

# 9. 检查执行 iconhash args file shodan 是否正常
@test "Check iconhash args file shodan" {
  ${ICONHASH_BIN} -file ${TEST_DATA}/favicon.ico -shodan > ${BATS_TMP_DIRNAME}/${BATS_TEST_NAME}.golden
  run golden_diff
  echo "${output}"
  [ $status -eq 0 ]
}

# 10. 检查执行 iconhash args file no fofa shodan 是否正常
@test "Check iconhash args file no fofa shodan" {
  ${ICONHASH_BIN} -file ${TEST_DATA}/favicon.ico -shodan -fofa=false > ${BATS_TMP_DIRNAME}/${BATS_TEST_NAME}.golden
  run golden_diff
  echo "${output}"
  [ $status -eq 0 ]
}

# 11. 检查执行 iconhash args file no fofa no shodan 是否正常
@test "Check iconhash args file no fofa no shodan" {
  ${ICONHASH_BIN} -file ${TEST_DATA}/favicon.ico -shodan=false -fofa=false > ${BATS_TMP_DIRNAME}/${BATS_TEST_NAME}.golden
  run golden_diff
  echo "${output}"
  [ $status -eq 0 ]
}

# 12. 检查执行 iconhash args file no fofa 是否正常
@test "Check iconhash args file no fofa" {
  ${ICONHASH_BIN} -file ${TEST_DATA}/favicon.ico  -fofa=false > ${BATS_TMP_DIRNAME}/${BATS_TEST_NAME}.golden
  run golden_diff
  echo "${output}"
  [ $status -eq 0 ]
}

# 13. 检查执行 iconhash args file no fofa 是否正常
@test "Check iconhash args file no fofa" {
  ${ICONHASH_BIN} -file ${TEST_DATA}/favicon.ico  -fofa=false > ${BATS_TMP_DIRNAME}/${BATS_TEST_NAME}.golden
  run golden_diff
  echo "${output}"
  [ $status -eq 0 ]
}

# 13. 检查执行 iconhash args other file 是否正常
@test "Check iconhash args other file" {
  ${ICONHASH_BIN} -file ${TEST_DATA}/dede.ico  -fofa=false > ${BATS_TMP_DIRNAME}/${BATS_TEST_NAME}.golden
  run golden_diff
  echo "${output}"
  [ $status -eq 0 ]
}

# 14. 检查执行 iconhash url 是否正常
@test "Check iconhash other url" {

  ${ICONHASH_BIN} https://raw.githubusercontent.com/Becivells/iconhash/master/test/data/favicon.ico > ${BATS_TMP_DIRNAME}/${BATS_TEST_NAME}.golden
  run golden_diff
  echo "${output}"

  [ $status -eq 0 ]
}

# 13. 检查执行 iconhash args other file uint32 是否正常
@test "Check iconhash args other file uint32" {
  ${ICONHASH_BIN} -file ${TEST_DATA}/dede.ico  -fofa=false -uint32 > ${BATS_TMP_DIRNAME}/${BATS_TEST_NAME}.golden
  run golden_diff
  echo "${output}"
  [ $status -eq 0 ]
}