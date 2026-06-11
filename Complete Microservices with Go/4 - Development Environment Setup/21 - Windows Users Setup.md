# WSL & Go Setup: Short & Sweet Summary

Bro, here is the exact breakdown of the Windows setup lecture. Zero stress, just the core facts on what we did and why!

### 1. Installing WSL (The Fake Linux)
* **What we did:** We opened PowerShell and ran `wsl --install` to put a mini Ubuntu OS inside Windows [cite: 724].
* **Why we did it:** Because the self-taught developer is a Mac user, and coding these cloud tools natively on Windows is a headache. WSL gives us a perfect, hidden Linux environment to work in [cite: 721, 726].

### 2. Using `wget` (The Linux Downloader)
* **What we did:** We used the command `wget` to grab the Go setup file [cite: 797, 815].
* **Why we did it:** `wget` is a classic Linux tool. We used it to download the file directly from the internet right into our terminal without ever opening Chrome [cite: 811, 812].

### 3. Setting Up `GOROOT` (The Engine Room)
* **What we did:** We extracted the Go folder and moved it to `/usr/local`. Then we added `export GOROOT=/usr/local/go` to our `.bashrc` settings file [cite: 801, 822, 823].
* **Why we did it:** To tell the blind terminal exactly where we parked the core Go compiler software so it knows where to run the engine from [cite: 803, 822].

### 4. Setting Up `GOPATH` (Your Workspace)
* **What we did:** We added `export GOPATH=$HOME/go` to our settings [cite: 825]. 
* **Why we did it:** Go is strict and refuses to mix internet libraries with its core engine [cite: 838, 839]. This line tells Go: *"Whenever I download a third-party package (like a web framework), save it safely in this dedicated home folder!"* [cite: 843].

### 5. The Master Shortcut (`$PATH`)
* **What we did:** We added `export PATH=$GOPATH/bin:$GOROOT/bin:$PATH` to our settings.
* **Why we did it:** So we don't have to type out the ridiculously long folder address every time. This maps Go to the terminal's master shortcut list so we can just type `go` from anywhere [cite: 827, 830].