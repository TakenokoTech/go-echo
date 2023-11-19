items=(
  "core"
  "todos"
  "users"
)

for item in "${items[@]}" ; do
    mockgen \
      -source="./domain/${item}/repository.go" \
      -destination="./domain/${item}/repository_mock.go" \
      -package="${item}"
done
