# Checkpoint trainer

⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⡀⠠⠤⠀⣀⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀  
⠀⠀⠀⠀⣀⢤⡒⠉⠁⠀⠒⢂⡀⠀⠀⠀⠈⠉⣒⠤⣀⠀⠀⠀⠀  
⠀⠀⣠⠾⠅⠈⠀⠙⠀⠀⠀⠈⠀⠀⢀⣀⣓⡀⠉⠀⠬⠕⢄⠀⠀  
⠀⣰⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡤⠶⢦⡀⠑⠀⠀⠀⠀⠈⢧⠀  
⠀⡇⠀⠀⠀⠀⠀⢤⣀⣀⣀⣀⡀⢀⣀⣀⠙⠀⠀⠀⠀⠀⠀⢸⡄  
⠀⢹⡀⠀⠀⠀⠀⡜⠁⠀⠀⠙⡴⠁⠀⠀⠱⡄⠀⠀⠀⠀⠀⣸⠀  
⠀⠀⠱⢄⡀⠀⢰⣁⣒⣒⣂⣰⣃⣀⣒⣒⣂⢣⠀⠀⠀⢀⡴⠁⠀  
⠀⠀⠀⠀⠙⠲⢼⡀⠀⠙⠀⢠⡇⠀⠛⠀⠀⣌⣀⡤⠖⠉⠀⠀⠀  
⠀⠀⠀⠀⠀⠀⢸⡗⢄⣀⡠⠊⠈⢦⣀⣀⠔⡏⠀⠀⠀⠀⠀⠀⠀  
⠀⠀⠀⠀⠀⠀⠈⡇⠀⢰⠁⠀⠀⠀⢣⠀⠀⣷⠀⠀⠀⠀This app will choose a serie of exercices for you, no more excuses.  
⠀⠀⠀⠀⣠⠔⠊⠉⠁⡏⠀⠀⠀⠀⠘⡆⠤⠿⣄⣀⠀⠀⠀⠀⠀  
⠀⠀⠀⠀⣧⠸⠒⣚⡩⡇⠀⠀⠀⠀⠀⣏⣙⠒⢴⠈⡇⠀⠀⠀⠀  
⠀⠀⠀⠀⠈⠋⠉⠀⠀⢳⡀⠀⠀⠀⣸⠁⠈⠉⠓⠚⠁⠀⠀⠀⠀  
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠓⠛⠛

## Usage

```sh
git clone https://github.com/nicgen/checkpoint-trainer.git
cd checkpoint-trainer
go run app/main.go
# your tests are now created, go to the test directory and begin your work
cd test/
# good luck with your revisions 
```

Notes:
- some exercices have placeholders to avoid errors, just erase them.
- `PrintRune` is already integrated
- before running any test, uncomment the tested function in your main.go
- once in the `test` directory, run `go run .`

## Architecture

```sh
├── app
│   ├── .subjects
│   ├── go.mod
│   └── main.go
├── fetch_files.sh
├── README.md
└── test
```

- `.subjects` contains all the subjects files (md) and templates (go)
- `test` is the folder where the app will gather your test files (md and go)
- `fetch_files.sh` is the script that can fetch and adapt the files from the edu01 repo.

Note: **the test folder will be entirely erased after each run**.

## Participate

If you find some errors or a better way to handle some exercises, please fork-it and submit a pull request.

## Attribution

Thanks to RAFTANDJANI DJIHADI aka Rafta for the idea and part of the code.  
All the exercises and subjects proposed belong to [01Edu](https://github.com/01-edu/public).
