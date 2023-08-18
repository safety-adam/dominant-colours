# Dominant Colors

This code finds the top 3 dominant colors in JPEG image.

# Usage

Modify the following line in `main.go` to read the required file.

``` golang
colors, err := FindDominantColors("cat.jpeg")
```

Run the follow in the root of the project directory.

``` zsh
% go get
% go run .
```