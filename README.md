# Easy way to decrypt your md5 hashed encrypted passwords

### How to use?
* Download the program (binary)
* Aassign executable permissions to is
* Optionally - add it in your path to make it easier to use

### Commands 
* wget https://github.com/npsoni88/md5-decrypter/raw/master/md5-cli
* chmod +x md5-cli
* sudo mv md5-cli /usr/local/bin/md5-cli

### Usage
Simply execute the app with the hash, for example.

> md5-cli 5d41402abc4b2a76b9719d911017c592

You will receive this output
> Successfully connected to the hash base!
> The plain text password for your hash is => hello⏎

This tells you that the decrypted plain text is "hello". Enjoy!