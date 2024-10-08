#! /usr/bin/env bash

# NEW subjects
arr=(CamelToSnakeCase FifthAndSkip)

# CHANGELOG
# 09/09/2024 added: CamelToSnakeCase FifthAndSkip ExpandStr CountCharacter FromTo PrintRevComb ReverseStrCap
# 08/02/2024 added: itoa atoi
# 07/31/2024 added: digitlen itoabase notdecimal concatslice iscapitalized revconcatalternate

# base PATH
base=/home/nic/dev/golang/go-checkpoint-trainer/app/.subjects

# FUNC
lower() {
  printf '%s' "${1,,}"
}

for elmnt in "${arr[@]}"; do
  printf '%s\n' "$elmnt"
  lower="$(lower $elmnt)"
  file="$base/$lower"

  if [ -f "$file.go" ]; then
    echo "$elmnt.go already exists"
  else 
    echo "$elmnt.go does not exist."
    if curl "https://raw.githubusercontent.com/01-edu/public/master/subjects/$lower/main.go" -o "$file.go"; then
      echo "$elmnt.go retrieved"
    else
      echo "$elmnt.go failed"
    fi
  fi

  # replace occurences
  echo "clean $elmnt.go"
  sed -i 's/   "piscine".*//g' "$file.go"
  sed -i 's/	fmt.Println(piscine./	fmt.Println(/g' "$file.go"
  sed -i "s/func main() {/func $elmnt() {/g" "$file.go"

  sleep 1

  if [ -f "$file.md" ]; then
    echo "$elmnt.md already exists"
  else
    if curl "https://raw.githubusercontent.com/01-edu/public/master/subjects/$lower/README.md" -o "$file.md"; then
      echo "$elmnt.md retrieved"
    else
      echo "$elmnt.md failed"
    fi
  fi
  sleep 1
  echo "------------------------------------------"
done

