# Go Installation

Follow these steps to install and verify Go on your system.

## 1. Download and Install Go

Go to the official download page: [https://go.dev/dl](https://go.dev/dl)

- **macOS:** Use the `.pkg` installer or run `brew install go`.
- **Windows:** Use the `.msi` installer or run `choco install golang`.

> Both installers automatically set up your environment.

### Linux:
```
tar -C /usr/local -xzf go1.20.5.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> $HOME/.bash_profile
source $HOME/.bash_profile
```

> Use `sudo` if necessary when writing to `/usr/local`.

## 2. Verify the Installation

Open your terminal and run:
```
go version
```

Example output:
```
go version go1.20.5 linux/amd64
```

That’s it — Go is now installed and ready to use.


## 📂 Code

```go
// 01_hello_world.go
package main

import "fmt"

func main() {
	// Hello :)
	fmt.Println("Hello, World!")
}
```
