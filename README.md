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

# Example

Input image

<img src="cat.jpeg" width=300px>

Output

``` shell
Color 1: #C99165 https://www.color-hex.com/color/C99165
Color 2: #3C3538 https://www.color-hex.com/color/3C3538
Color 3: #1D1208 https://www.color-hex.com/color/1D1208
```